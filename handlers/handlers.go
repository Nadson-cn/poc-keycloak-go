package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"keycloak-service/auth"
	"keycloak-service/config"
	"keycloak-service/utils"
)

func HomeHandler(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte(`<html><body>
		<a href="/login">Login com Keycloak</a>
		</body></html>`))
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	url := config.AppConfig.OAuth2Config.AuthCodeURL("randomstate")
	http.Redirect(w, r, url, http.StatusFound)
}

func CallbackHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Código não encontrado", http.StatusBadRequest)
		return
	}

	token, err := config.AppConfig.OAuth2Config.Exchange(context.Background(), code)
	if err != nil {
		http.Error(w, "Erro ao trocar código por token", http.StatusInternalServerError)
		return
	}

	idToken, ok := token.Extra("id_token").(string)
	if !ok {
		http.Error(w, "Erro ao trocar código por token", http.StatusInternalServerError)
		return
	}

	resp := map[string]interface{}{
		"access_token":  token.AccessToken,
		"refresh_token": token.RefreshToken,
		"expires_in":    token.Expiry,
		"id_token":      idToken,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func GetTokenHandler(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GetToken()
	if err != nil {
		http.Error(w, "Erro ao obter token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"access_token": token})
}

func ValidateTokenHandler(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Token não fornecido", http.StatusUnauthorized)
		return
	}

	token, err := auth.ValidateToken(authHeader)
	if err != nil {
		http.Error(w, "Token inválido", http.StatusUnauthorized)
		return
	}

	var claims map[string]interface{}
	if err := token.Claims(&claims); err != nil {
		http.Error(w, "Erro ao obter claims", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"valid": true,
		"token": claims,
	})
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	utils.LogInfo("Health check realizado")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
