all:
	go build

package-arch:
	./contrib/build_tarball.sh

clean:
	go clean
	rm -rf build
