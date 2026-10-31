'use strict';

const router = require('express').Router();
const { safe } = require('../config');
const layout = require('../views/layout');

router.get('/frame',(req,res)=>res.send(layout('Clickjacking target',`<p>Демонстрационная кнопка:</p><button id="confirm">Подтвердить действие</button><p id="status"></p><script ${safe?`nonce="${res.locals.nonce}"`:''}>document.getElementById('confirm').addEventListener('click',()=>document.getElementById('status').textContent='Кнопка нажата');</script>`)));

module.exports = router;
