package board

import (
	"github.com/gin-gonic/gin"
	"github.com/olahol/go-imageupload"
	"github.com/vmihailenco/msgpack"
	"fmt"
	"github.com/zeebe-io/zbc-go/zbc"
	"io/ioutil"
	"github.com/davecgh/go-spew/spew"
	"time"
	"net/http"
	//todo: remove
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

func Run() {
	r := gin.Default()
	board := NewBoard()

	r.GET("/", gin.WrapH(http.FileServer(http.Dir("/tmp/watermarking"))))

	r.GET("/upload", func(c *gin.Context) {
		c.File("board/templates/upload.html")
	})
	
	r.POST("/upload", func(c *gin.Context) {
		img, err := imageupload.Process(c.Request, "file")

		if err != nil {
			panic(err)
		}


		type Foo struct {
			Bar string `msgpack:"imagePath"`
		}

		t := time.Now()
		imgPath := fmt.Sprintf("/tmp/watermarking/%s-%s", t.Format("20060102150405"), img.Filename)
		ioutil.WriteFile(imgPath, img.Data, 0644)
		fmt.Println(imgPath)
		i := Foo{Bar: imgPath}
		spew.Dump(i)
		payload, err := msgpack.Marshal(&i)
		spew.Dump(payload)

		if err != nil {
			panic(err)
		}

		instance := zbc.NewWorkflowInstance("watermark", -1, payload)
		fmt.Println(instance)
		resp, err := board.client.CreateWorkflowInstance("default-topic", instance);
		if err != nil {
			panic(err)
		}
		fmt.Println(resp)
	})

	r.Run(":5000")
}