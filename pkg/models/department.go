package models

type Department struct {
	ID       int    `gorm:"column:id" json:"id"`
	Name     string `gorm:"column:name" json:"name"`
	ParentID int    `gorm:"column:parentId" json:"parentId"`
	DepPath  string `gorm:"column:depPath" json:"depPath"`
	Enabled  bool   `gorm:"column:enabled" json:"enabled"`
	IsParent bool   `gorm:"column:isParent" json:"isParent"`

	Children []*Department `json:"children"`
}

func GetAllDepartments(parentid int) []*Department {
	deps := make([]*Department, 0)
	DB.Table("department").Select("*").Where("parentId = ?", parentid).Find(&deps)

	for _, dep := range deps {
		if !dep.IsParent {
			continue
		}
		parentid = dep.ID
		dep.Children = GetAllDepartments(parentid)
	}

	return deps
}
