const { getLogger } = require('log4js');
const express = require('express');

const router = express.Router();
const JWTService = require('../service/JWTService');
const authSchema = require('../schema/authSchema');

const logger = getLogger('AuthController');

router.post('/auth', (req, res, next) => {
   const { error, value } = authSchema.validate(req.body);
   if (error) return next(error);
   const encoded = JWTService.sign({
      login: value.login,
      iat: Date.now(),
   }, { expiresIn: '30d' });
   logger.log('auth', value.login);
   res.cookie('Authorization', `Bearer ${encoded}`);
   res.send('ok');
});

module.exports = router;
