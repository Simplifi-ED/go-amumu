# go-amumu
A go CLI app for sending and receiving emails using Microsoft Graph API

## Set enviroment variables in .env file
```
CLIENT_ID=< ID >
TENANT_ID=< ID >
CLIENT_SECRET=< SECRET >
```
## Run the app
## Usage
```bash
Usage:
 amumu [subcommand] [options]
server options:
 -config string - Path to yaml config file
client options:
 -to string - The email address of the recipient
 -from string - The email address of the sender
 -subject string - The subject of the email
 -message string - The message body of the email
 -channel boolean - Send to MS Teams channel (default=false)
```
## Examples
### As an smtp server
> [!NOTE]  
> The command server expect a yaml file called amumu-config.yaml in /etc/amumu-config.yaml otherwise, path should be specified with -config flag

```
amumu server -config=/tmp/config.yaml
```
```yaml
host: "0.0.0.0:1065"
domain: "simplified"
writetimeout: 10
readtimeout: 10
MaxMessageBytes: 1048576
maxRecipients: 50
allowInsecureAuth: true
```
### Send email
```
amumu client -to="mail@example.com" -from="mail@example.com" -subject="Example Subject" -message="Example Message"
```
### Send with Alert to Ms Teams channel
```
amumu client -to="mail@example.com" -from="mail@example.com" -subject="Example Subject" -message="Example Message" --channel=true
```
### Using config file
```bash
amumu client -config=/tmp/config.yaml
```
```yaml
to: "email@example.com"
from: "email@example.com"
subject: "test subject"
body: "Hello world"
channel: 
  title: "IT Challenge Alerts"
  text: "Hello world test"
``` 

![Alt text](slsa/SLSA-Badge-full-level1.svg)