package productservice

//go:generate oapi-codegen --package productservice -o productservice.gen.go openapi.yaml

type ProductSlice []Product

func (products ProductSlice) Where(filter func(p Product) bool) ProductSlice {
	result := []Product{}
	for _, sub := range products {
		if filter(sub) {
			result = append(result, sub)
		}
	}
	return result
}

func (products ProductSlice) First() Product {
	return products[0]
}
