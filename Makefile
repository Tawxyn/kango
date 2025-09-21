APP     ?= kango
CMD     ?= .
BIN     ?= bin/$(APP)

.PHONY: build run clean

build:
	@mkdir -p bin
	@echo ">> building $(BIN)"
	@go build -o $(BIN) $(CMD)

run: build
	@echo ">> running $(BIN)"
	@$(BIN) $(ARGS)

clean:
	@echo ">> cleaning"
	@rm -rf bin
