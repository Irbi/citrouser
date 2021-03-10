package model

type Permission struct {

	ID	uint 	`gorm:"primary_key" json:"-"`

	ViewClientsList				bool 	`json:"viewClientsList" gorm:"default:false;"`
	CreateClient				bool 	`json:"createClient" gorm:"default:false;"`
	EditClient					bool 	`json:"editClient" gorm:"default:false;"`
	DeleteClient				bool 	`json:"deleteClient" gorm:"default:false;"`
	InviteClient				bool 	`json:"inviteClient" gorm:"default:false;"`

	ViewAdvisorsList			bool 	`json:"viewAdvisorsList" gorm:"default:false;"`
	CreateAdvisor				bool 	`json:"createAdvisor" gorm:"default:false;"`
	ActivateAdvisor				bool 	`json:"activateAdvisor" gorm:"default:false;"`
	EditAdvisor					bool 	`json:"editAdvisor" gorm:"default:false;"`
	DeleteAdvisor				bool 	`json:"deleteAdvisor" gorm:"default:false;"`

	ConnectAdvsidorToClient		bool	`json:"connectAdvisorToClient" gorm:"default:false;"`

	ViewAdminsList				bool 	`json:"viewAdminsList" gorm:"default:false;"`
	CreateAdmin					bool 	`json:"createAdmin" gorm:"default:false;"`
	EditAdmin					bool 	`json:"editAdmin" gorm:"default:false;"`
	DeleteAdmin					bool 	`json:"deleteAdmin" gorm:"default:false;"`

	CreatePortfolio				bool	`json:"createPortfolio" gorm:"default:false;"`
	EditPortfolio				bool	`json:"editPortfolio" gorm:"default:false;"`
	ExportPortfolio				bool	`json:"exportPortfolio" gorm:"default:false;"`
	RequestOptimizePortfolio	bool	`json:"requestOptimizePortfolio" gorm:"default:false;"`
	OptimizePortfolio			bool	`json:"optimizePortfolio" gorm:"default:false;"`
	ViewPortfolioHistory		bool	`json:"viewPortfolioHistory" gorm:"default:false;"`

	CreatePortfolioDraft		bool	`json:"createPortfolioDraft" gorm:"default:false;"`
	EditPortfolioDraft			bool	`json:"editPortfolioDraft" gorm:"default:false;"`
	DeletePortfolioDraft		bool	`json:"deletePortfolioDraft" gorm:"default:false;"`

	CreateReport				bool	`json:"createReport" gorm:"default:false;"`
	DeleteReport				bool	`json:"deleteReport" gorm:"default:false;"`

	ManageAssets				bool	`json:"manageAssets" gorm:"default:false;"`
}

/**
0 -- checked false; mutable false
1 -- checked false; mutable true
2 -- checked true; mutable true
3 -- checked true; mutable false
 */
type PermissionPreset struct {
	ID	uint 	`gorm:"primary_key" json:"-"`

	ViewClientsList				uint 	`json:"viewClientsList" gorm:"default:0;"`
	CreateClient				uint 	`json:"createClient" gorm:"default:0;"`
	EditClient					uint 	`json:"editClient" gorm:"default:0;"`
	DeleteClient				uint 	`json:"deleteClient" gorm:"default:0;"`
	InviteClient				uint 	`json:"inviteClient" gorm:"default:0;"`

	ViewAdvisorsList			uint 	`json:"viewAdvisorsList" gorm:"default:0;"`
	CreateAdvisor				uint 	`json:"createAdvisor" gorm:"default:0;"`
	ActivateAdvisor				uint 	`json:"activateAdvisor" gorm:"default:0;"`
	EditAdvisor					uint 	`json:"editAdvisor" gorm:"default:0;"`
	DeleteAdvisor				uint 	`json:"deleteAdvisor" gorm:"default:0;"`

	ConnectAdvsidorToClient		uint	`json:"connectAdvisorToClient" gorm:"default:0;"`

	ViewAdminsList				uint 	`json:"viewAdminsList" gorm:"default:0;"`
	CreateAdmin					uint 	`json:"createAdmin" gorm:"default:0;"`
	EditAdmin					uint 	`json:"editAdmin" gorm:"default:0;"`
	DeleteAdmin					uint 	`json:"deleteAdmin" gorm:"default:0;"`

	CreatePortfolio				uint	`json:"createPortfolio" gorm:"default:0;"`
	EditPortfolio				uint	`json:"editPortfolio" gorm:"default:0;"`
	ExportPortfolio				uint	`json:"exportPortfolio" gorm:"default:0;"`
	RequestOptimizePortfolio	uint	`json:"requestOptimizePortfolio" gorm:"default:0;"`
	OptimizePortfolio			uint	`json:"optimizePortfolio" gorm:"default:0;"`
	ViewPortfolioHistory		uint	`json:"viewPortfolioHistory" gorm:"default:0;"`

	CreatePortfolioDraft		uint	`json:"createPortfolioDraft" gorm:"default:0;"`
	EditPortfolioDraft			uint	`json:"editPortfolioDraft" gorm:"default:0;"`
	DeletePortfolioDraft		uint	`json:"deletePortfolioDraft" gorm:"default:0;"`

	CreateReport				uint	`json:"createReport" gorm:"default:0;"`
	DeleteReport				uint	`json:"deleteReport" gorm:"default:0;"`

	ManageAssets				uint	`json:"manageAssets" gorm:"default:0;"`
}