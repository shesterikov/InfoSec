'use strict';

const router = require('express').Router();
const page = require('../views/page');
const target = require('../target');

router.get('/clickjack',(req,res)=>res.send(page('Clickjacking: прозрачный iframe',`<p>Попробуйте нажать на видимую область кнопки. В защищённой версии браузер заблокирует iframe.</p><div style="position:relative;width:600px;height:240px"><button style="position:absolute;top:190px;left:14px">Получить приз</button><iframe title="Локальная учебная цель" src="http://localhost:${target(req)}/frame" style="position:absolute;top:0;left:0;opacity:.08;border:1px dashed red"></iframe></div><p>Для наглядности iframe оставлен слегка видимым; откройте DevTools.</p>`)));

module.exports = router;
