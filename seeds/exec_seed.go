package seeds

func ExecSeed(count int) {
	for i := 1; i <= count; i++ {
		OrderSeed()

		OrderItemSeed(count)
	}
}
