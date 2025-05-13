#!/bin/sh

echo "等待 MySQL (${DB_HOST}:${DB_PORT}) 啟動..."

# 等資料庫 port 開啟
while ! nc -z "$DB_HOST" "$DB_PORT"; do
  sleep 1
done

echo "MySQL 已啟動，啟動 Go App..."
exec ./myapp