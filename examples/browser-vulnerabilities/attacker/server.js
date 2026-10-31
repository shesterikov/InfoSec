'use strict';

const app = require('./src/app');

app.listen(3001, '0.0.0.0', () => {
  console.log('Local attacker site on 3001');
});
