run: build
	make -j 3 air css templ

build:
	@go build -o bin/app .

air:
	air

css:
	npx tailwindcss -i ./views/css/app.css -o ./public/styles.css --watch

css-build:
	npx tailwindcss -i ./views/css/app.css -o ./public/styles.css

templ:
	templ generate --watch
