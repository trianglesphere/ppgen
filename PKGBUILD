# Maintainer: Joshua Gutow <jbgutow@gmail.com>
pkgname=ppgen
pkgver=0.0.1
pkgrel=1
pkgdesc='Pass phrase generator.'
arch=('x86_64')
# url=""
license=('GPL')
makedepends=('go')
source=("$pkgname-$pkgver.tar.gz")
md5sums=('82d7ace62a84bc01eb64f7794287835b')

prepare(){
  cd "$SRCDIR"
  mkdir -p build/
}

build() {
  cd "$SRCDIR"
  export CGO_CPPFLAGS="${CPPFLAGS}"
  export CGO_CFLAGS="${CFLAGS}"
  export CGO_CXXFLAGS="${CXXFLAGS}"
  export CGO_LDFLAGS="${LDFLAGS}"
  export GOFLAGS="-buildmode=pie -trimpath -mod=readonly -modcacherw"
  go build -o build .
}

package() {
  cd "$SRCDIR"
  install -Dm755 build/$pkgname "$pkgdir"/usr/bin/$pkgname
}
