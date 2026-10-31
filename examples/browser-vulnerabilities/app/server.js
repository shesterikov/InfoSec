'use strict';

const app = require('./src/app');
const { safe, port } = require('./src/config');

app.listen(port, '0.0.0.0', () => {
  console.log(`Lab target ${safe ? 'protected' : 'vulnerable'} listening on ${port}`);
});
