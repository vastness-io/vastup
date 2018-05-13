compile:
	@echo "Removing previously built binaries"
	@rm -rf _output/bin || true
	@mkdir -p _output/bin
	@go build -o _output/bin/vastup -v cmd/vastup/main.go
