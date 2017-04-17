// This trivial smart contract stores an integer value against a key.
// While storing , it deducts an amount from the input number.
// The deduction amount is set during the deployment.

package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// TrivalSmartContract -  Example of a simple start contratc implementation
type TrivalSmartContract struct {
}

// Starting point of the smart contract
func main() {
	err := shim.Start(new(TrivalSmartContract))
	if err != nil {
		fmt.Printf("Error starting TrivalSmartContract chaincode: %s", err)
	}
}

// Init resets all the things
func (t *TrivalSmartContract) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}

	_, err := strconv.Atoi(args[0])
	if err != nil {
		return nil, errors.New("Incorrect deduction amount. Should be an integer")
	}

	errStore := stub.PutState("deduct_amt", []byte(args[0]))

	if errStore != nil {
		return nil, errStore
	}

	return nil, nil
}

// Invoke invoke entry point to invoke a chaincode function
func (t *TrivalSmartContract) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "init" {
		return t.Init(stub, "init", args)
	} else if function == "deposite" {
		return t.deposite(stub, args)
	}
	fmt.Println("invoke did not find func: " + function)

	return nil, errors.New("Received unknown function invocation: " + function)
}

// Query is our entry point for queries
func (t *TrivalSmartContract) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)

	// Handle different functions
	if function == "read" { //read a variable
		return t.read(stub, args)
	}
	fmt.Println("query did not find func: " + function)

	return nil, errors.New("Received unknown function query: " + function)
}

// deposite - invoke function to write key/value pair after applying the business logic.
func (t *TrivalSmartContract) deposite(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var key string
	var err error
	fmt.Println("running deposite()")

	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2. name of the key and value to be deposited")
	}

	key = args[0] //rename for funsies
	value, err := strconv.Atoi(args[1])
	if err != nil {
		return nil, errors.New("Invalid value provided. It should be an integer")
	}

	deductStr, errGet := stub.GetState("deduct_amt")
	if errGet != nil {
		return nil, errors.New("Unable to retrieve the deduct_amt key")
	}
	deductVal, _ := strconv.Atoi(string(deductStr))
	storeVal := value - deductVal
	strToStore := strconv.Itoa(storeVal)

	errStore := stub.PutState(key, []byte(strToStore))

	if errStore != nil {
		return nil, errStore
	}
	return nil, nil
}

// read - query function to read key/value pair
func (t *TrivalSmartContract) read(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var key, jsonResp string
	var err error

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the key to query")
	}

	key = args[0]
	valAsbytes, err := stub.GetState(key)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
		return nil, errors.New(jsonResp)
	}

	return valAsbytes, nil
}
