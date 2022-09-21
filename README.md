## TimeShiBot Backend

### Описание
- TODO

### Запуск на VDS **required go >= 1.18*
```bash
git clone https://github.com/rombintu/timeshibot-backend.git
cd timeshibot-backend
go build -o bot cmd/main.go
export (#TODO)
PROD=true ./bot
```

### Запуск на Docker-compose
```bash
git clone https://github.com/rombintu/timeshibot-backend.git
cd timeshibot-backend
vim docker-compose.yaml
docker-compose up -d
docker-compose logs
```