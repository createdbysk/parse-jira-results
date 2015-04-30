var requireInjector,
    expect,
    sinon,
    linq;
requireInjector = require('library/test_utilities/requireInjector');
expect = require('expect.js');
sinon = require('sinon');
linq = require('linq');

describe('issue', function () {
    'use strict';
    var injector,
        issue,
        issueStatusExtractor,
        statuses,
        expectedStartDate,
        expectedEndDate,
        statusFilter;
        
    beforeEach(function (done) {
        injector = requireInjector.createInjector();
        expectedStartDate = 'date2';
        expectedEndDate = 'date5';
        statuses = [
            {
                date: 'date1',
                from: "Triage",
                to: "Open"
            },
            {
                date: expectedStartDate,
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
                date: expectedEndDate,
                from: "In Test",
                to: "Closed"
            }
        ];
        issue = {
            statuses: statuses
        }
        statusFilter = sinon.stub();
        issueStatusExtractor = sinon.stub();

        issueStatusExtractor.withArgs(issue, sinon.match.func).callsArgWith(1, null, statuses);
        done();
    });           

    describe('issue start date extractor', function () {
        'use strict';
        var issueStartDateExtractor;
            
        beforeEach(function (done) {
            statusFilter.withArgs(statuses, 
                                sinon.match(function (predicate) {
                                    if (typeof(predicate) !== 'function') {
                                        return false;
                                    }
                                    return predicate(statuses[1]);
                                }, 'Start Status Filter'),
                                sinon.match.func)
                        .callsArgWith(2, null, linq.make(statuses[1]));

            injector
                .mock('lib/statusFilter', statusFilter)
                .mock('lib/issueStatusExtractor', issueStatusExtractor)
                .mock('linq', linq)
                .require(['lib/issueStartDateExtractor'], function (theIssueStartDateExtractor) {
                    issueStartDateExtractor = theIssueStartDateExtractor;
                    done();
                });            
        });
        it('should extract the start for the given issue.', 
            function (done) {
                var issueStartDate;
                issueStartDateExtractor(issue, function (err, startDate) {
                    expect(startDate).to.eql(expectedStartDate);
                    done();
                });
            }
        );
    });

    describe('issue end date extractor', function () {
        'use strict';
        var issueEndDateExtractor;
            
        beforeEach(function (done) {
            statusFilter.withArgs(statuses, 
                                sinon.match(function (predicate) {
                                    if (typeof(predicate) !== 'function') {
                                        return false;
                                    }
                                    return predicate(statuses[4]);
                                }, 'End Status Filter'),
                                sinon.match.func)
                        .callsArgWith(2, null, linq.make(statuses[4]));

            injector
                .mock('lib/statusFilter', statusFilter)
                .mock('lib/issueStatusExtractor', issueStatusExtractor)
                .mock('linq', linq)
                .require(['lib/issueEndDateExtractor'], function (theIssueEndDateExtractor) {
                    issueEndDateExtractor = theIssueEndDateExtractor;
                    done();
                });            
        });
        it('should extract the end date for the given issue.', 
            function (done) {
                var issueEndDate;
                issueEndDateExtractor(issue, function (err, endDate) {
                    expect(endDate).to.eql(expectedEndDate);
                    done();
                });
            }
        );
    });
})
