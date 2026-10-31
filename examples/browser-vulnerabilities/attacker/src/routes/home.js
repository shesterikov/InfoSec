'use strict';

const router = require('express').Router();
const page = require('../views/page');

router.get('/',(req,res)=>res.send(page('Локальный атакующий сайт','<p><a href="/csrf?target=3000">CSRF → vulnerable</a> | <a href="/csrf?target=3002">CSRF → protected</a></p><p><a href="/clickjack?target=3000">Clickjacking → vulnerable</a> | <a href="/clickjack?target=3002">Clickjacking → protected</a></p><p><a href="/cors?target=3000">CORS → vulnerable</a> | <a href="/cors?target=3002">CORS → protected</a></p>')));

module.exports = router;
