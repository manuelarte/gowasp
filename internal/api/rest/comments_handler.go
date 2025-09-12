package rest

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/manuelarte/pagorminator"
	"github.com/manuelarte/ptrutils"

	"github.com/manuelarte/gowasp/internal/models"
	"github.com/manuelarte/gowasp/internal/posts/postcomments"
	"github.com/manuelarte/gowasp/internal/sliceutils"
)

const defaultPageRequestSize = 10

type CommentsHandler struct {
	service postcomments.Service
}

func NewComments(service postcomments.Service) *CommentsHandler {
	return &CommentsHandler{service: service}
}

func (h *CommentsHandler) GetPostComments(c *gin.Context, postID uint, params GetPostCommentsParams) {
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

	postComments, err := h.service.GetAllForPostID(c, postID, pageRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Code:    http.StatusBadRequest,
			Details: err,
			Message: "Error getting the post comments",
		})

		return
	}
	hourTime := time.Hour
	c.SetCookie("csrf", uuid.New().String(), int(hourTime),
		fmt.Sprintf("/posts/%d/comments", postID), "localhost", false, true)
	c.JSON(http.StatusOK, postPagePostCommentToDTO(postComments, pageRequest))
}

func (h *CommentsHandler) PostAPostComment(c *gin.Context, postID uint) {
	postCommentNew := PostCommentNew{}
	if err := c.BindJSON(&postCommentNew); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Code:    http.StatusBadRequest,
			Details: err,
			Message: "Error marshalling the post comment",
		})

		return
	}
	postComment := postCommentNewToDAO(postCommentNew, postID, time.Now())
	err := h.service.Create(c, &postComment)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Code:    http.StatusBadRequest,
			Details: err,
			Message: "Error creating the post comment",
		})

		return
	}
	c.JSON(http.StatusOK, postComment)
}

func postCommentNewToDAO(dto PostCommentNew, postID uint, postedAt time.Time) models.PostComment {
	return models.PostComment{
		PostedAt: postedAt,
		PostID:   postID,
		UserID:   dto.UserID,
		Comment:  dto.Comment,
	}
}

func postPagePostCommentToDTO(
	postComments []*models.PostComment,
	pageRequest *pagorminator.Pagination,
) PagePostComments {
	return PagePostComments{
		UnderscoreMetadata: PageMetadata{
			Page:       pageRequest.GetPage(),
			Size:       pageRequest.GetSize(),
			TotalCount: int(pageRequest.GetTotalElements()),
			TotalPages: pageRequest.GetTotalPages(),
		},
		Data: sliceutils.Transform(postComments, func(x *models.PostComment) PostComment {
			return PostComment{
				Comment:   x.Comment,
				CreatedAt: x.CreatedAt,
				ID:        x.ID,
				PostID:    x.PostID,
				PostedAt:  x.PostedAt,
				UpdatedAt: x.UpdatedAt,
				UserID:    x.UserID,
			}
		},
		),
	}
}
