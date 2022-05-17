package constant

var AuthExcludePath = map[string]struct{}{
	"POST /auth/register":   {},
	"POST /auth/login":      {},
	"GET /user/:id":         {},
	"GET /organization":     {},
	"GET /organization/:id": {},
	"GET /team/:id":         {},
}
