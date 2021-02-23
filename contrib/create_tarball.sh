#!/bin/bash
PKGNAME=ppgen
PKGVER="0.1.0"
SRCS="go.mod ppgen.go wordlists/*"

if [ $(basename `pwd`) = "contrib" ]; then do
	cd ..;
fi;

mkdir -p build

tar -czf build/$PKGNAME-$PKGVER.tar.gz $SRCS
md5sum $PKGNAME-$PKGVER.tar.gz
