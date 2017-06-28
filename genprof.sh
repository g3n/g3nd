#!/bin/sh
# To generate a profile execute "g3nd -profile g3nd.prof" for
# some time to collect profile data. Then exits "g3nd" and
# execute this script to generate the profile diagram in svg.
go tool pprof -svg $GOPATH/bin/g3nd g3nd.prof > g3nd.svg


