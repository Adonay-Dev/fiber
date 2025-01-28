// ⚡️ Fiber is an Express inspired web framework written in Go with ☕️
// 🤖 Github Repository: https://github.com/gofiber/fiber
// 📌 API Documentation: https://docs.gofiber.io

package fiber

import (
	"errors"

	"github.com/valyala/fasthttp"
)

type CustomCtx[T any] interface {
	Ctx[T]

	// Reset is a method to reset context fields by given request when to use server handlers.
	Reset(fctx *fasthttp.RequestCtx)

	// Methods to use with next stack.
	getMethodINT() int
	getIndexRoute() int
	getTreePath() string
	getDetectionPath() string
	getPathOriginal() string
	getValues() *[maxParams]string
	getMatched() bool
	setIndexHandler(handler int)
	setIndexRoute(route int)
	setMatched(matched bool)
	setRoute(route *Route)
}

func NewDefaultCtx(app *App[DefaultCtx]) *DefaultCtx {
	// return ctx
	return &DefaultCtx{
		// Set app reference
		app: app,
	}
}

func (app *App[TCtx]) newCtx() Ctx[TCtx] {
	var c Ctx[TCtx]

	if app.newCtxFunc != nil {
		c = app.newCtxFunc(app)
	} else {
		c = NewDefaultCtx(app)
	}

	return c
}

// AcquireCtx retrieves a new Ctx from the pool.
func (app *App[TCtx]) AcquireCtx(fctx *fasthttp.RequestCtx) Ctx[TCtx] {
	ctx, ok := app.pool.Get().(Ctx)

	if !ok {
		panic(errors.New("failed to type-assert to Ctx"))
	}
	ctx.Reset(fctx)

	return ctx
}

// ReleaseCtx releases the ctx back into the pool.
func (app *App[TCtx]) ReleaseCtx(c Ctx[TCtx]) {
	c.release()
	app.pool.Put(c)
}
