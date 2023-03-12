const Joi = require('joi');

module.exports = Joi.object({
   login: Joi.string().required(),
   password: Joi.string(),
});
