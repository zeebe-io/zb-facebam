package io.zeebe.demo.facebam.watermark;

import java.io.File;
import java.io.FileInputStream;
import java.io.IOException;
import java.io.InputStream;
import java.util.HashMap;
import java.util.Map;

import org.apache.commons.io.FilenameUtils;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;

import io.zeebe.client.TasksClient;
import io.zeebe.client.event.TaskEvent;
import io.zeebe.client.task.TaskHandler;

public class WatermarkImageHandler implements TaskHandler
{

    protected static final String INPUT_PATH_VAR = "imagePath";
    protected static final String OUTPUT_PATH_VAR = "watermarkPath";
    protected static final String WATERMARK_FILE = "zeebe-logo.png";

    protected ObjectMapper objectMapper = new ObjectMapper();
    protected WatermarkingService watermarkingService = new WatermarkingService();

    @Override
    public void handle(TasksClient client, TaskEvent task)
    {
        Map<String, Object> inputPayload = deserializePayload(task.getPayload());

        Object imageLocationValue = inputPayload.get(INPUT_PATH_VAR);
        if (imageLocationValue == null || !(imageLocationValue instanceof String))
        {
            throw new RuntimeException("Payload does not contain String variable '" + INPUT_PATH_VAR + "'");
        }

        String imageLocation = (String) imageLocationValue;
        File image = new File(imageLocation);

        InputStream watermarkStream = WatermarkImageHandler.class.getClassLoader().getResourceAsStream(WATERMARK_FILE);

        String waterMarkPath;
        try (FileInputStream imageStream = new FileInputStream(image))
        {
            WatermarkedImage watermarkedImage = watermarkingService.markImageFromStream(imageStream, watermarkStream);
            waterMarkPath = FilenameUtils.removeExtension(imageLocation) + "-watermarked." + watermarkedImage.getFormat();
            watermarkedImage.writeToFile(waterMarkPath);
        }
        catch (IOException e)
        {
            throw new RuntimeException(e);
        }

        inputPayload.put(OUTPUT_PATH_VAR, waterMarkPath);

        String outputPayload = serializePayload(inputPayload);

        client.complete(task)
            .payload(outputPayload)
            .execute();
    }

    private String serializePayload(Map<String, Object> deserializedPayload)
    {
        String outputPayload;
        try
        {
            outputPayload = objectMapper.writeValueAsString(deserializedPayload);
        } catch (JsonProcessingException e)
        {
            throw new RuntimeException("Could not generate output payload", e);
        }
        return outputPayload;
    }

    private Map<String, Object> deserializePayload(String serializedPayload)
    {
        Map<String, Object> payload;
        try
        {
            payload = (Map<String, Object>) objectMapper.readValue(serializedPayload, HashMap.class);
        } catch (IOException e)
        {
            throw new RuntimeException("Could not decode JSON", e);
        }
        return payload;
    }


}
