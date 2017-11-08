board:
	go run main.go service board

thumbnail:
	go run main.go service thumbnail

watermark:
	mvn -f watermarking/pom.xml package -DskipTests
	java -jar watermarking/target/watermarker.jar default-topic watermarker

.PHONY: board thumbnail watermark
