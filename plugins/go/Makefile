linux-release: clean
	# CGO_ENABLED=1 CGO_LDFLAGS="-static" go build -v -o ../goplugin.so
	CGO_ENABLED=1 GOARCH=amd64 GOOS=linux go build -v -o ../goplugin.so
.PHONY: clean
clean: 
	rm -rf ./tmp
	rm -rf .tmp
	rm -rf dist