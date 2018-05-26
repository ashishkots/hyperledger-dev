var express = require("express");
var bodyParser = require('body-parser');
var app = express();
var urlencodedParser = bodyParser.urlencoded({ extended: false });
//var KYC = require("./contract_server.js");
var router = express.Router();
var query_contract = require("./query.js");
var invoke = require("./invoke.js");
    

		app.post('/createUser', urlencodedParser, function (req, res) {
                  if (!req.body){
                    res.send({error: "no params given"});
                    console.log("error: no params given");
                  } 
                  else{
                    console.log(req.body);
                    invoke(req.body.func, req.body.key, req.body.docid, req.body.timestamp, req.body.dochash, req.body.owner);
                    //func, docid, timestamp, dochash, owner
                  
                    
                   
                  }
                  
                });


                app.get('/getUser/:key', function(req, res){
                    console.log(req.params.key);
                    query_contract("getUser",req.params.key);

                });

               

                

                // Tell express to use this router with /api before.
                // You can put just '/' if you don't want any sub path before routes.

                

                // Listen to this Port

                app.listen(3600,function(){
                  console.log("Live at Port 3600");
                });


