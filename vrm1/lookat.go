package vrm1

import "github.com/thara/go-vrm-migrate/vrm1/internal"

type LookAt = internal.VRMCVrmLookAtSchemaJson

type LookAtRangeMap = internal.VRMCVrmLookAtRangeMapSchemaJson

type LookAtType = internal.VRMCVrmLookAtSchemaJsonType

const (
	LookAtTypeBone       = internal.VRMCVrmLookAtSchemaJsonTypeBone
	LookAtTypeExpression = internal.VRMCVrmLookAtSchemaJsonTypeExpression
)
