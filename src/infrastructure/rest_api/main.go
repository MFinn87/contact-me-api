package rest

import (
	"context"

	huma "github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humagin"
)

// Adapts Gin Context to Huma Context, so Gin middleware is more compatible.
func adaptContext() func(ctx huma.Context, humaNext func(huma.Context)) {
	return func(ctx huma.Context, humaNext func(huma.Context)) {
		ginContext := humagin.Unwrap(ctx)
		copiedGinContext := ginContext.Copy()

		newContext := ctx.Context()

		cKeys := copiedGinContext.Keys

		for k, v := range cKeys {
			newContext = context.WithValue(newContext, k, v)
		}

		updatedHumaContext := huma.WithContext(ctx, newContext)
		humaNext(updatedHumaContext)
	}
}

func UseMiddleware(middlewares huma.Middlewares) huma.Middlewares {
	return append(middlewares, adaptContext())
}

func UseRoute[I any, O any](api huma.API, op huma.Operation, routeHandler func(ctx context.Context, input *I) (*O, error)) {
	operationToUse := huma.Operation{
		OperationID: op.Method + op.Path,
		Method:      op.Method,
		Summary:     op.Method + " " + op.Path,
		Path:        op.Path,
		Middlewares: UseMiddleware(op.Middlewares),
		Description: op.Description,
	}

	huma.Register(
		api,
		operationToUse,
		func(ctx context.Context, input *I) (*O, error) {
			return routeHandler(ctx, input)
		})
}
