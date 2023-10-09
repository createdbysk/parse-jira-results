import puppeteer from 'puppeteer';
import * as readline from 'readline';

Promise.resolve().then(
    () => {
        if (process.env.START_URL) {
            return process.env.START_URL;
        }
        throw new Error("Set START_URL environment variable to the start page.")
    }
).then(
    start_url => {
        puppeteer.connect({ browserURL: 'http://localhost:9222' })
        .then(async browser => {
            return browser.newPage()
                .then(async page => {
                    return page.setBypassCSP(true)
                        .then(() => page);
                })
                .then(async page => {
                    await page.setRequestInterception(true);
                    page.on('request', (req) => {
                        if (req.isNavigationRequest()) {
                            const headers = req.headers();
                            delete headers['content-security-policy'];
                            req.continue({ headers });
                        } else {
                            req.continue();
                        }
                    });
                    await page.goto(start_url);
                    return page;
                })
                .then(async page => {
                    // Wait for the link to appear
                    console.log("Wait for list of tickets.");
                    await page.waitForSelector('table#issuetable tr:nth-child(2) td.issuekey a').then(
                        async element => element?.click()
                    )
                    console.log("Clicked on JIRA ticket");
                    return page;
                })
                .then(async page => {
                    // We want to get to the transitions tab.
                    // The tabs change their html when they are active vs when they are not.
                    // Check for the all and transitions tabs to be inactive.
                    // If the all is inactive, always click on it first, which will render transitions inactive.
                    // If all is active, then transitions will already be inactive.
                    console.log("Wait for tab panels.");
                    await page.waitForSelector('a#all-tabpanel, a#transitions-summary-tabpanel', { timeout: 10000 });  // adjust the timeout as necessary
                    return page;
                })
                .then(async page => {
                    // Check if a#all-tabpanel exists
                    return await page.$('a#all-tabpanel').then(
                        async allTabExists => {
                            if (allTabExists) {
                                console.log("Click on All.");
                                await allTabExists.click();
                                await page.waitForSelector('a#transitions-summary-tabpanel');
                            }
                            return page;
                        }
                    )
                })
                .then(
                    // At this point, the transitions panel must be inactive.
                    async page => {
                        console.log("Click on transitions.");
                        await page.click('a#transitions-summary-tabpanel');
                        return page;
                    }
                )
                .then(async page => {
                    console.log("Wait for Open to In Progress")
                    // Look for the transition from Open to In Progress
                    await page.waitForFunction(() => {
                        // Find the "Open" element with lozenge structure
                        console.log("Look for Open -> In Progress")
                        const openElements = document.querySelectorAll('span.jira-issue-status-lozenge');
                        const openElement = Array.from(openElements).find(el => el.textContent?.trim() === "Open");

                        console.log("openElement", openElement)
                        if (!openElement) {
                            console.log("Did not find Open");
                            return false;
                        }
                        console.log("Found Open");
                        // Traverse to parent TD
                        const openTd = openElement.parentElement;
                        if (!openTd || openTd.tagName.toLowerCase() !== 'td') return false;
                    
                        // Check next TD for image (assuming "img" tag for ->)
                        const arrowTd = openTd.nextElementSibling;
                        console.log("arrowTd", arrowTd);
                        console.log("img selector", arrowTd?.querySelector('img'))
                        if (!arrowTd || !arrowTd.querySelector('img')) {
                            console.log("Did not find the arrow");
                            return false;
                        }
                    
                        // Check subsequent TD for "In Progress"
                        const inProgressTd = arrowTd.nextElementSibling;
                        if (!inProgressTd) return false;
                        const inProgressElement = inProgressTd.firstChild;
                        if (!inProgressElement || inProgressElement?.textContent?.trim() !== "In Progress") {
                            console.log("Did not find In Progress");
                            return false;
                        }
                        console.log("Found In Progress");
                        return true;
                    });
                })
                .then(() => readline.createInterface({
                    input: process.stdin,
                    output: process.stdout
                }))
                .then(rl => rl.question("Press any key to continue...", () => {
                    rl.close();
                }))
                .then(() => browser.disconnect());
        })    
    }
).catch(err => {
    console.error(err);
    process.exit(1);
});


