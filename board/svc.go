package board

import (
	"github.com/gin-gonic/gin"
	"github.com/olahol/go-imageupload"
	"fmt"
	"github.com/zeebe-io/zbc-go/zbc"
	"io/ioutil"
	"time"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type Board struct {
	client *zbc.Client
}

func NewBoard() *Board {
	zb, _ := zbc.NewClient("0.0.0.0:51015")
	return &Board{
		zb,
	}
}
type Payload struct {
	Image string `msgpack:"imagePath"`
}

func Run() {
 	 _ = os.Mkdir("/tmp/watermarking", os.ModePerm)

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
		
		payload := make(map[string]interface{})
		payload["imagePath"] = imgPath

		if err != nil {
			panic(err)
		}

		instance := zbc.NewWorkflowInstance("watermark", -1, payload)
		_, err = board.client.CreateWorkflowInstance("default-topic", instance)
		if err != nil {
			panic(err)
		}

		fmt.Println("Upload file to", imgPath)

		c.Redirect(302, "/images")
	})


	r.GET("/images/", gin.WrapH(http.StripPrefix("/images/", http.FileServer(http.Dir("/tmp/watermarking")))))

	r.GET("/images/:filename", func(ctx *gin.Context) {
		fileName := ctx.Param("filename")
		targetPath := filepath.Join("/tmp/watermarking", fileName)
		if !strings.HasPrefix(filepath.Clean(targetPath), "/tmp/watermarking") {
			ctx.String(403, ":(")
			return
		}
		
		ctx.File(targetPath)
	})

	r.Run(":5000")
}