// 自动生成模板InExCondition
package mimiciv

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 纳入排除条件 结构体  InExCondition
type InExCondition struct {
	global.GVA_MODEL
	MinAge *int `json:"minAge" form:"minAge" gorm:"default:18;column:min_age;comment:;" binding:"required"` //最小年龄
	MaxAge *int `json:"maxAge" form:"maxAge" gorm:"default:120;column:max_age;comment:;"`                   //最大年龄
}

// TableName 纳入排除条件 InExCondition自定义表名 in_ex_condition
func (InExCondition) TableName() string {
	return "in_ex_condition"
}
