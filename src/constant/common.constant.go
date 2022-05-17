package constant

var AuthExcludePath = map[string]struct{}{
	"/auth/register": {},
	"/auth/login":    {},
	"/user/:id":      {},
}
