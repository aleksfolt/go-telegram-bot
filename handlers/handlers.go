package handlers

import (
    "go-telegram-bot/src/api"
    "go-telegram-bot/src/bot"
    "log"
)

func init() {
    log.Println("Регистрация обработчиков...")
    bot.Command("/start")(StartHandler)
    bot.Command("/help")(HelpHandler)
}

func StartHandler(ctx api.Context) {
    ctx.SendMessage("Привет! Это тестовый бот 🤖")
}

func HelpHandler(ctx api.Context) {
    i := 0
    for i < 5 {
        ctx.SendMessage("Я умею выполнять команды!\n\nДоступные команды:\n/start - Приветствие\n/help - Список команд")
        i++
    }
}