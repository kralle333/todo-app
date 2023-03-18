package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"kristianpilegaard.dk/todo-app/pkg/todo"
	"net/http"
)

func (a *todoApp) createTodoList(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	req := struct {
		Title string `json:"title"`
	}{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeError(w, err, http.StatusInternalServerError)
	}
	if req.Title == "" {
		writeError(w, err, http.StatusBadRequest)
		return
	}
	list, err := a.todoComponent.CreateTodoList(r.Context(), req.Title)
	if err != nil {
		writeError(w, err, http.StatusInternalServerError)
		return
	}
	resp := struct {
		NewList todo.List `json:"new_list"`
	}{
		NewList: list,
	}
	writeJsonResponse(w, resp)
}

func (a *todoApp) getTodo(w http.ResponseWriter, r *http.Request) {
	id, err := readIDParam(r, todoIDKey)
	if err != nil {
		writeError(w, err, http.StatusBadRequest)
		return
	}

	list, err := a.todoComponent.GetTodoList(r.Context(), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeError(w, err, http.StatusNotFound)
			return
		}
		writeError(w, err, http.StatusInternalServerError)
		return
	}
	writeJsonResponse(w, list)
}

func (a *todoApp) getAllTodos(w http.ResponseWriter, r *http.Request) {
	lists, err := a.todoComponent.GetAllTodoLists(r.Context())
	if err != nil {
		writeError(w, err, http.StatusInternalServerError)
		return
	}
	resp := struct {
		List []todo.List `json:"list"`
	}{
		List: lists,
	}
	writeJsonResponse(w, resp)
}

func (a *todoApp) deleteTodo(w http.ResponseWriter, r *http.Request) {
	id, err := readIDParam(r, todoIDKey)
	if err != nil {
		writeError(w, err, http.StatusBadRequest)
		return
	}

	err = a.todoComponent.DeleteTodoList(r.Context(), id)
	if err != nil {
		writeError(w, err, http.StatusBadRequest)
		return
	}
	writeJsonResponse(w, "OK")
}
