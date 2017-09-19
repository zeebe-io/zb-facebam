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

type WrappedResponseWriter struct {
	gin.ResponseWriter
	writer http.ResponseWriter
}

func (w *WrappedResponseWriter) Write(data []byte) (int, error) {
	return w.writer.Write(data)
}

func (w *WrappedResponseWriter) WriteString(s string) (n int, err error) {
	return w.writer.Write([]byte(s))
}


type NextRequestHandler struct{
	c *gin.Context
}

// Run the next request in the middleware chain and return
// See: https://godoc.org/github.com/gin-gonic/gin#Context.Next
func (h *NextRequestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.c.Writer = &WrappedResponseWriter{h.c.Writer, w}
	h.c.Next()
}

// Wrap something that accepts an http.Handler, returns an http.Handler
func WrapHH(hh func(h http.Handler) http.Handler) gin.HandlerFunc {
	// Steps:
	// - create an http handler to pass `hh`
	// - call `hh` with the http handler, which returns a function
	// - call the ServeHTTP method of the resulting function to run the rest of the middleware chain

	return func(c *gin.Context) {
		hh(&NextRequestHandler{c}).ServeHTTP(c.Writer, c.Request)
	}
}

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