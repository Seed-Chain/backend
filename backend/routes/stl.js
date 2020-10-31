const express = require('express');
const Seed = require('../models/seedSchema');
const stlRouter = express.Router();


//Certifies data .
stlRouter.post('/testify', async (req,res)=>{

  Seed.findOneAndUpdate({ "SampleSecretCode" : req.body.SampleSecretCode },
    {
      SamplePassed : "Yes",
      SampleTestDate : "25/10/2020:12:30"
    }).then((data)=> {
      res.status(201).json({ msg : 'Seed :' + req.body.lotNumber+ ' certified successfully' } );
    })
    .catch((err)=> {
      console.log(err);
      res.status(422).json({ msg : err });
    })


});


stlRouter.post('/view', (req,res)=>{
  Seed.find({})
  .then((data)=>{
    res.send(data);
  })
  .catch((err)=>{
    res.send(err);
  })
});
module.exports = stlRouter;
