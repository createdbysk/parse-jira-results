FROM node:5.11.1-wheezy

RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app

COPY package.json /usr/src/app
RUN npm install

COPY lib /usr/src/app/lib
COPY bin /usr/src/app/bin

ENTRYPOINT ["node", "bin/overwriteJiraResultsInGoogleSheet.js"]
