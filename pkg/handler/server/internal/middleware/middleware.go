package middleware

type middleware struct {
	authURL string
}

func NewMiddleware(authURL string) *middleware {
	return &middleware{
		authURL: authURL,
	}
}
