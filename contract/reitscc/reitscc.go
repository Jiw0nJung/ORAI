package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type SmartContract struct {
}

type UserInfo struct {
	User  string `json:"user"`
	Token int    `json:"token"`
	Stake int    `json:"stake"`
	//Project_ref string `json:"project_ref"`
}

type ProjectInfo struct {
	ProjectTitle string `json:"projecttitle"`
	Token        int    `json:"token"`
	Stake        int    `json:"stake"`
}

func (t *SmartContract) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("SmartContract Init")
	args := stub.GetStringArgs()
	var A, B string                        // Entities
	var Atoken, Btoken, Astake, Bstake int // Asset holdings
	var err error

	if len(args) != 6 {
		return shim.Error("Incorrect number of arguments. Expecting 6")
	}

	// Initialize the chaincode
	A = args[0]
	Atoken, err = strconv.Atoi(args[1])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding")
	}

	Astake, err = strconv.Atoi(args[2])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding")
	}

	B = args[3]
	Btoken, err = strconv.Atoi(args[4])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding")
	}
	Bstake, err = strconv.Atoi(args[5])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding")
	}

	fmt.Printf("Usertoken = %d, Userstake = %d\n Projecttoken = %d Projectstake", Atoken, Astake, Btoken, Bstake)

	// Write the state to the ledger

	var user = UserInfo{User: A, Token: Atoken, Stake: Astake}
	userAsBytes, _ := json.Marshal(user)

	var project = ProjectInfo{ProjectTitle: B, Token: Btoken, Stake: Bstake}
	projectAsBytes, _ := json.Marshal(project)

	stub.PutState(A, userAsBytes)
	stub.PutState(B, projectAsBytes)

	return shim.Success(nil)
}

func (t *SmartContract) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("SmartContract Invoke")
	function, args := stub.GetFunctionAndParameters()
	if function == "addUser" {
		return t.addUser(stub, args)
	} else if function == "invest" {
		return t.invest(stub, args)
	} else if function == "query" {
		return t.query(stub, args)
	}

	return shim.Error("Invalid invoke function name. Expecting \"addUser\" \"invest\" \"query\"")
}

func (t *SmartContract) addUser(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 1 {
		return shim.Error("fail!")
	}
	var user = UserInfo{User: args[0], Token: 100, Stake: 0}
	userAsBytes, _ := json.Marshal(user)
	stub.PutState(args[0], userAsBytes)

	return shim.Success(nil)
}

func (t *SmartContract) invest(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var A, B string
	var X int
	var err error
	userinfo := UserInfo{}
	projectinfo := ProjectInfo{}

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	A = args[0]
	B = args[1]
	X, err = strconv.Atoi(args[2])
	if err != nil {
		return shim.Error("Invalid transaction amount, expecting a integer value")
	}
	//////////////////////////////////////////////////////////////////

	Aget, err := stub.GetState(A)
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if Aget == nil {
		return shim.Error("Entity not found")
	}

	json.Unmarshal(Aget, &userinfo)

	userinfo.Token -= X
	userinfo.Stake += X

	Aget, _ = json.Marshal(userinfo)

	//////////////////////////////////////////////////////////////////

	Bget, err := stub.GetState(B)
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if Bget == nil {
		return shim.Error("Entity not found")
	}

	json.Unmarshal(Bget, &projectinfo)

	projectinfo.Token -= X
	projectinfo.Stake += X

	Bget, _ = json.Marshal(projectinfo)

	fmt.Printf("UserToken = %d, UserStake = %d\n, ProjectToken = %d, ProjectStake = %d\n", userinfo.Token, userinfo.Stake, projectinfo.Token, projectinfo.Stake)

	err = stub.PutState(A, Aget)
	if err != nil {
		return shim.Error("Failed in putState")
	}

	err = stub.PutState(B, Bget)
	if err != nil {
		return shim.Error("Failed in putState")
	}

	return shim.Success(nil)
}

//////////////////////////////////////////////////////////////////

func (t *SmartContract) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	UserAsBytes, _ := stub.GetState(args[0])
	return shim.Success(UserAsBytes)
}

func main() {
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error starting SmartContract chaincode: %s", err)
	}
}
