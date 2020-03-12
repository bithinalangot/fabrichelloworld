package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type HelloWorld struct{}

func main() {
	err := shim.Start(new(HelloWorld))
	if err != nil {
		fmt.Printf("Error starting BlockchainID: %s", err)
	}
}

func (t *HelloWorld) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	// We could initialize.
	return nil, nil
}

func (t *HelloWorld) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	switch function {
	case "helloWorld":
		return t.helloWorld(stub, args)
	}

	return nil, nil
}

func (t *HelloWorld) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	val, _ := stub.GetState(args[0])
	if val == nil {
		return nil, fmt.Errorf("Key does not exist %v" + args[0])
	}
	return []byte(val), nil
}

func (t *HelloWorld) helloWorld(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	stub.PutState(args[0], []byte("Hello World"))
	return []byte("Value added successfully"), nil
}
