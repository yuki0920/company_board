## 起動

### Docker

```
# ビルド
docker-compose build
# 起動と同時にairでホットリロードする
docker-compose up -d
```

### ローカル

```bash
# 外部モジュールダウンロード
go mod download
# コールドリロード
go run main.go
# ホットリロード
air
```

## Seed データ

```bash
docker-compose run app go run seeds/main.go
```
