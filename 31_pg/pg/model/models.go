package model

type Tokens struct {
	//gorm.Model
	Id           uint   `json:"id,omitempty" gorm:"primaryKey"`
	UserId       uint   `json:"user_id,omitempty"`
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	Exp          int64  `json:"exp,omitempty"`
}
//
//type User struct {
//	gorm.Model
//	Id    uint
//	Login string
//	Name  string
//	Pass  string
//}
//
//type Task struct {
//	gorm.Model
//	Id          uint   `json:"id"`
//	UserId      uint   `json:"user_id"`
//	Title       string `json:"title"`
//	Description string `json:"description"`
//	//Completed   bool   `json:"completed" sql:"DEFAULT:false;type:boolean;index" gorm:"not null"`
//	StartDate   int64  `json:"start_date"`
//	EndDate     int64  `json:"end_date"`
//}
