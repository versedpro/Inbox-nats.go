package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/matryer/way"
	"github.com/nicolasparada/nakama/internal/service"
)

func (h *handler) users(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	search := q.Get("search")
	first, _ := strconv.Atoi(q.Get("first"))
	after := q.Get("after")
	uu, err := h.Users(r.Context(), search, first, after)
	if err != nil {
		respondErr(w, err)
		return
	}

	respond(w, uu, http.StatusOK)
}

func (h *handler) usernames(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	startingWith := q.Get("starting_with")
	first, _ := strconv.Atoi(q.Get("first"))
	after := q.Get("after")
	uu, err := h.Usernames(r.Context(), startingWith, first, after)
	if err != nil {
		respondErr(w, err)
		return
	}

	respond(w, uu, http.StatusOK)
}

func (h *handler) user(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	username := way.Param(ctx, "username")
	u, err := h.User(ctx, username)
	if err == service.ErrInvalidUsername {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	if err == service.ErrUserNotFound {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err != nil {
		respondErr(w, err)
		return
	}

	respond(w, u, http.StatusOK)
}

func (h *handler) updateAvatar(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	reader := http.MaxBytesReader(w, r.Body, service.MaxAvatarBytes)
	defer reader.Close()

	avatarURL, err := h.UpdateAvatar(r.Context(), reader)
	if err == service.ErrUnauthenticated {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	if err == service.ErrUnsupportedAvatarFormat {
		http.Error(w, err.Error(), http.StatusUnsupportedMediaType)
		return
	}

	if err != nil {
		respondErr(w, err)
		return
	}

	fmt.Fprint(w, avatarURL)
}

func (h *handler) toggleFollow(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	username := way.Param(ctx, "username")

	out, err := h.ToggleFollow(ctx, username)
	if err == service.ErrUnauthenticated {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	if err == service.ErrInvalidUsername {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	if err == service.ErrUserNotFound {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err == service.ErrForbiddenFollow {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	if err != nil {
		respondErr(w, err)
		return
	}

	respond(w, out, http.StatusOK)
}

func (h *handler) followers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	q := r.URL.Query()
	username := way.Param(ctx, "username")
	first, _ := strconv.Atoi(q.Get("first"))
	after := q.Get("after")
	uu, err := h.Followers(ctx, username, first, after)
	if err == service.ErrInvalidUsername {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	if err != nil {
		respondErr(w, err)
		return
	}

	respond(w, uu, http.StatusOK)
}

func (h *handler) followees(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	q := r.URL.Query()
	username := way.Param(ctx, "username")
	first, _ := strconv.Atoi(q.Get("first"))
	after := q.Get("after")
	uu, err := h.Followees(ctx, username, first, after)
	if err == service.ErrInvalidUsername {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	if err != nil {
		respondErr(w, err)
		return
	}

	respond(w, uu, http.StatusOK)
}
