package auth

import (
	"context"
	"errors"
	"strings"

	"keycloak-service/config"
	"keycloak-service/utils"

	"github.com/coreos/go-oidc"
	"golang.org/x/oauth2/clientcredentials"
)

type TokenResponse struct {
	AccessToken string `json:"access_token"`
}

func GetToken() (string, error) {
	config := clientcredentials.Config{
		TokenURL: config.TokenEndpoint,
		ClientID: config.ClientID,
	}

	token, err := config.Token(context.Background())
	if err != nil {
		utils.LogError("Erro ao obter token: " + err.Error())
		return "", err
	}

	return token.AccessToken, nil
}

func ValidateToken(tokenString string) (*oidc.IDToken, error) {
	if tokenString == "" {
		return nil, errors.New("token vazio")
	}

	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	tokenVerifier := config.AppConfig.Provider.Verifier(&oidc.Config{ClientID: config.ClientID})
	idToken, err := tokenVerifier.Verify(context.Background(), tokenString)

	if err != nil {
		utils.LogError("Token inválido: " + err.Error())
		return nil, err

	}

	return idToken, nil
}
