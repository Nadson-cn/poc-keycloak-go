package main

import (
	"fmt"
	"log"
	"net/http"

	"keycloak-service/config"
	"keycloak-service/handlers"
	"keycloak-service/utils"
)

func main() {

	if err := config.InitConfig(); err != nil {
		log.Fatalf("Erro ao inicializar configuração: %v", err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.HomeHandler)
	mux.HandleFunc("/login", handlers.LoginHandler)
	mux.HandleFunc("/callback", handlers.CallbackHandler)
	mux.HandleFunc("/get-token", handlers.GetTokenHandler)
	mux.HandleFunc("/validate-token", handlers.ValidateTokenHandler)
	mux.HandleFunc("/health", handlers.HealthCheckHandler)

	port := ":8081"
	utils.LogInfo("Servidor main iniciado na porta" + port)
	fmt.Println("Servidor main rodando em http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, mux))
}
