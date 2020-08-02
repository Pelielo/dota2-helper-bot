# Dota 2 Helper Bot

This is a Discord bot with helper functions for Dota 2 players. To add the bot, go to [Discord's bot page](https://top.gg/bot/739172033606451230).

## Commands

The available commands are:
- `-help`: Shows the available list of commands.
- `-toss`: Tosses a coin and outputs heads or tails.
- `-roll`: Randomly chooses a value between two numbers. Defaults to 0-100.
- `-lobby`: Creates a lobby with 10-12 randomly chosen players.
- `-lobby-roles`: Creates a lobby with 10-12 randomly chosen players and assigns each of them a role.


## Examples

1. **Toss a coin**

`-toss`

> tails

2. **Roll a number**

`-roll`

> 45

3. **Randomize a lobby for 10 players**

`-lobby Sumail Dendi Arteezy N0tail Puppey Miracle s4 Ceb fy XBOCT`

> **The Radiant**
> ```
> XBOCT
> Sumail
> Ceb
> Puppey
> s4
> ```
> **The Dire**
> ```
> fy
> Arteezy
> N0tail
> Dendi
> Miracle
> ```

4. **Randomize a lobby for 12 players with roles**

`-lobby-roles Sumail Dendi Arteezy N0tail Puppey Miracle s4 Ceb fy XBOCT w33 Peli`

> **The Radiant**
> ```
> Off lane - Sumail
> Hard support - Ceb
> Soft support - fy
> Safe lane - Dendi
> Mid lane - Arteezy
> Coach - w33
> ```
> **The Dire**
> ```
> Off lane - s4
> Hard support - N0tail
> Soft support - Miracle
> Safe lane - Puppey
> Mid lane - Peli
> Coach - XBOCT
> ```
