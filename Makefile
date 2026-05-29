GOPATH?=$(shell go env GOPATH)
build:
	go build -ldflags="-s -w" lgctl.go
	mv lgctl $(GOPATH)/bin/lgctl
	lgctl template clean
	lgctl template init
	$(if $(shell command -v upx), upx goctl)

mac:
	@echo ${GOPATH}
	GOOS=darwin go build -ldflags="-s -w" -o lgctl lgctl.go
	$(if $(shell command -v upx), upx goctl-darwin)
	mv lgctl $(GOPATH)/bin/lgctl

win:
	GOOS=windows go build -ldflags="-s -w" -o lgctl.exe lgctl.go
	$(if $(shell command -v upx), upx goctl.exe)
	mv lgctl.exe $(GOPATH)/bin/lgctl.exe

linux:
	GOOS=linux go build -ldflags="-s -w" -o lgctl lgctl.go
	$(if $(shell command -v upx), upx goctl-linux)
	mv lgctl $(GOPATH)/bin/lgctl

image:
	docker build --rm --platform linux/amd64 -t kevinwan/goctl:$(version) .
	docker tag kevinwan/goctl:$(version) kevinwan/goctl:latest
	docker push kevinwan/goctl:$(version)
	docker push kevinwan/goctl:latest
	docker build --rm --platform linux/arm64 -t kevinwan/goctl:$(version)-arm64 .
	docker tag kevinwan/goctl:$(version)-arm64 kevinwan/goctl:latest-arm64
	docker push kevinwan/goctl:$(version)-arm64
	docker push kevinwan/goctl:latest-arm64
