package backend

import (
	"context"
)

func (s *Server) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	// allowed_endpoints := map[string]bool{
	// 	"/backend.Backend/CreateUser":   true,
	// }

	// if allow, ok := allowed_endpoints[fullMethodName]; allow && ok {
	// 	return ctx, nil
	// }
	return ctx, nil

}
