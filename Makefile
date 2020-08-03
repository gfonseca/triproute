BINPATH=./bin
SERVER_BIN=$(BINPATH)/server
SERVER_SRC=cmd/server.go
CMD_BIN=$(BINPATH)/command
CMD_SRC=cmd/command.go
GOCMD=go
GOTEST=$(GOCMD) test -covermode=count -coverprofile=coverage.out ./pkg/...
GOCOVER=$(GOCMD) tool cover -html=coverage.out
GOBUILD_CMD=$(GOCMD) build -o ./$(CMD_BIN) -v ./$(CMD_SRC)
GOBUILD_SERVER=$(GOCMD) build -o ./$(SERVER_BIN) -v ./$(SERVER_SRC)

test:
	$(GOTEST)
cover: test
	$(GOCOVER)
build-cmd: test
	$(GOBUILD_CMD)
run-cmd: build-cmd
	$(CMD_BIN)
build-server: test
	$(GOBUILD_SERVER)
run-server: build-server
	$(SERVER_BIN)
clear:
	rm -fv $(BINPATH)/*


#docker 
docker-build:
	docker-compose build
docker-up:
	docker-compose up