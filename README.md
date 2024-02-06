# go-send
A go CLI app for sending emails using Microsoft Graph API

![Alt text](slsa/SLSA-Badge-full-level1.svg)

## export enviroment variables
```
export CLIENT_ID=< ID >
export TENANT_ID=< ID >
export CLIENT_SECRET=< SECRET >
```
## Run the app
### as an smtp server
```
./appname server
```
### send email with graph api
```
./appname client -to="mail@example.com" -from="mail@example.com" -subject="Example Subject" -message="Example Message"
```
### Send Alert to Ms Teams channel
```
./appname client -to="mail@example.com" -from="mail@example.com" -subject="Example Subject" -message="Example Message" --channel=true
```
> [!NOTE]  
> The flag channel is "false" by default
