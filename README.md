## 起動

### バックエンド

```bash
# コールドリロード
go run main.go
# ホットリロード
air
```

## Seed データ

Order/OrderItem は Seed を使う

```main.go
seeds.ExecSeed(30)
```

データ刷新時はテーブルを消す

```database/connect.go
database.Migrator().DropTable(&models.Order{})
database.Migrator().DropTable(&models.OrderItem{})
```
