package bot

import (
    "fmt"
    "sync"
    "go-telegram-bot/src/api"
)

var (
    globalRouter *Router
    once        sync.Once
)

type Router struct {
    handlers map[string]func(api.Context)
    mu       sync.RWMutex
}

func GetRouter() *Router {
    once.Do(func() {
        globalRouter = &Router{
            handlers: make(map[string]func(api.Context)),
        }
    })
    return globalRouter
}

func (r *Router) AddHandler(command string, handler func(api.Context)) {
    r.mu.Lock()
    defer r.mu.Unlock()
    fmt.Printf("Регистрируем команду: %s\n", command)
    r.handlers[command] = handler
}

func (r *Router) Dispatch(ctx api.Context) {
    r.mu.RLock()
    handler, exists := r.handlers[ctx.Update.Command]
    r.mu.RUnlock()

    if exists {
        go handler(ctx)
    }
}


func Command(command string) func(func(api.Context)) {
    return func(handler func(api.Context)) {
        GetRouter().AddHandler(command, handler)
    }
}