package vrm1

import "github.com/thara/go-vrm-migrate/vrm1/internal"

type MaterialsMToon = internal.VRMCMaterialsMtoonSchemaJson

type MaterialsMtoonOutlineWidthMode = internal.VRMCMaterialsMtoonSchemaJsonOutlineWidthMode

const (
	MaterialsMtoonOutlineWidthModeNone              = internal.VRMCMaterialsMtoonSchemaJsonOutlineWidthModeNone
	MaterialsMtoonOutlineWidthModeScreenCoordinates = internal.VRMCMaterialsMtoonSchemaJsonOutlineWidthModeScreenCoordinates
	MaterialsMtoonOutlineWidthModeWorldCoordinates  = internal.VRMCMaterialsMtoonSchemaJsonOutlineWidthModeWorldCoordinates
)
