package board

import (
	"github.com/gin-gonic/gin"
	"github.com/olahol/go-imageupload"
	"fmt"
	"sync"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/zeebe-io/zbc-go/zbc"
	"io/ioutil"
)

type Board struct {
	//*sync.WaitGroup
	//BucketName string
	//s3svc *s3.S3
	client *zbc.Client
}

func (b *Board) createWorkflowInstance() {
	instance := zbc.NewWorkflowInstance()
	msg, err := b.client.CreateWorkflowInstance("default-topic", instance)
}

//
//func (b *Board) uploadFileToS3(filePath, name string) {
//	defer b.Done()
//	fmt.Printf("[+] Spawning S3UPLOAD routine: %s\n", filePath)
//
//	file, err := os.Open(filePath)
//	defer file.Close()
//
//	stat, _ := file.Stat()
//	if stat.Size() == 0 { return } // file is empty. skip upload.
//
//	params := &s3.PutObjectInput{ Bucket: aws.String(b.BucketName), Key: aws.String(name), Body: file, }
//	_, err = b.s3svc.PutObject(params)
//
//	if err != nil {
//		panic(err)
//	}
//
//	return
//}

func NewBoard() *Board {
	zb, _ := zbc.NewClient("0.0.0.0:51015")
	return &Board{
		//&sync.WaitGroup{},
		//"facebam",
		//s3.New(session.Must(session.NewSession())),
		zb,
	}
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

		thumb, err := imageupload.ThumbnailPNG(img, 300, 300)

		if err != nil {
			panic(err)
		}

		imgPath := fmt.Sprintf("/tmp/%s", img.Filename)
		ioutil.WriteFile(imgPath, img.Data, 0644)
		board.Add(1)
		//go board.uploadFileToS3(imgPath, img.Filename)
		// TODO: imgPath to msgpack -> payload
		instance := zbc.NewWorkflowInstance()

		thumb.Write(c.Writer)
	})

	r.Run(":5000")
}