package models

type Department struct {
	Dname string `json:"name" gorm:"primaryKey"`
	Desc  string `json:"description" gorm:"size:255"`

	// 外鍵欄位
	ManagerID     string `json:"manager_id"`
	ViceManagerID string `json:"vice_manager_id"`

	// 關聯
	Manager     User `gorm:"foreignKey:ManagerID;references:Username"`
	ViceManager User `gorm:"foreignKey:ViceManagerID;references:Username"`
}

// 禁止外鍵
type UserExtension struct {
	User
	DepartmentID string     `json:"department_id"` // 存 Department.Dname，但不設 DB 外鍵
	Department   Department `gorm:"foreignKey:DepartmentID;references:Dname;constraint:OnDelete:SET NULL;"`
}
