VERSION:=$(shell sed 's/\./_/g' <<<$$(cut -d' ' -f3 <<<$$(go run main.go --version)))

dep:
	dep ensure

generate:
	protoc -I service/ service/service.proto --go_out=plugins=grpc:service

build_mac:
	GOOS=darwin go build -o bitnode_mac_$(VERSION) main.go

build_linux_x86_64:
	GOOS=linux go build -o bitnode_linux_$(VERSION) main.go

build_pi:
	GOARM=6 GOOS=linux GOARCH=arm go build -o bitnode_pi_$(VERSION) main.go

build: build_mac build_linux_x86_64 build_pi

all: dep generate
