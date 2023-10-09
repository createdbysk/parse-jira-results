import puppeteer from 'puppeteer';
import readline from 'readline';

const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout
});

const waitForKeystroke = () => {
    return new Promise(resolve => {
        rl.question("Press any key to continue...", () => {
            rl.close();
            resolve(null);
        });
    });
};

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
                if (process.env.START_URL) {
                    return page.goto(process.env.START_URL);
                } else {
                    throw new Error("Set START_URL to the start page.")
                }
            })
            .then(() => waitForKeystroke())
            .then(() => browser.disconnect());
    })
    .catch(err => {
        console.error(err);
    });
