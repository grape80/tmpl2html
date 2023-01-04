MAKEFLAGS += --no-builtin-rules
.SUFFIXES:

app := tmpl2html
versionFile := embed/version.txt
mainDir := cmd/$(app)
bin := bin/$(app)
logDir := log

gofiles := $(shell find . -type f -name '*.go' -print)
embed := $(shell find . -type f -path '*/embed/*' -print)
now = $(shell date '+%Y%m%d-%H%M%S')

gosetver:
	echo $(now) > $(versionFile)

gobuild: gosetver $(bin)

$(bin): $(embed) $(gofiles) 
	go build -ldflags '-s -w' -o $@ $(mainDir)/main.go

.PHONY: gotest
gotest:
	mkdir -p $(logDir)
	go test -v -cover -count=1 -coverprofile=$(logDir)/gocover-$(now).out > $(logDir)/gotest-$(now).log
	cat $(logDir)/gotest-$(now).log
	go tool cover -html=$(logDir)/gocover-$(now).out -o $(logDir)/gocover-$(now).html
	# open $(logDir)/gocover-$(now).html

goall: gotest gobuild
