BIN_NAME=bus-system

dep:
	@echo "> fetching dependencies..."
	@go mod tidy

install: dep
	@echo "> building binary..."
	@CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static" ' -o ./${BIN_NAME} main.go

run: dep install 
	@echo "> running applications..."
	@./${BIN_NAME} start

clean:
	@echo "> cleaning applications..."
	@rm ${BIN_NAME}

eal.PHONY: dep install run clean