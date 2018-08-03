package handler

import (
	"github.com/l1huanyu/suren"
	"github.com/labstack/echo"
	"net/http"
	"time"
	"fmt"
)

type (
	Handler struct {
		s *suren.Suren
	}
)

func New() *Handler {
	h := new(Handler)
	h.s = suren.New(APPID, SECRET, TOKEN)
	return h
}

func (h *Handler) ResponseWeChat(c echo.Context) error {
	echostr := c.QueryParam("echostr")
	if ok, err := h.s.CheckSignature(&suren.Signature{
		Signature: c.QueryParam("signature"),
		Timestamp: c.QueryParam("timestamp"),
		Nonce:     c.QueryParam("nonce"),
		Echostr:   echostr,
	}); ok && err != nil {
		return c.String(http.StatusOK, echostr)
	}
	return c.NoContent(http.StatusAccepted)
}

func (h *Handler) ReceiveMessage(c echo.Context) error {
	msgRx := new(suren.TextMsgRx)
	err := c.Bind(msgRx)
	if err != nil {
		return err
	}
	msgTx := &suren.TextMsgTx{
		ToUserName:   msgRx.FromUserName,
		FromUserName: msgRx.ToUserName,
		CreateTime:   int(time.Now().Unix()),
		MsgType:      suren.TEXT,
		Content:      fmt.Sprintf("收到消息\"%s\"。", msgRx.Content),
	}
	return c.XML(http.StatusOK, msgTx)
}