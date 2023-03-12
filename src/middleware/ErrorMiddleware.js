const { getLogger } = require('log4js');

const logger = getLogger('default');

module.exports = (err, req, res, next) => {
   if (err) {
      logger.error(`${req.ip} - ${req.method} ${req.url} - ${err.message} | XID(${req.cookies.token})`);
      return res.status(err.status || 500).json({ message: err.message });
   }
   next();
};
