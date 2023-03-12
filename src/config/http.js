const cookieParser = require('cookie-parser');
const express = require('express');

const httpApp = express();
const { getLogger } = require('log4js');
const ErrorMiddleware = require('../middleware/ErrorMiddleware');
const routes = require('./routes');

const logger = getLogger('default');

function init(port = 3000) {
   httpApp.use(express.json({ limit: '1mb' }));
   httpApp.use(express.urlencoded({ extended: true, limit: '1mb' }));
   httpApp.use(cookieParser());

   httpApp.use(routes);
   httpApp.use(ErrorMiddleware);

   return new Promise((resolve) => {
      httpApp.listen(port, () => resolve());
   });
}

module.exports = { httpApp, init };
