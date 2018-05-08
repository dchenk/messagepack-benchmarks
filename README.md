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

## Results
This is what I get on my workstation (MacOS 10.13.3; 3.6 GHz Intel Core i7; 32 GB 2400 MHz DDR4):
```
BenchmarkCustomDchenk1-8        10000000               217 ns/op             160 B/op          1 allocs/op
BenchmarkCustomTinylib1-8       10000000               218 ns/op             160 B/op          1 allocs/op
BenchmarkCustomDchenk2-8         1000000              1590 ns/op            1440 B/op          3 allocs/op
BenchmarkCustomTinylib2-8        1000000              1779 ns/op            1440 B/op          3 allocs/op
```
There is not a significant difference for small objects, but it gets noticeable with larger objects;
dchenk/msgp performs better than tinylib/msgp. Also, I've noticed many times on my computer that the
first couple tests that run always take a little bit more time than the subsequent tests, which means
the CPU begins to allocate more resources to the go benchmarking tool a couple seconds after it starts
up.

Keep in mind also that the dchenk/msgp code generator is faster, and the runtime library dchenk/msgp/msgp
is much smaller and leaner than the tinylib one--and safer, without unsafely converting memory holding
`[]byte` to `string` type. More info here: [https://github.com/dchenk/msgp](https://github.com/dchenk/msgp)

## Sizes of the two libraries
In case you care about the size of your program (you should), the tinylib/msgp/msgp library compiled is 552K but dchenk/msgp/msgp compiled is 519K.
