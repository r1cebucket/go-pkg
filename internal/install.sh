# postgresql
docker pull postgres:11.8
docker volume create postgres-test
docker run --name postgres-test --restart=always -e POSTGRES_PASSWORD=ricebucket -p 5432:5432 -v postgres-test:/var/lib/postgresql/data -d postgres:11.8
# /宿主机目录:/容器目录

# redis
docker run --name redis-test --restart=always -e POSTGRES_PASSWORD=ricebucket -p 6379:6379 -d redis