docker run --name postgis -e POSTGRES_USER=dev -e POSTGRES_PASSWORD=123456 -e POSTGRES_DB=dev_db -p 5432:5432 -d postgis/postgis

docker exec -it postgis psql -U dev -d dev_db

//
SELECT PostGIS_Version();

docker stop postgis
docker rm postgis