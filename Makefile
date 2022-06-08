build: ### build application
	go build main.go
.PHONY: build

install: ### install application
	go install
.PHONY: install

clean: ### clean folder after build
	rm -rf main
.PHONY: clean

publish: build install clean ### Install cli go use global
.PHONY: publish


