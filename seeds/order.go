package seeds

import (
	"math/rand"
	"strconv"

	"github.com/yuki0920/company_board/database"
	"github.com/yuki0920/company_board/models"
)

func OrderSeed() {
	date := strconv.Itoa(rand.Intn(29) + 1)

	order := models.Order{
		FirstName: "Foo",
		LastName:  "Bar",
		Email:     "test@example.com",
		CreatedAt: "2021-4-" + date,
		UpdatedAt: "2021-4-" + date,
	}

	database.DB.Create(&order)
}

func OrderItemSeed(id int) {
	orderId := rand.Intn(id)

	orderItem := models.OrderItem{
		OrderId:      uint(orderId),
		ProductTitle: "Title " + strconv.Itoa(orderId),
		Price:        float32(rand.Intn(99) + 1),
		Quantity:     uint(rand.Intn(9) + 1),
	}

	database.DB.Create(&orderItem)
}
