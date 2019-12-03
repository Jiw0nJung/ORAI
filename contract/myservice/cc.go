// 패키지 정의
package main

import (
	"encoding/json"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// 외부모듈 포함

//표준 I/O 라이브러리

// 피어 api

// 체인코드 데이터 정의
type User struct {
	UserID   string    `json:"user"`
	Token    int       `json:"token"`
	Projects []Project `json:"projects"`
}

type Project struct {
	ProjectTitle string `json:"projecttitle"`
	HoldingStake int    `json:"holdingstake"`
}

type ProjectInfo struct {
	ProjectTitle   string  `json:"projecttitle"`
	EntireStake    int     `json:"entirestake`
	ExpectedReturn float64 `json:"expectedreturn"`
}

func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	function, args := APIstub.GetFunctionAndParameters()

	if function == "addUser" {
		return s.addUser(APIstub, args)
	} else if function == "addProjectInfo" {
		return s.addProjectInfo(APIstub, args)
	} else if function == "buyToken" {
		return s.buyToken(APIstub, args)
	} else if function == "perchaseStake" {
		return s.perchaseStake(APIstub, args)
	} else if function == "query" {
		return s.query(APIstub, args)
	}
	return shim.Error("Invalid Smart Contract function name.")
}

func (s *SmartContract) addUser(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("fail!")
	}
	var user = UserRating{User: args[0], Average: 0}
	userAsBytes, _ := json.Marshal(user)
	APIstub.PutState(args[0], userAsBytes)

	return shim.Success(nil)
}
