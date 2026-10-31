'use strict';

const express = require('express');

const app = express();

app.use(require('./routes/home'));
app.use(require('./routes/csrf'));
app.use(require('./routes/clickjack'));
app.use(require('./routes/cors'));

module.exports = app;
