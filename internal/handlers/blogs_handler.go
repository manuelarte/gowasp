package handlers

import (
	"bufio"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ing-bank/ginerr/v3"
	"github.com/manuelarte/pagorminator"
	"gowasp/internal/services"
	"io"
	"net/http"
	"os"
	"strconv"
)

type BlogsHandler struct {
	BlogService services.BlogService
}

func (h *BlogsHandler) GetOnePage(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		code, response := ginerr.NewErrorResponse(c, err)
		c.JSON(code, response)
		return
	}
	blog, err := h.BlogService.GetById(c, id)
	if err != nil {
		code, response := ginerr.NewErrorResponse(c, err)
		c.JSON(code, response)
		return
	}
	c.HTML(http.StatusOK, "blogs/one.tpl", gin.H{"blog": blog})
}

func (h *BlogsHandler) GetStaticBlogFileByName(c *gin.Context) {
	name := c.Query("name")
	// CWE-918: Server-Side Request Forgery (SSRF) https://cwe.mitre.org/data/definitions/918.html
	file, err := os.Open(fmt.Sprintf("./resources/blogs/%s", name))
	if err != nil {
		code, response := ginerr.NewErrorResponse(c, err)
		c.JSON(code, response)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			code, response := ginerr.NewErrorResponse(c, err)
			c.JSON(code, response)
			return
		}
	}(file)
	// Get the file size
	stat, err := file.Stat()
	if err != nil {
		code, response := ginerr.NewErrorResponse(c, err)
		c.JSON(code, response)
		return
	}

	// Read the file into a byte slice
	bs := make([]byte, stat.Size())
	_, err = bufio.NewReader(file).Read(bs)
	if err != nil && err != io.EOF {
		code, response := ginerr.NewErrorResponse(c, err)
		c.JSON(code, response)
		return
	}

	c.Data(200, "text/plain", bs)
}

func (h *BlogsHandler) GetAll(c *gin.Context) {
	pageString := c.DefaultQuery("page", "0")
	page, _ := strconv.Atoi(pageString)
	sizeString := c.DefaultQuery("size", "10")
	size, _ := strconv.Atoi(sizeString)
	pageRequest, _ := pagorminator.PageRequest(page, size)
	pageBlogsResponse, err := h.BlogService.GetAll(c, pageRequest)
	if err != nil {
		code, response := ginerr.NewErrorResponse(c, err)
		c.JSON(code, response)
		return
	}
	c.JSON(200, pageBlogsResponse)
}
