'use strict';

module.exports = {
  safe: process.env.MODE === 'protected',
  port: Number(process.env.PORT || 3000),
};
