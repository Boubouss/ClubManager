package handler

import (
  "net/http"
  "ClubManager/frontend/internal/views/user"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
  user.User().Render(r.Context(), w)
}
