## prnsrv

### Why?

I regularly just want to dump the output from a program pointed at a server
endpoint. This prints out whatever is sent in the `prnsrv` logs.

### How to use?

1. cd into the repo
2. run either `go run prnsrv.go` or `go run prnsrv.go -port <port-number>`
3. if you want to install, run `go install`, then you can run it using
   `prnsrv` or `prnsrv -port <port-number>` as long as `GOBIN` is in your path.

### Other
There are better tools for this, this is just an easy thing for me to use.
