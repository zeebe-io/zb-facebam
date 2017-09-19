board:
	go run main.go service board

thumbnail:
	go run main.go service thumbnail

watermark:
	java -jar watermarking/target/watermarker.jar default-topic watermarker

.PHONY: board thumbnail watermark
