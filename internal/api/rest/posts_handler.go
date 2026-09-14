package rest

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golaxo/goqrius"
	"github.com/manuelarte/pagorminator/pagegeneric"
	"github.com/manuelarte/pagorminator/pagepagination"
	"github.com/manuelarte/ptrutils"

	"github.com/manuelarte/gowasp/internal/models"
	"github.com/manuelarte/gowasp/internal/posts"
	"github.com/manuelarte/gowasp/internal/sliceutils"
)

type PostsHandler struct {
	service posts.Service
}

func NewPosts(service posts.Service) PostsHandler {
	return PostsHandler{
		service: service,
	}
}

func (h PostsHandler) GetPostByID(c *gin.Context, postID uint) {
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

func (h PostsHandler) GetPosts(c *gin.Context, params GetPostsParams) {
	q, err := goqrius.Parse(ptrutils.DerefOr(params.Q, ""))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Code:    http.StatusBadRequest,
			Details: err.Error(),
			Message: "Error parsing the q parameter",
		})

		return
	}

	pageRequest, err := pagepagination.New(
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
	postPage, err := h.service.GetAll(c, q, pageRequest)
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

func orderFrom(sortingCriteria GetPostsParamsSort) pagegeneric.Order {
	switch sortingCriteria {
	case PostedAtasc:
		return pagegeneric.Asc("posted_at")
	case PostedAtdesc:
		return pagegeneric.Desc("posted_at")
	case Titleasc:
		return pagegeneric.Asc("title")
	case Titledesc:
		return pagegeneric.Desc("title")
	default:
		return pagegeneric.Desc("posted_at")
	}
}

func postPageRequestToDTO(posts []*models.Post, pageRequest *pagepagination.Pagination) PagePosts {
	totalElements, _ := pageRequest.TotalElements()

	return PagePosts{
		UnderscoreMetadata: PageMetadata{
			Page:       pageRequest.Page(),
			Size:       pageRequest.Size(),
			TotalCount: int(totalElements),
			TotalPages: pageRequest.TotalPages(),
		},
		Data: sliceutils.Transform(posts, postToDto),
	}
}

func postToDto(post *models.Post) Post {
	return Post{
		//#nosec G115
		Self:      Paths{}.GetPostByIDEndpoint.Path(strconv.Itoa(int(post.ID))),
		Content:   post.Content,
		CreatedAt: post.CreatedAt,
		ID:        post.ID,
		PostedAt:  post.PostedAt,
		Title:     post.Title,
		UserID:    post.UserID,
		UpdatedAt: post.UpdatedAt,
	}
}
