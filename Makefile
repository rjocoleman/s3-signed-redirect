.PHONY: go all build clean vendor package-darwin package-linux package-windows release

all: build

build:
	go build

darwin: *.go
	GOOS=darwin GOARCH=amd64 go build

linux: *.go
	GOOS=linux GOARCH=amd64 go build

windows: *.go
	GOOS=windows GOARCH=386 go build

clean:
	rm -f s3-signed-redirect

package-darwin: darwin
	tar zcf s3-signed-redirect-darwin-amd64.tar.gz s3-signed-redirect

package-linux: linux
	tar zcf s3-signed-redirect-linux-amd64.tar.gz s3-signed-redirect

package-windows: windows
	zip s3-signed-redirect-win32-i386.zip -xi s3-signed-redirect.exe

release: package-linux package-darwin package-windows