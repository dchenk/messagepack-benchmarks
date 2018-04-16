# messagepack-benchmarks
Performance comparisons of the dchenk/msgp MessagePack implementation with rivals

1. Download: `go get -u github.com/dchenk/messagepack-benchmarks`
2. Move into the directory: `cd $(go env GOPATH)/src/github.com/dchenk/messagepack-benchmarks`
2. Set things up: `make setup`
4. Run: `make bench`

Try to have nothing else running on your computer so that the tests are not impacted by other stuff
your computer is doing.

**Running `make setup` will download the latest version of github.com/dchenk/msgp, install it where
your Go binaries are installed ($GOPATH/bin), and immediately rename it from `msgp` to `msgp_dchenk`.
To be able to use the `msgp` code generator in your projects, you need to install the CLI tool again
like so: `cd $(go env GOPATH)/src/github.com/dchenk/msgp && go install`. Always try to keep up with
the latest release.**
