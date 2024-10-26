package mimiciv

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/mimiciv"
	mimicivReq "github.com/flipped-aurora/gin-vue-admin/server/model/mimiciv/request"
)

type InExConditionService struct{}

// CreateInExCondition 创建纳入排除条件记录
// Author [yourname](https://github.com/yourname)
func (InExCdtService *InExConditionService) CreateInExCondition(InExCdt *mimiciv.InExCondition) (err error) {
	err = global.GVA_DB.Create(InExCdt).Error
	return err
}

// DeleteInExCondition 删除纳入排除条件记录
// Author [yourname](https://github.com/yourname)
func (InExCdtService *InExConditionService) DeleteInExCondition(ID string) (err error) {
	err = global.GVA_DB.Delete(&mimiciv.InExCondition{}, "id = ?", ID).Error
	return err
}

// DeleteInExConditionByIds 批量删除纳入排除条件记录
// Author [yourname](https://github.com/yourname)
func (InExCdtService *InExConditionService) DeleteInExConditionByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]mimiciv.InExCondition{}, "id in ?", IDs).Error
	return err
}

// UpdateInExCondition 更新纳入排除条件记录
// Author [yourname](https://github.com/yourname)
func (InExCdtService *InExConditionService) UpdateInExCondition(InExCdt mimiciv.InExCondition) (err error) {
	err = global.GVA_DB.Model(&mimiciv.InExCondition{}).Where("id = ?", InExCdt.ID).Updates(&InExCdt).Error
	return err
}

// GetInExCondition 根据ID获取纳入排除条件记录
// Author [yourname](https://github.com/yourname)
func (InExCdtService *InExConditionService) GetInExCondition(ID string) (InExCdt mimiciv.InExCondition, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&InExCdt).Error
	return
}

// GetInExConditionInfoList 分页获取纳入排除条件记录
// Author [yourname](https://github.com/yourname)
func (InExCdtService *InExConditionService) GetInExConditionInfoList(info mimicivReq.InExConditionSearch) (list []mimiciv.InExCondition, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&mimiciv.InExCondition{})
	var InExCdts []mimiciv.InExCondition
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.MinAge != nil {
		db = db.Where("min_age > ?", info.MinAge)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&InExCdts).Error
	return InExCdts, total, err
}
func (InExCdtService *InExConditionService) GetInExConditionPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
