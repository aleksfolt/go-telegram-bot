package bot

import (
    "go-telegram-bot/src/api"
)

type Bot struct {
    client *api.Client
    router *Router
}

func New(token string) *Bot {
    return &Bot{
        client: api.NewClient(token),
        router: GetRouter(),
    }
}

func Start(token string) error {
    bot := New(token)
    
    updates := make(chan api.Update)
    go bot.client.GetUpdates(updates)

    for update := range updates {
        ctx := api.NewContext(bot.client, update)
        go bot.router.Dispatch(ctx)
    }
    return nil
}