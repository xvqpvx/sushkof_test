package controller

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io"
	"net/http"
	"sushkof_test/internal/data/response"
)

type HTMLController struct {
	templates *template.Template
}

func NewHTMLController() *HTMLController {
	templates := template.Must(template.ParseGlob("../../templates/*.html"))
	return &HTMLController{templates: templates}
}

func (c *HTMLController) RegisterHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := c.templates.ExecuteTemplate(w, "register.html", nil)
	if err != nil {
		return
	}
}

func (c *HTMLController) TaskHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := c.templates.ExecuteTemplate(w, "create_task.html", nil)
	if err != nil {
		return
	}
}

func (c *HTMLController) ListHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	resp, err := http.Get("http://localhost:8080/api/user/findAll")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var users []response.UserResponse
	err = json.Unmarshal(body, &users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = c.templates.ExecuteTemplate(w, "list.html", users)
	if err != nil {
		return
	}
}
