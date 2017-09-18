package io.zeebe.demo.zalando.image;

import java.io.IOException;
import java.io.InputStream;

import org.junit.Test;

import io.zeebe.demo.facebam.watermark.WatermarkedImage;
import io.zeebe.demo.facebam.watermark.WatermarkingService;

public class WatermarkingServiceTest
{

    @Test
    public void shouldWatermarkImage() throws IOException
    {
        // given
        WatermarkingService service = new WatermarkingService();

        InputStream imageStream = WatermarkingServiceTest.class.getClassLoader().getResourceAsStream("zeebe-logo.png");
        InputStream watermarkStream = WatermarkingServiceTest.class.getClassLoader().getResourceAsStream("zeebe-logo.png");

        // when
        WatermarkedImage watermarkedImage = service.markImageFromStream(imageStream, watermarkStream);

        watermarkedImage.writeToFile("out." + watermarkedImage.getFormat());

    }
}
