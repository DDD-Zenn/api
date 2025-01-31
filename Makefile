.PHONY: build

up:
	docker build -t ddd-api . && \
	docker-compose up -d

down:
	docker-compose down