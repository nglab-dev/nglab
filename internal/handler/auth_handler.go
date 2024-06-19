package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nglab-dev/nglab/web/views"
)

func LoginViewHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "", views.LoginIndex(views.AuthIndexPageData{}))
}

func LoginHandler(c *gin.Context) {
	// TODO: 登录逻辑
	c.Redirect(http.StatusFound, "/")
}
