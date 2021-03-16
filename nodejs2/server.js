const express = require('express');
const app = express();

var bodyParser = require('body-parser'); 



app.set('views', __dirname + '/views/colorlib-regform-8'); // tell the directory of webpage .ejs extention b/c run JavaScript within html page

app.use( express.static( 'views/colorlib-regform-8' ) );
app.set('view engine','ejs');



var searchOneEntry = require('./searchOne.js');
var loginUser = require('./login.js');
var addRecored = require('./addRecord.js');
var clientInfo = require('./addCLient.js');
var clientInfoByOne = require ('./searchOneCLientRecord.js');


//body parser
app.use(bodyParser.urlencoded({ extended: false }))
//app.use(bodyParser.json())




app.get('/',(req,res)=>{
  res.render('login',{notlogin:""});
  
  });

  

  app.get('/clientLoan',(req,res)=>{
    res.render('addClientLoan',{errorIsHere:""});
    
    });
  






  app.post('/loanInfo',(req,res)=>{
  
    console.log(req.body.cnic);
    console.log(req.body.loanamount);
    console.log(req.body.paymentplan)
    console.log(req.body.instalmentdate)

    console.log(req.body.loan)

    console.log("#######")

callback = function(queryResult0){

     
 if (queryResult0=="errorIsHere"){


  callback = function(queryResult1){
    var obj = JSON.parse(queryResult1);
     console.log(obj.CNIC);
     console.log(obj.first_name);
     console.log(obj.last_name);
     console.log(obj.credit_type);
     console.log(obj.loan_amount);
     console.log(obj.issue_loandate);
     console.log(obj.payment_plan);
     console.log(obj.instalment_per_month);
     console.log(obj.payment_date);
     console.log(obj.Credit_Score_Is);

     var loanIssuedInfo = {
      "CNIC":             obj.CNIC,
      "fname":            obj.first_name,
      "lname":            obj.last_name,
      "credit_Type":      obj.credit_type,
      "loan_Amount":      obj.loan_amount,
      "loan_Issue_Date":  obj.issue_loandate,
      "paymant_Plan":     obj.payment_plan,
      "Installment_Per_Month": obj.instalment_per_month,
      "paymant_Date":     obj.payment_date,
      "final_Score":      obj.Credit_Score_Is,}

      res.render('showClientRecord',{loanIssuedInfo,RecordAdded:"Recored is Added , Loan is issued to Client with this CNIC : "+ req.body.cnic});}
  clientInfoByOne.searchOneClient(req.body.cnic,callback);
     }
console.log("query Result is1 "+queryResult0);
    if(queryResult0==""){
      console.log("EMPTY QYERY Reuslt:")
      res.render('addClientLoan',{errorIsHere:" No prior Record Found with This CNIC : "+ req.body.cnic});}


    }   
clientInfo.loanClientInfo(req.body.cnic,req.body.loan,parseInt(req.body.loanamount), parseInt(req.body.paymentplan),req.body.instalmentdate,0,0,0,0,0,0,0,0,0,0,0,0,0,0,callback); 
});









  app.get('/clientLoanInfo',(req,res)=>{

    res.render('foundOneCLientRecord',{errorIsHere:""});
  });



  let CNICIs= ""
  let  fname =""
  let lastName = ""
  let credit_Type_Is = ""
  let loanAmount = ""
  let loanIssueDate = ""
  let paymantPlanIs = ""
  let instalmantPerMonth = ""
  let patmantDateIs = ""
  let creditScoreIs = ""
  let Jan1 = 0
  let feb1 = 0
  let mar1 = 0
  let aprl1 = 0 
  let may1 = 0 
  let jun1 = 0
  
  let july1 = 0
  let aug1 = 0
  let sep1 = 0
  let oct1 = 0 
  let nov1 = 0 
  let dec1 = 0   


  app.get('/foundCLientLoanRecord',(req,res)=>{
    console.log("Searching CNIC");
    console.log("Searching CNIC : "+ req.query.ClientSearchCnic +"  AND Type is : "+typeof (req.query.ClientSearchCnic));

    callback = function(queryResult1){
      console.log("yes")


      var obj = JSON.parse(queryResult1);
 
      console.log(obj);
      
      CNICIs = obj.CNIC
      fname = obj.first_name
      lastName = obj.last_name
      credit_Type_Is = obj.credit_type
      loanAmount   = obj.loan_amount
      loanIssueDate = obj.issue_loandate
      paymantPlanIs = obj.payment_plan
      instalmantPerMonth = obj.instalment_per_month
      patmantDateIs = obj.payment_date
      creditScoreIs = obj.Credit_Score_Is
      
      
      
      
      
      
      var months = obj.YearStructure;

      console.log("month")
      console.log(months);
      
      let monthObjectKeys= Object.values(months)
      console.log(monthObjectKeys[0])
      console.log(monthObjectKeys[1])
      
      jan1 = monthObjectKeys[0]
      feb1 = monthObjectKeys[1]
      mar1 = monthObjectKeys[2]
      aprl1 = monthObjectKeys[3]
      may1 = monthObjectKeys[4] 
      jun1 = monthObjectKeys[5]
      july1 = monthObjectKeys[6]
       aug1 = monthObjectKeys[7]
      sep1 = monthObjectKeys[8]
      oct1 = monthObjectKeys[9] 
       nov1 = monthObjectKeys[10] 
       dec1 = monthObjectKeys[11] 
      console.log("THIS IS JAN ON "+jan1 )



      

  
      let objLength = Object.keys(obj.YearStructure).length
      console.log("object length is "+ objLength)
        
         
         var loanIssuedInfo = {
           "CNIC":             obj.CNIC,
           "fname":            obj.first_name,
           "lname":            obj.last_name,
           "credit_Type":      obj.credit_type,
           "loan_Amount":      obj.loan_amount,
           "loan_Issue_Date":  obj.issue_loandate,
           "paymant_Plan":     obj.payment_plan,
           "Installment_Per_Month": obj.instalment_per_month,
           "paymant_Date":     obj.payment_date,
           "final_Score":      obj.Credit_Score_Is,
         }
        




console.log("length of loanIssuedInfo is : "+Object.keys(loanIssuedInfo).length)
res.render('foundCLientLoanRecord',{loanIssuedInfo,monthObjectKeys,objLength,foundRecord:"Recored Found"});

    }
    clientInfoByOne.searchOneClient(req.query.ClientSearchCnic,callback); 



  
  });  






  app.post('/InstalmantInfo',(req,res)=>{
  
    let paymantNo1 = req.body.paymantNo
    let paymantMonthIs = req.body.paymantMonth
    let payingAmountIs =  req.body.PayingAmount
  
    console.log("Payment no is"+paymantNo1)
    console.log("Paymanet date is "+paymantMonthIs)
    console.log("patmanet amount"+payingAmountIs)


 
    console.log("first month value is "+Jan1)
    callback = function(queryResult0){
  
      callback = function(queryResult1){
        var obj = JSON.parse(queryResult1);
   
        console.log(obj);
    
        CNICIs = obj.CNIC
        fname = obj.first_name
        lastName = obj.last_name
        credit_Type_Is = obj.credit_type
        loanAmount   = obj.loan_amount
        loanIssueDate = obj.issue_loandate
        paymantPlanIs = obj.payment_plan
        instalmantPerMonth = obj.instalment_per_month
        patmantDateIs = obj.payment_date
        creditScoreIs = obj.Credit_Score_Is
    
        
          //  console.log(Obj.YearStructure)
         
       var months = obj.YearStructure;
    
    console.log("month")
    console.log(months);
    
    let monthObjectKeys= Object.values(months)
    console.log("This is month on "+monthObjectKeys[0])
    console.log(monthObjectKeys[1])
    
  
    jan1 = monthObjectKeys[0]
    feb1 = monthObjectKeys[1]
    mar1 = monthObjectKeys[2]
    aprl1 = monthObjectKeys[3]
    may1 = monthObjectKeys[4] 
    jun1 = monthObjectKeys[5]
    july1 = monthObjectKeys[6]
   aug1 = monthObjectKeys[7]
  sep1 = monthObjectKeys[8]
  oct1 = monthObjectKeys[9] 
   nov1 = monthObjectKeys[10] 
   dec1 = monthObjectKeys[11] 
  
  
    console.log("THIS IS JAN ON "+jan1 )
       
     let objLength = Object.keys(obj.YearStructure).length
      console.log("object length is "+ objLength)
        
         
         var loanIssuedInfo = {
           "CNIC":             obj.CNIC,
           "fname":            obj.first_name,
           "lname":            obj.last_name,
           "credit_Type":      obj.credit_type,
           "loan_Amount":      obj.loan_amount,
           "loan_Issue_Date":  obj.issue_loandate,
           "paymant_Plan":     obj.payment_plan,
           "Installment_Per_Month": obj.instalment_per_month,
           "paymant_Date":     obj.payment_date,
           "final_Score":      obj.Credit_Score_Is,
         }
    
         console.log("length of loanIssuedInfo is : "+Object.keys(loanIssuedInfo).length)
         res.render('foundCLientLoanRecord',{loanIssuedInfo,monthObjectKeys,objLength,foundRecord:"Recored Found"});
    
        
    }
      clientInfoByOne.searchOneClient(CNICIs,callback); 
  
    
    }
    clientInfo.loanClientInfo(CNICIs, credit_Type_Is, parseInt(loanAmount), parseInt(paymantPlanIs),patmantDateIs,parseInt(paymantNo1),parseInt(payingAmountIs),parseInt(jan1),parseInt(feb1),parseInt(mar1),parseInt(aprl1),parseInt(may1),parseInt(jun1),parseInt(july1),parseInt(aug1),parseInt(sep1),parseInt(oct1),parseInt(nov1),parseInt(dec1),callback);
  

  });


