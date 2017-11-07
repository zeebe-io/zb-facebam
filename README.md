# zb-facebam

This project implements example application which consist of 3 services: board, thumbnail and watermarking.

The process which is implemented is the following:

![process](process.png)


### Requirements

* Java 1.8+
* Golang 1.5+
* Running Zeebe broker


### Setup

First create topic which will be used by our process.

```
zbctl create topic --name default-topic --partitions 1
```

then create workflow on the broker.

```
zbctl create workflow process.bpmn
```



To start microservices use ```make``` command. Open 3 terminals and in each of them start 1 service with the following commands:

```
make board
```

After starting this service, point your browser to ```localhost:5000``` and click on the upload photo button. Select some PNG photo and confirm the upload.



```
make thumbnail
```

This service will create thumbnail of the uploaded image.


```
make watermark
```

This service will watermark uploaded image.



*Note*: Board and thumbnail services are written in Go and watermark is build in Java, so you have to build them accordingly. To build Java service you can use the following command:


```bash
cd watermarking
mvn package -DskipTests
```