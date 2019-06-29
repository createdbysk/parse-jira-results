FROM mhart/alpine-node:8.11.4

RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app

COPY package.json /usr/src/app
COPY package-lock.json /usr/src/app
RUN apk add --update git
RUN npm install

COPY lib /usr/src/app/lib
COPY bin /usr/src/app/bin

ENTRYPOINT ["node", "bin/overwriteJiraResultsInGoogleSheet.js"]
