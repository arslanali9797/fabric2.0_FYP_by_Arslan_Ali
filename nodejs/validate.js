module.exports.add= function(cnic) {

    var CNICformat = /^[0-9]{5}[0-9]{7}[0-9]{1}$/;
     if(cnic.match(CNICformat))
    {   
        return 1
    } 
     else{
         throw Error ("CNIC No: is not Valid")
     }
   
     

}
