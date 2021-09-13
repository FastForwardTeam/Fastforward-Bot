const Discord = require('discord.io'),
  auth = require('./auth.json'),
  client = new Discord.Client();

client.on('ready', () => {
  console.log(`Logged in as ${client.user.tag}!`);
  console.log(client.user.presence.status);
})

client.on('message', msg => {

})

client.login(auth.token);
