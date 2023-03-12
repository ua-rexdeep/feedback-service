const jwt = require('jsonwebtoken');

class JWTService {
   #secret = null;

   constructor(secret) {
      this.#secret = secret;
   }

   decode(token) {
      if (!token || typeof token !== 'string') throw new jwt.JsonWebTokenError(`expected string, got ${typeof token}`);
      return jwt.verify(token, this.#secret);
   }

   sign(data, options) {
      if (data && (typeof data === 'object' || typeof data === 'string')) {
         return jwt.sign(data, this.#secret, options);
      } throw new Error(`data is invalid, expecret object or string, got: ${typeof data}`);
   }
}

module.exports = new JWTService('zalupa');
