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
 * specific language governming permissions and limitations
 * under the License.
 */

package main

import (
	

	"encoding/json"
	"fmt"
	//"strings"
	

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

	First_Name					string			`json:"first_name"`

	Last_Name 					string			`json:"last_name"`
	
	Credit_Type					  string 		`json:"credit_type"`

	Loan_Amount 				  float64 			`json:"loan_amount"`

	IssueLoanDate 				 string 		`json:"issue_loandate"`

	PaymentPlan					float64 			`json:"payment_plan"`

	InstalmentPerMonth 			float64 			`json:"instalment_per_month"`
	
	PaymentDate  				string  		`json:"payment_date"`

	Credit_Score_Is				int				`json:"Credit_Score_Is"`

	YearStructure 			  PlanYear2			
	}
/* 
type PlanYear1 struct {
	JAN,FEB,MAR,ARL,MAY,JUNE,JULY,AUG,SEP,OCT,NOV,DEC float64
	
}
 */
type PlanYear2 struct {
	JAN1,FEB1,MAR1,ARL1,MAY1,JUNE1,JULY1,AUG1,SEP1,OCT1,NOV1,DEC1,JAN2,FEB2,MAR2,ARL2,MAY2,JUNE2,JULY2,AUG2,SEP2,OCT2,NOV2,DEC2,JAN3,FEB3,MAR3,ARL3,MAY3,JUNE3,JULY3,AUG3,SEP3,OCT3,NOV3,DEC3,JAN4,FEB4,MAR4,ARL4,MAY4,JUNE4,JULY4,AUG4,SEP4,OCT4,NOV4,DEC4,JAN5,FEB5,MAR5,ARL5,MAY5,JUNE5,JULY5,AUG5,SEP5,OCT5,NOV5,DEC5 float64
	
}
 




