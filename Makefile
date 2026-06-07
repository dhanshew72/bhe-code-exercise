.PHONY: test
test: 
	go test ./... -timeout=30s

.PHONY: test-fuzz
test-fuzz:
	go test ./... -fuzz=. -fuzztime=30s
