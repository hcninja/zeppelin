# Config
BINARY=zeppelin
VERSION=v0.2.0
TARGET=all
BUILD_TIME=`date +%FT%T%z`
LDFLAGS=-ldflags="\
	-s \
	-w \
	-X main.version=${VERSION} \
	-X main.buildTime=${BUILD_TIME}"


.DEFAULT_GOAL: $(BINARY)

.PHONY: all
$(TARGET):
	GOOS=darwin GOARCH=amd64 go build ${LDFLAGS} -o ${BINARY}_darwin
	GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o ${BINARY}_linux
	GOOS=windows GOARCH=amd64 go build ${LDFLAGS} -o ${BINARY}.exe
	tar czvf ${BINARY}_darwin.tgz ${BINARY}_darwin
	tar czvf ${BINARY}_linux.tgz ${BINARY}_linux
	zip -9 ${BINARY}_windows.zip ${BINARY}.exe

.PHONY: packed
packed:
	upx --ultra-brute -9 -o ${BINARY}_darwin_upx ${BINARY}_darwin 
	upx --ultra-brute -9 -o ${BINARY}_linux_upx ${BINARY}_linux
	upx --ultra-brute -9 -o ${BINARY}_upx.exe ${BINARY}.exe 
	tar czvf ${BINARY}_darwin_upx.tgz ${BINARY}_darwin_upx
	tar czvf ${BINARY}_linux_upx.tgz ${BINARY}_linux_upx
	zip -9 ${BINARY}_windows_upx.zip ${BINARY}_upx.exe

.PHONY: macos
macos:
	GOOS=darwin GOARCH=amd64 go build ${LDFLAGS} -o ${BINARY}_darwin

.PHONY: clean
clean:
	rm -rf ${BINARY}.exe ${BINARY}_darwin ${BINARY}_linux 
	rm -rf ${BINARY}_upx.exe ${BINARY}_darwin_upx ${BINARY}_linux_upx

	rm -rf ${BINARY}_darwin.tgz ${BINARY}_linux.tgz ${BINARY}_windows.zip
	rm -rf ${BINARY}_darwin_upx.tgz ${BINARY}_linux_upx.tgz ${BINARY}_windows_upx.zip