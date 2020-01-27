# **TimeBot**

TimeBot is a Google Sheet integrated Slackbot used for timekeeping. TimeBot keeps track of current hours while providing information about a users workschedule.

---

## **Features**

- ~~Clock-In~~
- ~~Clock-Out~~
- ~~Weekly Hours~~
- ~~Pay Period Hours~~

---

## **Project Status**

- Project is currently under development. 
- Initial setup and implementation is in progress.

---

## **Setup** 

> Note: Start with offical documentation below

**[Offical Documentation](https://target.github.io/flottbot-docs/basics/slack/)**

### 1. Setup up https://api.slack.com/apps with Following Information

#### OAuth & Permissions

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

#### Event Subscriptions

- Request URL 
	- ``` http://example.com/slack_events/v1/mybot-v1_events ``` 
       
    > (Note: Replace 'http://example.com' with URL will come from Ngrok)

- Subscribe to bot events
	- message.channels
	- message.groups
	- message.im
	- message.mpim

### 2. Setup for Docker

#### Set Environment Variables

- SLACK_TOKEN from 'OAuth & Permissions'
- SLACK_VERIFICATION_TOKEN from 'Basic Information'
- SLACK_EVENTS_CALLBACK_PATH equal to '/slack_events/v1/mybot-v1_events'
- SLACK_INTERACTIONS_CALLBACK_PATH equal to '/slack_events/v1/mybot-v1_events'

#### Run Docker Command in Directory Holding 'config' Folder

- ``` docker run --rm --name mybot --env SLACK_TOKEN=$SLACK_TOKEN --env SLACK_VERIFICATION_TOKEN=$SLACK_VERIFICATION_TOKEN --env SLACK_EVENTS_CALLBACK_PATH=$SLACK_EVENTS_CALLBACK_PATH --env SLACK_INTERACTIONS_CALLBACK_PATH=$SLACK_INTERACTIONS_CALLBACK_PATH -p 3000:3000 -v "$PWD"/config:/config target/flottbot:latest "./flottbot" ```
> You can also download the correct flottbot release file for your OS [here](https://github.com/target/flottbot/releases) and run it in the directory above your /config/ folder to start TimeBot

### 3. Using Ngrok 

#### Run Ngrok Command

- ``` Ngrok http 3000 ```

#### Add Ngrok URL into Event Subscriptions

- Request URL Example
	- ``` http://8a0221b8.ngrok.io/slack_events/v1/mybot-v1_events ```

---

## **Contributers** 

- [RJ Souza](https://github.com/Empyreus)
- [Jacob Brolliar](https://github.com/MrDr-Professor)
- [Daniel Hunt](https://github.com/DanHunt27)