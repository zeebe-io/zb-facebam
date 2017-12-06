package zbc

import (
	"github.com/zeebe-io/zbc-go/zbc/zbmsgpack"
	"math/rand"

	"time"
)

type TopicRoundRobin struct {
	index int
	topic string
}

type topologyManager struct {
	*transportManager

	topologyRequest *requestWrapper
	topologyWorkload chan *requestWrapper

	topicRoundRobin map[string][]int
	bootstrapAddr string
	cluster             *zbmsgpack.ClusterTopology
}

func (tm *topologyManager) partitionID(topic string) (uint16, error) {
	leaders, ok := tm.cluster.TopicLeaders[topic]

	if !ok {
		tm.refreshTopology()

		leaders, ok = tm.cluster.TopicLeaders[topic]

		if !ok {
			return 0, errTopicLeaderNotFound
		}
	}

	//tm[topic].lastRoundRobinIndex++
	//if tm.lastRoundRobinIndex == len(leaders) {
	//	tm.lastRoundRobinIndex = 0
	//}

	// TODO: zbc-go/issues#40 + zbc-go/issues#48
	return leaders[0].PartitionID, nil
}

func (tm *topologyManager) initTopology() (*zbmsgpack.ClusterTopology, error) {
	resp, err := tm.executeRequest(tm.topologyRequest)

	topology := newResponseHandler().unmarshalTopology(resp)
	tm.cluster = topology
	return topology, err
}


func (tm *topologyManager) refreshTopology() (*zbmsgpack.ClusterTopology, error) {
	rand.Seed(time.Now().Unix())
	request := tm.topologyRequest
	request.addr = tm.cluster.Brokers[rand.Int()%len(tm.cluster.Brokers)].Addr()
	resp, err := tm.executeRequest(request)
	topology := newResponseHandler().unmarshalTopology(resp)
	tm.cluster = topology
	return topology, err
}


func (tm *topologyManager) getDestinationAddr(msg *Message) (string, error) {
	if msg.isTopologyMessage() && tm.cluster == nil {
		return tm.bootstrapAddr, nil
	}

	partitionId := msg.forPartitionId()
	if partitionId == nil {
		return "", brokerNotFound
	}

	for _, topicLeaders := range tm.cluster.TopicLeaders {
		for _, leader := range topicLeaders {
			if leader.PartitionID == *partitionId {
				addr := leader.Broker.Addr()
				return addr, nil
			}
		}
	}
	return "", brokerNotFound
}


func (tm *topologyManager) executeRequest(request *requestWrapper) (*Message, error) {
	addr, err := tm.getDestinationAddr(request.payload)

	if err == brokerNotFound {
		return nil, brokerNotFound
	}
	request.addr = addr
	tm.topologyWorkload <- request

	select {

	case resp := <-request.responseCh:
		return resp, nil

	case err := <-request.errorCh:
		if tm.cluster != nil {
			tm.refreshTopology()
		}
		return nil, err

	}
}

func (tm *topologyManager) topologyTicker() {
	for {
		select {
		case <-time.After(TopologyRefreshInterval * time.Second):
			if time.Since(tm.cluster.UpdatedAt) > TopologyRefreshInterval*time.Second {
				tm.topologyWorkload <- tm.topologyRequest
			}
			break
		}
	}
}

func (tm *topologyManager) topologyWorker() {
	for {
		select {
		case request := <-tm.topologyWorkload:
			tm.execTransport(request)
		}
	}
}

func newTopologyManager(bootstrapAddr string) *topologyManager {
	tm := &topologyManager{
		newTransportManager(),
		newRequestWrapper(newRequestFactory().topologyRequest()),
		make(chan *requestWrapper, requestQueueSize),
		make(map[string][]int),
		bootstrapAddr,
		nil,
	}

	go tm.topologyWorker()
	go tm.topologyTicker()

	tm.initTopology()
	return tm
}
