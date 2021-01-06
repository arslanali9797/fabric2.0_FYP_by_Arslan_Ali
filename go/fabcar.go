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
	"strings"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
  //  "github.com/hyperledger/fabric/core/chaincode/shim/ext/cid"


)

// SmartContract provides functions for managing a Records
type SmartContract struct {
	contractapi.Contract
}

//  describes basic details 
type AssetInfo struct {

	CNIC 	 					  string 	  `json:"CNIC"` 
	
	F_name   					  string 	  `json:"fname"` //ok
	
	L_name   					  string 	  `json:"lname"`//ok

	Gender	 					  string 	  `json:"gender"`//ok
	
	Age 	 				 	  int     	  `json:"age"` //ok
	
	Marital_status 				  string      `json:"marital_status"` //ok
	
	Value_Of_Assets				  int		  `json:"value_of_assets"` //ok
	
	Deposit_Per_Month 			  int         `json:"deposit_per_month"`//ok
	
	Withdraw_Per_Month 			  int  	      `json:"withdraw_per_month"`//ok
	
	Has_Been_Saving_For_Years 	  int         `json:"has_been_saving_for_years"`//ok
	
	Transaction_Per_Month_Average int 		  `json:"transaction_per_month_average"` //ok
	
	Type_Of_Bussiness			  string 	  `json:"type_of_bussiness"`//ok

	Saving_Amount				  int 		  `json:"saving_amount"` //ok
	
	Bank1_Score					  int		   `json:"bank1_score"`

	Bank2_Score					  int		   `json:"bank2_score"`

	Final_Score 				  int		   `json:"finale_score"`
	
	MspOrg						string			`json:"msporg"`

}		

type UserDetails struct {

	Id 			string 		`json:"id"`
	Password 	string 		`json:"password"`

}

type ClientPrivateData struct {
	CNIC 	 					  string 	   `json:"CNIC"` 
	
	Credit_Type					  string 		`json:"credit_type"`

	Loan_Amount 				  float64 			`json:"loan_amount"`

	IssueLoanDate 				 string 		`json:"issue_loandate"`

	PaymentPlan					float64 			`json:"payment_plan"`

	InstalmentPerMonth 			float64 			`json:"instalment_per_month"`
	
	PaymentDate  				string  		`json:"payment_date"`
}





// InitLedger adds a base set of Records to the ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	assetInfos := []AssetInfo{
		AssetInfo{CNIC: "4550503812465" , F_name: "Arslan", L_name: "ALi", Age: 20, Gender: "Male", Marital_status:"single", Value_Of_Assets:1000000, Deposit_Per_Month:1, Withdraw_Per_Month:2, Has_Been_Saving_For_Years:2, Transaction_Per_Month_Average:15, Type_Of_Bussiness:"agriculture",Saving_Amount:2000000,Bank1_Score:434,Bank2_Score:300, Final_Score:500,MspOrg:"org1MSP"},
		AssetInfo{CNIC: "5440503812223" , F_name: "Azhar",  L_name: "Ali", Age: 30, Gender: "Male", Marital_status:"single", Value_Of_Assets:1000000, Deposit_Per_Month:1, Withdraw_Per_Month:2, Has_Been_Saving_For_Years:2, Transaction_Per_Month_Average:15, Type_Of_Bussiness:"agriculture",Saving_Amount:200000,Bank1_Score:234,Bank2_Score:500, Final_Score:300,MspOrg:"org2MSP"},	
	}
	for _, assetInfo := range assetInfos {
		assetJSON, err := json.Marshal(assetInfo)
		if err != nil {
			return err
		  }	

		err1 := ctx.GetStub().PutState(assetInfo.CNIC, assetJSON)
		if err1 != nil {
			return fmt.Errorf("Failed to put to world state. %s", err.Error())}
		}

	userDetails := []UserDetails{
		UserDetails {Id: "B1007org1msp" , Password: "b1007" },
		UserDetails {Id: "B2007org2msp" , Password: "b2007" },	
	}
	for _, userInfo := range userDetails{
		userJSON,err :=json.Marshal(userInfo)
		if err !=nil{
			return err
		}
		err2 := ctx.GetStub().PutState(userInfo.Id, userJSON)
		if err2 !=nil{
			return fmt.Errorf("Failed to register the user with thier id and password. %s",err.Error())}

		}
	

	return nil
}

