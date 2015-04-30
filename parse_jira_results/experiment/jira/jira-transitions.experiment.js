/**
 * Run npm install -g jira before you try this.
 * Abandoned because this did not work with the authentication scheme.
 */
var program,
    JiraApi,
    jira;

program = require('commander');
jira = require('jira');

program
    .version('0.0.1')
    .usage('[options] url')
    .option('-u --user [value]', 'username')
    .option('-p --password [value]', 'password')
    .parse(process.argv);

JiraApi = require('jira').JiraApi;
console.error(program.user, program.password);
console.error(program.args[0]);
jira = new JiraApi('https', program.args[0], 443, program.user, program.password, '2.0', true, false);

jira.searchJira("project%3DSW&expand=changelog&maxResults=500", {fields: ['priority', 'created']}, function (results) {
    console.error('RESULTS: ', results);
});
