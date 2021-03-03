package model

type Portfolio struct {
	BaseModel
	SourceID 		uint 				`json:"-"`
	Source 			*PortfolioSource 	`gorm:"foreignkey:SourceID;"`
	ClientID 		uint	 			`json:"-"`
	Client			*User				`gorm:"foreignkey:UserID;"`
	Status			string				`json:"status" gorm:"varchar(255);"`
	AssigneeID	 	uint 				`json:"-"`
	Assignee	 	*User 				`gorm:"foreignkey:UserID;"`
}
