const { validate } = require('uuid');
const { getLogger } = require('log4js');
const express = require('express');
const AuthMiddleware = require('../middleware/AuthMiddleware');

const { getFeedback, createFeedback, getUserFeedbacks, SerializeFeedback } = require('../service/FeedbackService');
const feedbackSchema = require('../schema/feedbackSchema');
const { RedisClient } = require('../config/redis');
const { SendProducerFeedback } = require('../service/KafkaService');

const logger = getLogger('FeedbackController');

const router = express.Router();

router.get('/', AuthMiddleware, async (req, res, next) => {
   try {
      const page = 'page' in req.query ? req.query.page : 1;
      const r = await getUserFeedbacks(req.user.login, +page);
      res.json(r.map((f) => SerializeFeedback(f)));
   } catch (e) {
      next(e);
   }
});

router.get('/:feedbackId', AuthMiddleware, async (req, res, next) => {
   logger.log(req.headers.authorization, req.params);

   if ('feedbackId' in req.params && typeof req.params.feedbackId === 'string' && validate(req.params.feedbackId)) {
      try {
         const feedback = await getFeedback(req.params.feedbackId);

         if (feedback) {
            res.json(SerializeFeedback(feedback));
         } else res.status(400).json({ error: 'No feedback with id' });
      } catch (e) {
         next(e);
      }
   } else {
      res.status(400).json({
         error: 'invalid feedbackId',
      });
   }
});

router.post('/', AuthMiddleware, async (req, res, next) => {
   const { error, value } = feedbackSchema.validate(req.body);

   if (error) {
      return res.status(400).json({ error: error.details[0].message });
   }
   try {
      const feedback = await createFeedback(req.user, value);
      const feedbackWithId = { ...value, id: feedback.id };
      logger.log('Inserted feedback', feedback.id);
      RedisClient.setex(`feedback/${feedback.id}`, 3600, JSON.stringify(value));
      SendProducerFeedback(feedbackWithId);
      return res.status(201).json(feedbackWithId);
   } catch (e) {
      next(e);
   }
});

module.exports = router;
