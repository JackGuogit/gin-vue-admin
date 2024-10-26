package mimiciv

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ InExConditionRouter }

var InExCdtApi = api.ApiGroupApp.MimicivApiGroup.InExConditionApi
