var express = require("express");
var bodyParser = require('body-parser');
var app = express();
var urlencodedParser = bodyParser.urlencoded({ extended: false });
//var KYC = require("./contract_server.js");
var router = express.Router();
var query_contract = require("./query.js");
var invoke = require("./invoke.js");
    

		app.post('/invoke', urlencodedParser, function (req, res) {
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

        app.post('/changeAccount', urlencodedParser, function (req, res) {
                  if (!req.body){
                    res.send({error: "no params given"});
                    console.log("error: no params given");
                  } 
                  else{
                    console.log(req.body);
                    unlockAccount();

                    //bytes32 newUserHash, bytes32 pubKey
                    var getData = contract.changeAccount.getData(req.body.newUserHash, req.body.pubKey);
                    
                    var gasNeeded = contract.changeAccount.estimateGas(req.body.newUserHash, req.body.pubKey
                        ,{ from: walletAddress });


                    var trxAddr = web3.eth.sendTransaction({to:contractAddress, from:walletAddress, data: getData, 
                        gas: gasNeeded, gasPrice: "180000000000"},function(error,result){
                            if(error){
                                res.json({error: error});
                            }
                            else{
                                res.json({result: result});
                            }
                        });
                  }
                  
                });



                app.get('/query', function(req, res){
                    query_contract();

                });

                app.get('/balanceOf/:userAddress', function(req, res){

                    contract.balanceOf(req.params.userAddress,function(error,result){
                        if(error){
                                res.json({error: error});
                            }
                            else{
                                res.json({result: result});
                            }
                    });

                });


                

                // Tell express to use this router with /api before.
                // You can put just '/' if you don't want any sub path before routes.

                

                // Listen to this Port

                app.listen(3600,function(){
                  console.log("Live at Port 3600");
                });

function unlockAccount(){
        web3.personal.unlockAccount(walletAddress,"vasainc..");
    }

