package api

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "strconv"
)

func (c *Client) GetUpdates(updates chan Update) {
    url := TelegramApiUrl + c.token + "/getUpdates"
    var lastUpdateId int

    for {
        resp, err := http.Get(url + "?offset=" + strconv.Itoa(lastUpdateId+1) + "&timeout=30")
        if err != nil {
            continue
        }

        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            continue
        }
        resp.Body.Close()

        fmt.Println("Получен ответ от Telegram API:", string(body))

        var result struct {
            Result []Update `json:"result"`
        }
        err = json.Unmarshal(body, &result)
        if err != nil {
            continue
        }

        for _, update := range result.Result {
            if update.UpdateId > lastUpdateId {
                update.Command = update.Message.Text
                chatID := update.Message.Chat.Id
                fmt.Println("Получен chat_id:", chatID)
                updates <- update
                lastUpdateId = update.UpdateId
            }
        }
    }
}