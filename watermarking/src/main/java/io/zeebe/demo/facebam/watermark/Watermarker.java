package io.zeebe.demo.facebam.watermark;

import java.time.Duration;
import java.util.Properties;
import java.util.Scanner;

import io.zeebe.client.ClientProperties;
import io.zeebe.client.ZeebeClient;
import io.zeebe.client.task.TaskSubscription;

public class Watermarker
{

    public static final int LOCK_TIME_SECS = 20;

    public static void main(String[] args)
    {
        if (args.length < 2 || args.length > 3)
        {
            StringBuilder sb = new StringBuilder();
            sb.append("Reads a file from file system, watermarks it and writes the result back to file system next to the original image.\n\n");
            sb.append("Usage: java -jar watermarker.jar <topic> <lock_owner> (<broker-contact-point=localhost:51015>)");
            System.out.println(sb.toString());
        }

        String topic = args[0];
        String lockOwner = args[1];

        Properties properties = new Properties();
        if (args.length == 3)
        {
            properties.put(ClientProperties.BROKER_CONTACTPOINT, args[2]);
        }

        ZeebeClient client = ZeebeClient.create(properties);
        client.connect();

        TaskSubscription subscription = client.tasks().newTaskSubscription(topic)
            .handler(new WatermarkImageHandler())
            .taskType("watermark")
            .lockOwner(lockOwner)
            .lockTime(Duration.ofSeconds(LOCK_TIME_SECS))
            .open();

        try (Scanner scanner = new Scanner(System.in))
        {
            while (scanner.hasNextLine())
            {
                final String nextLine = scanner.nextLine();
                if (nextLine.contains("exit")
                    || nextLine.contains("close")
                    || nextLine.contains("quit")
                    || nextLine.contains("halt")
                    || nextLine.contains("shutdown")
                    || nextLine.contains("stop"))
                {
                    subscription.close();
                    System.exit(0);
                }
            }
        }

    }

}
