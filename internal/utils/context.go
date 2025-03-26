package utils

import "context"

type ContextKey string

const (
	CtxKeyToken ContextKey = "Token-Claim"
)

func GetClaimFromContext(ctx context.Context) *TokenClaims {
	tc := ctx.Value(CtxKeyToken)
	if tc == nil {
		return nil
	}
	return tc.(*TokenClaims)
}
