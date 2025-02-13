package handlers

import (
	"bufio"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ing-bank/ginerr/v3"
	"io"
	"os"
)

type BlogsHandler struct {
}

func (h *BlogsHandler) GetBlogFileByName(c *gin.Context) {
	name := c.Query("name")
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
