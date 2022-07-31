<div align="center">
<h1> <img src="https://fastforward.team/img/branding.png" width="400" /> </h1>
<p> <h3> <strong> FastForward Bot: Our Discord Toolbox </strong> </h3> </p>
<p align="center">
<img alt="GitHub Workflow Status" src="https://img.shields.io/github/workflow/status/FastForwardTeam/Fastforward-Bot/Go?label=Status&style=for-the-badge">
<a href="https://discord.com/invite/RSAf7b5njt"><img alt="Discord" src="https://img.shields.io/discord/876622516607656006?label=Join%20our%20discord&logo=discord&style=for-the-badge"></a>
</p>
<p> <a href="#about">About</a> | <a href="#compiling--running">Compiling & Running</a> | <a href="#support">Support</a> </p>
</div>


# About
This is the FastForward utility box, a bot used mainly on our Discord server. It contains some functions such has verifying string hashes & more. Written in Golang.

Made specific for slash (/) commands.

# Compiling & Running

## Prerequisites

- [Go](https://go.dev/dl) ≥ 1.18.4
- [Node.js](https://nodejs.org/en/download/) ≥ 16.16.0
- [Npm](https://www.npmjs.com/package/npm) ≥ 8.12.1
- [hash-detector-cli](https://www.npmjs.com/package/hash-detector-cli) ≥ 1.2.6

## Preparing the environment
1. Install Go, Node.js and Npm.
The instructions are on their linked pages.

2. Install hash-detector-cli with Npm
```sh-session
npm install -g hash-detector-cli
```

3. Clone this repository
```sh-session
git clone git@github.com:fastforwardteam/fastforward-bot.git
cd fastforward-bot
```

4. Edit the .env.example file and rename it to .env
```sh-session
# Use your preferred editor.
nano .env.example
# Then, rename it
mv .env.example .env
```

5. Run or compile it.  
To run, use `go run`
```sh-session
go run .
```

To build, use `go build`
```sh-session
go build
```

Go should automatically fetch all dependencies.

# Support
If you need help, join our discord by clicking on the badge or [here](https://discord.com/invite/RSAf7b5njt).
