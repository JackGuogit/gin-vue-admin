package mimiciv

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ InExConditionApi }

var InExCdtService = service.ServiceGroupApp.MimicivServiceGroup.InExConditionService
