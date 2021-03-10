package model

type Portfolio struct {
	BaseModel
	ClientID 		uint	 			`json:"-"`
	Client			*User				`gorm:"foreignkey:UserID;"`
	Status			string				`json:"status" gorm:"varchar(255);"`
	State			string				`json:"status" gorm:"varchar(255);"`
	AssigneeID	 	uint 				`json:"-"`
	Assignee	 	*User 				`gorm:"foreignkey:UserID;"`
}