app.post('/Login_Verified',(req,res)=>{

  callback = function(queryResult) {
    console.log("result query is :"+queryResult);
    if(queryResult=="true"){
      res.render('Home');
    }
    else{
      res.render('login',{notlogin:"Please make sure the ID or Password is correct"});

    }
  }
  
  console.log("id is :"+req.body.id)
  console.log("password is :"+req.body.password);
  loginUser.isLogin(callback,req.body.id,req.body.password);
  
  });


  app.get('/home',(req,res)=>{

    res.render('Home')
  });

app.get('/add',(req,res)=>{
  res.render('Form',{result:""})

})



app.post('/clientInfo',(req,res)=>{


console.log("THIS IS DATA = "+JSON.stringify(req.body));


var clientINfo ={
  "CNIC": req.body.cnic,
  "fname": req.body.fname,
  "lname":req.body.lname,
  "age" :req.body.age,
  "gender": req.body.gender,
  "martialStatue" :req.body.status,
  "valueOfAssets":req.body.assetsValue,
  "depositPerMonth":req.body.deposit,  
  "withdrawPerMonth" :req.body.withdrawPerMonth,
  "savingYears": req.body.savingForYear,
  "transactionPerMonth": req.body.transactionPerMonth,
  "typeOfBussiness":req.body.JobType,
  "savingAmount": req.body.savingAmount,}
  
 console.log("CNIC is "+clientINfo.CNIC+" type is "+typeof clientINfo.CNIC);
 console.log("Fname is "+clientINfo.fname+" type is "+typeof clientINfo.fname);
 console.log("lname is "+clientINfo.lname+" type is "+typeof clientINfo.lname);
 console.log("age is "+clientINfo.age+" type is "+typeof  parseInt(clientINfo.age));
 console.log("Gender is "+clientINfo.gender+" type is "+typeof clientINfo.gender);
 console.log("status is "+clientINfo.martialStatue+" type is "+typeof clientINfo.martialStatue);
 console.log("assets value is "+clientINfo.valueOfAssets+" type is "+typeof parseInt(clientINfo.valueOfAssets));
 console.log("depositPerMonth value is "+clientINfo.depositPerMonth+" type is "+typeof parseInt(clientINfo.depositPerMonth));
 console.log("withdrawPerMonth is "+clientINfo.withdrawPerMonth+" type is "+typeof parseInt(clientINfo.withdrawPerMonth));
 console.log("savingYears is "+clientINfo.savingYears+" type is "+typeof parseInt(clientINfo.savingYears));
 console.log("transactionPerMonth is "+clientINfo.transactionPerMonth+" type is "+typeof parseInt(clientINfo.transactionPerMonth));
 console.log("typeOfBussiness is "+clientINfo.typeOfBussiness+" type is "+typeof clientINfo.typeOfBussiness);
 console.log("savingAmount is "+clientINfo.savingAmount+" type is "+typeof parseInt(clientINfo.savingAmount) );
 
 var ffinal=0;

 callback = function(queryResult) { 
 if (queryResult == '200') {
    callback = function(queryResult) { 
     var obj = JSON.parse(queryResult);
    
     clientINfo.finalScoreIs =obj.finale_score

     console.log("###############Final Score is ##########"+clientINfo.finalScoreIs);
     res.render('showResult',{clientINfo,success:'Record is successfuly added'});


   }
   
   searchOneEntry.queryOne(callback,clientINfo.CNIC);





  }



if (queryResult == 'yes1'){
   
  console.log("Record is Updated");
   }
   

if (queryResult == 'no2'){


  callback = function(queryResult) { 
    var obj = JSON.parse(queryResult);
   
    clientINfo.finalScoreIs =obj.finale_score

    console.log("###############Final Score is ##########"+clientINfo.finalScoreIs);
    res.render('showResult',{clientINfo,success:'Record is successfuly added'});

  }
  
  searchOneEntry.queryOne(callback,clientINfo.CNIC);

}

  

 
if (queryResult == 'no4'){
   
  console.log("Cnic is not matched");
   }
    
 if (queryResult == 'no1'){
  res.render('Form',{result:"'The Record with this CNIC : "+clientINfo.CNIC +" is already exist in the Organization"});
  }


 if (queryResult == 'no3'){
      callback = function(queryResult) { 
        var obj = JSON.parse(queryResult);
       
        clientINfo.finalScoreIs =obj.finale_score
  
        console.log("###############Final Score is ##########"+clientINfo.finalScoreIs);
        res.render('showResult',{clientINfo,success:'Record is successfuly added'});
  
      }
      
      searchOneEntry.queryOne(callback,clientINfo.CNIC);
    
    }
 
}

addRecored.addClient(clientINfo.CNIC,clientINfo.fname,clientINfo.lname,clientINfo.age,clientINfo.gender,
  clientINfo.martialStatue,clientINfo.valueOfAssets,clientINfo.depositPerMonth,clientINfo.withdrawPerMonth,
  clientINfo.savingYears,clientINfo.transactionPerMonth,clientINfo.typeOfBussiness,clientINfo.savingAmount,callback);


 
   

});





