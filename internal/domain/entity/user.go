package entity

type UserEntity struct {
	//model.BaseModelEntity
	DisplayName string
	Avatar      string
	Phone       string
	Password    string
	Email       *string
}

func (*UserEntity) TableName() string {
	return "users"
}
