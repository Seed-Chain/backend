const express = require('express');
const farmerRouter = express.Router();
const request = require('request');
const Seed = require('../models/seedSchema');

farmerRouter.get('/blockchain',async (req,res)=>{

request({
  url : ' https://39E4CBC73AA5428CB4CEBC030B46BB8F.blockchain.ocp.oraclecloud.com:443/restproxy1/bcsgw/rest/v1/transaction/getTxID?channel=default',
  headers: {
    'Authorization': 'Basic YW50b3ByaW5jZTAwMUBnbWFpbC5jb206Y291cG9uVGFyZ2V0XzAx'
  },
  rejectUnauthorized : false
}, function(err,res) {
  if(err)
  {
    console.log(err);
  }
  else {
    console.log(JSON.parse(res.body));
  }
});

});



farmerRouter.post('/queryByLot', (req,res)=>{

  console.log('as');
  res.send(req.body.LotNumber);


});


farmerRouter.get('/queryByTag',async (req,res)=>{


});

farmerRouter.get('/queryBySeason',async (req,res)=>{


});

farmerRouter.get('/queryBy',async (req,res)=>{


});




module.exports = farmerRouter;
