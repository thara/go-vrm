package vrm0

import "github.com/thara/go-vrm-migrate/vrm0/internal"

type FirstPerson = internal.VrmFirstpersonSchemaJson

type FirstPersonLookAtType = internal.VrmFirstpersonSchemaJsonLookAtTypeName

const (
	FirstPersonLookAtTypeBlendShape = internal.VrmFirstpersonSchemaJsonLookAtTypeNameBlendShape
	FirstPersonLookAtTypeBone       = internal.VrmFirstpersonSchemaJsonLookAtTypeNameBone
)

type FirstPersonDegreeMap = internal.VrmFirstpersonDegreemapSchemaJson

type FirstPersonMeshAnnotation = internal.VrmFirstpersonMeshannotationSchemaJson
