#TimeBot Setup

1. Setup up https://api.slack.com/apps with Following Information

## OAuth & Permissions

- Bot Token Scopes
	- app_mentions:read
	- channels:history
	- chat:write
	- groups:history
	- groups:write
	- im:history
	- im:write
	- mpim:history
	- users:read

## Event Subscriptions

- Request URL 
	- ``` http://example.com/slack_events/v1/mybot-v1_events ``` (Note this URL will come from Ngrok)

- Subscribe to bot events
	- message.channels
	- message.groups
	- message.im
	- message.mpim

2. Setup for Docker

## Set Environment Variables

- SLACK_TOKEN from 'OAuth & Permissions'
- SLACK_VERIFICATION_TOKEN from 'Basic Information'
- SLACK_EVENTS_CALLBACK_PATH equal to '/slack_events/v1/mybot-v1_events'
- SLACK_INTERACTIONS_CALLBACK_PATH equal to '/slack_events/v1/mybot-v1_events'

## Run Docker Command in Directory Holding 'config' Folder

- ``` docker run --rm --name mybot --env SLACK_TOKEN=$SLACK_TOKEN --env SLACK_VERIFICATION_TOKEN=$SLACK_VERIFICATION_TOKEN --env SLACK_EVENTS_CALLBACK_PATH=$SLACK_EVENTS_CALLBACK_PATH --env SLACK_INTERACTIONS_CALLBACK_PATH=$SLACK_INTERACTIONS_CALLBACK_PATH -p 3000:3000 -v "$PWD"/config:/config target/flottbot:latest "./flottbot" ```

3. Using Ngrok 

## Run Ngrok Command

- ``` Ngrok http 3000 ```

## Add Ngrok URL into Event Subscriptions

- Request URL Example
	- ``` http://8a0221b8.ngrok.io/slack_events/v1/mybot-v1_events ```



