// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameCCities = "ms_c_cities"

// CCities mapped from table <ms_c_cities>
type CCities struct {
	ID           int64   `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true" json:"id"`
	Code         *string `gorm:"column:code;type:varchar(255)" json:"code"`
	Name         *string `gorm:"column:name;type:varchar(255)" json:"name"`
	ProvinceCode *string `gorm:"column:province_code;type:varchar(255)" json:"province_code"`
	Leter        string  `gorm:"column:leter;type:varchar(4);not null;comment:首字母" json:"leter"`        // 首字母
	Spelling     string  `gorm:"column:spelling;type:varchar(32);not null;comment:全拼" json:"spelling"`  // 全拼
	Acronym      string  `gorm:"column:acronym;type:varchar(16);not null;comment:首字母缩写" json:"acronym"` // 首字母缩写
}

// TableName CCities's table name
func (*CCities) TableName() string {
	return TableNameCCities
}
