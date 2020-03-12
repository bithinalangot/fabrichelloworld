package main

import (
	"fmt"
	"os"
)

const CONTRACTPATH = "github.com/blockchain_id/contract/helloworld"
const CHAINCODEID = "didauthidv1"

func main() {
	c := Client{}
	switch os.Args[1] {
	case "deploy":
		{ // deploy <endpoint>
			c.Deploy(CONTRACTPATH, CHAINCODEID, os.Args[2])
		}
	case "helloWorld":
		{
			c.HelloWorld(CHAINCODEID, os.Args[2])
		}
	case "helloWorldQuery":
		{
			c.HelloWorldQuery(CHAINCODEID, os.Args[2])
		}
	}
	fmt.Printf("Done\n")
}
