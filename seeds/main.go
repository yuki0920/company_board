package main

import (
	"math/rand"
	"strconv"

	"github.com/yuki0920/company_board/database"
	"github.com/yuki0920/company_board/models"
)

func main() {
	database.Connect()

	PermissionSeed()
	OrderSeed()
}

func PermissionSeed() {
	database.DB.Migrator().DropTable(&models.Permission{})
	database.DB.AutoMigrate(&models.Permission{})

	all_permissions := []string{"view_users", "view_roles", "view_products", "view_orders", "edit_products", "edit_users", "edit_roles", "edit_orders"}

	for _, permissionName := range all_permissions {
		permission := models.Permission{
			Name: permissionName,
		}

		database.DB.Create(&permission)
	}
}

// うまくいかない
func RoleSeed() {
	database.DB.Migrator().DropTable(&models.Role{})
	database.DB.AutoMigrate(&models.Role{})

	permissions := make([]models.Permission, 8)
	for i := range permissions {

		permissions[i] = models.Permission{
			Id: uint(i + 1),
		}
	}

	admin := models.Role{
		Name:        "Admin User",
		Permissions: permissions,
	}

	database.DB.Create(&admin)
}

func OrderSeed() {
	database.DB.Migrator().DropTable(&models.Order{})
	database.DB.AutoMigrate(&models.Order{})
	database.DB.Migrator().DropTable(&models.OrderItem{})
	database.DB.AutoMigrate(&models.OrderItem{})

	for i := 1; i <= 30; i++ {
		date := strconv.Itoa(i)

		order := models.Order{
			FirstName: "Foo",
			LastName:  "Bar",
			Email:     "test@example.com",
			CreatedAt: "2021-4-" + date,
			UpdatedAt: "2021-4-" + date,
		}

		database.DB.Create(&order)

		orderItem := models.OrderItem{
			OrderId:      uint(i),
			ProductTitle: "Title " + strconv.Itoa(i),
			Price:        float32(rand.Intn(99) + 1),
			Quantity:     uint(rand.Intn(9) + 1),
		}

		database.DB.Create(&orderItem)
	}
}
