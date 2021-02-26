module github.com/mercadolibre/sre_challenge/mockdata

go 1.14

require (
	github.com/bxcodec/faker/v3 v3.6.0
	github.com/go-redis/redis/v8 v8.6.0 // indirect
	github.com/jlciappolino/sre_challenge/apitools v0.0.1
	github.com/vmihailenco/msgpack/v5 v5.2.0 // indirect
)

//delete when push the changes on apitools repo on github
replace github.com/jlciappolino/sre_challenge/apitools v0.0.1 => ../apitools

