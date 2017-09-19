package board

import (
	"github.com/gin-gonic/gin"
	"github.com/olahol/go-imageupload"
	"github.com/vmihailenco/msgpack"
	"fmt"
	"github.com/zeebe-io/zbc-go/zbc"
	"io/ioutil"
	"time"
)

type Board struct {
	client *zbc.Client
}


func NewBoard() *Board {
	zb, _ := zbc.NewClient("127.0.0.1:51015")
	return &Board{
		zb,
	}
}
type Payload struct {
	Image string `msgpack:"imagePath"`
}

func Run() {
	r := gin.Default()
	board := NewBoard()

	r.GET("/", func(c *gin.Context) {
		c.File("board/templates/upload.html")
	})

	r.POST("/upload", func(c *gin.Context) {
		img, err := imageupload.Process(c.Request, "file")

		if err != nil {
			panic(err)
		}



		t := time.Now()
		imgPath := fmt.Sprintf("/tmp/watermarking/%s-%s", t.Format("20060102150405"), img.Filename)
		ioutil.WriteFile(imgPath, img.Data, 0644)
		payload, err := msgpack.Marshal(&Payload{Image: imgPath})

		if err != nil {
			panic(err)
		}

		instance := zbc.NewWorkflowInstance("watermark", -1, payload)
		_, err = board.client.CreateWorkflowInstance("default-topic", instance);
		if err != nil {
			panic(err)
		}

		fmt.Println("Upload file to", imgPath)
	})

	r.Run(":5000")
}