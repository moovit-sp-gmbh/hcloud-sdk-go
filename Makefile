.PHONY: build test

test:
	go test ./...

test+dev:
	$(MAKE) -f Makefile.tests test