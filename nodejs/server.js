const express = require('express');
const app = express();

app.set('views', __dirname + '/views/colorlib-regform-8'); // tell the directory of webpage .ejs extention b/c run JavaScript within html page
app.use( express.static( 'views/colorlib-regform-8' ) );

app.set('view engine','ejs');

var query = require('./query.js');


app.get('/',(req,res)=>{
res.render('Home');
});


app.get('/find',(req,res)=>{
    console.log('##################### 1 ##################');
    callback = function(queryResult) { // this check user account is still active or not
        console.log(queryResult);
        var obj = JSON.parse(queryResult);

        console.log("typee of ",typeof obj);
        
        console.log('#############');
        console.log("object length ",obj.length);
        console.log('###########################',obj[0],'#########################');
        console.log('###########################',obj[1],'#########################');
      }
      query.queryAll(callback);
 


//     query.queryAll(function(err,result)
//     {
//     if(err){
//         console.log(err)
//       }else{
//          console.log('############### this is result ###########');
//          console.log(result);
//          console(typeof result);
//          console.log(result[1]);
//       }
    
      

// });

 res.render('Find');
    });







app.listen(4000);
console.log("Server is running on port 4000");
