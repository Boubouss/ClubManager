package handler

import (
    "net/http"
    "ClubManager/frontend/internal/views/home"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
  home.Home().Render(r.Context(), w)
}
