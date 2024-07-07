jwt_key:
	ssh-keygen -t rsa -b 4096 -m PEM -f starter_jwtRS256.key

jwt_pub:
	openssl rsa -in starter_jwtRS256.key -pubout -outform PEM -out starter_jwtRS256.key.pub

jwt: jwt_key jwt_pub

env:
	cp ./.env.example ./.env

dev:
	docker compose -f ./docker-compose.yml up -d

dev-build:
	docker compose -f ./docker-compose.yml up -d --build

dev-down:
	docker compose -f ./docker-compose.yml down
