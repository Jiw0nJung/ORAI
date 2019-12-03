// 패키지 정의
package main

// 외부모듈 포함
import (
	"encoding/json"
	"fmt" //표준 I/O 라이브러리

	"github.com/hyperledger/fabric/core/chaincode/shim" //심 인터페이스 api
	"github.com/hyperledger/fabric/protos/peer"         // 피어 api
)

// 체인코드 데이터 정의
type SimpleAsset struct {
}

type Key struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// init 함수 정의
func (t *SimpleAsset) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

// invoke 함수 정의
func (t *SimpleAsset) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	fn, args := stub.GetFunctionAndParameters()

	var result string

	if fn == "create" {
		result, _ = create(stub, args)
	} else if fn == "query" {
		result, _ = query(stub, args)
	} else if fn == "modify" {
		result, _ = modify(stub, args)
	} else {
		return shim.Error("Not supported function name")
	}
	return shim.Success([]byte(result))
}

// query 함수 정의
func query(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 1 {
		return "", fmt.Errorf("Incorrect arguments. Expecting a key")
	}

	value, err := stub.GetState(args[0])
	if err != nil {
		return "", fmt.Errorf("Failed in getState: %s", args[0])
	}
	if value == nil {
		return "", fmt.Errorf("Asset not found: %s", args[0])
	}
	return string(value), nil
}

// create 함수 정의
func create(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 2 {
		return "", fmt.Errorf("Incorrect arguments, Expecting a key and a value")
	}

	var key = Key{Name: args[0], Value: args[1]}
	keyAsBytes, _ := json.Marshal(key)
	err := stub.PutState(args[0], keyAsBytes)
	if err != nil {
		return "", fmt.Errorf("Failed in putState: %s", args[0])
	}
	return "tx is summited", nil
}

// modify 함수 정의
func modify(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 2 {
		return "", fmt.Errorf("Incorrect arguments, Expecting a key and a value")
	}

	value, err := stub.GetState(args[0])

	if value == nil {
		return "", fmt.Errorf("Asset not found: %s", args[0])
	}
	var key = Key{}
	json.Unmarshal(value, &key)
	key.Value = args[1]

	keyAsBytes, _ := json.Marshal(key)

	err = stub.PutState(args[0], keyAsBytes)
	if err != nil {
		return "", fmt.Errorf("Failed in putState: %s", args[0])
	}
	return "tx is summited", nil

}
func main() {
	if err := shim.Start(new(SimpleAsset)); err != nil {
		fmt.Printf("Error starting SimpleAsset chaincode: %s", err)
	}
}
