package rest

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/manuelarte/pagorminator"
	"github.com/manuelarte/ptrutils"

	"github.com/manuelarte/gowasp/internal/models"
	"github.com/manuelarte/gowasp/internal/posts"
	"github.com/manuelarte/gowasp/internal/sliceutils"
)

type PostsHandler struct {
	service posts.Service
}

func NewPosts(service posts.Service) *PostsHandler {
	return &PostsHandler{
		service: service,
	}
}

func (h *PostsHandler) GetPostById(c *gin.Context, postID uint) {
	post, err := h.service.GetByID(c, postID)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Code:    http.StatusBadRequest,
			Details: err,
			Message: "Error retrieving the post",
		})
	}

	c.JSON(http.StatusOK, postToDto(&post))
}

func (h *PostsHandler) GetPosts(c *gin.Context, params GetPostsParams) {
	pageRequest, err := pagorminator.PageRequest(
		ptrutils.DerefOr(params.Page, 0),
		ptrutils.DerefOr(params.Size, defaultPageRequestSize),
		orderFrom(ptrutils.DerefOr(params.Sort, PostedAtdesc)),
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

func orderFrom(sortingCriteria GetPostsParamsSort) pagorminator.Order {
	switch sortingCriteria {
	case PostedAtasc:
		return pagorminator.MustOrder("posted_at", pagorminator.ASC)
	case PostedAtdesc:
		return pagorminator.MustOrder("posted_at", pagorminator.DESC)
	case Titleasc:
		return pagorminator.MustOrder("title", pagorminator.ASC)
	case Titledesc:
		return pagorminator.MustOrder("title", pagorminator.DESC)
	default:
		return pagorminator.MustOrder("posted_at", pagorminator.DESC)
	}
}

func postPageRequestToDTO(posts []*models.Post, pageRequest *pagorminator.Pagination) PagePosts {
	return PagePosts{
		UnderscoreMetadata: PageMetadata{
			Page:       pageRequest.GetPage(),
			Size:       pageRequest.GetSize(),
			TotalCount: int(pageRequest.GetTotalElements()),
			TotalPages: pageRequest.GetTotalPages(),
		},
		Data: sliceutils.Transform(posts, postToDto),
	}
}

func postToDto(post *models.Post) Post {
	return Post{
		Self:      Paths{}.GetPostByIdEndpoint.Path(strconv.Itoa(int(post.ID))),
		Content:   post.Content,
		CreatedAt: post.CreatedAt,
		ID:        post.ID,
		PostedAt:  post.PostedAt,
		Title:     post.Title,
		UserID:    post.UserID,
		UpdatedAt: post.UpdatedAt,
	}
}
