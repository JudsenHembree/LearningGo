
package Item

import ("fmt")

type Itemer interface{
	PrintItem()
	LoadItem()
}

type Item struct{
	GTIN string
	Name string
	Description string
	Retail float64
	Wholesale float64
	Quantity int64
}

func (item Item) PrintItem(){
	fmt.Println(item.GTIN, item.Name, item.Wholesale, item.Retail, item.Quantity, item.Description)
}