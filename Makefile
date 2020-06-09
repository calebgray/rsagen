GOROOT=`go env GOROOT`

default:
	mkdir -p build
	GOOS=js GOARCH=wasm go build -o build/rsagen.wasm -ldflags="-s -w" .
	cp "${GOROOT}/misc/wasm/wasm_exec.js" build
	cp "${GOROOT}/misc/wasm/wasm_exec.html" build
