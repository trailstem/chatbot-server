up:
		docker-compose -f dockers/docker-compose.yml build --no-cache && docker-compose.yml docker compose up -d
down:
		docker-compose -f dockers/docker-compose.yml down