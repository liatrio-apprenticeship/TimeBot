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

> Note: Offical Framework documentation below

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
	- im:read
	- im:write
	- mpim:history
	- users:read
    - users:read.email

#### Event Subscriptions

- Request URL 
	- ``` http://example.com/slack_events/v1/mybot-v1_events ``` 
       
    > (Note: Replace 'http://example.com' with URL, Ngrok is an option to create URL, Bot may need to be running to save URL)

- Subscribe to bot events
	- message.channels
	- message.groups
	- message.im
	- message.mpim

### 2. Initialize Flottbot Locally

1. Clone Github Repository

	- ``` git clone https://github.com/liatrio-apprenticeship/TimeBot ```

2. Update and Add ./flottbot-volume/ files

	- Update bot.yml with Slack Tokens
		- ${SLACK_EVENTS_CALLBACK_PATH}       -> /slack_events/v1/mybot-v1_events
		- ${SLACK_INTERACTIONS_CALLBACK_PATH} -> /slack_events/v1/mybot-v1_events
		- ${SLACK_TOKEN}                      -> Slack Token from OAuth & Permissions

		> Slack_Token should start with 'xoxb-'

		- ${SLACK_VERIFICATION_TOKEN}         -> Slack Verification Token from Basic Information

	- Google Sheets API credentials
		- Download the Configuration.json file from https://developers.google.com/sheets/api/quickstart/go
		- Place this file in the flottbot-volume directory

3. Initial TimeBot Start

	- Start TimeBot with Docker
		- ``` docker-compose up ```
	
	- Initial Bot Setup Commands
		- Once bot is running setup a slack direct message with bot
		- Type 'sheet-print', you will receive a url to authenticate with. Do with with the email account associated with bot.
		- Once you receive a token, type 'sheet-token your_token_here' to authenticate the slack bot

**Bot Setup Complete!!**
	
---

## **Notes** 

- Ngrok Command
	- ``` ngrok http 3000 ```

- Example Slack Event Subscription URL
	- http://8a0221b8.ngrok.io/slack_events/v1/mybot-v1_events

---

## **Contributers** 

- [RJ Souza](https://github.com/Empyreus)
- [Jacob Brolliar](https://github.com/MrDr-Professor)
- [Daniel Hunt](https://github.com/DanHunt27)
