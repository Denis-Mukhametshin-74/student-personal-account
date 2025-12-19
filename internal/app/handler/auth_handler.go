package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"student-portal/internal/app/dto"
	"student-portal/internal/app/service"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req dto.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Неверный формат данных", http.StatusBadRequest)
		return
	}

	response, err := h.authService.Login(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req dto.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Неверный формат данных", http.StatusBadRequest)
		return
	}

	response, err := h.authService.Register(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (h *AuthHandler) OAuthGoogle(w http.ResponseWriter, r *http.Request) {
	state := "random_state_string"

	url := h.authService.GetGoogleOAuthURL(state)

	http.Redirect(w, r, url, http.StatusFound)
}

func (h *AuthHandler) OAuthGoogleCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	state := r.URL.Query().Get("state")

	if state != "random_state_string" {
		http.Error(w, "Invalid state", http.StatusBadRequest)
		return
	}

	if code == "" {
		http.Error(w, "Code not provided", http.StatusBadRequest)
		return
	}

	response, err := h.authService.HandleGoogleCallback(code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	html := fmt.Sprintf(`
    <!DOCTYPE html>
    <html>
    <head>
        <title>Авторизация...</title>
        <script>
            try {
                // Сохраняем токен и данные студента
                localStorage.setItem('token', '%s');
                localStorage.setItem('student', '%s');
                
                console.log('OAuth успешен, токен сохранен');
                
                // Редирект на главную страницу
                window.location.href = '/index.html';
            } catch (error) {
                console.error('Ошибка при сохранении токена:', error);
                document.body.innerHTML = '<h1>Ошибка авторизации</h1><p>' + error.message + '</p>';
            }
        </script>
    </head>
    <body>
        <p>Авторизация успешна. Перенаправляем...</p>
    </body>
    </html>
    `,
		response.Token,
		escapeJSON(response.Student))

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(html))
}

func escapeJSON(v interface{}) string {
	data, err := json.Marshal(v)
	if err != nil {
		return "{}"
	}
	str := string(data)
	str = strings.ReplaceAll(str, `\`, `\\`)
	str = strings.ReplaceAll(str, `"`, `\"`)
	str = strings.ReplaceAll(str, "\n", `\n`)
	str = strings.ReplaceAll(str, "\r", `\r`)
	return str
}
