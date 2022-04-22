APP = $(shell basename ${CURDIR})

TAG = $(shell git log --pretty=format:"%cd.%h" --date=short -1)

.PHONY: all build build-pro run bra

build:
	goctl api go --api ./api/datac.api  --dir=./ --home=./data/template -style=goZero

build-pro:
	go build -o ./bin/linux/${APP} ./cmd/web

run:
	./bin/win/${APP}

wire:
	go run github.com/google/wire/cmd/wire

# bra 热更新用于调试程序
bra:
	bra run