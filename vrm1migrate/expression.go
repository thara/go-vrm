package vrm1migrate

import (
	"github.com/qmuntal/gltf"
	"github.com/thara/go-vrm-migrate/vrm0"
	"github.com/thara/go-vrm-migrate/vrm1"
)

func migrateExpression(doc *gltf.Document, ext0 *vrm0.VRMExtension, meshToNode map[uint32]int) *vrm1.Expressions {
	var expressions vrm1.Expressions

	for _, group := range ext0.BlendShapeMaster.BlendShapeGroups {
		var name string
		if group.Name != nil {
			name = *group.Name
		}
		isBinary := group.IsBinary != nil && *group.IsBinary

		presetName := vrm0.BlendshapeGroupPresetNameUnknown
		if group.PresetName != nil {
			presetName = *group.PresetName
		}
		// https://github.com/vrm-c/UniVRM/blob/8864846a7f13ffcd6f516a7cdd304b50bf30d71a/Assets/VRM10/Runtime/Migration/MigrationVrmExpression.cs#L22
		if presetName == vrm0.BlendshapeGroupPresetNameUnknown {
			presetName = vrm0.BlendshapeGroupPresetName(name)
		}

		expression := vrm1.Expression{
			IsBinary: isBinary,
		}
		expression.MorphTargetBinds = toMorphTargetBinds(group.Binds, meshToNode)

		toMaterialColorBinds(doc, group.MaterialValues, &expression)

		switch presetName {
		case vrm0.BlendshapeGroupPresetNameAngry:
			expressions.Preset.Angry = &expression
		case vrm0.BlendshapeGroupPresetNameBlink:
			expressions.Preset.Blink = &expression
		case vrm0.BlendshapeGroupPresetNameBlinkL:
			expressions.Preset.BlinkLeft = &expression
		case vrm0.BlendshapeGroupPresetNameBlinkR:
			expressions.Preset.BlinkRight = &expression
		case vrm0.BlendshapeGroupPresetNameE:
			expressions.Preset.Ee = &expression
		case vrm0.BlendshapeGroupPresetNameFun:
			expressions.Preset.Relaxed = &expression
		case vrm0.BlendshapeGroupPresetNameI:
			expressions.Preset.Ih = &expression
		case vrm0.BlendshapeGroupPresetNameJoy:
			expressions.Preset.Happy = &expression
		case vrm0.BlendshapeGroupPresetNameLookdown:
			expressions.Preset.LookDown = &expression
		case vrm0.BlendshapeGroupPresetNameLookleft:
			expressions.Preset.LookLeft = &expression
		case vrm0.BlendshapeGroupPresetNameLookright:
			expressions.Preset.LookRight = &expression
		case vrm0.BlendshapeGroupPresetNameLookup:
			expressions.Preset.LookUp = &expression
		case vrm0.BlendshapeGroupPresetNameNeutral:
			expressions.Preset.Neutral = &expression
		case vrm0.BlendshapeGroupPresetNameO:
			expressions.Preset.Oh = &expression
		case vrm0.BlendshapeGroupPresetNameSorrow:
			expressions.Preset.Sad = &expression
		case vrm0.BlendshapeGroupPresetNameU:
			expressions.Preset.Ou = &expression
		default:
			if _, ok := expressions.Custom[name]; !ok {
				expressions.Custom[name] = expression
			}
		}
	}
	return &expressions
}

func toMorphTargetBinds(binds []vrm0.BlendshapeBind, meshToNode map[uint32]int) []vrm1.ExpressionMorphTargetBind {
	var result []vrm1.ExpressionMorphTargetBind

	for _, src := range binds {
		var dst vrm1.ExpressionMorphTargetBind

		node, ok := meshToNode[uint32(*src.Mesh)]
		if !ok {
			continue
		}
		dst.Node = node
		dst.Index = *src.Index
		dst.Weight = *src.Weight * 0.01

		result = append(result, dst)
	}
	return result
}

func toMaterialColorBinds(doc *gltf.Document, materialBinds []vrm0.BlendshapeMaterialbind, expression *vrm1.Expression) {
	for _, b := range materialBinds {
		materialName := *b.MaterialName

		var materialIndex int
		var material *gltf.Material
		for i, m := range doc.Materials {
			if m.Name == materialName {
				materialIndex = i
				material = m
				break
			}
		}
		if material == nil {
			continue
		}

		propertyName := *b.PropertyName

		switch propertyName {
		case "_MainTex_ST":
			scale, offset := verticalFlipScaleOffset(
				pt(b.TargetValue[0], b.TargetValue[1]),
				pt(b.TargetValue[2], b.TargetValue[3]))
			for _, bind := range expression.TextureTransformBinds {
				if bind.Material == materialIndex {
					expression.TextureTransformBinds = append(
						expression.TextureTransformBinds,
						vrm1.ExpressionTextureTransformBind{
							Material: materialIndex,
							Scale:    []float64{scale.x, scale.y},
							Offset:   []float64{offset.x, offset.y},
						})
					break
				}
			}
		case "_MainTex_ST_S":
			scale, offset := verticalFlipScaleOffset(
				pt(b.TargetValue[0], 1),
				pt(b.TargetValue[2], 0))

			exists := false
			for _, bind := range expression.TextureTransformBinds {
				if bind.Material == materialIndex {
					exists = true
					break
				}
			}
			if !exists {
				expression.TextureTransformBinds = append(
					expression.TextureTransformBinds,
					vrm1.ExpressionTextureTransformBind{
						Material: materialIndex,
						Scale:    []float64{scale.x, scale.y},
						Offset:   []float64{offset.x, offset.y},
					})

			}
		case "_MainTex_ST_T":
			scale, offset := verticalFlipScaleOffset(
				pt(1, b.TargetValue[1]),
				pt(0, b.TargetValue[3]))
			for _, bind := range expression.TextureTransformBinds {
				if bind.Material == materialIndex {
					expression.TextureTransformBinds = append(
						expression.TextureTransformBinds,
						vrm1.ExpressionTextureTransformBind{
							Material: materialIndex,
							Scale:    []float64{scale.x, scale.y},
							Offset:   []float64{offset.x, offset.y},
						})
					break
				}
			}
		default:
			colorType, ok := toMaterialColorType(propertyName)
			if ok {
				expression.MaterialColorBinds = append(
					expression.MaterialColorBinds,
					vrm1.ExpressionMaterialColorBind{
						Material:    materialIndex,
						Type:        colorType,
						TargetValue: b.TargetValue,
					})
			}
		}
	}
}

func toMaterialColorType(src string) (dst vrm1.ExpressionMaterialColorBindType, ok bool) {
	switch src {
	case "_Color":
		return vrm1.ExpressionMaterialColorBindTypeColor, true
	case "_EmissionColor":
		return vrm1.ExpressionMaterialColorBindTypeEmissionColor, true
	case "_RimColor":
		return vrm1.ExpressionMaterialColorBindTypeRimColor, true
	case "_OutlineColor":
		return vrm1.ExpressionMaterialColorBindTypeOutlineColor, true
	case "_ShadeColor":
		return vrm1.ExpressionMaterialColorBindTypeShadeColor, true
	}
	return
}

type point struct {
	x, y float64
}

func pt(x, y float64) point {
	return point{x, y}
}

func verticalFlipScaleOffset(s, o point) (scale, offset point) {
	return pt(s.x, s.y), pt(o.x, 1.0-o.y-s.y)
}
