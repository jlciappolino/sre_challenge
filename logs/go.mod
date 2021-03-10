module github.com/mercadolibre/sre_challenge/logs

go 1.14

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/go-redis/redis/v8 v8.6.0
	github.com/google/uuid v1.2.0
	github.com/jlciappolino/sre_challenge/apitools v0.0.0-20210226183241-75d384afcc2a
)

replace github.com/jlciappolino/sre_challenge/apitools => ../apitools