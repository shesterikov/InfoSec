'use strict';

const helmet = require('helmet');
const crypto = require('crypto');
const { safe } = require('../config');

module.exports = (req, res, next) => {
  if (!safe) {
    res.setHeader(
      'Content-Security-Policy',
      "default-src * data: 'unsafe-inline' 'unsafe-eval'; frame-ancestors *",
    );
    return next();
  }

  res.locals.nonce = crypto.randomBytes(16).toString('base64');
  helmet({
    contentSecurityPolicy: {
      useDefaults: false,
      directives: {
        'default-src': ["'self'"],
        'script-src': ["'self'", `'nonce-${res.locals.nonce}'`],
        'style-src': ["'self'", "'unsafe-inline'"],
        'img-src': ["'self'"],
        'object-src': ["'none'"],
        'base-uri': ["'self'"],
        'frame-ancestors': ["'none'"],
      },
    },
    frameguard: { action: 'deny' },
    referrerPolicy: { policy: 'strict-origin-when-cross-origin' },
  })(req, res, next);
};
