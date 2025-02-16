package handlers

import (
    "go-telegram-bot/src/api"
    "go-telegram-bot/src/bot"
    "log"
	"time"
)

func init() {
    log.Println("регистрация обработчиков...")
    bot.Command("/ping")(AsyncHandler)
}

func AsyncHandler(ctx api.Context) {
    msgID, err := ctx.SendMessage("начинаю работу...")
    if err != nil {
        log.Println("ошибка отправки сообщения:", err)
        return
    }

    go func() {
        time.Sleep(2 * time.Second)
        ctx.EditMessage(msgID, "первая задача готова!")
    }()

    go func() {
        time.Sleep(3 * time.Second)
        ctx.SendMessage("вторая задача готова!")
    }()
}