// InitLedger adds a base set of Records to the ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	assetInfos := []AssetInfo{
		AssetInfo{CNIC: "4550503812465" , F_name: "Arslan", L_name: "ALi", Age: 20, Gender: "Male", Marital_status:"single", Value_Of_Assets:1000000, Deposit_Per_Month:1, Withdraw_Per_Month:2, Has_Been_Saving_For_Years:2, Transaction_Per_Month_Average:15, Type_Of_Bussiness:"agriculture",Saving_Amount:2000000,Bank1_Score:434,Bank2_Score:300, Final_Score:500,MspOrg:"org1MSP"},
		AssetInfo{CNIC: "4550503812467" , F_name: "Azhar",  L_name: "Ali", Age: 30, Gender: "Male", Marital_status:"single", Value_Of_Assets:1000000, Deposit_Per_Month:1, Withdraw_Per_Month:2, Has_Been_Saving_For_Years:2, Transaction_Per_Month_Average:15, Type_Of_Bussiness:"agriculture",Saving_Amount:200000,Bank1_Score:234,Bank2_Score:500, Final_Score:300,MspOrg:"org1MSP"},	
	}
	for _, assetInfo := range assetInfos {
		assetJSON, err := json.Marshal(assetInfo)
		if err != nil {
			return err
		  }	

		err1 := ctx.GetStub().PutState(assetInfo.CNIC+"org1msp", assetJSON)
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

func (s *SmartContract) AddRecord(ctx contractapi.TransactionContextInterface, CNIC string, fname string,lname string, age int,gender string,marital_status string,value_of_assets int,deposit_per_month int,withdraw_per_month int ,has_been_saving_for_years int,transaction_per_month_average int, type_of_bussiness string,saving_amount int,bank1_score int ,bank2_score int ,final_Score int,msporg string) string {
	

	if(msporg=="org1msp"){
		ClientRecord , err :=s.ReadAssetOrg(ctx,CNIC,msporg)


		if (ClientRecord != nil){ //yes found in B1
			return "no1"}

		
			if err !=nil{
				
		ClientRecord2 , err2 :=s.ReadAssetOrg(ctx,CNIC,"org2msp")
		
		
		
		if (ClientRecord2 != nil){ //Update the data from B2 (SOME fields)
		
			msporg="org2msp"

			if (ClientRecord2.CNIC == CNIC){

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
				
				var bank1_scoreN  = genderIs+ageIs+maritalStatus+assetValue+depositPerMmonth+withdrawPerMonth+hasBeenSavingForYears+transactionPerMonthAverage+typeOfBussiness+savingAmount
				var final_ScoreN = (bank1_scoreN+ClientRecord2.Bank2_Score)/2

				assetInfo := AssetInfo{

					CNIC:									CNIC,
					F_name:									ClientRecord2.F_name, 
					L_name:									ClientRecord2.L_name, 
					Age: 									ClientRecord2.Age,
					Gender: 								ClientRecord2.Gender,
					Marital_status: 						ClientRecord2.Marital_status,
					Value_Of_Assets: 						value_of_assets,
					Deposit_Per_Month: 						deposit_per_month,
					Withdraw_Per_Month: 					withdraw_per_month,
					Has_Been_Saving_For_Years:				has_been_saving_for_years,
					Transaction_Per_Month_Average:  		transaction_per_month_average,
					Type_Of_Bussiness: 						type_of_bussiness,
					Saving_Amount:  						saving_amount,
					Bank1_Score: 							bank1_scoreN, 
					Bank2_Score: 							ClientRecord2.Bank2_Score, 
					Final_Score: 							final_ScoreN,
					MspOrg:									msporg,
				
				}
				assetInfoBytes, _ := json.Marshal(assetInfo)
				
				err3 := ctx.GetStub().PutState(CNIC+msporg, assetInfoBytes)
				
				if(err3 != nil){

					return "yes1"
				}





			return "no2"}
			
			 return "no4"

			
		}

		if err2 !=nil{ //add complete record
	
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

	if (msporg=="org1msp"){
	 	bank1_scoreN = genderIs+ageIs+maritalStatus+assetValue+depositPerMmonth+withdrawPerMonth+hasBeenSavingForYears+transactionPerMonthAverage+typeOfBussiness+savingAmount
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
				
				err3 := ctx.GetStub().PutState(CNIC+msporg, assetInfoBytes)
				
				if(err3 != nil){

					return "yes"
				}
}
}
}





//now for org2 




if(msporg=="org2msp"){
	ClientRecord , err :=s.ReadAssetOrg(ctx,CNIC,msporg)


	if (ClientRecord != nil){ //yes found in B1
		return "no1"}

	
		if err !=nil{
			
			
	ClientRecord2 , err2 :=s.ReadAssetOrg(ctx,CNIC,"org1msp")

	if (ClientRecord2 != nil){ //Update the data from B2 (SOME fields)
		
		
		if (ClientRecord2.CNIC == CNIC){
		
			msporg="org1msp"
			
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
				
				var bank2_scoreN  = genderIs+ageIs+maritalStatus+assetValue+depositPerMmonth+withdrawPerMonth+hasBeenSavingForYears+transactionPerMonthAverage+typeOfBussiness+savingAmount
				var final_ScoreN = (ClientRecord2.Bank1_Score+bank2_scoreN)/2

				assetInfo := AssetInfo{

					CNIC:									CNIC,
					F_name:									ClientRecord2.F_name, 
					L_name:									ClientRecord2.L_name, 
					Age: 									ClientRecord2.Age,
					Gender: 								ClientRecord2.Gender,
					Marital_status: 						ClientRecord2.Marital_status,
					Value_Of_Assets: 						value_of_assets,
					Deposit_Per_Month: 						deposit_per_month,
					Withdraw_Per_Month: 					withdraw_per_month,
					Has_Been_Saving_For_Years:				has_been_saving_for_years,
					Transaction_Per_Month_Average:  		transaction_per_month_average,
					Type_Of_Bussiness: 						type_of_bussiness,
					Saving_Amount:  						saving_amount,
					Bank1_Score: 							ClientRecord2.Bank1_Score, 
					Bank2_Score: 							bank2_scoreN, 
					Final_Score: 							final_ScoreN,
					MspOrg:									msporg,
				
				}
				assetInfoBytes, _ := json.Marshal(assetInfo)
				
				err3 := ctx.GetStub().PutState(CNIC+msporg, assetInfoBytes)
				
				if(err3 != nil){

					return "yes1"
				}


		
		
			return "no2"}

	return "no4"
	
}

	if err2 !=nil{ //add complete record

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

if (msporg=="org2msp"){
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
			
			err3 := ctx.GetStub().PutState(CNIC+msporg, assetInfoBytes)
			
			if(err3 != nil){

				return "yes"
			}



}
}
}












				






/* 


			if(msporg=="org2msp"){
				ClientRecord2 , err2 :=s.ReadAssetOrg(ctx,CNIC,msporg)
	
				if err2 !=nil{

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

	if (msporg=="org1msp"){
	 	bank1_scoreN = genderIs+ageIs+maritalStatus+assetValue+depositPerMmonth+withdrawPerMonth+hasBeenSavingForYears+transactionPerMonthAverage+typeOfBussiness+savingAmount
	}else if (msporg=="org2msp"){
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
				
				err := ctx.GetStub().PutState(CNIC+msporg, assetInfoBytes)
				
				if(err != nil){

					return "yes"
				}
	
				}
				if(ClientRecord2 != nil){

						return "no2";
		
				}

		}

		return "no3"






 */


		return "no3"

	

 
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
 

func (s *SmartContract) ClientPrivateRecord(ctx contractapi.TransactionContextInterface, CNICloan string,credit_type string,loan_amount float64 ,issue_loandate string,payment_plan float64,payment_date string,month float64,amount float64,jan1 float64, feb1 float64 ,mar1 float64,aprl1 float64,may1 float64 ,jun1 float64,jul1 float64,aug1 float64 , sep1 float64,oct1 float64 ,nov1 float64 , dec float64,orgloan string,msporg string) string {

	clientLoan, err4 := s.ReadAssetOrg(ctx, CNICloan, msporg)

	if  clientLoan.CNIC!=CNICloan{
		return "recordNotFound"}	

	if  clientLoan.CNIC==CNICloan{

		clientRecordfound , err1 :=s.GetOneClientRecord(ctx,CNICloan,orgloan)

		var instalment_per_month = paymantPlanF(loan_amount,payment_plan)
		
		

		if err1 !=nil{
		clientInfo := ClientPrivateData{
		CNIC:				CNICloan,				//cinc unique with cnic+orgMSPloan 
		First_Name:			clientLoan.F_name,
		Last_Name: 			clientLoan.L_name,
		Credit_Type:		credit_type,			
		Loan_Amount:		loan_amount,
		IssueLoanDate:		issue_loandate,
		PaymentPlan:		payment_plan,
		InstalmentPerMonth:	instalment_per_month,	
		PaymentDate:		payment_date,
		Credit_Score_Is:	clientLoan.Final_Score,	
		YearStructure:		PlanYear2{JAN1:0,FEB1:0,MAR1:0,ARL1:0,MAY1:0,JUNE1:0,JULY1:0,AUG1:0,SEP1:0,OCT1:0,NOV1:0,DEC1:0,JAN2:0,FEB2:0,MAR2:0,ARL2:0,MAY2:0,JUNE2:0,JULY2:0,AUG2:0,SEP2:0,OCT2:0,NOV2:0,DEC2:0,JAN3:0,FEB3:0,MAR3:0,ARL3:0,MAY3:0,JUNE3:0,JULY3:0,AUG3:0,SEP3:0,OCT3:0,NOV3:0,DEC3:0,JAN4:0,FEB4:0,MAR4:0,ARL4:0,MAY4:0,JUNE4:0,JULY4:0,AUG4:0,SEP4:0,OCT4:0,NOV4:0,DEC4:0,JAN5:0,FEB5:0,MAR5:0,ARL5:0,MAY5:0,JUNE5:0,JULY5:0,AUG5:0,SEP5:0,OCT5:0,NOV5:0,DEC5:0},
	}
	//JAN1:0,FEB1:0,MAR1:0,ARL1:0,MAY1:0,JUNE1:0,JULY1:0,AUG1:0,SEP1:0,OCT1:0,NOV1:0,DEC1:0,JAN2:0,FEB2:0,MAR2:0,ARL2:0,MAY2:0,JUNE2:0,JULY2:0,AUG2:0,SEP2:0,OCT2:0,NOV2:0,DEC2:0,JAN3:0,FEB3:0,MAR3:0,ARL3:0,MAY3:0,JUNE3:0,JULY3:0,AUG3:0,SEP3:0,OCT3:0,NOV3:0,DEC3:0,JAN4:0,FEB4:0,MAR4:0,ARL4:0,MAY4:0,JUNE4:0,JULY4:0,AUG4:0,SEP4:0,OCT4:0,NOV4:0,DEC4:0,JAN5:0,FEB5:0,MAR5:0,ARL5:0,MAY5:0,JUNE5:0,JULY5:0,AUG5:0,SEP5:0,OCT5:0,NOV5:0,DEC5:0
if(month == 0){
		clientInfoBytes, _ := json.Marshal(clientInfo)
		err :=ctx.GetStub().PutState(CNICloan+orgloan, clientInfoBytes)	
		if(err !=nil){
		return "yes1"}
		}
		return "errorIsHere"

	}

	if(month == 1){
	if clientRecordfound != nil{
		clientLoan.Final_Score  =  clientLoan.Final_Score+3
	//	var monthis float64 = month 
		var amountIs float64 = amount
		
		clientInfo := ClientPrivateData{
			CNIC:				CNICloan,				//cinc unique with cnic+orgMSPloan 
			First_Name:			clientLoan.F_name,
			Last_Name: 			clientLoan.L_name,
			Credit_Type:		credit_type,			
			Loan_Amount:		loan_amount,
			IssueLoanDate:		issue_loandate,
			PaymentPlan:		payment_plan,
			InstalmentPerMonth:	instalment_per_month,	
			PaymentDate:		payment_date,
			Credit_Score_Is:	clientLoan.Final_Score,	
			YearStructure:		PlanYear2{JAN1:amountIs,FEB1:0,MAR1:0,ARL1:0,MAY1:0,JUNE1:0,JULY1:0,AUG1:0,SEP1:0,OCT1:0,NOV1:0,DEC1:0,JAN2:0,FEB2:0,MAR2:0,ARL2:0,MAY2:0,JUNE2:0,JULY2:0,AUG2:0,SEP2:0,OCT2:0,NOV2:0,DEC2:0,JAN3:0,FEB3:0,MAR3:0,ARL3:0,MAY3:0,JUNE3:0,JULY3:0,AUG3:0,SEP3:0,OCT3:0,NOV3:0,DEC3:0,JAN4:0,FEB4:0,MAR4:0,ARL4:0,MAY4:0,JUNE4:0,JULY4:0,AUG4:0,SEP4:0,OCT4:0,NOV4:0,DEC4:0,JAN5:0,FEB5:0,MAR5:0,ARL5:0,MAY5:0,JUNE5:0,JULY5:0,AUG5:0,SEP5:0,OCT5:0,NOV5:0,DEC5:0},
		}
		//JAN1:0,FEB1:0,MAR1:0,ARL1:0,MAY1:0,JUNE1:0,JULY1:0,AUG1:0,SEP1:0,OCT1:0,NOV1:0,DEC1:0,JAN2:0,FEB2:0,MAR2:0,ARL2:0,MAY2:0,JUNE2:0,JULY2:0,AUG2:0,SEP2:0,OCT2:0,NOV2:0,DEC2:0,JAN3:0,FEB3:0,MAR3:0,ARL3:0,MAY3:0,JUNE3:0,JULY3:0,AUG3:0,SEP3:0,OCT3:0,NOV3:0,DEC3:0,JAN4:0,FEB4:0,MAR4:0,ARL4:0,MAY4:0,JUNE4:0,JULY4:0,AUG4:0,SEP4:0,OCT4:0,NOV4:0,DEC4:0,JAN5:0,FEB5:0,MAR5:0,ARL5:0,MAY5:0,JUNE5:0,JULY5:0,AUG5:0,SEP5:0,OCT5:0,NOV5:0,DEC5:0
	
			clientInfoBytes, _ := json.Marshal(clientInfo)
			err :=ctx.GetStub().PutState(CNICloan+orgloan, clientInfoBytes)	
			if(err !=nil){
			return "yes1"}
		
		
		
		return "errorIsHere"
	}else{
		return "alreadyIssued"
	}
}//end

//month=2
if(month == 2){
	if clientRecordfound != nil{
		
	//	var monthis float64 = month 
		
	clientLoan.Final_Score  =  clientLoan.Final_Score+6
		var jan11 float64= jan1
		
		var amountIs float64 = amount
		clientInfo := ClientPrivateData{
			CNIC:				CNICloan,				//cinc unique with cnic+orgMSPloan 
			First_Name:			clientLoan.F_name,
			Last_Name: 			clientLoan.L_name,
			Credit_Type:		credit_type,			
			Loan_Amount:		loan_amount,
			IssueLoanDate:		issue_loandate,
			PaymentPlan:		payment_plan,
			InstalmentPerMonth:	instalment_per_month,	
			PaymentDate:		payment_date,
			Credit_Score_Is:	clientLoan.Final_Score ,	
			YearStructure:		PlanYear2{JAN1:jan11,FEB1:amountIs,MAR1:0,ARL1:0,MAY1:0,JUNE1:0,JULY1:0,AUG1:0,SEP1:0,OCT1:0,NOV1:0,DEC1:0,JAN2:0,FEB2:0,MAR2:0,ARL2:0,MAY2:0,JUNE2:0,JULY2:0,AUG2:0,SEP2:0,OCT2:0,NOV2:0,DEC2:0,JAN3:0,FEB3:0,MAR3:0,ARL3:0,MAY3:0,JUNE3:0,JULY3:0,AUG3:0,SEP3:0,OCT3:0,NOV3:0,DEC3:0,JAN4:0,FEB4:0,MAR4:0,ARL4:0,MAY4:0,JUNE4:0,JULY4:0,AUG4:0,SEP4:0,OCT4:0,NOV4:0,DEC4:0,JAN5:0,FEB5:0,MAR5:0,ARL5:0,MAY5:0,JUNE5:0,JULY5:0,AUG5:0,SEP5:0,OCT5:0,NOV5:0,DEC5:0},
		}
		//JAN1:0,FEB1:0,MAR1:0,ARL1:0,MAY1:0,JUNE1:0,JULY1:0,AUG1:0,SEP1:0,OCT1:0,NOV1:0,DEC1:0,JAN2:0,FEB2:0,MAR2:0,ARL2:0,MAY2:0,JUNE2:0,JULY2:0,AUG2:0,SEP2:0,OCT2:0,NOV2:0,DEC2:0,JAN3:0,FEB3:0,MAR3:0,ARL3:0,MAY3:0,JUNE3:0,JULY3:0,AUG3:0,SEP3:0,OCT3:0,NOV3:0,DEC3:0,JAN4:0,FEB4:0,MAR4:0,ARL4:0,MAY4:0,JUNE4:0,JULY4:0,AUG4:0,SEP4:0,OCT4:0,NOV4:0,DEC4:0,JAN5:0,FEB5:0,MAR5:0,ARL5:0,MAY5:0,JUNE5:0,JULY5:0,AUG5:0,SEP5:0,OCT5:0,NOV5:0,DEC5:0
	
			clientInfoBytes, _ := json.Marshal(clientInfo)
			err :=ctx.GetStub().PutState(CNICloan+orgloan, clientInfoBytes)	
			if(err !=nil){
			return "yes1"}
		
		
		
		return "errorIsHere"
	}else{
		return "alreadyIssued"
	}
}//end


//month=3
if(month == 3){
	if clientRecordfound != nil{
		
	//	var monthis float64 = month 
	clientLoan.Final_Score =clientLoan.Final_Score+9
		var jan11 float64= jan1
		var feb11 float64 = feb1

		var amountIs float64 = amount
		clientInfo := ClientPrivateData{
			CNIC:				CNICloan,				//cinc unique with cnic+orgMSPloan 
			First_Name:			clientLoan.F_name,
			Last_Name: 			clientLoan.L_name,
			Credit_Type:		credit_type,			
			Loan_Amount:		loan_amount,
			IssueLoanDate:		issue_loandate,
			PaymentPlan:		payment_plan,
			InstalmentPerMonth:	instalment_per_month,	
			PaymentDate:		payment_date,
			Credit_Score_Is:	clientLoan.Final_Score,	
			YearStructure:		PlanYear2{JAN1:jan11,FEB1:feb11,MAR1:amountIs,ARL1:0,MAY1:0,JUNE1:0,JULY1:0,AUG1:0,SEP1:0,OCT1:0,NOV1:0,DEC1:0,JAN2:0,FEB2:0,MAR2:0,ARL2:0,MAY2:0,JUNE2:0,JULY2:0,AUG2:0,SEP2:0,OCT2:0,NOV2:0,DEC2:0,JAN3:0,FEB3:0,MAR3:0,ARL3:0,MAY3:0,JUNE3:0,JULY3:0,AUG3:0,SEP3:0,OCT3:0,NOV3:0,DEC3:0,JAN4:0,FEB4:0,MAR4:0,ARL4:0,MAY4:0,JUNE4:0,JULY4:0,AUG4:0,SEP4:0,OCT4:0,NOV4:0,DEC4:0,JAN5:0,FEB5:0,MAR5:0,ARL5:0,MAY5:0,JUNE5:0,JULY5:0,AUG5:0,SEP5:0,OCT5:0,NOV5:0,DEC5:0},
		}
		//JAN1:0,FEB1:0,MAR1:0,ARL1:0,MAY1:0,JUNE1:0,JULY1:0,AUG1:0,SEP1:0,OCT1:0,NOV1:0,DEC1:0,JAN2:0,FEB2:0,MAR2:0,ARL2:0,MAY2:0,JUNE2:0,JULY2:0,AUG2:0,SEP2:0,OCT2:0,NOV2:0,DEC2:0,JAN3:0,FEB3:0,MAR3:0,ARL3:0,MAY3:0,JUNE3:0,JULY3:0,AUG3:0,SEP3:0,OCT3:0,NOV3:0,DEC3:0,JAN4:0,FEB4:0,MAR4:0,ARL4:0,MAY4:0,JUNE4:0,JULY4:0,AUG4:0,SEP4:0,OCT4:0,NOV4:0,DEC4:0,JAN5:0,FEB5:0,MAR5:0,ARL5:0,MAY5:0,JUNE5:0,JULY5:0,AUG5:0,SEP5:0,OCT5:0,NOV5:0,DEC5:0
	
			clientInfoBytes, _ := json.Marshal(clientInfo)
			err :=ctx.GetStub().PutState(CNICloan+orgloan, clientInfoBytes)	
			if(err !=nil){
			return "yes1"}
		
		
		
		return "errorIsHere"
	}else{
		return "alreadyIssued"
	}
}//end


if(month == 4){
	if clientRecordfound != nil{
		
	//	var monthis float64 = month 
	clientLoan.Final_Score =clientLoan.Final_Score+12
		var jan11 float64= jan1
		var feb11 float64 = feb1
		var mar11 float64 = mar1

		var amountIs float64 = amount
		clientInfo := ClientPrivateData{
			CNIC:				CNICloan,				//cinc unique with cnic+orgMSPloan 
			First_Name:			clientLoan.F_name,
			Last_Name: 			clientLoan.L_name,
			Credit_Type:		credit_type,			
			Loan_Amount:		loan_amount,
			IssueLoanDate:		issue_loandate,
			PaymentPlan:		payment_plan,
			InstalmentPerMonth:	instalment_per_month,	
			PaymentDate:		payment_date,
			Credit_Score_Is:	clientLoan.Final_Score,	
			YearStructure:		PlanYear2{JAN1:jan11,FEB1:feb11,MAR1:mar11,ARL1:amountIs,MAY1:0,JUNE1:0,JULY1:0,AUG1:0,SEP1:0,OCT1:0,NOV1:0,DEC1:0,JAN2:0,FEB2:0,MAR2:0,ARL2:0,MAY2:0,JUNE2:0,JULY2:0,AUG2:0,SEP2:0,OCT2:0,NOV2:0,DEC2:0,JAN3:0,FEB3:0,MAR3:0,ARL3:0,MAY3:0,JUNE3:0,JULY3:0,AUG3:0,SEP3:0,OCT3:0,NOV3:0,DEC3:0,JAN4:0,FEB4:0,MAR4:0,ARL4:0,MAY4:0,JUNE4:0,JULY4:0,AUG4:0,SEP4:0,OCT4:0,NOV4:0,DEC4:0,JAN5:0,FEB5:0,MAR5:0,ARL5:0,MAY5:0,JUNE5:0,JULY5:0,AUG5:0,SEP5:0,OCT5:0,NOV5:0,DEC5:0},
		}
		//JAN1:0,FEB1:0,MAR1:0,ARL1:0,MAY1:0,JUNE1:0,JULY1:0,AUG1:0,SEP1:0,OCT1:0,NOV1:0,DEC1:0,JAN2:0,FEB2:0,MAR2:0,ARL2:0,MAY2:0,JUNE2:0,JULY2:0,AUG2:0,SEP2:0,OCT2:0,NOV2:0,DEC2:0,JAN3:0,FEB3:0,MAR3:0,ARL3:0,MAY3:0,JUNE3:0,JULY3:0,AUG3:0,SEP3:0,OCT3:0,NOV3:0,DEC3:0,JAN4:0,FEB4:0,MAR4:0,ARL4:0,MAY4:0,JUNE4:0,JULY4:0,AUG4:0,SEP4:0,OCT4:0,NOV4:0,DEC4:0,JAN5:0,FEB5:0,MAR5:0,ARL5:0,MAY5:0,JUNE5:0,JULY5:0,AUG5:0,SEP5:0,OCT5:0,NOV5:0,DEC5:0
	
			clientInfoBytes, _ := json.Marshal(clientInfo)
			err :=ctx.GetStub().PutState(CNICloan+orgloan, clientInfoBytes)	
			if(err !=nil){
			return "yes1"}
		
		
		
		return "errorIsHere"
	}else{
		return "alreadyIssued"
	}
}//end

//month  = 5
if(month == 5){
	if clientRecordfound != nil{
		clientLoan.Final_Score  =  clientLoan.Final_Score+15
		
		var jan11 float64= jan1
		var feb11 float64 = feb1
		var mar11 float64 = mar1
		var aprl11 float64 = aprl1
		

		var amountIs float64 = amount


		clientInfo := ClientPrivateData{
			CNIC:				CNICloan,				//cinc unique with cnic+orgMSPloan 
			First_Name:			clientLoan.F_name,
			Last_Name: 			clientLoan.L_name,
			Credit_Type:		credit_type,			
			Loan_Amount:		loan_amount,
			IssueLoanDate:		issue_loandate,
			PaymentPlan:		payment_plan,
			InstalmentPerMonth:	instalment_per_month,	
			PaymentDate:		payment_date,
			Credit_Score_Is:	clientLoan.Final_Score,	
			YearStructure:		PlanYear2{JAN1:jan11,FEB1:feb11,MAR1:mar11,ARL1:aprl11,MAY1:amountIs,JUNE1:0,JULY1:0,AUG1:0,SEP1:0,OCT1:0,NOV1:0,DEC1:0,JAN2:0,FEB2:0,MAR2:0,ARL2:0,MAY2:0,JUNE2:0,JULY2:0,AUG2:0,SEP2:0,OCT2:0,NOV2:0,DEC2:0,JAN3:0,FEB3:0,MAR3:0,ARL3:0,MAY3:0,JUNE3:0,JULY3:0,AUG3:0,SEP3:0,OCT3:0,NOV3:0,DEC3:0,JAN4:0,FEB4:0,MAR4:0,ARL4:0,MAY4:0,JUNE4:0,JULY4:0,AUG4:0,SEP4:0,OCT4:0,NOV4:0,DEC4:0,JAN5:0,FEB5:0,MAR5:0,ARL5:0,MAY5:0,JUNE5:0,JULY5:0,AUG5:0,SEP5:0,OCT5:0,NOV5:0,DEC5:0},
		}
		//JAN1:0,FEB1:0,MAR1:0,ARL1:0,MAY1:0,JUNE1:0,JULY1:0,AUG1:0,SEP1:0,OCT1:0,NOV1:0,DEC1:0,JAN2:0,FEB2:0,MAR2:0,ARL2:0,MAY2:0,JUNE2:0,JULY2:0,AUG2:0,SEP2:0,OCT2:0,NOV2:0,DEC2:0,JAN3:0,FEB3:0,MAR3:0,ARL3:0,MAY3:0,JUNE3:0,JULY3:0,AUG3:0,SEP3:0,OCT3:0,NOV3:0,DEC3:0,JAN4:0,FEB4:0,MAR4:0,ARL4:0,MAY4:0,JUNE4:0,JULY4:0,AUG4:0,SEP4:0,OCT4:0,NOV4:0,DEC4:0,JAN5:0,FEB5:0,MAR5:0,ARL5:0,MAY5:0,JUNE5:0,JULY5:0,AUG5:0,SEP5:0,OCT5:0,NOV5:0,DEC5:0
	
			clientInfoBytes, _ := json.Marshal(clientInfo)
			err :=ctx.GetStub().PutState(CNICloan+orgloan, clientInfoBytes)	
			if(err !=nil){
			return "yes1"}
		
		
		
		return "errorIsHere"
	}else{
		return "alreadyIssued"
	}
}//end




//month  = 6
if(month == 6){
	if clientRecordfound != nil{
		clientLoan.Final_Score  =  clientLoan.Final_Score+18
		
		var jan11 float64= jan1
		var feb11 float64 = feb1
		var mar11 float64 = mar1
		var aprl11 float64 = aprl1
		var may11 float64 = may1

		var amountIs float64 = amount


		clientInfo := ClientPrivateData{
			CNIC:				CNICloan,				//cinc unique with cnic+orgMSPloan 
			First_Name:			clientLoan.F_name,
			Last_Name: 			clientLoan.L_name,
			Credit_Type:		credit_type,			
			Loan_Amount:		loan_amount,
			IssueLoanDate:		issue_loandate,
			PaymentPlan:		payment_plan,
			InstalmentPerMonth:	instalment_per_month,	
			PaymentDate:		payment_date,
			Credit_Score_Is:	clientLoan.Final_Score,	
			YearStructure:		PlanYear2{JAN1:jan11,FEB1:feb11,MAR1:mar11,ARL1:aprl11,MAY1:may11,JUNE1:amountIs,JULY1:0,AUG1:0,SEP1:0,OCT1:0,NOV1:0,DEC1:0,JAN2:0,FEB2:0,MAR2:0,ARL2:0,MAY2:0,JUNE2:0,JULY2:0,AUG2:0,SEP2:0,OCT2:0,NOV2:0,DEC2:0,JAN3:0,FEB3:0,MAR3:0,ARL3:0,MAY3:0,JUNE3:0,JULY3:0,AUG3:0,SEP3:0,OCT3:0,NOV3:0,DEC3:0,JAN4:0,FEB4:0,MAR4:0,ARL4:0,MAY4:0,JUNE4:0,JULY4:0,AUG4:0,SEP4:0,OCT4:0,NOV4:0,DEC4:0,JAN5:0,FEB5:0,MAR5:0,ARL5:0,MAY5:0,JUNE5:0,JULY5:0,AUG5:0,SEP5:0,OCT5:0,NOV5:0,DEC5:0},
		}
	
			clientInfoBytes, _ := json.Marshal(clientInfo)
			err :=ctx.GetStub().PutState(CNICloan+orgloan, clientInfoBytes)	
			if(err !=nil){
			return "yes1"}
		
		
		
		return "errorIsHere"
	}else{
		return "alreadyIssued"
	}
}//end



//month  = 7
if(month == 7){
	if clientRecordfound != nil{
		clientLoan.Final_Score  =  clientLoan.Final_Score+21
		
		var jan11 float64= jan1
		var feb11 float64 = feb1
		var mar11 float64 = mar1
		var aprl11 float64 = aprl1
		var may11 float64 = may1
		var jun11 float64 = jun1 

		var amountIs float64 = amount


		clientInfo := ClientPrivateData{
			CNIC:				CNICloan,				//cinc unique with cnic+orgMSPloan 
			First_Name:			clientLoan.F_name,
			Last_Name: 			clientLoan.L_name,
			Credit_Type:		credit_type,			
			Loan_Amount:		loan_amount,
			IssueLoanDate:		issue_loandate,
			PaymentPlan:		payment_plan,
			InstalmentPerMonth:	instalment_per_month,	
			PaymentDate:		payment_date,
			Credit_Score_Is:	clientLoan.Final_Score,	
			YearStructure:		PlanYear2{JAN1:jan11,FEB1:feb11,MAR1:mar11,ARL1:aprl11,MAY1:may11,JUNE1:jun11,JULY1:amountIs,AUG1:0,SEP1:0,OCT1:0,NOV1:0,DEC1:0,JAN2:0,FEB2:0,MAR2:0,ARL2:0,MAY2:0,JUNE2:0,JULY2:0,AUG2:0,SEP2:0,OCT2:0,NOV2:0,DEC2:0,JAN3:0,FEB3:0,MAR3:0,ARL3:0,MAY3:0,JUNE3:0,JULY3:0,AUG3:0,SEP3:0,OCT3:0,NOV3:0,DEC3:0,JAN4:0,FEB4:0,MAR4:0,ARL4:0,MAY4:0,JUNE4:0,JULY4:0,AUG4:0,SEP4:0,OCT4:0,NOV4:0,DEC4:0,JAN5:0,FEB5:0,MAR5:0,ARL5:0,MAY5:0,JUNE5:0,JULY5:0,AUG5:0,SEP5:0,OCT5:0,NOV5:0,DEC5:0},
		}
	
			clientInfoBytes, _ := json.Marshal(clientInfo)
			err :=ctx.GetStub().PutState(CNICloan+orgloan, clientInfoBytes)	
			if(err !=nil){
			return "yes1"}
		
		
		
		return "errorIsHere"
	}else{
		return "alreadyIssued"
	}
}//end


//month  = 8
if(month == 8){
	if clientRecordfound != nil{
		clientLoan.Final_Score  =  clientLoan.Final_Score+24
		
		var jan11 float64= jan1
		var feb11 float64 = feb1
		var mar11 float64 = mar1
		var aprl11 float64 = aprl1
		var may11 float64 = may1
		var jun11 float64 = jun1
		var jul11 float64 = jul1 

		var amountIs float64 = amount


		clientInfo := ClientPrivateData{
			CNIC:				CNICloan,				//cinc unique with cnic+orgMSPloan 
			First_Name:			clientLoan.F_name,
			Last_Name: 			clientLoan.L_name,
			Credit_Type:		credit_type,			
			Loan_Amount:		loan_amount,
			IssueLoanDate:		issue_loandate,
			PaymentPlan:		payment_plan,
			InstalmentPerMonth:	instalment_per_month,	
			PaymentDate:		payment_date,
			Credit_Score_Is:	clientLoan.Final_Score,	
			YearStructure:		PlanYear2{JAN1:jan11,FEB1:feb11,MAR1:mar11,ARL1:aprl11,MAY1:may11,JUNE1:jun11,JULY1:jul11,AUG1:amountIs,SEP1:0,OCT1:0,NOV1:0,DEC1:0,JAN2:0,FEB2:0,MAR2:0,ARL2:0,MAY2:0,JUNE2:0,JULY2:0,AUG2:0,SEP2:0,OCT2:0,NOV2:0,DEC2:0,JAN3:0,FEB3:0,MAR3:0,ARL3:0,MAY3:0,JUNE3:0,JULY3:0,AUG3:0,SEP3:0,OCT3:0,NOV3:0,DEC3:0,JAN4:0,FEB4:0,MAR4:0,ARL4:0,MAY4:0,JUNE4:0,JULY4:0,AUG4:0,SEP4:0,OCT4:0,NOV4:0,DEC4:0,JAN5:0,FEB5:0,MAR5:0,ARL5:0,MAY5:0,JUNE5:0,JULY5:0,AUG5:0,SEP5:0,OCT5:0,NOV5:0,DEC5:0},
		}
	
			clientInfoBytes, _ := json.Marshal(clientInfo)
			err :=ctx.GetStub().PutState(CNICloan+orgloan, clientInfoBytes)	
			if(err !=nil){
			return "yes1"}
		
		
		
		return "errorIsHere"
	}else{
		return "alreadyIssued"
	}
}//end



//month  = 9
if(month == 9){
	if clientRecordfound != nil{
		clientLoan.Final_Score  =  clientLoan.Final_Score+27
		
		var jan11 float64= jan1
		var feb11 float64 = feb1
		var mar11 float64 = mar1
		var aprl11 float64 = aprl1
		var may11 float64 = may1
		var jun11 float64 = jun1
		var jul11 float64 = jul1 
		var aug11 float64 = aug1

		var amountIs float64 = amount


		clientInfo := ClientPrivateData{
			CNIC:				CNICloan,				//cinc unique with cnic+orgMSPloan 
			First_Name:			clientLoan.F_name,
			Last_Name: 			clientLoan.L_name,
			Credit_Type:		credit_type,			
			Loan_Amount:		loan_amount,
			IssueLoanDate:		issue_loandate,
			PaymentPlan:		payment_plan,
			InstalmentPerMonth:	instalment_per_month,	
			PaymentDate:		payment_date,
			Credit_Score_Is:	clientLoan.Final_Score,	
			YearStructure:		PlanYear2{JAN1:jan11,FEB1:feb11,MAR1:mar11,ARL1:aprl11,MAY1:may11,JUNE1:jun11,JULY1:jul11,AUG1:aug11,SEP1:amountIs,OCT1:0,NOV1:0,DEC1:0,JAN2:0,FEB2:0,MAR2:0,ARL2:0,MAY2:0,JUNE2:0,JULY2:0,AUG2:0,SEP2:0,OCT2:0,NOV2:0,DEC2:0,JAN3:0,FEB3:0,MAR3:0,ARL3:0,MAY3:0,JUNE3:0,JULY3:0,AUG3:0,SEP3:0,OCT3:0,NOV3:0,DEC3:0,JAN4:0,FEB4:0,MAR4:0,ARL4:0,MAY4:0,JUNE4:0,JULY4:0,AUG4:0,SEP4:0,OCT4:0,NOV4:0,DEC4:0,JAN5:0,FEB5:0,MAR5:0,ARL5:0,MAY5:0,JUNE5:0,JULY5:0,AUG5:0,SEP5:0,OCT5:0,NOV5:0,DEC5:0},
		}
	
			clientInfoBytes, _ := json.Marshal(clientInfo)
			err :=ctx.GetStub().PutState(CNICloan+orgloan, clientInfoBytes)	
			if(err !=nil){
			return "yes1"}
		
		
		
		return "errorIsHere"
	}else{
		return "alreadyIssued"
	}
}//end



//month  = 10
if(month == 10){
	if clientRecordfound != nil{
		clientLoan.Final_Score  =  clientLoan.Final_Score+30
		
		var jan11 float64= jan1
		var feb11 float64 = feb1
		var mar11 float64 = mar1
		var aprl11 float64 = aprl1
		var may11 float64 = may1
		var jun11 float64 = jun1
		var jul11 float64 = jul1 
		var aug11 float64 = aug1
		var sep11 float64 = sep1
		var amountIs float64 = amount


		clientInfo := ClientPrivateData{
			CNIC:				CNICloan,				//cinc unique with cnic+orgMSPloan 
			First_Name:			clientLoan.F_name,
			Last_Name: 			clientLoan.L_name,
			Credit_Type:		credit_type,			
			Loan_Amount:		loan_amount,
			IssueLoanDate:		issue_loandate,
			PaymentPlan:		payment_plan,
			InstalmentPerMonth:	instalment_per_month,	
			PaymentDate:		payment_date,
			Credit_Score_Is:	clientLoan.Final_Score,	
			YearStructure:		PlanYear2{JAN1:jan11,FEB1:feb11,MAR1:mar11,ARL1:aprl11,MAY1:may11,JUNE1:jun11,JULY1:jul11,AUG1:aug11,SEP1:sep11,OCT1:amountIs,NOV1:0,DEC1:0,JAN2:0,FEB2:0,MAR2:0,ARL2:0,MAY2:0,JUNE2:0,JULY2:0,AUG2:0,SEP2:0,OCT2:0,NOV2:0,DEC2:0,JAN3:0,FEB3:0,MAR3:0,ARL3:0,MAY3:0,JUNE3:0,JULY3:0,AUG3:0,SEP3:0,OCT3:0,NOV3:0,DEC3:0,JAN4:0,FEB4:0,MAR4:0,ARL4:0,MAY4:0,JUNE4:0,JULY4:0,AUG4:0,SEP4:0,OCT4:0,NOV4:0,DEC4:0,JAN5:0,FEB5:0,MAR5:0,ARL5:0,MAY5:0,JUNE5:0,JULY5:0,AUG5:0,SEP5:0,OCT5:0,NOV5:0,DEC5:0},
		}
	
			clientInfoBytes, _ := json.Marshal(clientInfo)
			err :=ctx.GetStub().PutState(CNICloan+orgloan, clientInfoBytes)	
			if(err !=nil){
			return "yes1"}
		
		
		
		return "errorIsHere"
	}else{
		return "alreadyIssued"
	}
}//end



//month  = 11
if(month == 11){
	if clientRecordfound != nil{
		clientLoan.Final_Score  =  clientLoan.Final_Score+33
		
		var jan11 float64= jan1
		var feb11 float64 = feb1
		var mar11 float64 = mar1
		var aprl11 float64 = aprl1
		var may11 float64 = may1
		var jun11 float64 = jun1
		var jul11 float64 = jul1 
		var aug11 float64 = aug1
		var sep11 float64 = sep1
		var oct11 float64 = oct1
		 


		var amountIs float64 = amount


		clientInfo := ClientPrivateData{
			CNIC:				CNICloan,				//cinc unique with cnic+orgMSPloan 
			First_Name:			clientLoan.F_name,
			Last_Name: 			clientLoan.L_name,
			Credit_Type:		credit_type,			
			Loan_Amount:		loan_amount,
			IssueLoanDate:		issue_loandate,
			PaymentPlan:		payment_plan,
			InstalmentPerMonth:	instalment_per_month,	
			PaymentDate:		payment_date,
			Credit_Score_Is:	clientLoan.Final_Score,	
			YearStructure:		PlanYear2{JAN1:jan11,FEB1:feb11,MAR1:mar11,ARL1:aprl11,MAY1:may11,JUNE1:jun11,JULY1:jul11,AUG1:aug11,SEP1:sep11,OCT1:oct11,NOV1:amountIs,DEC1:0,JAN2:0,FEB2:0,MAR2:0,ARL2:0,MAY2:0,JUNE2:0,JULY2:0,AUG2:0,SEP2:0,OCT2:0,NOV2:0,DEC2:0,JAN3:0,FEB3:0,MAR3:0,ARL3:0,MAY3:0,JUNE3:0,JULY3:0,AUG3:0,SEP3:0,OCT3:0,NOV3:0,DEC3:0,JAN4:0,FEB4:0,MAR4:0,ARL4:0,MAY4:0,JUNE4:0,JULY4:0,AUG4:0,SEP4:0,OCT4:0,NOV4:0,DEC4:0,JAN5:0,FEB5:0,MAR5:0,ARL5:0,MAY5:0,JUNE5:0,JULY5:0,AUG5:0,SEP5:0,OCT5:0,NOV5:0,DEC5:0},
		}
	
			clientInfoBytes, _ := json.Marshal(clientInfo)
			err :=ctx.GetStub().PutState(CNICloan+orgloan, clientInfoBytes)	
			if(err !=nil){
			return "yes1"}
		
		
		
		return "errorIsHere"
	}else{
		return "alreadyIssued"
	}
}//end




//month  = 12
if(month == 12){
	if clientRecordfound != nil{
		clientLoan.Final_Score  =  clientLoan.Final_Score+36
		
		var jan11 float64= jan1
		var feb11 float64 = feb1
		var mar11 float64 = mar1
		var aprl11 float64 = aprl1
		var may11 float64 = may1
		var jun11 float64 = jun1
		var jul11 float64 = jul1 
		var aug11 float64 = aug1
		var sep11 float64 = sep1
		var oct11 float64 = oct1
		 var nov11 float64 = nov1


		var amountIs float64 = amount


		clientInfo := ClientPrivateData{
			CNIC:				CNICloan,				//cinc unique with cnic+orgMSPloan 
			First_Name:			clientLoan.F_name,
			Last_Name: 			clientLoan.L_name,
			Credit_Type:		credit_type,			
			Loan_Amount:		loan_amount,
			IssueLoanDate:		issue_loandate,
			PaymentPlan:		payment_plan,
			InstalmentPerMonth:	instalment_per_month,	
			PaymentDate:		payment_date,
			Credit_Score_Is:	clientLoan.Final_Score,	
			YearStructure:		PlanYear2{JAN1:jan11,FEB1:feb11,MAR1:mar11,ARL1:aprl11,MAY1:may11,JUNE1:jun11,JULY1:jul11,AUG1:aug11,SEP1:sep11,OCT1:oct11,NOV1:nov11,DEC1:amountIs,JAN2:0,FEB2:0,MAR2:0,ARL2:0,MAY2:0,JUNE2:0,JULY2:0,AUG2:0,SEP2:0,OCT2:0,NOV2:0,DEC2:0,JAN3:0,FEB3:0,MAR3:0,ARL3:0,MAY3:0,JUNE3:0,JULY3:0,AUG3:0,SEP3:0,OCT3:0,NOV3:0,DEC3:0,JAN4:0,FEB4:0,MAR4:0,ARL4:0,MAY4:0,JUNE4:0,JULY4:0,AUG4:0,SEP4:0,OCT4:0,NOV4:0,DEC4:0,JAN5:0,FEB5:0,MAR5:0,ARL5:0,MAY5:0,JUNE5:0,JULY5:0,AUG5:0,SEP5:0,OCT5:0,NOV5:0,DEC5:0},
		}
	
			clientInfoBytes, _ := json.Marshal(clientInfo)
			err :=ctx.GetStub().PutState(CNICloan+orgloan, clientInfoBytes)	
			if(err !=nil){
			return "yes1"}
		
		
		
		return "errorIsHere"
	}else{
		return "alreadyIssued"
	}
}//end



}


 if clientLoan == nil{
return "recordNotFound"
}

if err4 != nil {
	return "errorInErr4"
}



 		return "yes"}






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

func (s *SmartContract) FindOneClientRecord(ctx contractapi.TransactionContextInterface, CNIC string ,orgloan string) (*ClientPrivateData, error) {

	record , err := s.GetOneClientRecord(ctx,CNIC,orgloan)

	if err != nil {
		return record, nil
	}

	if record != nil {
		return record,nil
	} 

	return record , nil
	

}



func (s *SmartContract) GetOneClientRecord(ctx contractapi.TransactionContextInterface, CNIC string ,orgloan string) (*ClientPrivateData, error) {
	
	var searchCNIC string = CNIC+orgloan
	assetJSON, err := ctx.GetStub().GetState(searchCNIC)
	if err != nil {
	  return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if assetJSON == nil {
	  return nil, fmt.Errorf("the asset %s does not exist", searchCNIC)
	}
  
	var asset ClientPrivateData
	err = json.Unmarshal(assetJSON, &asset)
	if err != nil {
	  return nil, err
	}
  
	return &asset, nil
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

	 /*  if((strings.Contains(asset.CNIC, "org1loan"))){
		assets = append(assets, &asset)
	} 
  	}
   */
	  assets = append(assets, &asset)
	}
	return assets, nil  }
  

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
func (s *SmartContract) ReadAsset(ctx contractapi.TransactionContextInterface, CNIC string,msporg string) (*AssetInfo, error) {

  showRecord1 , err1 :=s.ReadAssetOrg(ctx, CNIC, msporg)

  

  if showRecord1 == nil{
		
		showRecord2 , err2 :=s.ReadAssetOrg(ctx, CNIC, "org2msp" )
		
		
		if showRecord2 !=nil{

			return showRecord2, nil
		}
		if err2 != nil{
		 

			return showRecord1, nil
		}
		}else {	return showRecord1, nil}


		if err1 != nil{

			
			return nil, nil
		}
			
		return showRecord1, nil
}
	

	func (s *SmartContract) ReadAsset2(ctx contractapi.TransactionContextInterface, CNIC string,msporg string) (*AssetInfo, error) {

		
		showRecord1 , err1 :=s.ReadAssetOrg(ctx, CNIC, msporg)  

		if showRecord1 == nil{
			  
			  showRecord2 , err2 :=s.ReadAssetOrg(ctx, CNIC, "org1msp" )
			  
			  
			  if showRecord2 !=nil{
	  
				  return showRecord2, nil
			  }
			  if err2 != nil{
			   
	  
				  return showRecord1, nil
			  }
			  }else {	return showRecord1, nil}
	  
	  
			  if err1 != nil{
	  
				  
				  return nil, nil
			  }
				  
			  return showRecord1, nil
	  
		  }


  func (s *SmartContract) ReadAssetOrg(ctx contractapi.TransactionContextInterface, CNIC string ,msporg string) (*AssetInfo, error) {
	
	var searchCNIC string = CNIC+msporg
	assetJSON, err := ctx.GetStub().GetState(searchCNIC)
	if err != nil {
	  return nil, fmt.Errorf("failed to read from world state: %v", err)
	}

	if assetJSON == nil {
	  return nil, fmt.Errorf("the asset %s does not exist", searchCNIC)
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
		return `ID and Password is incorrect`
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