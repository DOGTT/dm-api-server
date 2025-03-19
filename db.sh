#!/bin/bash

# 定义 SQL 初始化脚本目录
SQL_INIT_DIR="./migration"

# 定义启动函数
start_services() {
  echo "Starting PostGIS container..."
  docker run --name postgis \
    -e POSTGRES_USER=dev \
    -e POSTGRES_PASSWORD=12345678 \
    -e POSTGRES_DB=dev_db \
    -p 5432:5432 \
    -d postgis/postgis

  echo "Starting MinIO container..."
  mkdir -p $PWD/data/minio/
  docker run -p 9000:9000 -p 9001:9001 --name minio \
    -e "MINIO_ACCESS_KEY=dev" \
    -e "MINIO_SECRET_KEY=12345678" \
    -v $PWD/data/minio/:/data \
    -d minio/minio server /data --console-address ":9001"

#   echo "Starting pgAdmin container..."
#   docker run -p 5433:80 \
#     -e "PGADMIN_DEFAULT_EMAIL=dev@dev.com" \
#     -e "PGADMIN_DEFAULT_PASSWORD=admin" \
#     -d dpage/pgadmin4

 # 等待 PostGIS 容器启动
  echo "Waiting for PostGIS to be ready..."
  sleep 10

  # 执行 SQL 初始化脚本
  execute_sql_init_scripts
  echo "All services started successfully!"
}

# 定义执行 SQL 初始化脚本的函数
execute_sql_init_scripts() {
  if [[ -d "$SQL_INIT_DIR" ]]; then
    echo "Executing SQL initialization scripts from $SQL_INIT_DIR..."
    for script in $(ls $SQL_INIT_DIR/*.sql | sort); do
      echo "Running script: $script"
      docker exec -i postgis psql -U dev -d dev_db < $script
      if [[ $? -eq 0 ]]; then
        echo "Script $script executed successfully."
      else
        echo "Failed to execute script $script."
        exit 1
      fi
    done
    echo "All SQL initialization scripts executed successfully."
  else
    echo "SQL initialization script directory $SQL_INIT_DIR not found."
  fi
}


# 定义停止函数
stop_services() {
  echo "Stopping and removing PostGIS container..."
  docker stop postgis 2>/dev/null
  docker rm postgis 2>/dev/null

  echo "Stopping and removing MinIO container..."
  docker stop minio 2>/dev/null
  docker rm minio 2>/dev/null

#   echo "Stopping and removing pgAdmin container..."
#   docker stop pgadmin4 2>/dev/null
#   docker rm pgadmin4 2>/dev/null

  echo "All services stopped and removed successfully!"
}

# 检查命令行参数
if [[ "$1" == "start" ]]; then
  start_services
elif [[ "$1" == "stop" ]]; then
  stop_services
else
  echo "Usage: $0 {start|stop}"
  exit 1
fi