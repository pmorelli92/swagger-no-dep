## Requires developer to install on his/her machine both swag and redoc-cli
## go install github.com/swaggo/swag/cmd/swag@v1.7.1
## npm install -g redoc-cli@0.12.3
swagger/upd:
	swag init -g main.go -o ./swagger
	cd swagger && rm docs.go swagger.json
	cd swagger && redoc-cli bundle swagger.yaml -o swagger.html
