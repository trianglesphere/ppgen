#!/bin/bash
PKGNAME=ppgen
PKGVER="0.0.1"
SRCS="go.mod ppgen.go internal/wordlists/*"

tar -czf $PKGNAME-$PKGVER.tar.gz $SRCS
md5sum $PKGNAME-$PKGVER.tar.gz
