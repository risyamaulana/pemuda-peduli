package interfaces

import "github.com/fasthttp/router"

// IApplication ...
type IApplication interface {
	Initialize(r *router.Router)
	Destroy()
}
