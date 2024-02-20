# go-amumu
A go CLI app for sending and receiving emails using Microsoft Graph API

## set enviroment variables in .env file
```
CLIENT_ID=< ID >
TENANT_ID=< ID >
CLIENT_SECRET=< SECRET >
```
## Run the app
## Usage
```
Usage:
 amumu [subcommand] [options]
server options:
 -port string
 -config Path
client options:
 -to string - The email address of the recipient
 -from string - The email address of the sender
 -subject string - The subject of the email
 -message string - The message body of the email
 -channel boolean - Send to MS Teams channel (default=false)
```
## Examples
### as an smtp server
> [!NOTE]  
> The command server expect a yaml file called amumu-config.yaml in /etc/amumu-config.yaml otherwise, path should be specified with -config flag

```
amumu server -port 2525
```
```
amumu server -config /tmp/config.yaml
```
### send email with graph api
```
amumu client -to="mail@example.com" -from="mail@example.com" -subject="Example Subject" -message="Example Message"
```
### Send Alert to Ms Teams channel
```
amumu client -to="mail@example.com" -from="mail@example.com" -subject="Example Subject" -message="Example Message" --channel=true
```

![Alt text](slsa/SLSA-Badge-full-level1.svg)