# Go Telegram Bot Framework

Легкий и удобный фреймворк для создания Telegram ботов на Go с поддержкой асинхронной обработки команд и middleware.

## Особенности

- Простой и понятный API
- Асинхронная обработка сообщений
- Поддержка команд и обработчиков
- Встроенный роутер для команд
- Поддержка редактирования сообщений
- Контекстная система для обработки запросов

## Установка

```bash
go get github.com/yourusername/go-telegram-bot
```

## Быстрый старт

1. Создайте нового бота через [@BotFather](https://t.me/BotFather) и получите токен

2. Создайте базовый бот:

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
        fmt.Printf("Ошибка при запуске бота: %v\n", err)
    }
}
```

3. Добавьте обработчики команд:

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
    ctx.SendMessage("Привет! Это тестовый бот 🤖")
}

func HelpHandler(ctx api.Context) {
    ctx.SendMessage("Доступные команды:\n/start - Приветствие\n/help - Список команд")
}
```

## Асинхронная обработка

Фреймворк поддерживает асинхронную обработку команд с использованием горутин:

```go
func AsyncHandler(ctx api.Context) {
    msgID, _ := ctx.SendMessage("начинаю работу...")

    go func() {
        time.Sleep(2 * time.Second)
        ctx.EditMessage(msgID, "первая задача готова!")
    }()

    go func() {
        time.Sleep(3 * time.Second)
        ctx.SendMessage("вторая задача готова!")
    }()
}
```

## Структура проекта

```
go-telegram-bot/
├── src/
│   ├── api/
│   │   ├── client.go    # Telegram API клиент
│   │   ├── context.go   # Контекст обработчика
│   │   └── updates.go   # Обработка обновлений
│   ├── bot/
│   │   ├── bot.go       # Основной класс бота
│   │   └── router.go    # Роутер команд
│   └── handlers/
│       └── handlers.go  # Обработчики команд
└── main.go              # Точка входа
```

## API Reference

### Создание бота

```go
bot := bot.New(token)
```

### Регистрация команд

```go
bot.Command("/command")(handler)
```

### Контекст обработчика

```go
type Context struct {
    client *Client
    Update Update
}

// Отправка сообщения
messageId, err := ctx.SendMessage(text)

// Редактирование сообщения
err := ctx.EditMessage(messageId, newText)
```

## Лицензия

MIT

## Вклад в развитие

Мы приветствуем ваш вклад в развитие проекта! Пожалуйста, создавайте issue или отправляйте pull request с улучшениями.

## Авторы

- [@aleksfolt](https://t.me/zxcfolt), [@TheRevil](https://t.me/TheRevil)

## Поддержка

Если у вас возникли вопросы или проблемы:

1. Создайте issue в репозитории
2. Свяжитесь с автором через [Telegram](https://t.me/zxcfolt)