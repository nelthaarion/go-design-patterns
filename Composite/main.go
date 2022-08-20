package main

import "fmt"

type Printer interface {
	Print()
}
type Product struct {
	Amount int
	Name   string
}

type Category struct {
	Name          string
	Products      []*Product
	SubCategories []*Category
}

func (p *Product) Print() {
	fmt.Printf("\tProduct name: %s , available: %d\n", p.Name, p.Amount)
}
func (c *Category) Print() {
	fmt.Printf("Category: %s\n\tDirect Products:\n", c.Name)
	for _, product := range c.Products {
		product.Print()
	}
	fmt.Println("Subcategories:")
	for i, subCategory := range c.SubCategories {
		fmt.Printf("%d - category: %s\n\tProducts:\n", i, subCategory.Name)
		for _, productInSubcategory := range subCategory.Products {
			productInSubcategory.Print()
		}
	}
}
func (c *Category) AddSubCategories(category *Category) {
	c.SubCategories = append(c.SubCategories, category)
}
func (c *Category) AddProduct(product *Product) {
	c.Products = append(c.Products, product)
}

func main() {
	tvCategory := &Category{Name: "Televisions"}
	samsung := &Category{Name: "Samsung"}
	x13s := &Product{Name: "X13S model", Amount: 10}
	samsung.AddProduct(x13s)

	lg := &Category{Name: "LG"}
	b341 := &Product{Name: "B341 model", Amount: 20}
	lg.AddProduct(b341)

	tvCategory.AddSubCategories(samsung)
	tvCategory.AddSubCategories(lg)

	tvCategory.Print()
}
