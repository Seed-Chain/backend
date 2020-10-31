// REST API for website request handling

const express = require('express');
const { mongoURI } = require('./config/mongo');
const app = express();
const cookieParser = require('cookie-parser');
const mongoose = require('mongoose');
const path = require('path');
const cors = require('cors')
app.use(express.json()); // To read post requests
app.use(express.urlencoded({ extended: true }));
app.use(cookieParser()); // To parse jwt



/* Routes handling */


/* Seed Processing Plant */
const sppRouter = require('./routes/spp');
app.use('/spp',sppRouter);

/* Seed Processing Agency */
const spaRouter = require('./routes/spa');
app.use('/spa',sppRouter);

/* Seed Testing Lab */
const stlRouter = require('./routes/stl');
app.use('/stl',stlRouter);

/* Seed Certification Agency */
const scaRouter = require('./routes/sca');
app.use('/sca',scaRouter);

/* Seed grower/farmer */
const farmerRouter = require('./routes/farmer');
app.use('/farm',farmerRouter);

app.set('view engine','ejs');                   //Templating engine
app.use(express.static('views'));
app.get('/home',(req,res)=>{

     res.render('home.ejs');
});

//React server render
// app.use(express.static(path.join(__dirname, 'seedfront/build')));
// app.get('/*', function (req, res) {
//    res.sendFile(path.join(__dirname, 'seedfront/build', 'index.html'));
// });


app.use(cors());


/* Mongodb Configuration */
mongoose.connect(mongoURI,{useNewUrlParser : true,useUnifiedTopology: true},()=>{
    console.log('successfully connected to database');
});


/* Server config */
module.exports = app.listen(5000,()=>{
    console.log('express server started');
});
