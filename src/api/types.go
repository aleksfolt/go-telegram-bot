package api

type Update struct {
    UpdateId int `json:"update_id"`
    Message  struct {
        MessageId int64  `json:"message_id"`
        Text      string `json:"text"`
        Chat      struct {
            Id int64 `json:"id"`
        } `json:"chat"`
    } `json:"message"`
    Command string
}

type Context struct {
    client *Client
    Update Update
}

func NewContext(client *Client, update Update) Context {
    return Context{
        client: client,
        Update: update,
    }
}

func (ctx Context) SendMessage(text string) (int64, error) {
    messageId, err := ctx.client.SendMessage(ctx.Update.Message.Chat.Id, text)
    if err != nil {
        return 0, err
    }
    return messageId, nil
}

func (ctx Context) EditMessage(messageId int64, text string) error {
    return ctx.client.EditMessage(ctx.Update.Message.Chat.Id, messageId, text)
}
