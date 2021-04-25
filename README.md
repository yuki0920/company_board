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

Order/OrderItem は Seed を使う

```
# main.go
seeds.ExecSeed(30)
```

データ刷新時はテーブルを消す

```
# database/connect.go
database.Migrator().DropTable(&models.Order{})
database.Migrator().DropTable(&models.OrderItem{})
```
