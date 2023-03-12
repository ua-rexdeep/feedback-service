const express = require('express');
const FeedbackController = require('../controller/FeedbackController');
const UserController = require('../controller/UserController');

const router = express.Router();

const routes = [
   {
      path: '/user',
      route: UserController,
   },
   {
      path: '/feedback',
      route: FeedbackController,
   },
];

routes.forEach((route) => {
   router.use(route.path, route.route);
});

module.exports = router;
