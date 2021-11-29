# OCM URL Verify Bot
Checks the url included in a discord message. Displays a verified status when the domain of the url attached is in the [verified-domains.txt](verified-domains.txt).

Only `https` is accepted.

## Permissions
Requires the following permissions to work on Discord
```
Scope: Bot
Bot Permissions: Send message
```

## Setup
```bash
$ export OCM_URL_DISCORD_API_KEY="YOUR_DISCORD_BOT_API_KEY"
```
## Run in bash (macOS) requires taskdev

```bash
$ task br
```

## Run in bash (other OSes)
```bash
$ go build -o bin/ocm-bot src/*
$ bin/ocm-bot
```

# Contributing
Feel free to make pull requests, especially to add more verified urls.