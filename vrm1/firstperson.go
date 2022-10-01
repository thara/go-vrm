package vrm1

import "github.com/thara/go-vrm-migrate/vrm1/internal"

type FirstPerson = internal.VRMCVrmFirstPersonSchemaJson

type FirstPersonMeshAnnotation = internal.VRMCVrmFirstPersonMeshAnnotationSchemaJson

type FirstPersonMeshAnnotationType = internal.VRMCVrmFirstPersonMeshAnnotationSchemaJsonType

const (
	FirstPersonMeshAnnotationTypeAuto            = internal.VRMCVrmFirstPersonMeshAnnotationSchemaJsonTypeAuto
	FirstPersonMeshAnnotationTypeBoth            = internal.VRMCVrmFirstPersonMeshAnnotationSchemaJsonTypeAuto
	FirstPersonMeshAnnotationTypeFirstPersonOnly = internal.VRMCVrmFirstPersonMeshAnnotationSchemaJsonTypeFirstPersonOnly
	FirstPersonMeshAnnotationTypeThirdPersonOnly = internal.VRMCVrmFirstPersonMeshAnnotationSchemaJsonTypeThirdPersonOnly
)
