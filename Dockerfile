FROM golang:alpine

COPY . /go/src/github.com/zeebe-io/zb-facebam
WORKDIR /go/src/github.com/zeebe-io/zb-facebam

RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o miniflow .

ENTRYPOINT ["./miniflow"]