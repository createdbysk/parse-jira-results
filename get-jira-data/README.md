# get-jira-data

## Debug with a google chrome instance

### Development on local host

On a Desktop Host, close all Google Chrome instances. Then run

        "/Applications/Google Chrome.app/Contents/MacOS/Google Chrome" --remote-debugging-port=9222 --disable-dev-shm-usage


### Development on a remote host

On a Desktop Host, close all Google Chrome instances. Then run

        ssh -fNR 9222:localhost:9222 user@remote_machine
        "/Applications/Google Chrome.app/Contents/MacOS/Google Chrome" --remote-debugging-port=9222 --disable-dev-shm-usage
