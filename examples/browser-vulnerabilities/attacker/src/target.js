'use strict';

module.exports = req => req.query.target === '3002' ? '3002' : '3000';
