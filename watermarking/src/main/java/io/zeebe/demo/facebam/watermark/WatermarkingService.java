package io.zeebe.demo.facebam.watermark;

import java.awt.Graphics2D;
import java.awt.image.BufferedImage;
import java.io.IOException;
import java.io.InputStream;
import java.util.Iterator;

import javax.imageio.ImageIO;
import javax.imageio.ImageReader;
import javax.imageio.stream.ImageInputStream;

import org.imgscalr.Scalr;
import org.imgscalr.Scalr.Method;

public class WatermarkingService
{

    public WatermarkedImage markImageFromStream(InputStream imageStream, InputStream watermarkStream)
    {
        ImageInputStream typedImageStream = readImageInputStream(imageStream);

        ImageReader reader = determineImageReader(typedImageStream);
        reader.setInput(typedImageStream);

        BufferedImage image = readImage(reader);
        BufferedImage watermark = readImage(watermarkStream);

        markImage(image, watermark);

        String formatName;
        try
        {
            formatName = reader.getFormatName();
        } catch (IOException e)
        {
            throw new RuntimeException("Could not determine image format", e);
        }

        return new WatermarkedImage(formatName, image);
    }

    protected ImageInputStream readImageInputStream(InputStream stream)
    {
        try
        {
            return ImageIO.createImageInputStream(stream);
        } catch (IOException e)
        {
            throw new RuntimeException(e);
        }
    }

    protected ImageReader determineImageReader(ImageInputStream stream)
    {
        Iterator<ImageReader> imageReaders = ImageIO.getImageReaders(stream);

        if (!imageReaders.hasNext())
        {
            throw new RuntimeException("Stream does not contain image content");
        }

        return imageReaders.next();
    }

    protected BufferedImage readImage(InputStream stream)
    {
        try
        {
            return ImageIO.read(stream);
        } catch (IOException e)
        {
            throw new RuntimeException("Could not read stream as image", e);
        }
    }

    protected BufferedImage readImage(ImageReader reader)
    {
        try
        {
            return reader.read(0);
        } catch (IOException e)
        {
            throw new RuntimeException("Could not read stream as image", e);
        }
    }

    protected void markImage(BufferedImage image, BufferedImage watermark)
    {
        int watermarkWidth = image.getWidth() / 3;
        double scalingFactor = (double) watermarkWidth / watermark.getWidth();
        int watermarkHeight = (int) (watermark.getHeight() * scalingFactor);

        Graphics2D graphics = (Graphics2D) image.getGraphics();

        BufferedImage resizedWatermark = Scalr.resize(watermark, Method.ULTRA_QUALITY, watermarkWidth, watermarkHeight);
        graphics.drawImage(
                resizedWatermark,
                image.getWidth() - watermarkWidth,
                image.getHeight() - watermarkHeight,
                null);

    }
}
