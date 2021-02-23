#!/bin/bash
PKGNAME=ppgen
PKGVER="0.1.0"
SRCS="go.mod ppgen.go wordlists/*"

if [ $(basename `pwd`) = "contrib" ]; then 
	cd ..;
fi;

mkdir -p build
tar -C . -czf build/$PKGNAME-$PKGVER.tar.gz $SRCS
md5sum build/$PKGNAME-$PKGVER.tar.gz
