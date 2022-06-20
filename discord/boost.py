from commands.calculator import load_calculator
from commands.nonGames import loadNonGamesCommands
import hikari
import lightbulb

from commands.mastermind import load_mastermind_commands


with open("./discord/TOKEN.txt", "r") as f:
    TOKEN = f.readline().strip()

bot = lightbulb.BotApp(
    TOKEN,
    default_enabled_guilds=(755001834418208840,),
    banner=None
)


@bot.listen(hikari.StartingEvent)
async def load_commands(event):
    loadNonGamesCommands(bot)
    load_calculator(bot)
    load_mastermind_commands(bot)


@bot.listen(hikari.StartedEvent)
async def on_ready(event) -> None:
    print(
        f"{bot.get_me()} - [python version] has connected to Discord! ( ﾉ ﾟｰﾟ)ﾉ"
    )


if __name__ == "__main__":
    bot.run(
        activity=hikari.Activity(
            type=hikari.ActivityType.WATCHING, name="Avishai doing the magic..."
        )
    )
