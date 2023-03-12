const { getLogger, configure } = require('log4js');
const { init } = require('./config/http');
require('./config/redis');

configure({
   appenders: { console: { type: 'console' } },
   categories: { default: { level: 'ALL', appenders: ['console'] } },
});
const logger = getLogger('default');

async function boot() {
   logger.log('Boot started');
   // try {
   //    await ConnectProducer().then(() => logger.log('Kafka connected'));
   // } catch (e) {
   //    logger.error('Kafka error', e.message);
   // }
   init().then(() => logger.log('Boot finish'));
}
boot();
