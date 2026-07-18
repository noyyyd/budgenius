PROJECT_DIR := $(abspath $(dir $(lastword $(MAKEFILE_LIST))))
DATA_DIR := $(PROJECT_DIR)/.data

build:
	wails build -tags webkit2_41 
.PHONY: build

run: clean
	mkdir $(DATA_DIR)
	wails dev -tags webkit2_41 -ldflags "-X main.dataDir=$(DATA_DIR)"
.PHONY: run 

clean:
	rm -r .data
.PHONY: clean 
