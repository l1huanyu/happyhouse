package handler

import (
	"crypto/sha1"
	"net/http"
	"sort"

	"github.com/labstack/echo"
	"strings"
	"github.com/labstack/gommon/log"
	"fmt"
	"io"
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
	sl := []string{TOKEN, timestamp, nonce}
	//升序排序
	sort.Strings(sl)
	//sha1加密
	s := sha1.New()
	_, err := io.WriteString(s, strings.Join(sl, ""))
	if err != nil {
		log.Error("io.WriteString Error: " + err.Error())
		return c.NoContent(http.StatusInternalServerError)
	}
	localSignature := string(s.Sum(nil))
	if localSignature == signature {
		return c.String(http.StatusOK, echostr)
	}
	log.Info(fmt.Sprintf("Check Signature Error: signature = %s, timestamp = %s, nonce = %s, echostr = %s, " +
		"localSignature = %s", signature, timestamp, nonce, echostr, localSignature))
	return c.NoContent(http.StatusAccepted)
}

func (h *Handler) ReceiveMessage(c echo.Context) error {
	return c.String(http.StatusOK, "hello, welcome huanyu happy house!")
}
