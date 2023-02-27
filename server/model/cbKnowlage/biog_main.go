// 自动生成模板BiogMain
package cbKnowlage

// BiogMain 结构体
type BiogMain struct {
	ID uint `json:"id" form:"id" gorm:"column:c_id;primary_key"` // 主键ID
	//ID       uint   `gorm:"primary_key"` // 主键ID
	PersonId *int   `json:"personId" form:"personId" gorm:"column:c_personid;comment:人物ID;"`
	Name     string `json:"name" form:"name" gorm:"column:c_name;comment:人物名称;"`
}

// TableName BiogMain 表名
func (BiogMain) TableName() string {
	return "BIOG_MAIN"
}
