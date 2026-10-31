'use strict';

const crypto = require('crypto');
const { safe } = require('../config');

const cookieName = safe ? 'lab_safe_session' : 'lab_vuln_session';
const sessions = new Map();

function parseCookies(header) {
  return Object.fromEntries(
    (header || '').split(';').map(value => value.trim()).filter(Boolean).map(value => {
      const separator = value.indexOf('=');
      return [value.slice(0, separator), value.slice(separator + 1)];
    }),
  );
}

module.exports = (req, res, next) => {
  let sid = parseCookies(req.headers.cookie)[cookieName];
  if (!sid || !sessions.has(sid)) {
    sid = crypto.randomBytes(24).toString('hex');
    sessions.set(sid, {
      email: 'student@example.test',
      csrf: crypto.randomBytes(24).toString('hex'),
    });
    res.cookie(cookieName, sid, {
      httpOnly: true,
      sameSite: 'lax',
      secure: false,
      path: '/',
    });
  }
  req.session = sessions.get(sid);
  next();
};
