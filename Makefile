.PHONY: build test

test:
	go test ./... -v

test+dev:
	$(MAKE) -f Makefile.tests test