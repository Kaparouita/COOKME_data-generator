package models

/*archive
type Product struct {
	Id               uint      `json:"ingridient_id" gorm:"primaryKey"`
	CreatedAt        time.Time `json:"create_at"`
	CategoryName     string    `json:"category_name"`
	GroupName        string    `json:"group_name"`
	Name             string    `json:"name"`
	Price            float64   `json:"price"`
	PriceKilo        float64   `json:"price_kilo"`
	DeletedPrice     float64   `json:"deleted_price"`
	DeletedPriceKilo float64   `json:"deleted_price_kilo"`
	ProductEst       string    `json:"product_est"`
	UrlName          string    `json:"url_name"`
	GroupProductsID  uint      `json:"-"`
}

type GroupProducts struct {
	Id       uint      `json:"group_products_id" gorm:"primaryKey"`
	Group    string    `json:"group"`
	Products []Product `json:"products,omitempty" gorm:"foreignKey:GroupProductsID"`
}

func (p *Product) PrintProduct() {
	fmt.Printf("ID: %d\n", p.Id)
	fmt.Printf("Category Name: %s\n", p.CategoryName)
	fmt.Printf("Group Name: %s\n", p.GroupName)
	fmt.Printf("Name: %s\n", p.Name)
	fmt.Printf("Price: %.2f\n", p.Price)
	fmt.Printf("Price per Kilo: %.2f\n", p.PriceKilo)
	fmt.Printf("Deleted Price: %.2f\n", p.DeletedPrice)
	fmt.Printf("Deleted Price per Kilo: %.2f\n", p.DeletedPriceKilo)
	fmt.Printf("Product Estimate: %s\n", p.ProductEst)
}

func (p *Product) AddGroupName(url string) {
	segments := strings.Split(url, "/")
	p.GroupName = segments[len(segments)-3]
}
*/