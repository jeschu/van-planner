.PHONY: build run test clean all

all: clean run

build:
	go build -o van-planner ./cmd/van-planner

run: build
	./van-planner

test:
	go test ./...

clean:
	rm -f van-planner