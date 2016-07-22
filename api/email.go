package api

import (
	"fmt"
	"net/http"
	"os"

	"maildeliver/service"
	"maildeliver/utils"

	"github.com/gin-gonic/gin"
)

type Email struct {
	Err interface{}
}

func (c *Email) send(context *gin.Context) {
	formatMsgs, err := c.initSendMessage(context)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"message": err,
		})
	}
	service.AllService.Email.Send(formatMsgs)
}

func (c *Email) initSendMessage(context *gin.Context) (formatMsg []service.Message, err error) {
	var formatMsgs []service.Message
	if err = context.BindJSON(&formatMsgs); err == nil {
		return formatMsgs, nil
	} else {
		return []service.Message{}, err
	}
}

func (c *Email) formatSendMessage(formatMsgs []service.Message) []service.Message {
	var tmpMsgQueue []service.Message
	for k, formatMsg := range formatMsgs {
		if formatMsg.IsSplitTo {
			var dest []string
			if len(formatMsg.To) > 1 {
				for _, v := range formatMsg.To {

				}
			}
		}
		delete(formatMsgs, k)
	}
}

func (c *Email) decideContentType(content_type int) string {

}

func initEmail(r *gin.Engine) {
	emailController := &Email{}
	emailed := r.Group("/email")
	emailed.POST("/send", emailController.send)
}

func NewAppError(where, message string, code int) *AppError {
	return &AppError{where, message, code}
}

type AppError struct {
	Where      string
	Message    string
	StatusCode int
}

func recoverError(w http.ResponseWriter) {
	if err := recover(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.WriteError(w, http.StatusInternalServerError, err.(error).Error())
	}
}
