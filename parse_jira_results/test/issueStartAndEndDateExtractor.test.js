var requireInjector,
    expect,
    sinon,
    linq;
requireInjector = require('library/test_utilities/requireInjector');
expect = require('expect.js');
sinon = require('sinon');
linq = require('linq');

describe('issue start and end date extractor', function () {
    'use strict';
    var injector,
        statuses,
        expectedStartDate,
        expectedEndDate,
        statusFilter,
        issueStartAndEndDateExtractor;
        
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

        statusFilter = sinon.stub();
        statusFilter.withArgs(statuses, 
                            sinon.match(function (predicate) {
                                if (typeof(predicate) !== 'function') {
                                    return false;
                                }
                                return predicate(statuses[1]);
                            }, 'Start Status Filter'),
                            sinon.match.func)
                    .callsArgWith(2, null, linq.make(statuses[1]));
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
            .mock('linq', linq)
            .require(['lib/issueStartAndEndDateExtractor'], function (theIssueStartAndEndDateExtractor) {
            issueStartAndEndDateExtractor = theIssueStartAndEndDateExtractor;
            done();
        });            
    });
    it('should extract the start and end date for the given issue.', 
        function (done) {
            var issueStartAndEndDate;
            issueStartAndEndDateExtractor(statuses, function (err, startDate, endDate) {
                expect([startDate, endDate]).to.eql([expectedStartDate, expectedEndDate]);
                done();
            });
        }
    );
});