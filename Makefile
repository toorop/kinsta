build:
	go build -o dist/kinsta

run: build
	dist/kinsta

deploy: build
	rsync -arvz dist/* root@toorop.fr:/var/www/kinsta.toorop.fr/
