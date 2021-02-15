package mysession

import "github.com/gofiber/fiber/v2/middleware/session"

// SessionStore ...
var (
	SessionStore *session.Store
)

// CreateStore ...
func CreateStore(){
	SessionStore = session.New()
}