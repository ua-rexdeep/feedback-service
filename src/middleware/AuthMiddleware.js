const { JsonWebTokenError } = require('jsonwebtoken');
const JWTService = require('../service/JWTService');

const AuthMiddleware = (req, res, next) => {
   if (req.headers.authorization && req.headers.authorization.startsWith('Bearer ')) {
      try {
         const key = req.headers.authorization.split(' ')[1];
         const decoded = JWTService.decode(key);
         if (typeof decoded === 'object' && 'login' in decoded) {
            req.user = { login: decoded.login };
            next();
         } else return next(new Error('wrong jwt content'));
      } catch (e) {
         return next(e);
      }
   } else return next(new JsonWebTokenError('Not bearer token'));
};

module.exports = AuthMiddleware;
