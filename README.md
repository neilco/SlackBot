# SlackBot
**The simplest way to create Slack bots**

## Installation

```
go get github.com/neilco/slackbot
```

## Getting Started

Create a new Slack Incoming WebHook integration and copy the `Token`.

```
package main

import (
	"github.com/neilco/slackbot"
)

func main() {
	bot := slack.NewBot("<Slack Team name>", "<Slack Integration Token>")
	bot.PublishText("Hello World")
}
```

### Changing the name & icon of the bot on the fly

```
bot.Name = "octobot"
bot.IconURL = "http://example.com/octobot.png"
```

### Sending a message to a specific channel

```
message := NewMessage("#development", false)
message.Text = "Hello Developers"

bot.Publish(message)
```

### Sending an auto-expanding hyperlinked message


```
message := NewMessage("#development", true)
message.Text = "Check out <this Go package|http://github.com/neilco/SlackBot>!"

bot.Publish(message)
```

For details on how you can configure the message posted to Slack, see the [`Message`](https://github.com/neilco/SlackBot/blob/master/slackbot.go#L24) type in the source.

### Contact

[Neil Cowburn](http://github.com/neilco)  
[@neilco](https://twitter.com/neilco)

## License

[MIT license](http://neil.mit-license.org)

Copyright (c) 2014 Neil Cowburn (http://github.com/neilco/)

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
