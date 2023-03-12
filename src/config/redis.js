const { getLogger } = require('log4js');
// eslint-disable-next-line import/no-extraneous-dependencies
const Redis = require('ioredis');

const logger = getLogger('default');
const client = new Redis({
   username: process.env.REDIS_USER,
   host: process.env.REDIS_HOST,
   password: process.env.REDIS_PASS,
   port: +process.env.REDIS_PORT,
});
let status = false;

client.on('error', (err) => {
   logger.error('Redis Client Error', err);
   status = false;
});
client.on('end', () => { status = false; });
client.on('ready', () => {
   logger.log('Redis ready');
   status = true;
});

module.exports = { RedisClient: client, status };
