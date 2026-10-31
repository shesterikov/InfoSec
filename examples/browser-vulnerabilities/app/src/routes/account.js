'use strict';

const router = require('express').Router();
const { safe } = require('../config');
const layout = require('../views/layout');
const escapeHtml = require('../utils/escape-html');
const crypto = require('crypto');

router.get('/account',(req,res)=>res.send(layout('CSRF: профиль',`<p>Текущий email: <strong>${escapeHtml(req.session.email)}</strong></p><form action="/account/email" method="post"><input type="email" name="email" required value="${escapeHtml(req.session.email)}">${safe?`<input type="hidden" name="csrf" value="${req.session.csrf}">`:''}<button>Изменить email</button></form>`)));
router.post('/account/email',(req,res)=>{
  if (safe) {
    const origin=req.get('origin');
    const expected=`${req.protocol}://${req.get('host')}`;
    if (origin && origin!==expected) return res.status(403).send('Blocked: Origin mismatch');
    if (typeof req.body.csrf!=='string' || !crypto.timingSafeEqual(Buffer.from(req.body.csrf.padEnd(48,' ').slice(0,48)),Buffer.from(req.session.csrf))) return res.status(403).send('Blocked: invalid CSRF token');
  }
  req.session.email=String(req.body.email||'').slice(0,120);
  res.redirect('/account');
});

module.exports = router;
