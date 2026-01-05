package users

type UserRelations struct {
	User
	// Carts     []ProductUser `gorm:"constraint:OnDelete:CASCADE"`
	// Addresses []Address     `gorm:"constraint:OnDelete:CASCADE"`
	// Orders    []Order       `gorm:"constraint:OnDelete:CASCADE"`
}
