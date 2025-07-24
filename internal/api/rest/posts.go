package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manuelarte/pagorminator"
	"github.com/manuelarte/ptrutils"

	"github.com/manuelarte/gowasp/internal/models"
	"github.com/manuelarte/gowasp/internal/posts"
	"github.com/manuelarte/gowasp/internal/sliceutils"
)

type Posts struct {
	service posts.Service
}

func NewPosts(service posts.Service) *Posts {
	return &Posts{
		service: service,
	}
}

func (h *Posts) GetPosts(c *gin.Context, params GetPostsParams) {
	pageRequest, err := pagorminator.PageRequest(
		ptrutils.DerefOr(params.Page, 0),
		ptrutils.DerefOr(params.Size, defaultPageRequestSize),
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Code:    http.StatusBadRequest,
			Details: err,
			Message: "Error creating the page request",
		})

		return
	}
	postPage, err := h.service.GetAll(c, pageRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Code:    http.StatusBadRequest,
			Details: err,
			Message: "Error retrieving the posts page",
		})

		return
	}
	dto := postPageRequestToDTO(postPage, pageRequest)
	c.JSON(http.StatusOK, dto)
}

func postPageRequestToDTO(posts []*models.Post, pageRequest *pagorminator.Pagination) PagePosts {
	return PagePosts{
		UnderscoreMetadata: PageMetadata{
			Page:       pageRequest.GetPage(),
			Size:       pageRequest.GetSize(),
			TotalCount: int(pageRequest.GetTotalElements()),
			TotalPages: pageRequest.GetTotalPages(),
		},
		Data: sliceutils.Transform(posts, func(x *models.Post) Post {
			return Post{
				Content:   x.Content,
				CreatedAt: x.CreatedAt,
				ID:        x.ID,
				PostedAt:  x.PostedAt,
				Title:     x.Title,
				UserID:    x.UserID,
				UpdatedAt: x.UpdatedAt,
			}
		},
		),
	}
}
