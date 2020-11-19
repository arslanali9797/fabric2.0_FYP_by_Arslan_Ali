/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing a car
type SmartContract struct {
	contractapi.Contract
}

// Car describes basic details of what makes up a car
type AssetInfo struct {
	CNIC 	 					  string 	  `json:"CNIC"` 
	F_name   					  string 	  `json:"fname"`
	L_name   					  string 	  `json:"lname"`
	Gender	 					  string 	  `json:"gender"`
	Age 	 				 	  string      `json:"age"`
	Marital_status 				  string      `json:"marital_status"`
	Value_Of_Assets				  int		  `json:"value_of_assets"`
	Deposit_Per_Month 			  int         `json:"deposit_per_month"`
	Withdraw_Per_Month 			  int  	      `json:"withdraw_per_month"`
	Has_Been_Saving_For_Years 	  int         `json:"has_been_saving_for_years"`
	Transaction_Per_Month_Average int 		  `json:"transaction_per_month_average"`
	Type_Of_Bussiness			  string 	  `json:"type_of_bussiness"`
	Score 						  int		   `json:"score"`
	

}		



// InitLedger adds a base set of cars to the ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	assetInfos := []AssetInfo{
		AssetInfo{CNIC: "4550503812465" , F_name: "Arslan", L_name: "ALi", Age: "20", Gender: "Male", Marital_status:"single", Value_Of_Assets:1000000, Deposit_Per_Month:1, Withdraw_Per_Month:2, Has_Been_Saving_For_Years:2, Transaction_Per_Month_Average:15, Type_Of_Bussiness:"agriculture", Score:200},
		AssetInfo{CNIC: "5440503812223" , F_name: "Azhar",  L_name: "Ali", Age: "30", Gender: "Male", Marital_status:"single", Value_Of_Assets:1000000, Deposit_Per_Month:1, Withdraw_Per_Month:2, Has_Been_Saving_For_Years:2, Transaction_Per_Month_Average:15, Type_Of_Bussiness:"agriculture", Score:300},
		
	}

	for _, assetInfo := range assetInfos {
		assetJSON, err := json.Marshal(assetInfo)
		if err != nil {
			return err
		  }	

		err1 := ctx.GetStub().PutState(assetInfo.CNIC, assetJSON)
		if err1 != nil {
			return fmt.Errorf("Failed to put to world state. %s", err.Error())
		}
		
	}

	return nil
}

func (s *SmartContract) Addrecord(ctx contractapi.TransactionContextInterface, CNIC string, fname string, lname string,  age string, gender string, marital_status string, value_of_assets int, deposit_per_month int, withdraw_per_month int, has_been_saving_for_years int, transaction_per_month_average int , type_of_bussiness string , score int ) error {
	assetInfos := AssetInfo{
		CNIC: CNIC,
		F_name: fname,
		L_name: lname,
		Age: age,
		Gender: gender,
		Marital_status: marital_status,
		Value_Of_Assets: value_of_assets,
		Deposit_Per_Month: deposit_per_month,
		Withdraw_Per_Month: withdraw_per_month,
		Has_Been_Saving_For_Years: has_been_saving_for_years,
		Transaction_Per_Month_Average: transaction_per_month_average,
		Type_Of_Bussiness: type_of_bussiness,
		Score: score,

	}

	recoedAsBytes, _ := json.Marshal(assetInfos)

	return ctx.GetStub().PutState(CNIC, recoedAsBytes)
}




// GetAllAssets returns all assets found in world state
func (s *SmartContract) GetAllAssets(ctx contractapi.TransactionContextInterface) ([]*AssetInfo, error) {
	// range query with empty string for startKey and endKey does an
	// open-ended query of all assets in the chaincode namespace.
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
	  return nil, err
	}
	defer resultsIterator.Close()
  
	var assets []*AssetInfo
	for resultsIterator.HasNext() {
	  queryResponse, err := resultsIterator.Next()
	  if err != nil {
		return nil, err
	  }
  
	  var asset AssetInfo
	  err = json.Unmarshal(queryResponse.Value, &asset)
	  if err != nil {
		return nil, err
	  }
	  assets = append(assets, &asset)
	}
  
	return assets, nil
  }
  

// ReadAsset returns the asset stored in the world state with given id.
func (s *SmartContract) ReadAsset(ctx contractapi.TransactionContextInterface, id string) (*AssetInfo, error) {
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
	  return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if assetJSON == nil {
	  return nil, fmt.Errorf("the asset %s does not exist", id)
	}
  
	var asset AssetInfo
	err = json.Unmarshal(assetJSON, &asset)
	if err != nil {
	  return nil, err
	}
  
	return &asset, nil
  }

//calculate credit Score 
  func (s *SmartContract) CalculateScore(ctx contractapi.TransactionContextInterface, id string) (*AssetInfo, error) {
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
	  return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if assetJSON == nil {
	  return nil, fmt.Errorf("the asset %s does not exist", id)
	}

	var asset AssetInfo
	err = json.Unmarshal(assetJSON, &asset)
	if err != nil {
	  return nil , err
	}

	
	
	asset.Score=10+gender("female")

	return &asset  ,nil
  }


  func gender(gender string)(int){

	if (gender=="male"){
		return 0
	}
	return 2
}
func main() {

	chaincode, err := contractapi.NewChaincode(new(SmartContract))

	if err != nil {
		fmt.Printf("Error create credit score chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting credit score chaincode: %s", err.Error())
	}
}