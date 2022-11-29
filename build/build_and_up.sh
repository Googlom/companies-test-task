docker build -t task-postgres:latest -f ./build/postgres/Dockerfile .
docker build -t task-companies-api:latest -f ./build/companies-api/Dockerfile .
docker build -t task-auth-service:latest -f ./build/auth-service/Dockerfile .

echo "Building docker images finished."

docker-compose -f ./build/docker-compose.yml --env-file .env up -d
