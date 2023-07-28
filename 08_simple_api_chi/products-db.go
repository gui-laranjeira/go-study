package main

import "fmt"

var memoryDb map[string]*Product

func BuildDb() {
	startProducts := make(map[string]string)
	startProducts["Iphone"] = "eletronics"
	startProducts["Macbook"] = "eletronics"
	startProducts["Ipad"] = "eletronics"
	startProducts["Blue Skinny Jeans"] = "clothing"
	startProducts["Liverpool T-Shirt"] = "clothing"
	startProducts["Red Dress"] = "clothing"
	startProducts["Nike Air Max"] = "shoes"
	startProducts["Adidas Superstar"] = "shoes"
	startProducts["Vans Old Skool"] = "shoes"
	startProducts["Coca-Cola"] = "beverages"
	startProducts["Pepsi"] = "beverages"
	startProducts["Refrigerator"] = "appliances"

	memoryDb = make(map[string]*Product)
	i := 0
	for k, v := range startProducts {
		id := fmt.Sprintf("%d", i)
		memoryDb[id] = &Product{
			ID:       id,
			Name:     k,
			Type:     v,
			Quantity: 100,
		}
		i++
	}
}
