## prnsrv

### Why?

I regularly just want to dump the output from a program pointed at a server
endpoint. This prints out whatever is sent in the `prnsrv` logs.

### How to use?

1. cd into the repo
2. run either `go run prnsrv.go` or `go install`
3. run `prnsrv` if you want this to listen on port 8000 or provide a port using
   the flag `-port <port-number>`.

### Other
There are better tools for this, this is just an easy thing for me to use.
