package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

type SmartContract struct {
}

type UserInfo struct {
	User         string `json:"user"`
	Car          string `json:"car"`
	Menufacturer string `json:"menufacturer"`
	Datas        []Data `json:"datas"`
}
type Data struct {
	Weather           string `json:"weather"`
	Lighting          string `json:"lighting"`
	RoadwaySurface    string `json:"roadwaysurface"`
	RoadwayConditions string `json:"roadwayconditions"`
	Mpc               string `json:"mpc"`
	Toc               string `json:"toc"`
	// Mpc - Movement Preceding Collision
	// Toc - Type Of Collision
}

func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	function, args := APIstub.GetFunctionAndParameters()

	if function == "addUser" {
		return s.addUser(APIstub, args)
	} else if function == "addAccidents" {
		return s.addAccidents(APIstub, args)
	} else if function == "viewAccidents" {
		return s.viewAccidents(APIstub, args)
	}
	return shim.Error("Invalid Smart Contract function name.")
}

func (s *SmartContract) addUser(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 3 {
		return shim.Error("fail!")
	}
	var user = UserInfo{User: args[0], Car: args[1], Menufacturer: args[2]}
	userAsBytes, _ := json.Marshal(user)
	APIstub.PutState(args[0], userAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) addAccidents(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 7 {
		return shim.Error("Incorrect number of arguments. Expecting 7")
	}
	// getState User
	userAsBytes, err := APIstub.GetState(args[0])
	if err != nil {
		jsonResp := "\"Error\":\"Failed to get state for " + args[0] + "\"}"
		return shim.Error(jsonResp)
	} else if userAsBytes == nil { // no State! error
		jsonResp := "\"Error\":\"User does not exist: " + args[0] + "\"}"
		return shim.Error(jsonResp)
	}
	// state ok
	user := UserInfo{}
	err = json.Unmarshal(userAsBytes, &user)
	if err != nil {
		return shim.Error(err.Error())
	}
	// create rate structure
	var Data = Data{Weather: args[1], Lighting: args[2], RoadwaySurface: args[3], RoadwayConditions: args[4], Mpc: args[5], Toc: args[6]}

	user.Datas = append(user.Datas, Data)

	userAsBytes, err = json.Marshal(user)

	APIstub.PutState(args[0], userAsBytes)

	return shim.Success([]byte("data is updated"))
}

func (s *SmartContract) viewAccidents(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	UserAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(UserAsBytes)
}

func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
