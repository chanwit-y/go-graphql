package entity

type Customer struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	CustomerID  int    `gorm:"type:INT" json:"customer_id"`
	Name        string `gorm:"type:VARCHAR(100)" json:"name"`
	DateOfBirth string `gorm:"type:DATE" json:"date_of_birth"`
	City        string `gorm:"type:VARCHAR(100)" json:"city"`
	ZipCode     string `gorm:"type:VARCHAR(10)" json:"zip_code"`
	Status      int    `gorm:"type:INT" json:"status"`
}