func (s *SmartContract) AddRecord(ctx contractapi.TransactionContextInterface, CNIC string, fname string,lname string, age int,gender string,marital_status string,value_of_assets int,deposit_per_month int,withdraw_per_month int ,has_been_saving_for_years int,transaction_per_month_average int, type_of_bussiness string,saving_amount int,bank1_score int ,bank2_score int ,final_Score int,msporg string) error {
	


	var genderIs= genderF(gender)
	var ageIs = ageF(age)
	var maritalStatus= martialStatusF(marital_status)
	var assetValue = assetsValueF(value_of_assets)	
	var depositPerMmonth = monthlyDepositsF(deposit_per_month)
	var withdrawPerMonth= monthlyWithdrawF(withdraw_per_month)
	var hasBeenSavingForYears = savingYearsF(has_been_saving_for_years)
	var transactionPerMonthAverage = monthlyTransactionF(transaction_per_month_average)
	var typeOfBussiness = bussinessTypeF(type_of_bussiness)
	var savingAmount = currentSavingAmountF(saving_amount) 
	var bank1_scoreN =0
	var bank2_scoreN = 0

	if (msporg=="org1MSP"){
	 	bank1_scoreN = genderIs+ageIs+maritalStatus+assetValue+depositPerMmonth+withdrawPerMonth+hasBeenSavingForYears+transactionPerMonthAverage+typeOfBussiness+savingAmount
	}else if (msporg=="org2MSP"){
		bank2_scoreN = genderIs+ageIs+maritalStatus+assetValue+depositPerMmonth+withdrawPerMonth+hasBeenSavingForYears+transactionPerMonthAverage+typeOfBussiness+savingAmount	
	}

	var final_ScoreN  = 0

	if(bank1_scoreN==0){
	 	final_ScoreN = bank2_scoreN 
	}else if (bank2_scoreN==0){
		final_ScoreN = bank1_scoreN
	}else {
		final_ScoreN = (bank2_scoreN+bank1_scoreN)/2
	}
	

	assetInfo := AssetInfo{

		CNIC:									CNIC,
		F_name:									fname, 
		L_name:									lname, 
		Age: 									age,
		Gender: 								gender,
		Marital_status: 						marital_status,
		Value_Of_Assets: 						value_of_assets,
		Deposit_Per_Month: 						deposit_per_month,
		Withdraw_Per_Month: 					withdraw_per_month,
		Has_Been_Saving_For_Years:				has_been_saving_for_years,
		Transaction_Per_Month_Average:  		transaction_per_month_average,
	    Type_Of_Bussiness: 						type_of_bussiness,
	    Saving_Amount:  						saving_amount,
		Bank1_Score: 							bank1_scoreN, 
		Bank2_Score: 							bank2_scoreN, 
		Final_Score: 							final_ScoreN,
		MspOrg:									msporg,
	
	}

	assetInfoBytes, _ := json.Marshal(assetInfo)

	return ctx.GetStub().PutState(CNIC, assetInfoBytes)

}


//gender Function
func genderF(gender string)(int){

	if (gender=="male"){
		return 20
	}
	return 50	
}

//Age FUnction
func ageF(age int )(int){

	if(age<=25){
	return 	20
}	else if (age>25 && age <=35){
		return 40
	} else if (age>35 && age<=46){	
		return 50
	}
	return 70
}


func martialStatusF(status string)(int){
	if (status=="bacheler"){
		return 30
	}else if (status=="widow" || status=="widower"){
		return 50
	}else if (status =="married"){
		return 100
	}

	return 0
} 


func assetsValueF(value int)(int){
	if(value<=300000){
		return 30
	} else if(value>300000 && value<=500000){
		return 50
	}else if(value>500000 && value<=700000){
		return 70}
		
	return 90
	}



func savingYearsF(year int )(int){

	if (year < 5){
		return 30 
	} else if (year >=5 && year <=9){
		return 50 
	}else if (year >= 10 && year <=12){
		return 70
	}
	return 80
}

 



func monthlyDepositsF(no int )(int){
	if (no <=1){
		return 30
	} else if (no>=2 && no <=3){
		return 50
	}
	return 100
}


func monthlyWithdrawF(withdrawNo int )(int){
	if (withdrawNo ==0){
		return 100
	}else if( withdrawNo ==1 ){
		return 60
	}else if(withdrawNo>=2 && withdrawNo<=3 ){
		return 40
	}
	return 30
}


func currentSavingAmountF(amount int )(int){
	if(amount<100000){
		return 30
	} else if(amount>=100000 && amount<=300000){
		return 50
	}else if(amount>300000 && amount<=500000){
		return 80}
		
	return 100
}




func monthlyTransactionF(trans int )(int){
	if(trans<10){
		return 30
	} else if (trans>=10 && trans<=20){
		return 40
	} else if (trans >20 && trans <50){
		return 60
	}
	return 80
}



