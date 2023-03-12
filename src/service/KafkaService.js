const { getLogger } = require('log4js');
const { SendProducer } = require('../config/kafka');

const logger = getLogger('default');

function SendProducerFeedback(feedback) {
   return SendProducer({
      topic: 'feedbacks',
      messages: [{ key: `feedback/${feedback.id}`, value: JSON.stringify(feedback) }],
   }).then(() => logger.log(`Send to producer feedback/${feedback?.id}`)).catch((e) => {
      logger.error('Error in sending feedback to producer ', e);
   });
}

module.exports = { SendProducerFeedback };
