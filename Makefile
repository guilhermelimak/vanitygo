PROJPATH=${GOPATH}/src/github.com/toxeus/go-secp256k1

default:
	go run *.go

build:
	go build -o vanitygo *.go

deps:
	go get github.com/akamensky/base58
	go get -u github.com/gizak/termui
	go get github.com/urfave/cli
	go get github.com/fatih/color
	make secp

secp:
	go get -d github.com/toxeus/go-secp256k1
	cd ${PROJPATH} && git submodule update --init # not needed for Go >= 1.6
	cd ${PROJPATH}/c-secp256k1 && ./autogen.sh && ./configure && make
	cd ${PROJPATH} && go install
	go install
