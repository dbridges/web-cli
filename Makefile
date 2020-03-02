.PHONY: all run run-with-race clean install uninstall

VERSION=0.0.1
GO_SRC = $(shell find . -iname '*.go' ! -iname "*test.go")
BINDIR?=/usr/local/bin
BINNAME?=web

all: $(BINNAME)

$(BINNAME): $(GO_SRC) dist
	go build -ldflags "-X main.Version=$(VERSION)" -o dist/$@

run:
	@go run $(GO_SRC)

run-with-race:
	@GORACE="log_path=race_log" go run -race *.go

clean:
	rm -f dist/*
	rm -f race_log.*

install: all
	mkdir -p $(BINDIR)
	install dist/$(BINNAME) $(BINDIR)/$(BINNAME)

dist:
	mkdir dist

uninstall:
	rm -f $(BINDIR)/$(BINNAME)
