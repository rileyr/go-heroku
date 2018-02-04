heroku: assets binary deploy cleanup

assets:
	statik --src=public/

binary:
	GOOS=linux GOARCH=amd64 go build -o app cmd/main.go

cleanup:
	rm -r statik && rm app

deploy:
	heroku container:push web --app $(HEROKU_APP)
