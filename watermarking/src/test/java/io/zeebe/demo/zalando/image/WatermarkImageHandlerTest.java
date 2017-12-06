package io.zeebe.demo.zalando.image;

import static io.zeebe.test.util.TestUtil.doRepeatedly;

import java.io.IOException;
import java.io.InputStream;
import java.net.URL;
import java.time.Duration;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.Properties;
import java.util.concurrent.CopyOnWriteArrayList;

import org.junit.After;
import org.junit.Assert;
import org.junit.Before;
import org.junit.Test;

import com.fasterxml.jackson.databind.ObjectMapper;

import io.zeebe.broker.Broker;
import io.zeebe.client.ZeebeClient;
import io.zeebe.client.event.WorkflowInstanceEvent;
import io.zeebe.demo.facebam.watermark.WatermarkImageHandler;

public class WatermarkImageHandlerTest
{

    private static final String TOPIC = "foo";

    private Broker broker;
    private ZeebeClient client;

    private ObjectMapper objectMapper;

    @Before
    public void setUp()
    {
        broker = new Broker((InputStream) null);
        client = ZeebeClient.create(new Properties());

        client.topics().create(TOPIC, 1).execute();

        objectMapper = new ObjectMapper();
    }


    @After
    public void tearDown()
    {
        broker.close();
    }

    // TODO: fails because of broker bug so that broker isn't able to decode payload length
    @Test
    public void shouldExecuteProcess() throws IOException
    {
        // given
        client.workflows()
            .deploy(TOPIC)
            .addResourceFromClasspath("process.bpmn")
            .execute();

        ClassLoader classLoader = WatermarkImageHandlerTest.class.getClassLoader();
        URL imageUrl = classLoader.getResource("zeebe-logo.png");

        Map<String, Object> payload = new HashMap<>();
        payload.put("imagePath", imageUrl.getPath());

        String serializedPayload = objectMapper.writeValueAsString(payload);

        client.workflows()
            .create(TOPIC)
            .bpmnProcessId("process")
            .latestVersion()
            .payload(serializedPayload)
            .execute();

        // when
        client.tasks().newTaskSubscription(TOPIC)
            .handler(new WatermarkImageHandler())
            .taskType("watermark")
            .lockOwner("foo")
            .lockTime(Duration.ofSeconds(20))
            .open();

        // then
        List<WorkflowInstanceEvent> events = new CopyOnWriteArrayList<>();

        client.topics().newSubscription(TOPIC)
            .workflowInstanceEventHandler(events::add)
            .name("test")
            .startAtHeadOfTopic()
            .open();

        WorkflowInstanceEvent completeEvent =
            doRepeatedly(() -> events.stream().filter(e -> "WORKFLOW_INSTANCE_COMPLETED".equals(e.getState())).findFirst())
                .until(o -> o.isPresent())
                .get();

        Map<String, Object> resultPayload = objectMapper.readValue(completeEvent.getPayload(), HashMap.class);

        Object watermarkPath = resultPayload.get("watermarkPath");
        Assert.assertNotNull(watermarkPath);
        System.out.println(watermarkPath);


    }


}
