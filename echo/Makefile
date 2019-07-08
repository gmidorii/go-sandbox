BIN=echo-server

.PHONY: install
install:
	# swag
	# realize

.PHONY: doc
doc: 
	swag i

.PHONY: build
build:
	go build -o $(BIN)

.PHONY: reload-local
reload-local:
	realize start