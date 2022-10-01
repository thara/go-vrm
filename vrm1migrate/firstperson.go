package vrm1migrate

import (
	"fmt"
	"strings"

	"github.com/qmuntal/gltf"
	"github.com/thara/go-vrm-migrate/vrm0"
	"github.com/thara/go-vrm-migrate/vrm1"
)

func migrateFirstPerson(doc *gltf.Document, ext0 *vrm0.VRMExtension) (*vrm1.LookAt, *vrm1.FirstPerson, error) {
	var lookAtTypeName vrm0.FirstPersonLookAtType
	if ext0.FirstPerson.LookAtTypeName == nil {
		lookAtTypeName = vrm0.FirstPersonLookAtTypeBone
	} else {
		lookAtTypeName = *ext0.FirstPerson.LookAtTypeName
	}

	defaultX := 90.0
	var defaultY float64
	switch lookAtTypeName {
	case vrm0.FirstPersonLookAtTypeBone:
		defaultY = 10.0
	case vrm0.FirstPersonLookAtTypeBlendShape:
		defaultY = 1.0
	default:
		return nil, nil, fmt.Errorf("unsupported firstPerson.lookAtTypeName: %s", lookAtTypeName)
	}

	var err error

	var lookAt vrm1.LookAt
	lookAt.Type = ptr(vrm1.LookAtType(lookAtTypeName))

	lookAt.RangeMapHorizontalInner = migrateLookAtRangeMap(ext0.FirstPerson.LookAtHorizontalInner, defaultX, defaultY)
	lookAt.RangeMapHorizontalOuter = migrateLookAtRangeMap(ext0.FirstPerson.LookAtHorizontalOuter, defaultX, defaultY)
	lookAt.RangeMapVerticalDown = migrateLookAtRangeMap(ext0.FirstPerson.LookAtVerticalDown, defaultX, defaultY)
	lookAt.RangeMapVerticalUp = migrateLookAtRangeMap(ext0.FirstPerson.LookAtVerticalUp, defaultX, defaultY)
	lookAt.OffsetFromHeadBone, err = migrateVec(
		ext0.FirstPerson.FirstPersonBoneOffset.X,
		ext0.FirstPerson.FirstPersonBoneOffset.Y,
		ext0.FirstPerson.FirstPersonBoneOffset.Z)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to migrate vrm0 firstPerson.firstPersonBoneOffset: %w", err)
	}

	var firstPerson vrm1.FirstPerson
	for _, m := range ext0.FirstPerson.MeshAnnotations {
		if meshIndex, ok := migrateFirstPersonMeshIndex(m.Mesh, doc); !ok {
			firstPerson.MeshAnnotations = append(firstPerson.MeshAnnotations, vrm1.FirstPersonMeshAnnotation{
				Node: meshIndex,
				Type: migrateFirstPersonType(m.FirstPersonFlag),
			})
		}
	}
	return &lookAt, &firstPerson, nil
}

func migrateLookAtRangeMap(degreeMap *vrm0.FirstPersonDegreeMap, defaultX, defaultY float64) *vrm1.LookAtRangeMap {
	// https://github.com/vrm-c/UniVRM/blob/8864846a7f13ffcd6f516a7cdd304b50bf30d71a/Assets/VRM10/Runtime/Migration/MigrationVrmFirstPersonAndLookAt.cs#L11

	if degreeMap != nil && degreeMap.XRange != nil && degreeMap.YRange != nil {
		return &vrm1.LookAtRangeMap{
			InputMaxValue: degreeMap.XRange,
			OutputScale:   degreeMap.YRange,
		}
	}

	return &vrm1.LookAtRangeMap{
		InputMaxValue: &defaultX,
		OutputScale:   &defaultY,
	}
}

func migrateFirstPersonMeshIndex(meshIndex *int, doc *gltf.Document) (int, bool) {
	// https://github.com/vrm-c/UniVRM/blob/8864846a7f13ffcd6f516a7cdd304b50bf30d71a/Assets/VRM10/Runtime/Migration/MigrationVrmFirstPersonAndLookAt.cs#L68
	if meshIndex == nil {
		return 0, false
	}

	mesh := uint32(*meshIndex)
	for i, n := range doc.Nodes {
		if n.Mesh != nil && *n.Mesh == mesh {
			return i, true
		}
	}
	return 0, false
}

func migrateFirstPersonType(v *string) vrm1.FirstPersonMeshAnnotationType {
	// https://github.com/vrm-c/UniVRM/blob/8864846a7f13ffcd6f516a7cdd304b50bf30d71a/Assets/VRM10/Runtime/Migration/MigrationVrmFirstPersonAndLookAt.cs#L48
	if v == nil {
		return vrm1.FirstPersonMeshAnnotationTypeAuto
	}
	switch strings.ToLower(*v) {
	case "auto":
		return vrm1.FirstPersonMeshAnnotationTypeAuto
	case "both":
		return vrm1.FirstPersonMeshAnnotationTypeBoth
	case "thirdpersononly":
		return vrm1.FirstPersonMeshAnnotationTypeThirdPersonOnly
	case "firstpersononly":
		return vrm1.FirstPersonMeshAnnotationTypeFirstPersonOnly
	}
	return vrm1.FirstPersonMeshAnnotationTypeAuto
}

func ptr[T any](v T) *T {
	return &v
}
