package api

import (
	"net/http"

	"potato/service"
	"potato/utils"

	"github.com/gin-gonic/gin"
)

type Email struct {
	Err interface{}
}

func (c *Email) send(context *gin.Context) {
	formatMsg, err := c.initSendMessage(context)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"message": err,
		})
	}
	service.AllService.Email.Send(formatMsg)
}

func (c *Email) initSendMessage(context *gin.Context) (formatMsg service.Message, err error) {
	if err = context.BindJSON(&formatMsg); err == nil {
		return formatMsg, nil
	} else {
		return service.Message{}, err
	}
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
