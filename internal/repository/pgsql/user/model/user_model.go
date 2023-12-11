package model_user

type User struct {
	Id        string `gorm:"column:id"`
	FirstName string `gorm:"column:first_name"`
	LastName  string `gorm:"column:last_name"`
	Email     string `gorm:"column:email"`
}

func (User) TableName() string {
	return "users"
}
