package models

type PageResponse[T any] struct {
	Data []T `json:"data"`
	//nolint:tagliatelle // good name
	Metadata PageMetadata `json:"_metadata"`
}

type PageMetadata struct {
	Page       int   `json:"page"`
	Size       int   `json:"size"`
	TotalCount int64 `json:"totalCount"`
	TotalPages int   `json:"totalPages"`
}

func Transform[T, Y any](original PageResponse[T], f func(t T) Y) PageResponse[Y] {
	data := make([]Y, len(original.Data))
	for i, item := range original.Data {
		data[i] = f(item)
	}

	return PageResponse[Y]{
		Data:     data,
		Metadata: original.Metadata,
	}
}
