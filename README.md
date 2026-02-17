<h1 style="color: #29d18d; text-align: center;">
F L O R O B O T<br/>
A simple funny Telegram Bot, written in Go<br/>
By @micheldevs<br/>
₍ᐢ. .ᐢ₎ ₊˚⊹♡
</h1>

Florobot is a simple implementation of a Telegram bot written in Go, focused in entertainment with commands to crack jokes, roast people if it gets insulted, answer questions with random facts, play to the russian roulette with friends or scrape contents of webs to retrieve it later through commands or notifications.

This bot has been designed as much modular as possible, so making any modification to the existing solution is easy. Feel free to modify, correct any part or play with it, this project was taken as a side project to learn Go, so it is likely that some parts do not follow the standard!

## Installation
Florobot is offered through two versions:

1. A standalone version (with SQLite as database)
2. A non-standalone version (with Postgres as database)

Both versions have been fully tested, and can be used/deployed using Docker or the binaries of the bot.

### Standalone installation

#### Installation with the binary (Windows/Ubuntu)
1. Navigate to https://github.com/micheldevs/florobot/releases/
2. Select a `standalone` version and download depending your OS the `florobot_win11_amd64.zip` or `florobot_linux_amd64.zip` file
3. Unzip the downloaded zip file in your desired directory
4. Copy the `example.env` file, rename it to `.env` file and set its variables according your preferences. Almost all these vars are already set, only being necessary to change the `TG_BOT_APITOKEN` by your Telegram bot token. For more information of the env variables, please check the [Bot configuration](#bot-configuration) section.
5. Run the binary of the bot (add execution permissions before in Linux), and test your bot instance by sending a `/start` to your TG bot.

    * `florobot.exe` - In Windows
    * `florobot` - In Linux

    If the bot has replied, you will already have the bot configured and ready to attend any request. If you desire that will attend any petition once your server/computer is booted up, do not forget to configure it as a Windows or Linux service and set the execution path as the place where the bot is! 

#### Installation with Docker
1. Navigate to https://github.com/micheldevs/florobot/releases/
2. Select a `standalone` version and download `Source code (zip)` file
3. Unzip the downloaded zip file in your desired directory
4. Open the `docker-compose.yaml` file and set the variables of the `environment` section according your preferences. Almost all these vars are already set, only being necessary to change the `TG_BOT_APITOKEN` by your Telegram bot token. For more information of the env variables, please check the [Bot configuration](#bot-configuration) section.
5. Open a terminal in the root directory of the bot, run `docker compose up -d` and test your bot instance by sending a `/start` to your TG bot.

    If the bot has replied, you will already have the bot configured and ready to attend any request. The docker compose file of the bot is already configured to work as a service if the Docker daemon is active, restarting if it finds any problem. If you desire to turn it off, just run `docker compose down`.

### Non-standalone installation

#### Installation with the binary (Windows/Ubuntu)
1. Navigate to https://github.com/micheldevs/florobot/releases/
2. Select a `non-standalone` version and download depending your OS the `florobot_win11_amd64.zip` or `florobot_linux_amd64.zip` file
3. Unzip the downloaded zip file in your desired directory
4. Copy the `example.env` file, rename it to `.env` file and set its variables according your preferences. Almost all these vars are already set, only being necessary to change the `TG_BOT_APITOKEN` by your Telegram bot token, and the database settings. The `non-standalone` version needs a Postgres instance and an empty database in order to work, so configure the database settings according your instance. For more information of the env variables, please check the [Bot configuration](#bot-configuration) section.
5. Run the binary of the bot (add execution permissions before in Linux), and test your bot instance by sending a `/start` to your TG bot.

    * `florobot.exe` - In Windows
    * `florobot` - In Linux

    If the bot has replied, you will already have the bot configured and ready to attend any request. If you desire that will attend any petition once your server/computer is booted up, do not forget to configure it as a Windows or Linux service and set the execution path as the place where the bot is! 

#### Installation with Docker
1. Navigate to https://github.com/micheldevs/florobot/releases/
2. Select a `non-standalone` version and download `Source code (zip)` file
3. Unzip the downloaded zip file in your desired directory
4. Open the `docker-compose.yaml` file and set the variables of the `environment` section according your preferences. Almost all these vars are already set, only being necessary to change the `TG_BOT_APITOKEN` by your Telegram bot token. The `non-standalone` version of the bot in Docker includes a Postgres container already configured to work and the internal credentials mapped in the `docker-compose.yaml` file. In case that it is wanted to use a Postgres instance served from another place, just set the database settings according your causistic and only serve the `florobot` service. For more information of the env variables, please check the [Bot configuration](#bot-configuration) section.

5. Open a terminal in the root directory of the bot, run `docker compose up -d` and test your bot instance by sending a `/start` to your TG bot.

    If the bot has replied, you will already have the bot configured and ready to attend any request. The docker compose file of the bot is already configured to work as a service if the Docker daemon is active, restarting if it finds any problem. If you desire to turn it off, just run `docker compose down`.

## Usage
### Before to run the bot
#### Bot configuration
The full configuration of the bot can be changed from its `.env` file or from the environment section of the `docker-compose.yaml` file, depending on the installation method that you have followed.

```
# BOT SETTINGS
# Telegram bot token
TG_BOT_APITOKEN="..."
# Maximum number of retries to send a message
TG_BOT_SEND_NUM_RETRIES=5
# Bot language (en or es only available for now!)
TG_BOT_LANGUAGE="es"
# Print full information of each request and response of the bot
TG_BOT_DEBUG="false"
# optional
# Cron expression to run background tasks of the bot, none background task will be executed by default, so feel free to comment it
TG_BOT_BACKGROUND_TASKS_CRON="*/3 * * * *"

# DATABASE SETTINGS (non-standalone only)
# IP or hostname where the Postgres instance is running
POSTGRES_HOST="localhost" 
# Postgres user of the database (required to have permissions to create and modify tables, and make SELECT, INSERT, DELETE or UPDATE operations)
POSTGRES_USER="postgres"
# Password of the Postgres user
POSTGRES_PASSWORD="secret"
# Postgres database name
POSTGRES_DB="florobot_db"
# Port where the Postgres instance is running
POSTGRES_PORT=5432
```
#### Translations and content
All the content along with the translations of the bot, can be found below the `assets/` folder. This folder has two subfolders:

* `csv/`: Where all the content available of the bot is served from csvs separated by '|', and can be populated with your jokes, mention reactions, and so on.... Details of each file below.
    * `jokes_keywords.csv`: Keywords that the bot has to find in order to interrupt a conversation with a joke, any keyword of this file will trigger a joke. The second column is the list of chatIds where the keyword is blacklisted, separated by commas (e.g. 000001,000002), it will not take effect if this column is left empty.
    * `jokes.csv`: Jokes of the bot, each joke is delimited by double-quotes.
    * `jokes_audios.csv`: Applause audios URLs, to simulate a public in a comedy show when the bot tells a joke. One URL per line and delimited with double quotes.
    * `mentions.csv`: Regular expressions that the bot has to match in a message, in order to reply with a random reaction. The second column is the list of possible reactions, separated by ';;'.
    * `questions.csv`: Replies of the questions asked to the bot. The second column is the type of reply, where:
        * `D`: Dunno type, the bot only sends a message replying to any question asked with the string of the left column.
        * `G`: GoogleURL type, the bot replies the person who has asked a question with the string of the left collumn, and then sends him to a Google URL (provided by `questions_searchs.csv`).
    * `questions_searchs.csv`: Base query URLs to search the question asked to the bot mentioning him. The % symbol is used to escape URL decode chars in the links, such as whitespaces (%20). Place %s where the original question asked by the user will be concatenated to the URL.
    * `roast_keywords.csv`: Keywords that the bot has to find along with a mention to its `@username` in order to mock someone. The second column is the list of chatIds where the keyword is blacklisted, separated by commas (e.g. 000001,000002), it will not take effect if this column is left empty.
    * `roast.csv`: List of roasts to mock someone that has mocked the bot first. The second column is the roast type, where:
        * `D`: Doxxed type, the bot sends the string of the left column and then sends random information, acting like he is doxxing the user that has mocked him.
        * `G`: GIF type, the bot sends the string of the left column and then sends a random gif (provided by `roast_gifs.csv`).
        * `I`: Imitate type, the bot imitates the last message sent by the user that is mocking him, replacing each vowel by an `i`, as it would be imitating the user to make fun of him. After that, he sends the string of the left column to finish the performance.
        * `N`: Normal type, the bot replies with the string of the left column, replying the insult of the user.
        * `P`: Poll type, the bot opens a yes/no poll using the string of the left column.
    
        Use %s in the roasts to use mentions to the user that has mocked the bot.
    * `roast_gifs.csv`: GIF URLs containing the GIFs that the bot will send to mock someone when a roast of type GIF will be triggered.
* `locale`: Where all the translations of the bot are stored as language json files.

Both folders have the same structure, being under them the `en/` or `es/` subfolders with the content csvs or translation jsons for english and spanish languages. 

#### Database files
The database files are created and kept once the bot is initialized by the first time. These database files are stored under the `apps/sqlite` folder for `standalone` installations (florobot.db sqlite file), and `apps/postgres` folder for the `non-standalone` dockerized installation. Keep safe these files if you are going to use the bot with the scraping capabilities or if you have multiple users/chats configured in the Chat table!

### Commands and features
This list corresponds to the available commands of the bot:
```
/start - Says hi!
/joke - Tells a random joke
/insult - Gives information about how to insult the bot
/question - Gives information about how to ask a question to the bot
/rrstart - Starts a russian roulette game
/rrjoinbot - Joins the bot to a started russian roulette game
/rrjoin - Joins the sender to a started russian roulette game
/rrclose - Closes the players of a started russian roulette game (min: 2 playes)
/rroll - Rolls the revolver chamber to play a turn of a russian roulette game
/listen - Asks a chat to listen, allowing to send messages once the chat is being listen. The available chats correspond to groupchats where the bot has entered to or invididual chats started by someone. (option only for admins)
/stoplisten - Stops the listening of the chat previously selected to listen.
/movspremiere - This command does not return any information, it is an example command to illustrate how the bot could be used to obtain information about things using scraping, such as for instance, which movies are playing in the cinema. To learn more, check the Scraping section.
```
In addition to the commands, the bot also can:

1. Tell a random joke (taken from `jokes.csv`) if it listens any joke keyword in a sent message (taken from `jokes_keywords.csv`). It can also post an audio with applauses (taken from `jokes_audios.csv`).
2. Reply to an insult with another one (taken from `roast.csv`) if it listens any insult (taken from `roast_keywords.csv`) along with its `@username`. It can also post a GIF (taken from `roast_gifs.csv`), in case the selected roast is GIF type.
3. Reply to a message with a question started or not by "¿" and ended with "?", asked along with its `@username`. Replies taken from `questions.csv` and URLs to reply questions from `questions_searchs.csv`.
4. Reply to those messages that match any regex of the first column `mentions.csv`, using a random text of the second column of the matched regex.
5. Control the russian roulette game flows, warning or expiring the inactive matches if 10 minutes passes.
6. Executes background tasks, such as sending a notification or scrape some web (see Background tasks section)

### For the developer
TODO

### License
Those ones who make any modification or use for fun without commercial use, can use the MIT license of the bot.

For other uses, this bot is licensed by a AGPL-3.0 license.