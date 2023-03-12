const Joi = require('joi');

module.exports = Joi.object({
   customer_name: Joi.string().required(),
   email: Joi.string().email().required(),
   feedback_text: Joi.string().required(),
   source: Joi.string().required(),
});
