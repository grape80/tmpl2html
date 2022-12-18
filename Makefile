MAKEFLAGS += --no-builtin-rules
.SUFFIXES:

app := tmpl2html
mainDir := cmd/$(app)
bin := bin/$(app)

gofiles := $(shell find . -type f -name '*.go' -print)

gobuild: $(bin)

$(bin): $(gofiles)
	go build -o $@ $(mainDir)/main.go

logDir := log
now = $(shell date '+%Y%m%d-%H%M%S')
.PHONY: gotest
gotest:
	mkdir -p $(logDir)
	go test -v -cover -count=1 -coverprofile=$(logDir)/gocover-$(now).out > $(logDir)/gotest-$(now).log
	cat $(logDir)/gotest-$(now).log
	go tool cover -html=$(logDir)/gocover-$(now).out -o $(logDir)/gocover-$(now).html
	open $(logDir)/gocover-$(now).html

goall: gobuild gotest
