// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package internal

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type VRMCMaterialsMtoonSchemaJson struct {
	// Extensions corresponds to the JSON schema field "extensions".
	Extensions interface{} `json:"extensions,omitempty"`

	// Extras corresponds to the JSON schema field "extras".
	Extras interface{} `json:"extras,omitempty"`

	// GiEqualizationFactor corresponds to the JSON schema field
	// "giEqualizationFactor".
	GiEqualizationFactor float64 `json:"giEqualizationFactor,omitempty"`

	// MatcapFactor corresponds to the JSON schema field "matcapFactor".
	MatcapFactor []float64 `json:"matcapFactor,omitempty"`

	// MatCap
	MatcapTexture interface{} `json:"matcapTexture,omitempty"`

	// OutlineColorFactor corresponds to the JSON schema field "outlineColorFactor".
	OutlineColorFactor []float64 `json:"outlineColorFactor,omitempty"`

	// OutlineLightingMixFactor corresponds to the JSON schema field
	// "outlineLightingMixFactor".
	OutlineLightingMixFactor float64 `json:"outlineLightingMixFactor,omitempty"`

	// OutlineWidthFactor corresponds to the JSON schema field "outlineWidthFactor".
	OutlineWidthFactor float64 `json:"outlineWidthFactor,omitempty"`

	// Outline
	OutlineWidthMode VRMCMaterialsMtoonSchemaJsonOutlineWidthMode `json:"outlineWidthMode,omitempty"`

	// OutlineWidthMultiplyTexture corresponds to the JSON schema field
	// "outlineWidthMultiplyTexture".
	OutlineWidthMultiplyTexture interface{} `json:"outlineWidthMultiplyTexture,omitempty"`

	// Rim
	ParametricRimColorFactor []float64 `json:"parametricRimColorFactor,omitempty"`

	// ParametricRimFresnelPowerFactor corresponds to the JSON schema field
	// "parametricRimFresnelPowerFactor".
	ParametricRimFresnelPowerFactor float64 `json:"parametricRimFresnelPowerFactor,omitempty"`

	// ParametricRimLiftFactor corresponds to the JSON schema field
	// "parametricRimLiftFactor".
	ParametricRimLiftFactor float64 `json:"parametricRimLiftFactor,omitempty"`

	// RenderQueueOffsetNumber corresponds to the JSON schema field
	// "renderQueueOffsetNumber".
	RenderQueueOffsetNumber int `json:"renderQueueOffsetNumber,omitempty"`

	// RimLightingMixFactor corresponds to the JSON schema field
	// "rimLightingMixFactor".
	RimLightingMixFactor float64 `json:"rimLightingMixFactor,omitempty"`

	// RimMultiplyTexture corresponds to the JSON schema field "rimMultiplyTexture".
	RimMultiplyTexture interface{} `json:"rimMultiplyTexture,omitempty"`

	// ShadeColorFactor corresponds to the JSON schema field "shadeColorFactor".
	ShadeColorFactor []float64 `json:"shadeColorFactor,omitempty"`

	// ShadeMultiplyTexture corresponds to the JSON schema field
	// "shadeMultiplyTexture".
	ShadeMultiplyTexture interface{} `json:"shadeMultiplyTexture,omitempty"`

	// Lighting
	ShadingShiftFactor float64 `json:"shadingShiftFactor,omitempty"`

	// ShadingShiftTexture corresponds to the JSON schema field "shadingShiftTexture".
	ShadingShiftTexture interface{} `json:"shadingShiftTexture,omitempty"`

	// ShadingToonyFactor corresponds to the JSON schema field "shadingToonyFactor".
	ShadingToonyFactor float64 `json:"shadingToonyFactor,omitempty"`

	// Specification version of VRMC_materials_mtoon
	SpecVersion string `json:"specVersion"`

	// enable depth buffer when `alphaMode` is `BLEND`
	TransparentWithZWrite bool `json:"transparentWithZWrite,omitempty"`

	// UvAnimationMaskTexture corresponds to the JSON schema field
	// "uvAnimationMaskTexture".
	UvAnimationMaskTexture interface{} `json:"uvAnimationMaskTexture,omitempty"`

	// UvAnimationRotationSpeedFactor corresponds to the JSON schema field
	// "uvAnimationRotationSpeedFactor".
	UvAnimationRotationSpeedFactor float64 `json:"uvAnimationRotationSpeedFactor,omitempty"`

	// UvAnimationScrollXSpeedFactor corresponds to the JSON schema field
	// "uvAnimationScrollXSpeedFactor".
	UvAnimationScrollXSpeedFactor float64 `json:"uvAnimationScrollXSpeedFactor,omitempty"`

	// UvAnimationScrollYSpeedFactor corresponds to the JSON schema field
	// "uvAnimationScrollYSpeedFactor".
	UvAnimationScrollYSpeedFactor float64 `json:"uvAnimationScrollYSpeedFactor,omitempty"`
}

