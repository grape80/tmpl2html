MAKEFLAGS += --no-builtin-rules
.SUFFIXES:
.DEFAULT_GOAL := help

app := tmpl2html
versionFile := embed/version.txt
mainDir := cmd/$(app)
binDir := bin
logDir := log
distDir := dist
gobuild_ldflags := '-s -w'

gofiles := $(shell find . -type f -name '*.go' -print)
embed := $(shell find . -type f -path '*/embed/*' -print)
version = $(shell cat $(versionFile))
now = $(shell date '+%Y%m%d-%H%M%S')

.PHONY: help ##       Print help.
help:
	@cat $(MAKEFILE_LIST) | grep '##' | grep -v 'MAKEFILE_LIST' | sed s/^.PHONY:/\ \ / | sed s/##//

.PHONY: appsetver ##  Set app version. This is for local.
appsetver:
	echo $(now) > $(versionFile)

.PHONY: build ##      Build for local.
build: appsetver $(binDir)/$(app)

$(binDir)/$(app): $(embed) $(gofiles)
	go build -ldflags $(gobuild_ldflags) -o $@ ./$(mainDir)

.PHONY: xbuild ##     Build for each platforms. See godist.list.
xbuild: cleandist resc
	./gox.sh $(gobuild_ldflags) $(mainDir) $(app) $(version) $(distDir)

.PHONY: resc ##       Compile rc files for windows.
resc: cleansyso
	x86_64-w64-mingw32-windres res/win/versioninfo.rc _versioninfo.syso

.PHONY: test ##       Run tests.
test:
	mkdir -p $(logDir)
	go test -v -cover -count=1 -coverprofile=$(logDir)/gocover-$(now).out > $(logDir)/gotest-$(now).log
	cat $(logDir)/gotest-$(now).log
	go tool cover -html=$(logDir)/gocover-$(now).out -o $(logDir)/gocover-$(now).html
	# open $(logDir)/gocover-$(now).html

.PHONY: all ##        Run build, test, xbuild.
all: build test xbuild

## Clean
.PHONY: cleansyso ##  remove syso files.
cleansyso:
	rm -vf *.syso

.PHONY: cleanbin ##   remove bin directory.
cleanbin:
	rm -rvf $(binDir)

.PHONY: cleandist ##  remove dist directory.
cleandist:
	rm -rvf $(distDir)

.PHONY: cleanlog ##   remove log directory.
cleanlog:
	rm -rvf $(logDir)

.PHONY: cleanall ##   run all clean targets. 
cleanall: cleansyso cleanbin cleandist cleanlog
