.PHONY: run dev build test verify-content lint
run:
	go run ./cmd/gopath
dev:
	cd web && npm run dev
build:
	cd web && npm run build
	go build -o gopath ./cmd/gopath
test:
	go test ./cmd/... ./internal/...
verify-content:
	go run ./cmd/verify
lint:
	go vet ./cmd/... ./internal/...
	@if command -v staticcheck >/dev/null; then staticcheck ./cmd/... ./internal/...; fi
