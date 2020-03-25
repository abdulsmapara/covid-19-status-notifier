<p align="center">
	<h1 align="center">Mattermost Virus Tracker Bot</h1>
	<h5 align="center">Bot that notifies Mattermost channel about the status of Covid-19 virus in India (by scanning the website - https://www.mohfw.gov.in/)</h5>
</p>


## Table of Content
- [About-the-bot](#about-the-bot)
- [Demonstration](#demonstration)
- [Features](#features)
- [Installation](#installation)
- [Info](#info)


## About-the-bot

On changes about status of Covid-19 virus on the official website - [https://www.mohfw.gov.in/](https://www.mohfw.gov.in/), the bot throws a message to Mattermost.

## Demonstration

## Features

- Get notifications on Mattermost
	* New Corona Virus cases happening in India
	* How many Indian nationals have Corona Virus per State?
	* How many deaths happened per State?
	* New States getting the virus
- Source is the official government website, so it is reliable
- Logs important events in ```info.log```
- Released under Apache 2.0 license

## Installation

- Go language required
- Mattermost account required
- Mattermost Webhook URL required
- Get the required libraries:
	* ```go get github.com/anaskhan96/soup```
	* ```go get github.com/olekukonko/tablewriter```
- Good to Go:
	* Clone this repository
    ``` 
    git clone https://github.com/abdulsmapara/mattermost-virus-tracker.git
    ```
	* ``` cd mattermost-virus-tracker```
	* Open ```consts.go``` and update ```<WEBHOOK-URL>``` with the webhook url
	* Build the bot
	```
	go build main.go consts.go utils.go
	```
	* Create a cron job to recieve updates on changes in status
	``` 
	crontab -e # Opens an editor
	# Run the bot every 5 minutes
	*/5 * * * * cd $PATH_TO_CLONE_DIR;./main
	```

## Info

Created by [@abdulsmapara](https://github.com/abdulsmapara)

