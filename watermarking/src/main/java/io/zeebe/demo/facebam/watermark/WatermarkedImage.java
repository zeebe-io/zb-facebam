package io.zeebe.demo.facebam.watermark;

import java.awt.image.BufferedImage;
import java.io.File;
import java.io.IOException;

import javax.imageio.ImageIO;

public class WatermarkedImage
{

    protected String format;
    protected BufferedImage image;

    public WatermarkedImage(String format, BufferedImage image)
    {
        this.format = format;
        this.image = image;
    }

    public String getFormat()
    {
        return format;
    }

    public void writeToFile(String path) throws IOException
    {
        ImageIO.write(image, format, new File(path));
    }
}
