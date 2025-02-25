# POC: Autenticação com Keycloak e OpenID Connect

## Descrição

Este projeto é uma **Prova de Conceito (POC)** para integrar um serviço Go com **Keycloak** utilizando **OpenID Connect (OIDC)** e **OAuth2**.

## 📌 Tecnologias Utilizadas

- **Go (Golang)**
- **Keycloak** (Identity Provider)
- **OpenID Connect (OIDC)**
- **OAuth2**
- **Docker** (para o ambiente do Keycloak)

## 📂 Estrutura do Projeto

```
poc-keycloak-go/
│── config/          # Configuração do Keycloak e OAuth2
│── handlers/        # Handlers para autenticação e validação de token
│── main.go          # Inicializa o servidor HTTP
│── Dockerfile       # Dockerfile para rodar o Keycloak
│── README.md        # Documentação do projeto
```

## 🚀 Como Executar

### 1️⃣ Clonar o repositório

```sh
git clone https://github.com/nadson-cn/poc-keycloak-go.git
cd poc-keycloak-go
```

### 2️⃣ Subir o Keycloak com Docker Compose

1. Certifique-se de ter o **Docker** e **Docker Compose** instalados.

```sh
docker-compose up -d
```

Isso iniciará um contêiner com o Keycloak rodando.

2. Acesse o Keycloak em: **http://localhost:8080**

3. Faça login com:

   - Usuário: `admin`
   - Senha: `admin`

4. Crie um **Realm** chamado `poc-keycloak` e um **Client** chamado `go-app` com:

   **Capability config**

   `Client authentication`: **On**

5. Em **Credentials**, anote o **Client Secret** para configurar no código.

---

### 3️⃣ Configurar e rodar a POC

1. Configure as variáveis no `config/config.go`:

   ```go
   const (
       KeycloakURL  = "http://localhost:8080/realms/myrealm"
       ClientID     = "go-app"
       ClientSecret = "YOUR_CLIENT_SECRET"
       RedirectURI  = "http://localhost:8081/callback"
   )
   ```

2. Instale as dependências:

   ```sh
   go mod tidy
   ```

3. Execute a aplicação:
   ```sh
   go run main.go
   ```

Agora o servidor estará rodando em **http://localhost:8081** 🎉

---

## 🛠️ Endpoints Disponíveis

### 🔐 **Login**

- **GET** `/login`
- Redireciona o usuário para o Keycloak para autenticação.

### 🔄 **Callback**

- **GET** `/callback?code=...`
- Troca o código de autorização pelo token JWT.

### ✅ **Validação de Token**

- **GET** `/validate-token`
- Recebe um Bearer token pelo header e valida no Keycloak.

---

## 📜 Licença

Este projeto é de uso livre para estudo e aprendizado.

---

## 📌 Referências

- [Documentação do Keycloak](https://www.keycloak.org/documentation/)
- [OpenID Connect](https://openid.net/connect/)
