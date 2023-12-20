// ⚡️ Fiber is an Express inspired web framework written in Go with ☕️
// 🤖 Github Repository: https://github.com/gofiber/fiber
// 📌 API Documentation: https://docs.gofiber.io

package routing

import "github.com/gofiber/fiber/v3"

// Register defines all router handle interface generate by Route().
type Register interface {
	All(handler fiber.Handler, middleware ...fiber.Handler) Register
	Get(handler fiber.Handler, middleware ...fiber.Handler) Register
	Head(handler fiber.Handler, middleware ...fiber.Handler) Register
	Post(handler fiber.Handler, middleware ...fiber.Handler) Register
	Put(handler fiber.Handler, middleware ...fiber.Handler) Register
	Delete(handler fiber.Handler, middleware ...fiber.Handler) Register
	Connect(handler fiber.Handler, middleware ...fiber.Handler) Register
	Options(handler fiber.Handler, middleware ...fiber.Handler) Register
	Trace(handler fiber.Handler, middleware ...fiber.Handler) Register
	Patch(handler fiber.Handler, middleware ...fiber.Handler) Register

	Add(methods []string, handler fiber.Handler, middleware ...fiber.Handler) Register

	Static(root string, config ...fiber.Static) Register

	Route(path string) Register
}

var _ (Register) = (*Registering)(nil)

// Registering struct
type Registering struct {
	app *fiber.App

	path string
}

// All registers a middleware route that will match requests
// with the provided path which is stored in register struct.
//
//	app.Route("/").All(func(c fiber.Ctx) error {
//	     return c.Next()
//	})
//	app.Route("/api").All(func(c fiber.Ctx) error {
//	     return c.Next()
//	})
//	app.Route("/api").All(handler, func(c fiber.Ctx) error {
//	     return c.Next()
//	})
//
// This method will match all HTTP verbs: GET, POST, PUT, HEAD etc...
func (r *Registering) All(handler fiber.Handler, middleware ...fiber.Handler) Register {
	r.app.register([]string{fiber.methodUse}, r.path, nil, handler, middleware...)
	return r
}

// Get registers a route for GET methods that requests a representation
// of the specified resource. Requests using GET should only retrieve data.
func (r *Registering) Get(handler fiber.Handler, middleware ...fiber.Handler) Register {
	r.app.Add([]string{fiber.MethodGet}, r.path, handler, middleware...)
	return r
}

// Head registers a route for HEAD methods that asks for a response identical
// to that of a GET request, but without the response body.
func (r *Registering) Head(handler fiber.Handler, middleware ...fiber.Handler) Register {
	return r.Add([]string{fiber.MethodHead}, handler, middleware...)
}

// Post registers a route for POST methods that is used to submit an entity to the
// specified resource, often causing a change in state or side effects on the server.
func (r *Registering) Post(handler fiber.Handler, middleware ...fiber.Handler) Register {
	return r.Add([]string{fiber.MethodPost}, handler, middleware...)
}

// Put registers a route for PUT methods that replaces all current representations
// of the target resource with the request payload.
func (r *Registering) Put(handler fiber.Handler, middleware ...fiber.Handler) Register {
	return r.Add([]string{fiber.MethodPut}, handler, middleware...)
}

// Delete registers a route for DELETE methods that deletes the specified resource.
func (r *Registering) Delete(handler fiber.Handler, middleware ...fiber.Handler) Register {
	return r.Add([]string{fiber.MethodDelete}, handler, middleware...)
}

// Connect registers a route for CONNECT methods that establishes a tunnel to the
// server identified by the target resource.
func (r *Registering) Connect(handler fiber.Handler, middleware ...fiber.Handler) Register {
	return r.Add([]string{fiber.MethodConnect}, handler, middleware...)
}

// Options registers a route for OPTIONS methods that is used to describe the
// communication options for the target resource.
func (r *Registering) Options(handler fiber.Handler, middleware ...fiber.Handler) Register {
	return r.Add([]string{fiber.MethodOptions}, handler, middleware...)
}

// Trace registers a route for TRACE methods that performs a message loop-back
// test along the r.Path to the target resource.
func (r *Registering) Trace(handler fiber.Handler, middleware ...fiber.Handler) Register {
	return r.Add([]string{fiber.MethodTrace}, handler, middleware...)
}

// Patch registers a route for PATCH methods that is used to apply partial
// modifications to a resource.
func (r *Registering) Patch(handler fiber.Handler, middleware ...fiber.Handler) Register {
	return r.Add([]string{fiber.MethodPatch}, handler, middleware...)
}

// Add allows you to specify multiple HTTP methods to register a route.
func (r *Registering) Add(methods []string, handler fiber.Handler, middleware ...fiber.Handler) Register {
	r.app.register(methods, r.path, nil, handler, middleware...)
	return r
}

// Static will create a file server serving static files
func (r *Registering) Static(root string, config ...fiber.Static) Register {
	r.app.registerStatic(r.path, root, config...)
	return r
}

// Route returns a new Register instance whose route path takes
// the path in the current instance as its prefix.
func (r *Registering) Route(path string) Register {
	// Create new group
	route := &Registering{app: r.app, path: fiber.getGroupPath(r.path, path)}

	return route
}
