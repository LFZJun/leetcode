#!/bin/sh

name=leetcode-cli
version=$1
configPath=./cmd/config
input=./cmd
output=./out/
go=go

if [ "$1" = "" ];then
    version=v0.3.0
fi


Build() {
    goarm=$4
    if [ "$4" = "" ];then
        goarm=7
    fi

    echo "Building $1..."
    export GOOS=$2 GOARCH=$3 GO386=sse2 CGO_ENABLED=0 GOARM=$4
    if [ $2 = "windows" ];then
        $go build -ldflags "-X main.Version=$version -s -w" -o "$output/$1/$name.exe" $input
        PackConfig $1
    else
        $go build -ldflags "-X main.Version=$version -s -w" -o "$output/$1/$name" $input
        PackConfig $1
    fi

    Pack $1
}

# config
PackConfig() {
    cp -r $configPath $output/$1/
}


# zip 打包
Pack() {
    cp README.md "$output/$1"

    cd $output
    zip -q -r "$1.zip" "$1"

    # 删除
    rm -rf "$1"

    cd ..
}

# OS X / macOS
Build $name-$version"-darwin-osx-amd64" darwin amd64

# Windows
Build $name-$version"-windows-x86" windows 386
Build $name-$version"-windows-x64" windows amd64

# Linux
Build $name-$version"-linux-amd64" linux amd64