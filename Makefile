build:
	CGO_ENABLED=0 go build -o dist/kinsta

run: build
	dist/kinsta

deploy: build
	rsync -arvz dist/* root@toorop.fr:/var/www/kinsta.toorop.fr/

## docker build
# docker build -t kinsta:v0.1.2 -f Dockerfile .

## push @scaleway
# source .scaleway
# docker login rg.fr-par.scw.cloud/rechdev -u nologin -p $SCW_SECRET_KEY
# docker tag kinsta:v0.1.2 rg.fr-par.scw.cloud/rechdev/kinsta:v0.1.2
# docker push rg.fr-par.scw.cloud/rechdev/kinsta:v0.1.2

# https://rechdevfmr01txm-kinsta.functions.fnc.fr-par.scw.cloud/?username=wesbos&images=5