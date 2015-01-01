var fs,
    linq,
    expect,
    requireInjector;
fs = require('fs');
linq = require('linq');
expect = require('expect.js');
requireInjector = require('library/test_utilities/requireInjector');

describe('parse jira results', function () {
    var injector;
    beforeEach(function (done) {
        injector = requireInjector.createInjector();   
        done();
    });

    describe('issue status extractor', function () {
        'use strict';
        var issue,
            issueStatusExtractor,
            createdDate;

        beforeEach(function (done) {
            createdDate = "2014-12-01T15:00:25.000+0000";
            issue = {
                 "fields": {
                    "created":"createdDate"
                 },
                 "changelog":{  
                    "histories":[  
                       {  
                          "created":"2014-12-01T15:58:25.000+0000",
                          "items":[  
                             {  
                                "field":"status",
                                "fieldtype":"jira",
                                "from":"10002",
                                "fromString":"Triage",
                                "to":"1",
                                "toString":"Open"
                             }
                            ]
                        },
                       {  
                          "created":"2014-12-10T19:41:38.000+0000",
                          "items":[  
                             {  
                                "field":"Component",
                                "fieldtype":"jira",
                                "from":null,
                                "fromString":null,
                                "to":"10650",
                                "toString":"Filters"
                             }
                          ]
                       },
                       {  
                          "created":"2014-12-18T12:12:21.000+0000",
                          "items":[  
                             {  
                                "field":"status",
                                "fieldtype":"jira",
                                "from":"10002",
                                "fromString":"Open",
                                "to":"6",
                                "toString":"Closed"
                             }
                          ]
                       }
                    ]
                }
            };
            injector.require(['lib/issueStatusExtractor'], function (theIssueStatusExtractor) {
                issueStatusExtractor = theIssueStatusExtractor;
                done();
            });
        });
        it('should extract the status', function (done) {
            var expectedStatuses;
            expectedStatuses =
                [
                    {
                        date: "createdDate",
                        from: "Created",
                        to: "Triage"
                    },
                    {
                        date: "2014-12-01T15:58:25.000+0000",
                        from: "Triage",
                        to: "Open"
                    },
                    {
                        date: "2014-12-18T12:12:21.000+0000",
                        from: "Open",
                        to: "Closed"
                    }
                ];

            issueStatusExtractor(issue, function (err, statuses) {
                expect(statuses.toArray()).to.eql(expectedStatuses);
                done();
            });
        });
    });

    describe('issue priority extractor', function () {
      'use strict';
      var issue,
          expectedPriority,
          issuePriorityExtractor;
      beforeEach(function (done) {
        expectedPriority = "thePrority";
        issue = {  
         "key":"issue-id",
         "fields":{  
            "priority":{  
               "self":"https://jira.dev.socialware.com/rest/api/2/priority/2",
               "iconUrl":"https://jira.dev.socialware.com/images/icons/priorities/critical.png",
               "name":expectedPriority,
               "id":"2"
            }
          }
        };
        injector.require(['lib/issuePriorityExtractor'], function (theIssuePriorityExtractor) {
            issuePriorityExtractor = theIssuePriorityExtractor;
            done();
        });
      });
      it('should return the priority', function (done) {
        var actualPriority;
        actualPriority = issuePriorityExtractor(issue);
        expect(actualPriority).to.be(expectedPriority);
        done();
      });
    });

    describe('issue created date extractor', function () {
      'use strict';
      var issue,
          expectedCreatedDate,
          issueCreatedDateExtractor;
      beforeEach(function (done) {
        expectedCreatedDate = "theCreatedDate";
        issue = {  
         "key":"issue-id",
         "fields":{  
            "created":expectedCreatedDate
          }
        };
        injector.require(['lib/issueCreatedDateExtractor'], function (theIssueCreatedDateExtractor) {
            issueCreatedDateExtractor = theIssueCreatedDateExtractor;
            done();
        });
      });
      it('should return the priority', function (done) {
        var actualDate;
        actualDate = issueCreatedDateExtractor(issue);
        expect(actualDate).to.be(expectedCreatedDate);
        done();
      });
    });

    describe('status filter', function () {
        'use strict';
        var statuses,
            statusFilter;
        beforeEach(function (done) {
            statuses = [
                {
                    date: "date1",
                    from: "Triage",
                    to: "Open"
                },
                {
                    date: "date2",
                    from: "Open",
                    to: "In Design"
                },
                {
                    date: "date3",
                    from: "In Design",
                    to: "In Progress"
                },
                {
                    date: "date4",
                    from: "In Progress",
                    to: "In Test"
                },
                {
                    date: "date5",
                    from: "In Test",
                    to: "Closed"
                }
            ];
            
            injector.require(['lib/statusFilter'], function (theStatusFilter) {
                statusFilter = theStatusFilter;
                done();
            });            
        });
        it('should filter the status given the predicate', function (done) {
            var expectedStatuses;
            expectedStatuses = [{
                date: "date2",
                from: "Open",
                to: "In Design"
            }];
            statusFilter(statuses, 
                function (status) {
                    return status.date === "date2";
                },
                function (err, filteredStatuses) {
                    expect(filteredStatuses.toArray()).to.eql(expectedStatuses);
                    done();
                }
            );
        });
    });

    describe('lead time calculator', function () {
        'use strict';
        var issue,
            leadTimeCalculator;
        beforeEach(function (done) {
            injector.require(['lib/leadTimeCalculator'], function (theLeadTimeCalculator) {
                leadTimeCalculator = theLeadTimeCalculator;
                done();
            });
        });
        it('should calculate lead time given 2 dates', function (done) {
            var expectedLeadTime,
                startDate,
                endDate,
                leadTime;
            expectedLeadTime = 18;
            startDate = "2014-12-01T15:58:25.000+0000";
            endDate = "2014-12-18T12:12:21.000+0000";
            leadTime = leadTimeCalculator(startDate, endDate);
            expect(leadTime).to.be(expectedLeadTime);
            done();
        });
        it('should return null if startDate is null', function (done) {
            var expectedLeadTime,
                startDate,
                endDate,
                leadTime;
            startDate = null;
            expectedLeadTime = null;
            endDate = "2014-12-18T12:12:21.000+0000";
            leadTime = leadTimeCalculator(startDate, endDate);
            expect(leadTime).to.be(expectedLeadTime);
            done();
        });
        it('should return null if endDate is null', function (done) {
            var expectedLeadTime,
                startDate,
                endDate,
                leadTime;
            startDate = "2014-12-18T12:12:21.000+0000";
            expectedLeadTime = null;
            endDate = null;
            leadTime = leadTimeCalculator(startDate, endDate);
            expect(leadTime).to.be(expectedLeadTime);
            done();
        });
    });
});

