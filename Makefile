DEPS=go list -f '{{range .TestImports}}{{.}} {{end}}' ./...

fmt:
	bash -c 'go list ./... | grep -v vendor | xargs -n1 go fmt'

test:
	bash -c 'go test -timeout=10s'
