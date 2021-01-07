package auth

import (
	"context"
	"fmt"
	"mmbt/constants"
	"mmbt/db"
	"mmbt/models"
	"net/http"
	"strings"
	"time"
)

// APIRouteMiddleware provides auth to a route handler
func APIRouteMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := strings.TrimSpace(r.Header.Get("Token"))

		if token == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		userLoginToken := &models.UserLoginToken{}

		err := db.DB.Model(userLoginToken).Where("token = ?", token).Select()
		if err != nil { //|| userLoginToken == nil {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		currentTime := time.Now()
		timeDiff := currentTime.Sub(userLoginToken.UpdatedAt)

		if timeDiff.Minutes() > 30 {
			// Session expired
			db.DB.Model(userLoginToken).WherePK().Delete()
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprintf(w, "Session timed out")
			return
		}

		userLoginToken.UpdatedAt = currentTime
		db.DB.Model(userLoginToken).WherePK().Update()

		ctx := context.WithValue(r.Context(), constants.RequestContextKey.AuthorizedUser, userLoginToken.UserID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
