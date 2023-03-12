const { Pool } = require('pg');

const pool = new Pool({
   user: process.env.DB_USER,
   password: process.env.DB_PASS,
   host: process.env.DB_HOST,
   port: +process.env.DB_PORT,
   database: process.env.DB_DATABSE,
   max: 20, // максимальна кількість з'єднань в пулі
   idleTimeoutMillis: 30000, // час очікування в мс до закриття з'єднання вільного з'єднання
   connectionTimeoutMillis: 2000, // час очікування в мс на підключення до бази даних
});

(async () => {
   const client = await pool.connect();
   try {
      await client.query('CREATE EXTENSION IF NOT EXISTS "uuid-ossp";');
      await client.query(`
       CREATE TABLE IF NOT EXISTS feedback (
         id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
         user_login VARCHAR(100) NOT NULL,
         customer_name VARCHAR(50) NOT NULL,
         email VARCHAR(100) NOT NULL,
         feedback_text TEXT NOT NULL,
         source VARCHAR(20) NOT NULL
       );
     `);
   } finally {
      client.release();
   }
})();

module.exports = pool;
