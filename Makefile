build:
	go build -o dist/validate

install_local:
	sudo mv dist/validate /usr/local/bin/