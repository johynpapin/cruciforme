package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/johynpapin/cruciforme/server/store"
)

type formResponse struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Verified bool   `json:"verified"`
}

type getFormsResponse struct {
	Forms []formResponse `json:"forms"`
}

type createFormRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

type createFormResponse struct {
	Form formResponse `json:"form"`
}

func (h *Handlers) GetFormsHandler(c *gin.Context) {
	userId := c.GetString("userId")

	forms, err := h.Store.GetFormsByUserId(userId)
	if err != nil {
		c.Error(fmt.Errorf("could not get the forms of the user: %w", err))
		c.Abort()
		return
	}

	response := &getFormsResponse{
		Forms: []formResponse{},
	}

	for _, form := range forms {
		response.Forms = append(response.Forms, formResponse{
			Id:       form.Id,
			Name:     form.Name,
			Email:    form.Email,
			Verified: form.Verified,
		})
	}

	c.JSON(http.StatusOK, response)
}

func (h *Handlers) CreateFormHandler(c *gin.Context) {
	var request createFormRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		NewAPIError(InvalidPayloadError, err).AbortContext(c)
		return
	}

	userId := c.GetString("userId")
	emailIsVerified := request.Email == userId

	form := &store.Form{
		Name:     request.Name,
		UserId:   c.GetString("userId"),
		Email:    request.Email,
		Verified: emailIsVerified,
	}

	if err := h.Store.CreateForm(form); err != nil {
		c.Error(fmt.Errorf("could not store the form: %w", err))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, &createFormResponse{
		Form: formResponse{
			Id:       form.Id,
			Name:     form.Name,
			Email:    request.Email,
			Verified: emailIsVerified,
		},
	})
}
