import puppeteer from 'puppeteer';
import readline from 'readline';

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
                .then(page => {
                    page.setRequestInterception(true);
                    page.on('request', (req) => {
                        if (req.isNavigationRequest()) {
                            const headers = req.headers();
                            delete headers['content-security-policy'];
                            req.continue({ headers });
                        } else {
                            req.continue();
                        }
                    });
                    return page.goto(start_url).then(() => page);
                })
                .then(async page => {
                    // Wait for the link to appear
                    await page.waitForSelector('table#issuetable tr:nth-child(2) td.issuekey a').then(
                        async element => element?.click()
                    )
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


