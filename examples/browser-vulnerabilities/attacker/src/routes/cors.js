'use strict';

const router = require('express').Router();
const page = require('../views/page');
const target = require('../target');

router.get('/cors',(req,res)=>res.send(page('CORS: чтение ответа API',`<button id="go">Выполнить fetch</button><pre id="result"></pre><script>document.getElementById('go').onclick=async()=>{try{const r=await fetch('http://localhost:${target(req)}/api/private',{credentials:'include'});document.getElementById('result').textContent=JSON.stringify(await r.json(),null,2)}catch(e){document.getElementById('result').textContent='Заблокировано браузером: '+e.message}}</script>`)));

module.exports = router;
