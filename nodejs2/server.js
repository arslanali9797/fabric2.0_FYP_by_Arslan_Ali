const express = require('express');
const app = express();

var bodyParser = require('body-parser'); 



app.set('views', __dirname + '/views/colorlib-regform-8'); // tell the directory of webpage .ejs extention b/c run JavaScript within html page

app.use( express.static( 'views/colorlib-regform-8' ) );
app.set('view engine','ejs');



var searchOneEntry = require('./searchOne.js');
var loginUser = require('./login.js');
var addRecored = require('./addRecord.js');


//body parser
app.use(bodyParser.urlencoded({ extended: false }))
//app.use(bodyParser.json())




app.get('/',(req,res)=>{
  res.render('login',{notlogin:""});
  
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
  "savingYears": req.body.savingAmount,
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
if (queryResult == 'no1'){
   
  res.render('Form',{result:"'The Record with this CNIC : "+clientINfo.CNIC +" is already exist in the Organization"});
  }

  if (queryResult == 'n2'){
    res.send("NO2 ok");
  }

  if (queryResult == 'no3'){
    res.send("NO3 OK")
  }
 
}

addRecored.addClient(clientINfo.CNIC,clientINfo.fname,clientINfo.lname,clientINfo.age,clientINfo.gender,
  clientINfo.martialStatue,clientINfo.valueOfAssets,clientINfo.depositPerMonth,clientINfo.withdrawPerMonth,
  clientINfo.savingYears,clientINfo.transactionPerMonth,clientINfo.typeOfBussiness,clientINfo.savingAmount,callback);


 
   

});





app.get('/find',(req,res)=>{

  
 res.render('Find');   
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
           // console.log(queryResult);
            var obj = JSON.parse(queryResult);
            
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
    
     res.render('FindOne',{data});
     //res.render('Find');
    }
    searchOneEntry.queryOne(callback,cnic);

    });


app.get('*',(req,res)=>{
res.send('invalid request .')

});










app.listen(5000);
console.log("Server is running on port 5000");
