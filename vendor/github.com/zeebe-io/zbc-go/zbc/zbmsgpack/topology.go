package zbmsgpack

import (
	"encoding/json"
	"fmt"
	"time"
)

// Broker is used to hold broker contact information.
type Broker struct {
	Host string `msgpack:"host"`
	Port uint64 `msgpack:"port"`
}

func (t *Broker) String() string {
	b, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return fmt.Sprintf("json marshaling failed\n")
	}
	return fmt.Sprintf("%+v", string(b))
}

func (t *Broker) Addr() string {
	return fmt.Sprintf("%s:%d", t.Host, t.Port)
}

// TopicLeader is used to hold information about the relation of broker and topic/partitions.
type TopicLeader struct {
	Broker

	TopicName   string `msgpack:"topicName"`
	PartitionID uint16 `msgpack:"partitionId"`
}

func (t *TopicLeader) String() string {
	b, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return fmt.Sprintf("json marshaling failed\n")
	}
	return fmt.Sprintf("%+v", string(b))
}

// TopologyRequest is used to make a topology request.
type TopologyRequest struct{}

func (t *TopologyRequest) String() string {
	b, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return fmt.Sprintf("json marshaling failed\n")
	}
	return fmt.Sprintf("%+v", string(b))
}

// ClusterTopologyResponse is used to parse the topology response from the broker.
type ClusterTopologyResponse struct {
	Brokers      []Broker      `msgpack:"brokers"`
	TopicLeaders []TopicLeader `msgpack:"topicLeaders"`
}

func (t *ClusterTopologyResponse) String() string {
	b, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return fmt.Sprintf("json marshaling failed\n")
	}
	return fmt.Sprintf("%+v", string(b))
}

// ClusterTopology is structure used by the client object to hold information about cluster.
type ClusterTopology struct {
	TopicLeaders map[string][]TopicLeader `msgpack:"topicLeaders"`
	Brokers      []Broker                 `msgpack:"brokers"`
	UpdatedAt    time.Time                `msgpack:"-"`
}

func (t *ClusterTopology) String() string {
	b, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return fmt.Sprintf("json marshaling failed\n")
	}
	return fmt.Sprintf("%+v", string(b))
}
