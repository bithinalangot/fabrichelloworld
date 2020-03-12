package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func (c *Client) Deploy(contractPath, filePath, endpoint string) {
	data := makeDeployRequest("deploy", contractPath, "Init", "")
	rs := c.postRequest(data, endpoint)
	f, _ := os.Create(filePath)
	defer f.Close()
	f.WriteString(rs.Result.Message)
}

func (c *Client) HelloWorld(chaincodeIDFile, endpoint string) {
	chaincodeID, _ := ioutil.ReadFile(chaincodeIDFile)
	data := makeRequest("invoke", string(chaincodeID), "helloWorld", "hello")
	rs := c.postRequest(data, endpoint)
	fmt.Printf("Transaction ID:%v\n", rs.Result.Message)
}

func (c *Client) HelloWorldQuery(chaincodeIDFile, endpoint string) {
	chaincodeID, _ := ioutil.ReadFile(chaincodeIDFile)
	data := makeRequest("query", string(chaincodeID), "none", "hello")
	rs := c.postRequest(data, endpoint)
	fmt.Printf("%v\n", rs.Result.Message)
}
