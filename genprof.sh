#!/bin/sh
# To generate a profile execute "g3nd -cpuprofile g3nd.prof" for
# some time to collect profile data. Then exits "g3nd" and
# execute this script to generate the profile diagram in svg.
# For Go1.9+ only the profile file is necessary
go tool pprof -svg g3nd.prof > g3nd.svg


