package jwtHelper

import (
	"context"
	"net/http"

	"github.com/lestrrat-go/jwx/jwt"

	"github.com/go-chi/jwtauth/v5"

	api "web-server/api"
	errorHelper "web-server/helper/error"
	errorModel "web-server/model/error"
)

type ctxKey int

const (
	ctxClaims ctxKey = iota
	ctxRefreshToken
)

// ClaimsFromCtx retrieves the parsed AppClaims from request context.
func ClaimsFromCtx(ctx context.Context) AppClaims {
	return ctx.Value(ctxClaims).(AppClaims)
}

// RefreshTokenFromCtx retrieves the parsed refresh token from context.
func RefreshTokenFromCtx(ctx context.Context) string {
	return ctx.Value(ctxRefreshToken).(string)
}

// Authenticator is a default authentication middleware to enforce access from the
// Verifier middleware request context values. The Authenticator sends a 401 Unauthorized
// response for any unverified tokens and passes the good ones through.
func Authenticator(next http.Handler) http.Handler {
	authenticator := func(w http.ResponseWriter, r *http.Request) *errorModel.AppError {
		token, claims, err := jwtauth.FromContext(r.Context())

		if err != nil {
			return errorHelper.ErrUnauthorized(err)
		}

		if err := jwt.Validate(token); err != nil {
			return errorHelper.ErrTokenExpired(err)
		}

		// Token is authenticated, parse claims
		var c AppClaims
		err = c.ParseClaims(claims)
		if err != nil {
			return errorHelper.ErrAccessTokenInvalid(err)
		}

		// Set AppClaims on context
		ctx := context.WithValue(r.Context(), ctxClaims, c)
		next.ServeHTTP(w, r.WithContext(ctx))

		return nil
	}
	return api.AppHandler(authenticator)
}

// AuthenticateRefreshJWT checks validity of refresh tokens and is only used for
// access token refresh and logout requests.
// It responds with 401 Unauthorized for invalid or expired refresh tokens.
func AuthenticateRefreshJWT(next http.Handler) http.Handler {
	authenticateRefreshJWT := func(w http.ResponseWriter, r *http.Request) *errorModel.AppError {
		token, claims, err := jwtauth.FromContext(r.Context())
		if err != nil {
			return errorHelper.ErrUnauthorized(err)
		}

		if err := jwt.Validate(token); err != nil {
			return errorHelper.ErrTokenExpired(err)
		}

		// Token is authenticated, parse refresh token string
		var c RefreshClaims
		err = c.ParseClaims(claims)
		if err != nil {
			return errorHelper.ErrRefreshTokenInvalid(err)
		}
		// Set refresh token string on context
		ctx := context.WithValue(r.Context(), ctxRefreshToken, c.Token)
		next.ServeHTTP(w, r.WithContext(ctx))
		return nil
	}
	return api.AppHandler(authenticateRefreshJWT)
}
