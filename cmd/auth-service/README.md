Simple mini-authentication service that have only one endpoint that issues JWT tokens on `GET /login` request without any user and password (for simplicity purpose). Server listens at `:8090`. 

`HMAC_SECRET` env var must be set to the same value as for companies api server. 