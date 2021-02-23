all:
	go build

tarball:
	./contrib/create_tarball.sh

clean:
	go clean
	rm -rf build
