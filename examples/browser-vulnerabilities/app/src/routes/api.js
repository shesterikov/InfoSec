'use strict';

const router = require('express').Router();
const { safe } = require('../config');

router.get('/api/private',(req,res)=>{
  const origin=req.get('origin');
  if (!safe && origin) {
    res.setHeader('Access-Control-Allow-Origin',origin);
    res.setHeader('Access-Control-Allow-Credentials','true');
    res.setHeader('Vary','Origin');
  }
  if (safe && origin === 'http://localhost:3002') {
    res.setHeader('Access-Control-Allow-Origin',origin);
    res.setHeader('Access-Control-Allow-Credentials','true');
    res.setHeader('Vary','Origin');
  }
  res.json({demo:'LOCAL TRAINING DATA ONLY',email:req.session.email,mode:safe?'protected':'vulnerable'});
});

module.exports = router;
