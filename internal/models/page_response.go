package models

type PageResponse[T any] struct {
	Data     []T          `json:"data"`
	Metadata PageMetadata `json:"_metadata"`
}

type PageMetadata struct {
	Page       int   `json:"page"`
	Size       int   `json:"size"`
	TotalCount int64 `json:"totalCount"`
	TotalPages int   `json:"totalPages"`
}
