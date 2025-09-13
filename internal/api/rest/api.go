package rest

type API struct {
	*UsersHandler
	*PostsHandler
	*CommentsHandler
}
