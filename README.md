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
## Usage
```
Usage:
 ./main [subcommand] [options]
server options:
 -port string
client options:
 -to string - The email address of the recipient
 -from string - The email address of the sender
 -subject string - The subject of the email
 -message string - The message body of the email
 -channel boolean - Send to MS Teams channel (default=false)
```
## Examples
### as an smtp server
```
./appname server -port 2525
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
