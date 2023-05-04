build:
	go build -o out/validate

install:
	sudo mv out/validate /usr/local/bin/