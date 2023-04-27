package models

type Menu struct {
	ID        int    `gorm:"column:id" json:"id"`
	Url       string `gorm:"column:url" json:"url"`
	Path      string `gorm:"column:path" json:"path"`
	Component string `gorm:"column:component" json:"component"`
	Name      string `gorm:"column:name" json:"name"`
	IconCls   string `gorm:"column:iconCls" json:"iconCls"`
	ParentID  int    `gorm:"column:parentId" json:"parentId"`
	Enabled   bool   `gorm:"column:enabled" json:"enabled"`
	Meta      Meta   `gorm:"embedded" json:"meta"`

	Children []Menu `json:"children"`
}

type Meta struct {
	KeepAlive   bool `gorm:"column:keepAlive" json:"keepAlive"`
	RequireAuth bool `gorm:"column:requireAuth" json:"requireAuth"`
}

func QueryMenu(id int64) []*Menu {
	menus := make([]*Menu, 0)
	DB.Table("menu").Select("menu.*").Where("parentId = 1").Order("id").Find(&menus)

	submenus := make([]*Menu, 0)
	DB.Table("menu").Select("menu.*").Where("parentId > 1").Order("id").Find(&submenus)

	for _, menu := range menus {
		for _, submenu := range submenus {
			if menu.ID == submenu.ParentID {
				menu.Children = append(menu.Children, *submenu)
			}
		}
	}

	return menus
}
