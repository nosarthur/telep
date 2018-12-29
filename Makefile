.PHONY: proto install clean

PROTO_DIR=messages

all: proto tserver
	go build
proto:
	protoc -I $(PROTO_DIR) --go_out=plugins=grpc:$(PROTO_DIR) $(PROTO_DIR)/*proto
tserver: server/main.go
	go build -o tserver server/main.go
install:
	go install
clean:
	rm tserver telep
