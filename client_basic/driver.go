package main

import (
	"io/ioutil"
	"strings"
)

type Client struct {
	startIdx, endIdx     int      // for the key range
	nThreads             int      // number of threads
	nOutstandingRequests int      // number of outstanding requests
	nRequestsPerSecond   int      // rate, invert to see the next request
	hosts                []string // all host to send requests
	duration             int      // how long to run, in seconds

	chaincodeID                 []byte
	pubkeys, privkeys, hashkeys [][]byte
}

// ns: number of servers
func NewClient(sid, eid, nt, nor, rps int, hostFile string, ns int, duration int) *Client {
	h, _ := ioutil.ReadFile(hostFile)
	hosts := strings.Split(string(h), "\n")
	return &Client{
		startIdx:             sid,
		endIdx:               eid,
		nThreads:             nt,
		nOutstandingRequests: nor,
		nRequestsPerSecond:   rps,
		hosts:                hosts[:ns],
		duration:             duration,
	}
}
