build:
	git pull
	docker compose build
	docker compose up -d --no-deps