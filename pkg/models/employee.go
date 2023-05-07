package models

import (
	"time"
)

type QueryEmployee struct {
	Page int64  `form:"page"`
	Size int64  `form:"size"`
	Name string `form:"name"`
}

type Employee struct {
	ID             int                  `json:"id" gorm:"column:id"`
	Name           string               `json:"name" gorm:"column:name"`
	Gender         string               `json:"gender" gorm:"column:gender"`
	Birthday       time.Time            `json:"birthday" gorm:"column:birthday"`
	IDCard         string               `json:"idCard" gorm:"column:idCard"`
	Wedlock        string               `json:"wedlock" gorm:"column:wedlock"`
	NationID       int                  `json:"nationId" gorm:"column:nationId"`
	NativePlace    string               `json:"nativePlace" gorm:"column:nativePlace"`
	PoliticID      int                  `json:"politicId" gorm:"column:politicId"`
	Email          string               `json:"email" gorm:"column:email"`
	Phone          string               `json:"phone" gorm:"column:phone"`
	Address        string               `json:"address" gorm:"column:address"`
	DepartmentID   int                  `json:"departmentId" gorm:"column:departmentId"`
	JobLevelID     int                  `json:"jobLevelId" gorm:"column:jobLevelId"`
	PosID          int                  `json:"posId" gorm:"column:posId"`
	EngageForm     string               `json:"engageForm" gorm:"column:engageForm"`
	TiptopDegree   string               `json:"tiptopDegree" gorm:"column:tiptopDegree"`
	Specialty      string               `json:"specialty" gorm:"column:specialty"`
	School         string               `json:"school" gorm:"column:school"`
	BeginDate      time.Time            `json:"beginDate" gorm:"column:beginDate"`
	WorkState      string               `json:"workState" gorm:"column:workState"`
	WorkID         string               `json:"workID" gorm:"column:workID"`
	ContractTerm   float64              `json:"contractTerm" gorm:"column:contractTerm"`
	ConversionTime time.Time            `json:"conversionTime" gorm:"column:conversionTime"`
	NotWorkDate    time.Time            `json:"notWorkDate" gorm:"column:notWorkDate"`
	BeginContract  time.Time            `json:"beginContract" gorm:"column:beginContract"`
	EndContract    time.Time            `json:"endContract" gorm:"column:endContract"`
	WorkAge        int                  `json:"workAge" gorm:"column:workAge"`
	Nation         Nation               `json:"nation" gorm:"embedded"`
	Politicsstatus Politicsstatus       `json:"politicsstatus" gorm:"embedded"`
	Department     DepartmentSimpleInfo `json:"department" gorm:"embedded"`
	JobLevel       JobLevel             `json:"jobLevel" gorm:"embedded"`
	Position       Position             `json:"position" gorm:"embedded"`
}

type Nation struct {
	ID   int    `json:"id" gorm:"column:nid"`
	Name string `json:"name" gorm:"column:nname"`
}

type Politicsstatus struct {
	ID   int    `json:"id" gorm:"column:pid"`
	Name string `json:"name" gorm:"column:pname"`
}

type DepartmentSimpleInfo struct {
	ID   int    `json:"id" gorm:"column:did"`
	Name string `json:"name" gorm:"column:dname"`
}

type JobLevel struct {
	ID   int    `json:"id" gorm:"column:jid"`
	Name string `json:"name" gorm:"column:jname"`
}

type Position struct {
	ID   int    `json:"id" gorm:"column:posid"`
	Name string `json:"name" gorm:"column:posname"`
}

type Nation2 struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Position2 struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type JobLevel2 struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Politicsstatus2 struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func GetEmployees(query QueryEmployee) (error, *ResponsePage) {
	var offset int64
	if query.Page != 0 && query.Size != 0 {
		offset = (query.Page - 1) * query.Size
	}

	emps := make([]*Employee, 0)
	var total int64

	DB.Select("e.*, p.id as pid, p.name as pname, n.id as nid, n.name as nname, d.id as did, d.name as dname, j.id as jid, j.name as jname, pos.id as posid, pos.name as posname").
		Table("employee e").
		Joins("inner join nation n on e.nationId = n.id").
		Joins("inner join politicsstatus p on e.politicId = p.id").
		Joins("inner join department d on e.departmentId = d.id").
		Joins("inner join joblevel j on e.jobLevelId = j.id").
		Joins("inner join position pos on e.posId = pos.id").
		Where("e.name like ?", "%"+query.Name+"%").
		Count(&total).
		Offset(offset).Limit(query.Size).
		Find(&emps)

	return nil, &ResponsePage{
		Total: total,
		Data:  emps,
	}
}

func GetAllNations() (error, []*Nation2) {
	nations := make([]*Nation2, 0)
	DB.Select("*").Table("nation").Find(&nations)

	return nil, nations
}

func GetAllPositions() []*Position2 {
	positions := make([]*Position2, 0)
	DB.Select("*").Table("position").Find(&positions)

	return positions
}

func GetAllJobLevels() []*JobLevel2 {
	joblevels := make([]*JobLevel2, 0)
	DB.Select("*").Table("joblevel").Find(&joblevels)

	return joblevels
}

func GetAllPoliticsstatus() []*Politicsstatus2 {
	pss := make([]*Politicsstatus2, 0)
	DB.Select("*").Table("politicsstatus").Find(&pss)

	return pss
}

func DeleteEmployee(id int) bool {
	emp := Employee{ID: id}
	DB.Delete(&emp).Table("employee")

	return true
}
