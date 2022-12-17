MAKEFLAGS += --no-builtin-rules
.SUFFIXES:

app := tmpl2html
mainDir := cmd/$(app)
bin := bin/$(app)

gofiles := $(shell find . -type f -name '*.go' -print)

gobuild: $(bin)

$(bin): $(gofiles)
	go build -o $@ $(mainDir)/main.go
