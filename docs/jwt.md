# Intro
This document describes all of the processing of JWT in this code. <br>
It is just a handbook for future project that will rely on this one. 

First chapter focuses on JWT in general, then the second section desribes code itself.

## JWT
### What it is
First, JSON Web Token looks like this
```sh
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcmVzQXQiOjE2OTM0MzY3MTksImxvZ2luIjoiYWRtaW4ifQ.Bf44wPT_hS-BC1c2r3ZCIlLIqeKQjs0xwyPwX4SXtM4
```
It consists of three parts:
![](img/jwt-structure.png)

- header:
    - consist information about the type of token and algorithm used
- payload:
    - the actual encrypted data in the format of JSON
- signature:
    - ciphertext is signed with the SECRET known only to the server

### The usage
JWT is a mechanism of an authorization (it says wheter user has permission to access resource or not).

The usage is as shown below.

First user proves that he/she is authorized (simply by password). If so, server sends JWT that user can later use to prove its identity.

Then in the subsequent HTTP requests user sends token. Server then checks the validity of this token, reads payload data (user information) from it and decide to send resource or not.

![](img/JWT_tokens_EN.png)

## Code
We can distinguish two functionalities to be implemented on the server side related to the JWT.

- Creation of JWT token for the user
- Parsing of JWT token sent by the user along with resource request
- Validation of parsed JWT token due to its authorization purposes

### Creation of JWT token

#### `api/jwt.go`
```golang
func CreateJWT(login string) (string, error) {
	claims := &jwt.MapClaims{
		"login":     login,
		"expiresAt": time.Now().Add(time.Minute * 15).Unix(),
	}

	secret := "SECRET"
	signingKey := []byte(secret)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString(signingKey)

	return tokenStr, err
}
```
Function is called with the parameter `login` which simply identifies the user who is trying to authorize himself/herself -> so the token is created for user with such login.

`claims` is put in the ***payload*** part of the token.

`secret` is the secret known to the server with which it signs the token.

At the end function creates the token with `claims` and signs it with `secret`.
### Parsing of JWT token
#### `api/jwt.go`
```golang
func ValidateJWT(tokenStr string) (*jwt.Token, error) {
	secret := "SECRET"

	return jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secret), nil
	})
}
```
This function receives the string sent by the user in header of HTTP packet containing encrypted token.
Then it returns parsed token with single check in between.

### Validation of parsed token
#### `api/server.go`
```golang
func withJWTAuth(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.Header.Get("x-jwt-token")

		token, err := ValidateJWT(tokenStr)
		if err != nil {
			WriteJSON(w, http.StatusOK, fmt.Sprintf("error: %s", err.Error()))
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if ok != true {
			WriteJSON(w, http.StatusOK, fmt.Sprintf("error: %s", err.Error()))
			return
		}

		if token.Valid != true {
			WriteJSON(w, http.StatusOK, fmt.Sprintf("error: %s", err.Error()))
			return
		}

		expiresAtFloat := claims["expiresAt"].(float64)
		expiresAtTime := time.Unix(int64(expiresAtFloat), 0)

		if time.Now().After(expiresAtTime) {
			WriteJSON(w, http.StatusOK, "error: permission denied")
			return
		}

		// at the end: call the given function
		handlerFunc(w, r)
	}
}
```
This function implements the Decorator Design Pattern.

It decorates handler function with the validation of JWT token sent by the browser.

It calls the previously described `ValidateJWT` function.

Here is the place in which you can unpack the `claims` (token payload) and check it for the purpose of authorization.
