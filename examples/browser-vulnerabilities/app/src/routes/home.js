'use strict';

const router = require('express').Router();
const { safe } = require('../config');
const layout = require('../views/layout');

router.get('/',(req,res)=>res.send(layout('Browser Security Lab',`<p class="note">Режим: <strong>${safe?'PROTECTED':'VULNERABLE'}</strong>. Данные хранятся в памяти контейнера.</p><p>Пройдите сценарии из README. Приложение использует только фиктивные данные.</p>`)));

module.exports = router;