describe('for later', function () {
    var resultsJSON;
        resultsJSON = 
        {  
           "issues":[  
            {
                 "key":"issue-id",
                 "changelog":{  
                    "histories":[  
                       {  
                          "created":"2014-12-01T15:58:25.000+0000",
                          "items":[  
                             {  
                                "field":"status",
                                "fieldtype":"jira",
                                "from":"10002",
                                "fromString":"Triage",
                                "to":"1",
                                "toString":"Open"
                             },
                             {  
                                "field":"assignee",
                                "fieldtype":"jira",
                                "from":null,
                                "fromString":null,
                                "to":"womack",
                                "toString":"James Wommack"
                             }
                          ]
                       },
                       {  
                          "created":"2014-12-10T19:41:38.000+0000",
                          "items":[  
                             {  
                                "field":"Component",
                                "fieldtype":"jira",
                                "from":null,
                                "fromString":null,
                                "to":"10650",
                                "toString":"Filters"
                             }
                          ]
                       },
                       {  
                          "created":"2014-12-10T19:41:42.000+0000",
                          "items":[  
                             {  
                                "field":"labels",
                                "fieldtype":"jira",
                                "from":null,
                                "fromString":"",
                                "to":null,
                                "toString":"blocking_filter"
                             }
                          ]
                       },
                       {  
                          "created":"2014-12-12T14:42:26.000+0000",
                          "items":[  
                             {  
                                "field":"status",
                                "fieldtype":"jira",
                                "from":"1",
                                "fromString":"Open",
                                "to":"10002",
                                "toString":"Triage"
                             }
                          ]
                       },
                       {  
                          "id":"181954",
                          "author":{  
                             "self":"https://jira.dev.socialware.com/rest/api/2/user?username=vparsons",
                             "name":"vparsons",
                             "emailAddress":"vparsons@socialware.com",
                             "avatarUrls":{  
                                "16x16":"https://jira.dev.socialware.com/secure/useravatar?size=xsmall&ownerId=vparsons&avatarId=11216",
                                "24x24":"https://jira.dev.socialware.com/secure/useravatar?size=small&ownerId=vparsons&avatarId=11216",
                                "32x32":"https://jira.dev.socialware.com/secure/useravatar?size=medium&ownerId=vparsons&avatarId=11216",
                                "48x48":"https://jira.dev.socialware.com/secure/useravatar?ownerId=vparsons&avatarId=11216"
                             },
                             "displayName":"Viki Parsons",
                             "active":true
                          },
                          "created":"2014-12-12T15:32:30.000+0000",
                          "items":[  
                             {  
                                "field":"status",
                                "fieldtype":"jira",
                                "from":"10002",
                                "fromString":"Triage",
                                "to":"1",
                                "toString":"Open"
                             },
                             {  
                                "field":"Fix Version",
                                "fieldtype":"jira",
                                "from":null,
                                "fromString":null,
                                "to":"10096",
                                "toString":"Intermediate Build"
                             }
                          ]
                       },
                       {  
                          "created":"2014-12-12T19:43:36.000+0000",
                          "items":[  
                             {  
                                "field":"status",
                                "fieldtype":"jira",
                                "from":"1",
                                "fromString":"Open",
                                "to":"3",
                                "toString":"In Progress"
                             }
                          ]
                       },
                       {  
                          "id":"182015",
                          "author":{  
                             "self":"https://jira.dev.socialware.com/rest/api/2/user?username=skumar",
                             "name":"skumar",
                             "emailAddress":"skumar@socialware.com",
                             "avatarUrls":{  
                                "16x16":"https://jira.dev.socialware.com/secure/useravatar?size=xsmall&avatarId=10119",
                                "24x24":"https://jira.dev.socialware.com/secure/useravatar?size=small&avatarId=10119",
                                "32x32":"https://jira.dev.socialware.com/secure/useravatar?size=medium&avatarId=10119",
                                "48x48":"https://jira.dev.socialware.com/secure/useravatar?avatarId=10119"
                             },
                             "displayName":"Satish Kumar",
                             "active":true
                          },
                          "created":"2014-12-12T19:43:51.000+0000",
                          "items":[  
                             {  
                                "field":"status",
                                "fieldtype":"jira",
                                "from":"3",
                                "fromString":"In Progress",
                                "to":"10020",
                                "toString":"In Design"
                             }
                          ]
                       },
                       {  
                          "id":"182048",
                          "created":"2014-12-13T02:44:04.000+0000",
                          "items":[  
                             {  
                                "field":"status",
                                "fieldtype":"jira",
                                "from":"10020",
                                "fromString":"In Design",
                                "to":"10003",
                                "toString":"In Review"
                             }
                          ]
                       },
                       {  
                          "created":"2014-12-13T02:50:03.000+0000",
                          "items":[  
                             {  
                                "field":"status",
                                "fieldtype":"jira",
                                "from":"10003",
                                "fromString":"In Review",
                                "to":"10005",
                                "toString":"Ready for Testing"
                             }
                          ]
                       },
                       {  
                          "created":"2014-12-15T18:23:32.000+0000",
                          "items":[  
                             {  
                                "field":"status",
                                "fieldtype":"jira",
                                "from":"10005",
                                "fromString":"Ready for Testing",
                                "to":"3",
                                "toString":"In Progress"
                             }
                          ]
                       },
                       {  
                          "created":"2014-12-16T16:00:39.000+0000",
                          "items":[  
                             {  
                                "field":"status",
                                "fieldtype":"jira",
                                "from":"3",
                                "fromString":"In Progress",
                                "to":"10005",
                                "toString":"Ready for Testing"
                             }
                          ]
                       },
                       {  
                          "created":"2014-12-18T08:49:20.000+0000",
                          "items":[  
                             {  
                                "field":"status",
                                "fieldtype":"jira",
                                "from":"10005",
                                "fromString":"Ready for Testing",
                                "to":"10002",
                                "toString":"Triage"
                             },
                          ]
                       },
                       {  
                          "created":"2014-12-18T12:12:21.000+0000",
                          "items":[  
                             {  
                                "field":"status",
                                "fieldtype":"jira",
                                "from":"10002",
                                "fromString":"Triage",
                                "to":"6",
                                "toString":"Closed"
                             }
                          ]
                       }
                    ]
                }
            }
        ]};

});