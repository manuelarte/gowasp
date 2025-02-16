package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/ing-bank/ginerr/v3"
	"github.com/manuelarte/pagorminator"
	"gowasp/internal/models"
	"gowasp/internal/services"
	"strconv"
)

type BlogCommentsHandler struct {
	BlogCommentService services.BlogCommentService
}

func (h *BlogCommentsHandler) GetBlogComments(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		code, response := ginerr.NewErrorResponse(c, err)
		c.JSON(code, response)
		return
	}
	pageString := c.DefaultQuery("page", "0")
	page, _ := strconv.Atoi(pageString)
	sizeString := c.DefaultQuery("size", "10")
	size, _ := strconv.Atoi(sizeString)
	pageRequest, _ := pagorminator.PageRequest(page, size)
	pageResponse, err := h.BlogCommentService.GetAllForBlog(c, uint(id), pageRequest)
	if err != nil {
		code, response := ginerr.NewErrorResponse(c, err)
		c.JSON(code, response)
		return
	}
	c.JSON(200, pageResponse)
}

func (h *BlogCommentsHandler) CreateBlogComment(c *gin.Context) {
	//id, err := strconv.Atoi(c.Param("id"))
	//if err != nil {
	//	code, response := ginerr.NewErrorResponse(c, err)
	//	c.JSON(code, response)
	//	return
	//}
	blogComment := models.BlogComment{}
	if err := c.BindJSON(&blogComment); err != nil {
		code, response := ginerr.NewErrorResponse(c, err)
		c.JSON(code, response)
		return
	}
	// create
	// vulnerabilities, csrf, template injection, not validating that the user who created comment is the one authenticated
	err := h.BlogCommentService.Create(c, &blogComment)
	if err != nil {
		code, response := ginerr.NewErrorResponse(c, err)
		c.JSON(code, response)
		return
	}
	c.JSON(200, blogComment)
}
