const { getLogger } = require('log4js');
const pool = require('../config/database');
const { RedisClient } = require('../config/redis');

const logger = getLogger('default');
const page_limit = 15;

const getFeedback = async (id) => {
   const cached = await RedisClient.get(`feedback/${id}`);
   if (cached) {
      logger.log(`returned cached data for feedback/${id}`);
      return JSON.parse(cached);
   }
   const client = await pool.connect();
   try {
      const res = await client.query('SELECT * FROM feedback WHERE id = $1', [id]);
      return res.rows[0];
   } finally {
      client.release();
   }
};

const getUserFeedbacks = async (user_login, page = 1) => {
   const client = await pool.connect();
   if (page < 1 || page > 25565) throw new Error('Invalid page');
   try {
      const res = await client.query(
         'SELECT * FROM feedback WHERE user_login = $1 ORDER BY id DESC LIMIT $2 OFFSET $3;',
         [user_login, page_limit, page_limit * (page - 1)],
      );
      return res.rows;
   } finally {
      client.release();
   }
};

async function createFeedback(creator, { customer_name, email, feedback_text, source }) {
   // const feedbackId = uuidv4();

   const query = {
      text: 'INSERT INTO feedback(user_login, customer_name, email, feedback_text, source) VALUES($1, $2, $3, $4, $5) RETURNING id',
      values: [creator.login, customer_name, email, feedback_text, source],
   };

   try {
      const result = await pool.query(query);
      logger.log(`New feedback with id ${result.rows[0].id} has been added to the database`);
      return result.rows[0];
   } catch (error) {
      logger.error(`Error occurred while creating feedback: ${error}`);
      throw error;
   }
}

function SerializeFeedback(feedback) {
   return {
      customer_name: feedback.customer_name,
      email: feedback.email,
      feedback_text: feedback.feedback_text,
      source: feedback.source,
   };
}

module.exports = { createFeedback, getFeedback, getUserFeedbacks, SerializeFeedback };
