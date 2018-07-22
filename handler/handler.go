package handler

import (
	"crypto/sha1"
	"net/http"
	"sort"

	"github.com/labstack/echo"
)

type (
	Handler struct {
	}
)

func (h *Handler) CheckSignature(c echo.Context) error {
	signature := c.QueryParam("signature")
	timestamp := c.QueryParam("timestamp")
	nonce := c.QueryParam("nonce")
	echostr := c.QueryParam("echostr")
	arr := []string{TOKEN, timestamp, nonce}
	//升序排序
	sort.Strings(arr)
	//拼接字符串
	s := arr[0] + arr[1] + arr[2]
	//sha1加密
	sha := sha1.New()
	_, err := sha.Write([]byte(s))
	if err != nil {
		return err
	}
	shas := string(sha.Sum(nil))
	if shas == signature {
		return c.String(http.StatusOK, echostr)
	}
	return c.NoContent(http.StatusAccepted)
}

func (h *Handler) ReceiveMessage(c echo.Context) error {
	return c.String(http.StatusOK, "hello, welcome huanyu happy house!")
}
