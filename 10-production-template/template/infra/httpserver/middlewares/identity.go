package middlewares

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	"template/infra/httpserver/utils"
)

type userContextKey struct{}
type accessShipsContextKey struct{}

type OwnUser struct {
	Username       string `json:"username"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	ID             string `json:"id"`
	Admin          bool   `json:"admin"`
	UserLvl        string `json:"user_lvl"`
	OrganizationID string `json:"organization_id"`
}

func UserHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		raw := r.Header.Get("x-user")
		if raw == "" {
			utils.WriteError(w, http.StatusUnauthorized, "invalid user")
			return
		}

		var user OwnUser
		if err := json.Unmarshal([]byte(raw), &user); err != nil {
			slog.Error("invalid x-user header", "err", err)
			utils.WriteError(w, http.StatusUnauthorized, "invalid user")
			return
		}

		ctx := context.WithValue(r.Context(), userContextKey{}, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func UserFromContext(ctx context.Context) (OwnUser, bool) {
	user, ok := ctx.Value(userContextKey{}).(OwnUser)
	return user, ok
}

func AccessShips(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		raw := r.Header.Get("x-access-ship")
		if raw == "" {
			utils.WriteError(w, http.StatusUnauthorized, "invalid ship header")
			return
		}

		var ships []string
		if err := json.Unmarshal([]byte(raw), &ships); err != nil {
			slog.Error("invalid x-access-ship header", "err", err)
			utils.WriteError(w, http.StatusUnauthorized, "invalid ship header")
			return
		}

		ctx := context.WithValue(r.Context(), accessShipsContextKey{}, ships)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func AccessShipsFromContext(ctx context.Context) ([]string, bool) {
	ships, ok := ctx.Value(accessShipsContextKey{}).([]string)
	return ships, ok
}