func bussinessTypeF(Bname string)(int ){
	if (Bname == "govt"){
		return 30
	}else if (Bname == "private"){
		return 50
	}else if (Bname=="bussiness"){
		return 100
	}
	return 80
}
 

func (s *SmartContract) ClientPrivateRecord(ctx contractapi.TransactionContextInterface, CNICloan string,credit_type string,loan_amount float64 ,issue_loandate string,payment_plan float64,payment_date string,orgloan string) error {
	

	clientLoan, err := s.ReadAsset(ctx, CNICloan)

	if err != nil {
		return fmt.Errorf(`CLient Does not Exist`);
	}

	if  clientLoan.CNIC==CNICloan{

		var instalment_per_month = paymantPlanF(loan_amount,payment_plan)

		clientInfo := ClientPrivateData{
		CNIC:				CNICloan+orgloan,				//cinc unique with cnic+orgMSPloan 
		Credit_Type:		credit_type,			
		Loan_Amount:		loan_amount,
		IssueLoanDate:		issue_loandate,
		PaymentPlan:		payment_plan,
		InstalmentPerMonth:	instalment_per_month,	
		PaymentDate:		payment_date,	
	}
	
		clientInfoBytes, _ := json.Marshal(clientInfo)
		return ctx.GetStub().PutState(CNICloan+orgloan, clientInfoBytes)

	}

		return fmt.Errorf (`CLient Exist`);


	
}

func paymantPlanF( amount float64,year float64)(float64){

	if (year == 1){
		return amount/12
	}else if (year==2){
		return amount/24
	}else if (year==3){
		return amount/36
	}else if (year==4){
		return amount/48
	}
	return amount/60
}


func (s *SmartContract) GetAllClientLoan(ctx contractapi.TransactionContextInterface) ([]*ClientPrivateData, error) {

	
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
	  return nil, err
	}
	defer resultsIterator.Close()
  

	var assets []*ClientPrivateData
	for resultsIterator.HasNext() {
	  queryResponse, err := resultsIterator.Next()
	  if err != nil {
		return nil, err
	  }
  
	  var asset ClientPrivateData
	  err = json.Unmarshal(queryResponse.Value, &asset)
	  if err != nil {
		return nil, err
	  }

	  if((strings.Contains(asset.CNIC, "org"))){
		assets = append(assets, &asset)
	} 
	  
	}
  
	return assets, nil
  }
  


















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
  
  func (s *SmartContract) GetAllUsers(ctx contractapi.TransactionContextInterface) ([]*UserDetails, error) {
	// range query with empty string for startKey and endKey does an
	// open-ended query of all assets in the chaincode namespace.
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
	  return nil, err
	}
	defer resultsIterator.Close()
  

	var assets []*UserDetails
	for resultsIterator.HasNext() {
	  queryResponse, err := resultsIterator.Next()
	  if err != nil {
		return nil, err
	  }
  
	  var asset UserDetails
	  err = json.Unmarshal(queryResponse.Value, &asset)
	  if err != nil {
		return nil, err
	  }
	  assets = append(assets, &asset)
	}
  
	return assets, nil
  }





  


// ReadAsset returns the asset stored in the world state with given id.
func (s *SmartContract) ReadAsset(ctx contractapi.TransactionContextInterface, CNIC string) (*AssetInfo, error) {
	assetJSON, err := ctx.GetStub().GetState(CNIC)
	if err != nil {
	  return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if assetJSON == nil {
	  return nil, fmt.Errorf("the asset %s does not exist", CNIC)
	}
  
	var asset AssetInfo
	err = json.Unmarshal(assetJSON, &asset)
	if err != nil {
	  return nil, err
	}
  
	return &asset, nil
  }
  




// ReadAsset returns the asset stored in the world state with given id.
func (s *SmartContract) ReadUser(ctx contractapi.TransactionContextInterface, Id string) (*UserDetails, error) {
	
	assetJSON, err := ctx.GetStub().GetState(Id)
	if err != nil {
	  return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if assetJSON == nil {
	  return nil, fmt.Errorf("the asset %s does not exist", Id)
	}
  
	var asset UserDetails
	err = json.Unmarshal(assetJSON, &asset)
	if err != nil {
	  return nil, err
	}
  
	return &asset, nil
  }


  func (s *SmartContract) LoginUser(ctx contractapi.TransactionContextInterface, Id string,Password string) string {

	user, err := s.ReadUser(ctx, Id)

	if err != nil {
		return `The error is in LoginQuery`
	}
	
  check := ``
  if user.Id == Id && user.Password == Password {	// check the Value
    check = `true`
  } else {
    check = `fasle`
  }

	return check





}





//calculate credit Score 
  


  
 

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