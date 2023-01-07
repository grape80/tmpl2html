#!/bin/sh

if [[ $# != 5 ]]; then
    exit 1
fi

LDFLAGS=$1
MAINDIR=$2
APP=$3
VERSION=$4
DIST=$5

GOOS=darwin GOARCH=arm64 && go build -ldflags "$LDFLAGS" -o $DIST/$GOOS-$GOARCH/$APP ./$MAINDIR && tar zcvf $DIST/$APP-$VERSION.$GOOS-$GOARCH.tar.gz -C $DIST/$GOOS-$GOARCH $APP
GOOS=windows GOARCH=amd64 && go build -ldflags "$LDFLAGS" -o $DIST/$GOOS-$GOARCH/$APP.exe ./$MAINDIR && zip -j $DIST/$APP-$VERSION.$GOOS-$GOARCH.zip $DIST/$GOOS-$GOARCH/$APP.exe
