# messagepack-benchmarks
Performance comparisons of the dchenk/msgp MessagePack implementation with rivals

1. Download: `go get -u github.com/dchenk/messagepack-benchmarks`
2. Move into the directory: `cd $(go env GOPATH)/src/github.com/dchenk/messagepack-benchmarks`
2. Set things up: `make setup`
4. Run: `make bench`

Try to have nothing else running on your computer so that the tests are not impacted by other stuff
your computer is doing.
