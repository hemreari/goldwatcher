# Goldwatcher

Simple web scrapper for local gold prices in Antalya/Turkey city and Telegram bot that shows recent prices. <br>
To use it simple send an `/anlik` message to the bot on Telegram: https://t.me/goldwatcherbot

## Before Running the App

- Get a Telegram Access Token for the bot: https://t.me/BotFather
- Add acquired token to the .env file (**<user_tgtoken>**)

## Installation

### Docker

- Run docker compose file with this command: `docker compose up -d`

##### Disclamer

compose-file uses default Postgresql configuration values. I strongly advise you to change them for security reasons.

### From The Source

- Clone the repository
- Build the app `go build`
- Run the app `./goldwatcher`

### Prediction

There is a prediction module written in Python. You can get a prediction for all entities for next 30 days. The prediction module uses Facebook's [Prophet](https://facebook.github.io/prophet/) library. <br>
To run the predictions follow this steps:

- `sh cd prediction`
- `sh pip install -r requirements.txt`
- `sh python3 main.py`

Make sure you have python v3.x installed on your machine.

##### Legal Notice

[**EN**] This bot (Goldwatcher) is developed only for education and informing purposes. I (the developer) and other contributers of this project can not be held responsible for any kind of invesment losts. <br>
Please visit http://akodkur.com to get most recent price values. <br>

[**TR**]
Bu bot (Goldwatcher) herhangi bir yatirim tavsiyesi vermez. Dogacak zararlardan dolayi bu projenin gelistiricileri sorumlu tutulamaz. <br>
En guncel fiyatlara ulasmak icin lutfen http://akodkur.com websitesini ziyaret ediniz.
