docker run --name postgis -e POSTGRES_USER=dev -e POSTGRES_PASSWORD=12345678 -e POSTGRES_DB=dev_db -p 5432:5432 -d postgis/postgis

docker exec -it postgis psql -U dev -d dev_db

//
SELECT PostGIS_Version();

docker stop postgis
docker rm postgis

mkdir -p $PWD/data/minio/
docker run -p 9000:9000 -p 9001:9001 --name minio \
  -e "MINIO_ACCESS_KEY=dev" \
  -e "MINIO_SECRET_KEY=12345678" \
  -v $PWD/data/minio/:/data \
  minio/minio server /data --console-address ":9001"

docker stop minio
docker rm minio