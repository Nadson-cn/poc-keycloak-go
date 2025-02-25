package config

import (
	"context"
	"log"
	"os"

	"github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
)

type Config struct {
	Provider     *oidc.Provider
	OAuth2Config oauth2.Config
}

var AppConfig *Config

func InitConfig() error {
	ctx := context.Background()
	provider, err := oidc.NewProvider(ctx, KeycloakURL)
	if err != nil {
		return err
	}

	AppConfig = &Config{
		Provider: provider,
		OAuth2Config: oauth2.Config{
			ClientID:     ClientID,
			ClientSecret: ClientSecret,
			RedirectURL:  RedirectURI,
			Endpoint:     provider.Endpoint(),
			Scopes:       []string{oidc.ScopeOpenID, "profile", "email", "roles"},
		},
	}

	log.Println("Configuração OIDC carregada com sucesso")
	return nil
}

var (
	KeycloakURL      = getEnv("KEYCLOAK_URL", "http://localhost:8080/realms/demo")
	RedirectURI      = getEnv("REDIRECT_URI", "http://localhost:8081/callback")
	Realm            = getEnv("KEYCLOAK_REALM", "poc-keycloak ")
	ClientID         = getEnv("KEYCLOAK_CLIENT_ID", "go-app")
	ClientSecret     = getEnv("KEYCLOAK_CLIENT_SECRET", "JTu4t4N96rcAPJkuMWlnNV5qLxOAkWIR")
	TokenEndpoint    = KeycloakURL + "/realms/" + Realm + "/protocol/openid-connect/token"
	UserInfoEndpoint = KeycloakURL + "/realms/" + Realm + "/protocol/openid-connect/userinfo"
)

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
