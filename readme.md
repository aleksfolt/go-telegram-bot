# Go Telegram Bot Framework

A lightweight and convenient framework for creating Telegram bots in Go with support for asynchronous command handling and middleware.

## Features

- Simple and intuitive API
- Asynchronous message processing
- Command and handler support
- Built-in command router
- Message editing support
- Context-based request handling system

## Installation

```bash
go get github.com/yourusername/go-telegram-bot
```

## Quick Start

1. Create a new bot through [@BotFather](https://t.me/BotFather) and get the token

2. Create a basic bot:

```go
package main

import (
    "fmt"
    "go-telegram-bot/src/bot"
    _ "go-telegram-bot/handlers"
)

func main() {
    token := "YOUR_BOT_TOKEN"
    fmt.Println("Bot started")
    
    if err := bot.Start(token); err != nil {
        fmt.Printf("Error starting bot: %v\n", err)
    }
}
```

3. Add command handlers:

```go
package handlers

import (
    "go-telegram-bot/src/api"
    "go-telegram-bot/src/bot"
)

func init() {
    bot.Command("/start")(StartHandler)
    bot.Command("/help")(HelpHandler)
}

func StartHandler(ctx api.Context) {
    ctx.SendMessage("Hello! This is a test bot 🤖")
}

func HelpHandler(ctx api.Context) {
    ctx.SendMessage("Available commands:\n/start - Greeting\n/help - Command list")
}
```

## Asynchronous Processing

The framework supports asynchronous command processing using goroutines:

```go
func AsyncHandler(ctx api.Context) {
    msgID, _ := ctx.SendMessage("starting work...")

    go func() {
        time.Sleep(2 * time.Second)
        ctx.EditMessage(msgID, "first task completed!")
    }()

    go func() {
        time.Sleep(3 * time.Second)
        ctx.SendMessage("second task completed!")
    }()
}
```

## Project Structure

```
go-telegram-bot/
├── src/
│   ├── api/
│   │   ├── client.go    # Telegram API client
│   │   ├── context.go   # Handler context
│   │   └── updates.go   # Update processing
│   ├── bot/
│   │   ├── bot.go       # Main bot class
│   │   └── router.go    # Command router
│   └── handlers/
│       └── handlers.go  # Command handlers
└── main.go              # Entry point
```

## API Reference

### Creating a Bot

```go
bot := bot.New(token)
```

### Registering Commands

```go
bot.Command("/command")(handler)
```

### Handler Context

```go
type Context struct {
    client *Client
    Update Update
}

// Sending a message
messageId, err := ctx.SendMessage(text)

// Editing a message
err := ctx.EditMessage(messageId, newText)
```

## License

MIT

## Contributing

We welcome your contributions to the project! Please create an issue or submit a pull request with improvements.

## Authors

- [@aleksfolt](https://t.me/zxcfolt), [@TheRevil](https://t.me/TheRevil)

## Support

If you have any questions or issues:

1. Create an issue in the repository
2. Contact the author via [Telegram](https://t.me/zxcfolt)