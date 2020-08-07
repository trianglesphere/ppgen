all:
	go build

ver=0.0.1
rel=2
clean:
	go clean
	rm -rf src/ pkg/ ppgen-$(ver).tar.gz ppgen-$(ver)-$(rel)-x86_64.pkg.tar.zst
