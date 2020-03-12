package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	pb "github.com/hyperledger/fabric/protos"
)

const REQUEST_PREFIX = `{
  "jsonrpc": "2.0",
  "id": 123,
  "method": `

const REQUEST_PARAMS = `,
  "params": {
    "type": 1,
    "chaincodeID": {
      "name": `

const DEPLOY_REQUEST_PARAMS = `,
  "params": {
    "type": 1,
    "chaincodeID": {
      "path": `

const REQUEST_ARGS = `},
  "ctorMsg": {
    "function": `

const FUNCTION_ARGS = `,
    "args": [`

const SUFFIX = `]
  }
  }
}`

func makeDeployRequest(method, chaincodePath, function string, args ...string) []byte {
	return []byte(REQUEST_PREFIX + `"` + method + `"` + DEPLOY_REQUEST_PARAMS + `"` + chaincodePath + `"` + REQUEST_ARGS +
		`"` + function + `"` + FUNCTION_ARGS + `"` + args[0] + `"` + SUFFIX)
}

func makeRequest(method, chaincodeID, function string, args ...string) []byte {
	tmp := ""
	for i := 0; i < len(args)-1; i++ {
		tmp = tmp + `"` + args[i] + `", `
	}
	tmp = tmp + `"` + args[len(args)-1] + `"`
	return []byte(REQUEST_PREFIX + `"` + method + `"` + REQUEST_PARAMS + `"` + chaincodeID + `"` + REQUEST_ARGS +
		`"` + function + `"` + FUNCTION_ARGS + tmp + SUFFIX)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
func toInt(arg string) int {
	if i, e := strconv.Atoi(arg); e != nil {
		panic("Error converting to int ")
	} else {
		return i
	}
}

func (c *Client) nextRequest() time.Time {
	iv := 1000 * 1000 * 1000 / c.nRequestsPerSecond
	return time.Now().Add(time.Duration(iv))
}

func (c *Client) postRequest(data []byte, endpoint string) rpcResponse {
	res, err := http.Post("http://"+endpoint+":7050/chaincode", "application/json", bytes.NewReader(data))
	check(err)
	defer res.Body.Close()
	var rs rpcResponse
	check(json.NewDecoder(res.Body).Decode(&rs))
	return rs
}

type strArgs struct {
	Function string
	Args     []string
}

type rpcRequest struct {
	Jsonrpc *string           `json:"jsonrpc,omitempty"`
	Method  *string           `json:"method,omitempty"`
	Params  *pb.ChaincodeSpec `json:"params,omitempty"`
	ID      *int64            `json:"id,omitempty"`
}

type rpcID struct {
	StringValue *string `json: "omitempty"`
	IntValue    *int64  `json: "omitempty"`
}

type rpcResponse struct {
	Jsonrpc string     `json:"jsonrpc,omitempty"`
	Result  *rpcResult `json:"result,omitempty"`
	Error   *rpcError  `json:"error,omitempty"`
	ID      *int64     `json:"id"`
}

// rpcResult defines the structure for an rpc sucess/error result message.
type rpcResult struct {
	Status  string    `json:"status,omitempty"`
	Message string    `json:"message,omitempty"`
	Error   *rpcError `json:"error,omitempty"`
}

// rpcError defines the structure for an rpc error.
type rpcError struct {
	// A Number that indicates the error type that occurred. This MUST be an integer.
	Code int64 `json:"code,omitempty"`
	// A String providing a short description of the error. The message SHOULD be
	// limited to a concise single sentence.
	Message string `json:"message,omitempty"`
	// A Primitive or Structured value that contains additional information about
	// the error. This may be omitted. The value of this member is defined by the
	// Server (e.g. detailed error information, nested errors etc.).
	Data string `json:"data,omitempty"`
}
