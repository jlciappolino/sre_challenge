module github.com/jlciappolino/sre_challenge/currency_conversions

go 1.15

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/jlciappolino/sre_challenge/apitools v0.0.0-20210226183241-75d384afcc2a
)

replace github.com/jlciappolino/sre_challenge/apitools => ../apitools
