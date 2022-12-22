binary = gmj
version = v0.1.0

release:
	GOOS=windows GOARCH=amd64 go build -o ./bin/$(binary)_$(version)_windows_amd64
	GOOS=linux GOARCH=amd64 go build -o ./bin/$(binary)_$(version)_linux_amd64
	GOOS=darwin GOARCH=amd64 go build -o ./bin/$(binary)_$(version)_darwin_amd64
