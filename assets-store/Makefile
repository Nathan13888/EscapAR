build:
	docker build -t assets-store .

run:
	docker run --rm -v $$PWD/nginx.conf:/etc/nginx/nginx.conf:ro  -p 5000:80 assets-store:latest

