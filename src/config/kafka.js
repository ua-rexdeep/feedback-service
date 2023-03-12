// eslint-disable-next-line import/no-extraneous-dependencies
const { Kafka } = require('kafkajs');

const kafka = new Kafka({
   clientId: 'app',
   brokers: ['localhost:9092'],
});
const KafkaProducer = kafka.producer();
let status = false;

KafkaProducer.on('producer.network.request_timeout', async () => {
   status = false;
   await KafkaProducer.disconnect();
});

async function ConnectProducer() {
   await KafkaProducer.connect();
   status = true;
}

async function SendProducer(record) {
   if (!status) await ConnectProducer();
   return KafkaProducer.send(record);
}

module.exports = { KafkaProducer, SendProducer };
