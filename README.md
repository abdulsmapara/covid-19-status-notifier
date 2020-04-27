<p align="left">
	<h1 align="left">Covid-19 Status Notifier</h1>
	<h4 align="left">Bot that notifies Mattermost/Slack channel about the status of Covid-19 virus in India (by scanning the website - https://www.mohfw.gov.in/)</h4>
</p>


## Table of Content
- [About-the-bot](#about-the-bot)
- [Demonstration](#demonstration)
- [Functionalities](#functionalities)
- [Installation](#installation)


## About-the-bot

On changes about status of Covid-19 virus on the official website - [https://www.mohfw.gov.in/](https://www.mohfw.gov.in/), the bot throws a message to Mattermost/Slack channel.

## Demonstration

#### Demo on Mattermost Channel
![Image-Demo](https://github.com/abdulsmapara/Github-Media/blob/master/screenshot1.1.png)

#### Notification by Slack
![Image-Demo](https://github.com/abdulsmapara/Github-Media/blob/master/covid-19-status-notifier/slack-notification.png)

## Functionalities

- NOTIFICATIONS on MATTERMOST/SLACK on:
	* Increase in Covid-19 cases in India statewise
	* Increase in cured/discharged/migrated cases in India statewise
	* Increase in deaths per state
	* New states getting the virus
- RELIABLE: Source is the official government website, so it is reliable
- Logs important events in ```info.log```
- Do you want to add any new functionality or modify existing ones ?
	* Please go through [CONTRIBUTING.md](https://github.com/abdulsmapara/covid-19-status-notifier/blob/master/CONTRIBUTING.md) and contribute to this repository

## Installation

- Go language required
- Mattermost/Slack account required
- Mattermost/Slack Webhook URL required
- Get the required libraries:
	* ```go get github.com/anaskhan96/soup```
	* ```go get github.com/olekukonko/tablewriter```
- Good to Go:
	* Clone this repository
    ``` 
    git clone https://github.com/abdulsmapara/covid-19-status-notifier.git
    ```
	* ``` cd covid-19-status-notifier```
	* Set ```TO_SLACK``` and/or ```TO_MATTERMOST``` to ```true``` depending on where you want to send the notification 
	* Open ```consts.go``` and update ```SLACK_WEBHOOK_URL``` and/or ```MATTERMOST_WEBHOOK_URL``` with the webhook url available
	* Build the bot
	```bash
	go build main.go consts.go utils.go
	```
	* Create a cron job to recieve updates on changes in status
	```bash 
	crontab -e # Opens an editor
	# Run the bot every 10 minutes
	*/10 * * * * cd $PATH_TO_CLONE_DIR;./main
	```
