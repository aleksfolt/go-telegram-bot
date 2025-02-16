package api

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
    "net/url"
)

const TelegramApiUrl = "https://api.telegram.org/bot"

type Client struct {
    token string
}

type MessageResponse struct {
    Ok     bool `json:"ok"`
    Result struct {
        MessageId int64 `json:"message_id"`
    } `json:"result"`
}

func NewClient(token string) *Client {
    return &Client{
        token: token,
    }
}

func (c *Client) SendMessage(chatId int64, text string) (int64, error) {
    response, err := http.PostForm(
        fmt.Sprintf("%s%s/sendMessage", TelegramApiUrl, c.token),
        url.Values{
            "chat_id": {fmt.Sprintf("%d", chatId)},
            "text":    {text},
        },
    )
    if err != nil {
        return 0, err
    }
    defer response.Body.Close()

    var messageResponse MessageResponse
    if err := json.NewDecoder(response.Body).Decode(&messageResponse); err != nil {
        return 0, err
    }

    if !messageResponse.Ok {
        return 0, fmt.Errorf("failed to send message")
    }

    return messageResponse.Result.MessageId, nil
}

func (c *Client) EditMessage(chatID int64, messageID int64, text string) error {
    fmt.Println("Попытка редактирования сообщения в чате:", chatID, "ID сообщения:", messageID)

    url := fmt.Sprintf("%s%s/editMessageText", TelegramApiUrl, c.token)

    body, _ := json.Marshal(map[string]interface{}{
        "chat_id":    chatID,
        "message_id": messageID,
        "text":       text,
    })

    fmt.Println("Запрос в Telegram API:", url)
    fmt.Println("Тело запроса:", string(body))

    resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
    if err != nil {
        fmt.Println("Ошибка при редактировании сообщения:", err)
        return err
    }
    defer resp.Body.Close()

    fmt.Println("Статус-код ответа от Telegram API:", resp.StatusCode)

    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("ошибка редактирования: Telegram API вернул код %d", resp.StatusCode)
    }

    fmt.Println("✅ Сообщение успешно отредактировано в чате", chatID)
    return nil
}
