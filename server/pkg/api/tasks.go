package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"kristianpilegaard.dk/todo-app/pkg/todo"
	"net/http"
)

func (a *todoApp) getAllTasks(w http.ResponseWriter, r *http.Request) {
	todoListID, err := readIDParam(r, todoIDKey)
	if err != nil {
		return
	}
	tasks, err := a.todoComponent.GetAllTasks(r.Context(), todoListID)
	if err != nil {
		writeError(w, err, http.StatusInternalServerError)
		return
	}
	resp := struct {
		Tasks []todo.Task `json:"tasks"`
	}{
		Tasks: tasks,
	}
	writeJsonResponse(w, resp)
}

func (a *todoApp) deleteTask(w http.ResponseWriter, r *http.Request) {
	taskID, err := readIDParam(r, taskIDKey)
	if err != nil {
		writeError(w, err, http.StatusBadRequest)
	}
	err = a.todoComponent.DeleteTask(r.Context(), taskID)
	if err != nil {
		writeError(w, err, http.StatusBadRequest)
		return
	}
	writeJsonResponse(w, "ok")
}

func (a *todoApp) addTask(w http.ResponseWriter, r *http.Request) {
	todoID, err := readIDParam(r, todoIDKey)
	if err != nil {
		writeError(w, err, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	req := struct {
		Title string `json:"title"`
	}{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeError(w, err, http.StatusInternalServerError)
		return
	}
	if req.Title == "" {
		writeError(w, err, http.StatusBadRequest)
		return
	}

	_, err = a.todoComponent.AddTask(r.Context(), todoID, req.Title)
	if err != nil {
		writeError(w, err, http.StatusInternalServerError)
		return
	}
	tasks, err := a.todoComponent.GetAllTasks(r.Context(), todoID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		writeError(w, err, http.StatusInternalServerError)
		return
	}
	resp := struct {
		Tasks []todo.Task `json:"tasks"`
	}{
		Tasks: tasks,
	}
	writeJsonResponse(w, resp)
}

func (a *todoApp) markCompleted(w http.ResponseWriter, r *http.Request) {
	taskID, err := readIDParam(r, taskIDKey)
	if err != nil {
		writeError(w, err, http.StatusBadRequest)
		return
	}
	err = a.todoComponent.MarkTaskCompleted(r.Context(), taskID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeError(w, err, http.StatusNotFound)
			return
		}
		writeError(w, err, http.StatusInternalServerError)
	}
	task, err := a.todoComponent.GetTask(r.Context(), taskID)
	writeJsonResponse(w, task)
}
