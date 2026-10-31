'use strict';

const express = require('express');
const security = require('./middleware/security');
const session = require('./middleware/session');

const app = express();

app.disable('x-powered-by');
app.use(express.urlencoded({ extended: false }));
app.use(express.json());
app.use(security);
app.use(session);

app.use(require('./routes/home'));
app.use(require('./routes/xss'));
app.use(require('./routes/account'));
app.use(require('./routes/frame'));
app.use(require('./routes/api'));

module.exports = app;
