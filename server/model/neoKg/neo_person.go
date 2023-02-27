// 自动生成模板BiogMain
package neoKg

// NeoPerson 结构体
type NeoPerson struct {
	//ID       uint   `gorm:"primary_key"` // 主键ID
	PersonId  int    `json:"personId" form:"personId" `
	Name      string `json:"name" form:"name"`
	IndexYear int    `json:"indexYear" form:"indexYear"`
}
