build:
	wails build -tags webkit2_41
.PHONY: build

run: build
	./build/bin/budgenius
.PHONY: run