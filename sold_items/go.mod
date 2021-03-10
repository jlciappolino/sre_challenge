module github.com/mercadolibre/sre_challenge/sold_items

go 1.15

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/go-redis/redis/v8 v8.6.0
	github.com/jlciappolino/sre_challenge/apitools v0.0.0-20210226161205-bcc6b04b757b
)

replace github.com/jlciappolino/sre_challenge/apitools => ../apitools
