module github.com/jlciappolino/sre_challenge/users

go 1.15

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/jlciappolino/sre_challenge v0.0.0-20210127143830-60a6f0be9d4b // indirect
	github.com/jlciappolino/sre_challenge/apitools v0.0.0-00010101000000-000000000000 // indirect
)

replace github.com/jlciappolino/sre_challenge/apitools => ../apitools
