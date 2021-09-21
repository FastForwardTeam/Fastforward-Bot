// compile to js with `npm run tsc`
// ts is better lol


import DiscordJS, { Intents } from 'discord.js'
import dotenv from 'dotenv'
dotenv.config()

const client = new DiscordJS.Client({
    intents: [
        Intents.FLAGS.GUILDS,
        Intents.FLAGS.GUILD_MESSAGES
    ]
})

client.on('ready', () => {
    console.log('Bot Online!')
})

client.on('messageCreate', (message) => {
    if (message.content === "ping") {
        message.reply({
            content: 'Pong!',
        })
    }
})






client.login(process.env.TOKEN)