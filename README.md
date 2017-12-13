```sh
go get github.com/gorilla/mux

#rest_api_single - GET for home + all employees
go build rest_api_single.go
./rest_api_single
curl -X GET localhost:9999/
curl -X GET localhost:9999/emp

#rest_api - GET/POST/DELETE
go build rest_api.go
./rest_api
curl -X GET localhost:9999/
curl -X GET localhost:9999/emp
curl -X GET localhost:9999/emp/<id>


```