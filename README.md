# Intro
### Welcome to the [@GetGeniusBot](https://t.me/GetGeniusBot) official repository!
You might wonder what does this bot do. Basically, gather track information on [Genius](https://genius.com). The original version was in JavaScript but I've changed my mind and re-did it in go.
Some features may be missing from the original version.
Using it is pretty simple: just type `@GetGeniusBot` followed by the name of the track you're looking for. Currently you can get track bios and streaming platforms.
Lyrics are not supported as they're a hard thing do deal with.

# Locally testing
This bot requires 3 tokens to run:
- Your Telegram bot token, through [@BotFather](https://t.me/BotFather)
- A [Genius API token](https://genius.com/api-clients)
- A [SongLink](https://odesli.co/) token. You gotta e-mail them for that, and for that reason this token is optional. **Just make sure to take off `&key=<key>` at `/src/handler/handle_callback`.**

## Config file
All those tokens described above must be in a config.json file at the root of the project. Also look for the `config.json.example` for more keys to add. *Yes, they're required too*. Once that's done the bot should work normally.

# Contributing
All kind of contributions are welcome! Just make sure to use the right buttons when providing feedback, such as the Issue tab for issues and the Pull Request tab for improvements/feature requests.
