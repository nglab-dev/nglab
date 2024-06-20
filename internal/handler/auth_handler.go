package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nglab-dev/nglab/internal/model"
	"github.com/nglab-dev/nglab/pkg/v"
	"github.com/nglab-dev/nglab/web/views"
)

var loginSchema = v.Schema{
	"username": v.Rules(v.Min(5), v.Max(20)),
	"password": v.Rules(v.Required),
}

var signupSchema = v.Schema{
	"username": v.Rules(v.Min(5), v.Max(20)),
	"password": v.Rules(
		v.ContainsSpecial,
		v.ContainsUpper,
		v.Min(7),
		v.Max(50),
	),
}

func HandleLoginView(c *gin.Context) {
	c.HTML(http.StatusOK, "", views.LoginView(views.AuthIndexPageData{}))
}

func HandleSignupView(c *gin.Context) {
	c.HTML(http.StatusOK, "", views.SignupView(views.SignupIndexPageData{}))
}

func HandleLogin(c *gin.Context) {
	req := &model.LoginRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Redirect(http.StatusFound, "/")
}

func HandleSignup(c *gin.Context) {
	var values views.SignupFormValues
	errors, ok := v.Request(c.Request, &values, signupSchema)

	fmt.Printf("errors: %v\n", errors)

	if !ok {
		c.HTML(http.StatusOK, "", views.SignupForm(values, errors))
		return
	}
}
