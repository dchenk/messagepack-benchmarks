.PHONY: dchenk_msgp tinylib_msgp setup bench gen

GOPATH = $(shell go env GOPATH)

dchenk_msgp:
	go get -u github.com/dchenk/msgp
	-mv $(GOPATH)/bin/msgp $(GOPATH)/bin/msgp_dchenk

tinylib_msgp:
	go get -u github.com/tinylib/msgp
	-mv $(GOPATH)/bin/msgp $(GOPATH)/bin/msgp_tinylib

setup: dchenk_msgp tinylib_msgp gen

gen:
	go generate

bench:
	go test -run NNN -bench Custom -benchmem
