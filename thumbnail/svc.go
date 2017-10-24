package thumbnail

import (
	"github.com/vmihailenco/msgpack"
	"github.com/zeebe-io/zbc-go/zbc"
	"log"
	"github.com/olahol/go-imageupload"
	"fmt"
	"io/ioutil"
	"strings"
)

type Payload struct {
	Image string `msgpack:"imagePath"`
	Watermark string `msgpack:"watermarkPath"`
}

func processTask(msg *zbc.SubscriptionEvent) {
	var payload Payload

	err := msgpack.Unmarshal(msg.Task.Payload, &payload)
	if err != nil {
		panic(err)
	}

	bytes, err := ioutil.ReadFile(payload.Watermark)
	if err != nil {
		panic(err)
	}

	image := &imageupload.Image{
		Filename:    payload.Watermark,
		ContentType: "image/png",
		Data:        bytes,
		Size:        len(bytes),
	}

	thumb, err := imageupload.ThumbnailPNG(image, 300, 300)
	parts := strings.Split(image.Filename, "/")
	filename := parts[len(parts) - 1]
	parts = strings.Split(filename, ".")
	thumbPath := fmt.Sprintf("/tmp/watermarking/%s-thumb.%s", parts[0], parts[len(parts) - 1])

	ioutil.WriteFile(thumbPath, thumb.Data, 0644)

	log.Printf("Saved thumbnail to %s\n", thumbPath)
}

func Run() {
	client, _ := zbc.NewClient("127.0.0.1:51015")

	subscriptionCh, subInfo, err := client.TaskConsumer("default-topic", "thumbnailer", "thumbnail")

	if err != nil {
		panic(err)
	}


	credits := subInfo.Credits

	log.Println("Subscription opened with", credits, "Credits")
	log.Println("Waiting for events ....")
	for {
		select {
		case message := <-subscriptionCh:
			credits--;

			processTask(message)
			response, err := client.CompleteTask(message)

			if err != nil {
				log.Println("Completing a task went wrong.")
				log.Println(err)
			}

			if response.State == zbc.TaskCompleted {
				log.Println("Task completed successfully.")
			} else {
				log.Println("Task not completed.")
			}

			if credits < 1 {
				response, err := client.IncreaseTaskSubscriptionCredits(subInfo);

				if err != nil {
					log.Println("Increasing task credits went wrong.")
					log.Println(err)

				} else {
					credits = response.Credits
					log.Println("Increased task credits to", credits)
				}
			}

			break
		}
	}


}