app.get('/find',(req,res)=>{

  
 res.render('Find',{notFound:""});   
});





 app.get('/findOne',(req,res)=>{

  var cnic = req.query.search;
  console.log("CNIC IS : "+cnic +" AND type is :"+typeof cnic);
  

console.log("search with this cnic :"+req.query.search)



  var Fname='';
  var Lname='';
  var CNIC='';
  var bank1=0;
  var bank2=0;
  var final=0;


           callback = function(queryResult) { 
            console.log("1")
            console.log("Type of Query Result:"+typeof queryResult)

            if(queryResult==""){
              console.log("EMPTY QYERY Reuslt:")
              res.render('Find',{notFound:"Record Not Found"});


            }else{
            var obj = JSON.parse(queryResult);
            console.log("2")  
            console.log("Type is "+typeof obj);

              if (Object.keys(obj).length === 0) {
                console.log("empty")              } 
                else {
                  console.log("not empty")
              }

            console.log("Query Result is : "+obj);
            Fname=obj.fname
            Lname=obj.lname
            CNIC=obj.CNIC
            bank1=obj.bank1_score
            bank2=obj.bank2_score
            final=obj.finale_score
    var data ={
        "name": Fname+" "+Lname,
        "CNIC":CNIC,
        "bank1":bank1,
        "bank2":bank2,
        "final":final
       }

    console.log("found Result");
    //console.log("this is data ");
    
     res.render('FindOne',{data,notFound:""});
     //res.render('Find');

    }
  }
    searchOneEntry.queryOne(callback,cnic);

    });


app.get('*',(req,res)=>{
res.send('invalid request .')

});










app.listen(5000);
console.log("Server is running on port 5000");