var enumValues_VRMCMaterialsMtoonSchemaJsonOutlineWidthMode = []interface{}{
	"none",
	"worldCoordinates",
	"screenCoordinates",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *VRMCMaterialsMtoonSchemaJsonOutlineWidthMode) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_VRMCMaterialsMtoonSchemaJsonOutlineWidthMode {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_VRMCMaterialsMtoonSchemaJsonOutlineWidthMode, v)
	}
	*j = VRMCMaterialsMtoonSchemaJsonOutlineWidthMode(v)
	return nil
}

type VRMCMaterialsMtoonSchemaJsonOutlineWidthMode string

const VRMCMaterialsMtoonSchemaJsonOutlineWidthModeNone VRMCMaterialsMtoonSchemaJsonOutlineWidthMode = "none"
const VRMCMaterialsMtoonSchemaJsonOutlineWidthModeScreenCoordinates VRMCMaterialsMtoonSchemaJsonOutlineWidthMode = "screenCoordinates"
const VRMCMaterialsMtoonSchemaJsonOutlineWidthModeWorldCoordinates VRMCMaterialsMtoonSchemaJsonOutlineWidthMode = "worldCoordinates"

// UnmarshalJSON implements json.Unmarshaler.
func (j *VRMCMaterialsMtoonSchemaJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["specVersion"]; !ok || v == nil {
		return fmt.Errorf("field specVersion: required")
	}
	type Plain VRMCMaterialsMtoonSchemaJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if v, ok := raw["giEqualizationFactor"]; !ok || v == nil {
		plain.GiEqualizationFactor = 0.9
	}
	if v, ok := raw["matcapFactor"]; !ok || v == nil {
		plain.MatcapFactor = []float64{
			0,
			0,
			0,
		}
	}
	if v, ok := raw["outlineColorFactor"]; !ok || v == nil {
		plain.OutlineColorFactor = []float64{
			0,
			0,
			0,
		}
	}
	if v, ok := raw["outlineLightingMixFactor"]; !ok || v == nil {
		plain.OutlineLightingMixFactor = 1
	}
	if v, ok := raw["outlineWidthFactor"]; !ok || v == nil {
		plain.OutlineWidthFactor = 0
	}
	if v, ok := raw["outlineWidthMode"]; !ok || v == nil {
		plain.OutlineWidthMode = "none"
	}
	if v, ok := raw["parametricRimColorFactor"]; !ok || v == nil {
		plain.ParametricRimColorFactor = []float64{
			0,
			0,
			0,
		}
	}
	if v, ok := raw["parametricRimFresnelPowerFactor"]; !ok || v == nil {
		plain.ParametricRimFresnelPowerFactor = 1
	}
	if v, ok := raw["parametricRimLiftFactor"]; !ok || v == nil {
		plain.ParametricRimLiftFactor = 0
	}
	if v, ok := raw["renderQueueOffsetNumber"]; !ok || v == nil {
		plain.RenderQueueOffsetNumber = 0
	}
	if v, ok := raw["rimLightingMixFactor"]; !ok || v == nil {
		plain.RimLightingMixFactor = 0
	}
	if v, ok := raw["shadeColorFactor"]; !ok || v == nil {
		plain.ShadeColorFactor = []float64{
			1,
			1,
			1,
		}
	}
	if v, ok := raw["shadingShiftFactor"]; !ok || v == nil {
		plain.ShadingShiftFactor = 0
	}
	if v, ok := raw["shadingToonyFactor"]; !ok || v == nil {
		plain.ShadingToonyFactor = 0.9
	}
	if v, ok := raw["transparentWithZWrite"]; !ok || v == nil {
		plain.TransparentWithZWrite = false
	}
	if v, ok := raw["uvAnimationRotationSpeedFactor"]; !ok || v == nil {
		plain.UvAnimationRotationSpeedFactor = 0
	}
	if v, ok := raw["uvAnimationScrollXSpeedFactor"]; !ok || v == nil {
		plain.UvAnimationScrollXSpeedFactor = 0
	}
	if v, ok := raw["uvAnimationScrollYSpeedFactor"]; !ok || v == nil {
		plain.UvAnimationScrollYSpeedFactor = 0
	}
	*j = VRMCMaterialsMtoonSchemaJson(plain)
	return nil
}
