package middleware

import (
	"net/http"
)

// RequireAuth intercepta la petición nativa y valida el token con Python
func RequireAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 1. Extraemos el Token
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, `{"error": "Acceso denegado: Token no proporcionado"}`, http.StatusUnauthorized)
			return
		}

		// 2. Nos comunicamos con el contenedor de Python (auth-service)
		req, err := http.NewRequest("GET", "http://auth-service:8000/api/auth/verify", nil)
		if err != nil {
			http.Error(w, `{"error": "Error interno al preparar la verificación"}`, http.StatusInternalServerError)
			return
		}
		req.Header.Add("Authorization", authHeader)

		// 3. Ejecutamos la llamada HTTP
		client := &http.Client{}
		resp, err := client.Do(req)

		// 4. Si Python rechaza el token, bloqueamos
		if err != nil || resp.StatusCode != http.StatusOK {
			http.Error(w, `{"error": "Acceso denegado: Token inválido o expirado"}`, http.StatusUnauthorized)
			return
		}

		// 5. ¡Luz Verde! Pasamos la petición a tu manejador original
		next.ServeHTTP(w, r)
	}
}
