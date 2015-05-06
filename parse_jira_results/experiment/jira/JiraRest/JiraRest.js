'use strict';
var request,
    requestProgress,
    callback,
    progressCallback,
    options;

request = require('request');
requestProgress = require('request-progress');
options = {
  url: "https://jira.dev.socialware.com/rest/api/latest/search?jql=project%3DSW&expand=changelog&startAt=1&maxResults=50",
  auth: {
    user: '',
    password: ''
  },
  strictSSL: false
};
console.log('here');
callback = function (error, response, body) {
  if (!error && response.statusCode == 200) {
    var info = JSON.parse(body);
    console.log(JSON.stringify(info));
  }
  else {
    console.log(error);
  }
}
progressCallback = function (state) {
    console.log(JSON.stringify(state));
};
requestProgress(request(options, callback), {throttle:1000, delay:1000})
  .on('progress', progressCallback)
  .on('error', function (err) {
    console.log(err)
  });
