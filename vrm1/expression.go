package vrm1

import "github.com/thara/go-vrm-migrate/vrm1/internal"

type Expressions = internal.VRMCVrmExpressionsSchemaJson

type ExpressionsCustom = internal.VRMCVrmExpressionsSchemaJsonCustom

type ExpressionsPreset = internal.VRMCVrmExpressionsSchemaJsonPreset

type Expression = internal.VRMCVrmExpressionsExpressionSchemaJson

type ExpressionMaterialColorBind = internal.VRMCVrmExpressionsExpressionMaterialColorBindSchemaJson

type ExpressionMaterialColorBindType = internal.VRMCVrmExpressionsExpressionMaterialColorBindSchemaJsonType

const (
	ExpressionMaterialColorBindTypeColor         = internal.VRMCVrmExpressionsExpressionMaterialColorBindSchemaJsonTypeColor
	ExpressionMaterialColorBindTypeEmissionColor = internal.VRMCVrmExpressionsExpressionMaterialColorBindSchemaJsonTypeEmissionColor
	ExpressionMaterialColorBindTypeShadeColor    = internal.VRMCVrmExpressionsExpressionMaterialColorBindSchemaJsonTypeShadeColor
	ExpressionMaterialColorBindTypeMatcapColor   = internal.VRMCVrmExpressionsExpressionMaterialColorBindSchemaJsonTypeMatcapColor
	ExpressionMaterialColorBindTypeRimColor      = internal.VRMCVrmExpressionsExpressionMaterialColorBindSchemaJsonTypeRimColor
	ExpressionMaterialColorBindTypeOutlineColor  = internal.VRMCVrmExpressionsExpressionMaterialColorBindSchemaJsonTypeOutlineColor
)

type ExpressionMorphTargetBind = internal.VRMCVrmExpressionsExpressionMorphTargetBindSchemaJson

type ExpressionOverrideBlink = internal.VRMCVrmExpressionsExpressionSchemaJsonOverrideBlink

const (
	ExpressionMorphTargetBindNone  = internal.VRMCVrmExpressionsExpressionSchemaJsonOverrideBlinkNone
	ExpressionMorphTargetBindBlock = internal.VRMCVrmExpressionsExpressionSchemaJsonOverrideBlinkBlock
	ExpressionMorphTargetBindBlend = internal.VRMCVrmExpressionsExpressionSchemaJsonOverrideBlinkBlend
)

type ExpressionOverrideLookAt = internal.VRMCVrmExpressionsExpressionSchemaJsonOverrideLookAt

const (
	ExpressionOverrideLookAtNone  = internal.VRMCVrmExpressionsExpressionSchemaJsonOverrideLookAtNone
	ExpressionOverrideLookAtBlock = internal.VRMCVrmExpressionsExpressionSchemaJsonOverrideLookAtBlock
	ExpressionOverrideLookAtBlend = internal.VRMCVrmExpressionsExpressionSchemaJsonOverrideLookAtBlend
)

type ExpressionOverrideMouth = internal.VRMCVrmExpressionsExpressionSchemaJsonOverrideMouth

const (
	ExpressionOverrideMouthNone  = internal.VRMCVrmExpressionsExpressionSchemaJsonOverrideMouthNone
	ExpressionOverrideMouthBlock = internal.VRMCVrmExpressionsExpressionSchemaJsonOverrideMouthBlock
	ExpressionOverrideMouthBlend = internal.VRMCVrmExpressionsExpressionSchemaJsonOverrideMouthBlend
)

type ExpressionTextureTransformBind = internal.VRMCVrmExpressionsExpressionTextureTransformBindSchemaJson