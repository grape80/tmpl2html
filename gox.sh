#!/bin/sh

if [ $# != 5 ]; then
    exit 1
fi

LDFLAGS=$1
MAINDIR=$2
APP=$3
VERSION=$4
DIST=$5
DISTLIST=godist.list

for target in `cat $DISTLIST | grep -v '#'`
do
    arr=(${target//\// })
    os=${arr[0]}
    arch=${arr[1]}

    if [ "$os" == 'windows' ]; then
        rename 's/^_//' *.syso && \
        GOOS=$os GOARCH=$arch go build -ldflags "$LDFLAGS" -o $DIST/$os-$arch/$APP.exe ./$MAINDIR  && \
        zip -j $DIST/$APP-$VERSION.$os-$arch.zip $DIST/$os-$arch/$APP.exe && \
        rename 's/^/_/' *.syso
    else
        GOOS=$os GOARCH=$arch go build -ldflags "$LDFLAGS" -o $DIST/$os-$arch/$APP ./$MAINDIR && \
        tar zcvf $DIST/$APP-$VERSION.$os-$arch.tar.gz -C $DIST/$os-$arch $APP
    fi
done
