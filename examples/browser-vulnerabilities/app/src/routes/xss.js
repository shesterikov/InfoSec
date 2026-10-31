'use strict';

const router = require('express').Router();
const { safe } = require('../config');
const layout = require('../views/layout');
const escapeHtml = require('../utils/escape-html');

const comments = [];

router.get('/search',(req,res)=>{
  const q=String(req.query.q||'');
  res.send(layout('Reflected XSS',`<form><input name="q" value="${safe?escapeHtml(q):q}"><button>Искать</button></form><p>Результат: ${safe?escapeHtml(q):q}</p>`));
});
router.get('/comments',(req,res)=>res.send(layout('Stored XSS',`<form method="post"><textarea name="comment" required></textarea><br><button>Добавить</button></form><ul>${comments.map(c=>`<li>${safe?escapeHtml(c):c}</li>`).join('')}</ul>`)));
router.post('/comments',(req,res)=>{comments.push(String(req.body.comment||'').slice(0,2000));res.redirect('/comments');});
router.get('/dom',(req,res)=>{
  const js=safe ? 'document.getElementById("name").textContent = new URLSearchParams(location.search).get("name") || "";' : 'document.getElementById("name").innerHTML = new URLSearchParams(location.search).get("name") || "";';
  res.send(layout('DOM-based XSS',`<p>Привет, <span id="name"></span></p><script ${safe?`nonce="${res.locals.nonce}"`:''}>${js}</script>`));
});

module.exports = router;
