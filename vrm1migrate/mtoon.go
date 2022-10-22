package vrm1migrate

import (
	"image/color"
	"log"

	"github.com/mandykoh/prism/linear"
	"github.com/mandykoh/prism/srgb"
	"github.com/qmuntal/gltf"
	"github.com/qmuntal/gltf/ext/texturetransform"
	"github.com/qmuntal/gltf/ext/unlit"
	"github.com/thara/go-vrm-migrate/vrm0"
	"github.com/thara/go-vrm-migrate/vrm1"
	"golang.org/x/image/colornames"
)

func migrateMToonMaterial(doc *gltf.Document, ext0 *vrm0.VRMExtension) {
	// https://github.com/vrm-c/UniVRM/blob/8864846a7f13ffcd6f516a7cdd304b50bf30d71a/Assets/VRM10/Runtime/Migration/Materials/MigrationMToonMaterial.cs#L17

	type mToonPair struct {
		mtoonValue *vrm0XMToonValue
		material   *gltf.Material
	}

	var sourceMaterials []mToonPair
	for i, m := range doc.Materials {
		material := ext0.MaterialProperties[i]

		shaderName := ""
		if material.Shader != nil {
			shaderName = *ext0.MaterialProperties[i].Shader
		}

		var p mToonPair
		if shaderName == "VRM/MToon" {
			p = mToonPair{
				mtoonValue: newVRM0XMToonValue(&material),
				material:   m,
			}
		}
		sourceMaterials = append(sourceMaterials, p)
	}

	transparentRenderQueues := newSortedSet[int]()
	transparentZWriteRenderQueues := newSortedSet[int]()

	for _, v := range sourceMaterials {
		mtoon := v.mtoonValue
		if mtoon == nil {
			continue
		}
		switch mtoon.definition.rendering.renderMode {
		case mtoonRenderModeOpaque, mtoonRenderModeCutout:
			//NOP
		case mtoonRenderModeTransparent:
			transparentRenderQueues.add(mtoon.definition.rendering.renderQueueOffsetNumber)
		case mtoonRenderModeTransparentWithZWrite:
			transparentZWriteRenderQueues.add(mtoon.definition.rendering.renderQueueOffsetNumber)
		default:
			//TODO error
		}
	}

	defaultTransparentQueue := 0
	transparentRenderQueueMap := map[int]int{}
	for _, src := range transparentRenderQueues.reverse() {
		transparentRenderQueueMap[src] = defaultTransparentQueue
		defaultTransparentQueue--
	}
	defaultTransparentZWriteQueue := 0
	transparentZWriteRenderQueueMap := map[int]int{}
	for _, src := range transparentZWriteRenderQueues.slice() {
		transparentZWriteRenderQueueMap[src] = defaultTransparentZWriteQueue
		defaultTransparentZWriteQueue++
	}

	//TODO next https://github.com/vrm-c/UniVRM/blob/8864846a7f13ffcd6f516a7cdd304b50bf30d71a/Assets/VRM10/Runtime/Migration/Materials/MigrationMToonMaterial.cs#L74

	for _, p := range sourceMaterials {
		if p.mtoonValue == nil {
			continue
		}
		mtoon := p.mtoonValue
		gltfMaterial := p.material

		doc.ExtensionsUsed = append(doc.ExtensionsUsed, unlit.ExtensionName)
		doc.Extensions[unlit.ExtensionName] = unlit.Unlit{}

		dst := vrm1.MaterialsMToon{
			SpecVersion: "1.0",
		}

		var texScale, texOffset *vec2
		if mtoon.textureIndexMap.mainTex != nil {
			if offsetScales, ok := mtoon.textureOffsetScales["_MainTex"]; ok {
				texScale = &vec2{offsetScales[2], offsetScales[2]}
				texOffset = &vec2{offsetScales[0], offsetScales[0]}
			}
		}

		exportTextureTransform := func(extensions gltf.Extensions) {
			if texScale != nil && texOffset != nil {
				offset := vec2{texScale.x, 1.0 - texOffset.y - texScale.y}
				extensions[texturetransform.ExtensionName] = &texturetransform.TextureTranform{
					Offset: [2]float32{float32(offset.x), float32(offset.y)},
					Scale:  [2]float32{float32(texScale.x), float32(texScale.y)},
				}
			}
		}

		// Rendering
		switch mtoon.definition.rendering.renderMode {
		case mtoonRenderModeOpaque:
			gltfMaterial.AlphaMode = gltf.AlphaOpaque
			gltfMaterial.AlphaCutoff = ptr(float32(0.5))
			dst.TransparentWithZWrite = false
			dst.RenderQueueOffsetNumber = 0
		case mtoonRenderModeCutout:
			gltfMaterial.AlphaMode = gltf.AlphaMask
			gltfMaterial.AlphaCutoff = ptr(float32(mtoon.definition.color.cutoutThresholdValue))
			dst.TransparentWithZWrite = false
			dst.RenderQueueOffsetNumber = 0
		case mtoonRenderModeTransparent:
			gltfMaterial.AlphaMode = gltf.AlphaBlend
			gltfMaterial.AlphaCutoff = ptr(float32(0.5))
			dst.TransparentWithZWrite = false
			dst.RenderQueueOffsetNumber = clamp(transparentRenderQueueMap[mtoon.definition.rendering.renderQueueOffsetNumber], -9, 0)
		case mtoonRenderModeTransparentWithZWrite:
			gltfMaterial.AlphaMode = gltf.AlphaBlend
			gltfMaterial.AlphaCutoff = ptr(float32(0.5))
			dst.TransparentWithZWrite = true
			dst.RenderQueueOffsetNumber = clamp(transparentRenderQueueMap[mtoon.definition.rendering.renderQueueOffsetNumber], 0, 9)
		default:
			//TODO error
		}
		switch mtoon.definition.rendering.cullMode {
		case mtoonCullModeOff:
			gltfMaterial.DoubleSided = true
		case mtoonCullModeFront, mtoonCullModeBack:
			gltfMaterial.DoubleSided = false
		default:
			//TODO error
		}

		gltfMaterial.PBRMetallicRoughness.BaseColorFactor = ptr(toFloat4[float32](srgb.LineariseColor(mtoon.definition.color.litColor)))
		if mtoon.textureIndexMap.mainTex != nil {
			tex := &gltf.TextureInfo{
				Index: uint32(*mtoon.textureIndexMap.mainTex),
			}
			gltfMaterial.PBRMetallicRoughness.BaseColorTexture = tex
			exportTextureTransform(tex.Extensions)
		}
		shadeColorFactor := toFloat4[float64](srgb.LineariseColor(mtoon.definition.color.shadeColor))
		dst.ShadeColorFactor = shadeColorFactor[:]
		if mtoon.textureIndexMap.shadeTexture != nil {
			tex := &gltf.TextureInfo{
				Index: uint32(*mtoon.textureIndexMap.shadeTexture),
			}
			dst.ShadeMultiplyTexture = tex
			exportTextureTransform(tex.Extensions)
		}
		//NOTE: https://github.com/vrm-c/UniVRM/blob/8864846a7f13ffcd6f516a7cdd304b50bf30d71a/Assets/VRM10/Runtime/Migration/Materials/MigrationMToonMaterial.cs#L185
		if mtoon.textureIndexMap.mainTex != nil && mtoon.textureIndexMap.shadeTexture == nil {
			tex := &gltf.TextureInfo{
				Index: uint32(*mtoon.textureIndexMap.mainTex),
			}
			dst.ShadeMultiplyTexture = tex
			exportTextureTransform(tex.Extensions)
		}
		if mtoon.textureIndexMap.bumpMap != nil {
			tex := &gltf.NormalTexture{
				Index: ptr(uint32(*mtoon.textureIndexMap.bumpMap)),
				Scale: ptr(float32(mtoon.definition.lighting.normal.normalScaleValue)),
			}
			gltfMaterial.NormalTexture = tex
			exportTextureTransform(tex.Extensions)
		}

		min, max := getShadingRange0x(mtoon.definition.lighting.litAndShadeMixing.shadingToonyValue, mtoon.definition.lighting.litAndShadeMixing.shadingShiftValue)
		dst.ShadingShiftFactor = clamp(float64(max)+float64(min)+0.5-1, -1, +1)
		dst.ShadingToonyFactor = clamp(2.0-(float64(max)-float64(min))*0.5, 0, +1)

		dst.GiEqualizationFactor = clamp[float64](mtoon.definition.lighting.lightingInfluence.giIntensityValue, 0, 1)

		// TODO: linear to linear? https://github.com/vrm-c/UniVRM/blob/8864846a7f13ffcd6f516a7cdd304b50bf30d71a/Assets/VRM10/Runtime/Migration/Materials/MigrationMToonMaterial.cs#L237
		gltfMaterial.EmissiveFactor = toFloat3[float32](mtoon.definition.emission.emissionColor)
		if mtoon.textureIndexMap.emissionMap != nil {
			tex := &gltf.TextureInfo{
				Index: uint32(*mtoon.textureIndexMap.emissionMap),
			}
			gltfMaterial.EmissiveTexture = tex
			exportTextureTransform(tex.Extensions)
		}

		if mtoon.textureIndexMap.sphereAdd != nil {
			tex := &gltf.TextureInfo{
				Index: uint32(*mtoon.textureIndexMap.sphereAdd),
			}
			dst.MatcapTexture = tex
		}
		parametricRimColorFactor := toFloat3[float64](srgb.LineariseColor(mtoon.definition.rim.rimColor))
		dst.ParametricRimColorFactor = parametricRimColorFactor[:]
		dst.ParametricRimFresnelPowerFactor = mtoon.definition.rim.rimFresnelPowerValue
		dst.ParametricRimLiftFactor = mtoon.definition.rim.rimLiftValue
		if mtoon.textureIndexMap.rimTexture != nil {
			tex := &gltf.TextureInfo{
				Index: uint32(*mtoon.textureIndexMap.rimTexture),
			}
			dst.RimMultiplyTexture = tex
			exportTextureTransform(tex.Extensions)
		}
		dst.RimLightingMixFactor = mtoon.definition.rim.rimLightingMixValue

		// Outline https://github.com/vrm-c/UniVRM/blob/8864846a7f13ffcd6f516a7cdd304b50bf30d71a/Assets/VRM10/Runtime/Migration/Materials/MigrationMToonMaterial.cs#L284
		const (
			centimeterToMeter = 0.01
			oneHundredth      = 0.01
		)
		switch mtoon.definition.outline.outlineWidthMode {
		case mtoonOutlineWidthModeNone:
			dst.OutlineWidthMode = vrm1.MaterialsMtoonOutlineWidthModeNone
			// dst.OutlineWidthFactor = nil
		case mtoonOutlineWidthModeWorldCoordinates:
			dst.OutlineWidthMode = vrm1.MaterialsMtoonOutlineWidthModeWorldCoordinates
			dst.OutlineWidthFactor = mtoon.definition.outline.outlineWidthValue * centimeterToMeter
		case mtoonOutlineWidthModeScreenCoordinates:
			dst.OutlineWidthMode = vrm1.MaterialsMtoonOutlineWidthModeScreenCoordinates
			dst.OutlineWidthFactor = mtoon.definition.outline.outlineWidthValue * oneHundredth * 0.5
		default:
			//TODO error
		}

		if mtoon.textureIndexMap.outlineWidthTexture != nil {
			tex := &gltf.TextureInfo{
				Index: uint32(*mtoon.textureIndexMap.outlineWidthTexture),
			}
			dst.OutlineWidthMultiplyTexture = tex
			exportTextureTransform(tex.Extensions)
		}
		outlineColorFactor := toFloat3[float64](srgb.LineariseColor(mtoon.definition.outline.outlineColor))
		dst.OutlineColorFactor = outlineColorFactor[:]
		switch mtoon.definition.outline.outlineColorMode {
		case mtoonOutlineColorModeFixedColor:
			dst.OutlineLightingMixFactor = 0
		case mtoonOutlineColorModeMixedLighting:
			dst.OutlineLightingMixFactor = mtoon.definition.outline.outlineLightingMixValue
		default:
			//TODO error
		}

		if mtoon.textureIndexMap.uvAnimMaskTexture != nil {
			tex := &gltf.TextureInfo{
				Index: uint32(*mtoon.textureIndexMap.uvAnimMaskTexture),
			}
			dst.UvAnimationMaskTexture = tex
			exportTextureTransform(tex.Extensions)
		}
		dst.UvAnimationRotationSpeedFactor = mtoon.definition.textureOption.uvAnimationRotationSpeedValue
		dst.UvAnimationScrollXSpeedFactor = mtoon.definition.textureOption.uvAnimationScrollXSpeedValue
		const invertY = -1.0
		dst.UvAnimationScrollYSpeedFactor = mtoon.definition.textureOption.uvAnimationScrollYSpeedValue * invertY

		gltfMaterial.Extensions[vrm1.ExtensionNameMaterialsMToon] = dst

		addExtensionsUsed(doc, vrm1.ExtensionNameMaterialsMToon)
		addExtensionsUsed(doc, texturetransform.ExtensionName)
	}
}

type vrm0XMToonValue struct {
	definition          mtoonDefinition
	textureOffsetScales map[string][]float64

	// https://github.com/vrm-c/UniVRM/blob/1d7898a04fe1f2e218873d3af3aa926b99666638/Assets/VRM10/Runtime/Migration/Materials/Vrm0XMToonTextureIndexMap.cs
	textureIndexMap struct {
		// glTF
		mainTex     *int
		bumpMap     *int
		emissionMap *int
		// VRMC_materials_mtoon
		shadeTexture         *int
		receiveShadowTexture *int
		shadingGradeTexture  *int
		rimTexture           *int
		sphereAdd            *int
		outlineWidthTexture  *int
		uvAnimMaskTexture    *int
	}
}

func newVRM0XMToonValue(material *vrm0.Material) *vrm0XMToonValue {
	// https://github.com/vrm-c/UniVRM/blob/1d7898a04fe1f2e218873d3af3aa926b99666638/Assets/VRM10/Runtime/Migration/Materials/Vrm0XMToonValue.cs#L22
	def := mtoonDefinition{}
	offsetScale := map[string][]float64{}

	for k, v := range material.VectorProperties {
		switch k {
		case "_Color":
			def.color.litColor = convertColorSpace(v, colorSpaceSRGB, colorSpaceSRGB)
		case "_ShadeColor":
			def.color.shadeColor = convertColorSpace(v, colorSpaceSRGB, colorSpaceSRGB)
		case "_EmissionColor":
			def.emission.emissionColor = convertColorSpace(v, colorSpaceLinear, colorSpaceLinear)
		case "_RimColor":
			def.rim.rimColor = convertColorSpace(v, colorSpaceSRGB, colorSpaceSRGB)
		case "_OutlineColor":
			def.outline.outlineColor = convertColorSpace(v, colorSpaceSRGB, colorSpaceSRGB)
		case "_MainTex",
			"_ShadeTexture",
			"_BumpMap",
			"_EmissionMap",
			"_OutlineWidthTexture",
			"_ReceiveShadowTexture",
			"_RimTexture",
			"_ShadingGradeTexture",
			"_SphereAdd",
			"_UvAnimMaskTexture":

			offsetScale[k] = v
		}
	}

	for k, v := range material.FloatProperties {
		switch k {
		case "_BlendMode":
			def.rendering.renderMode = mtoonRenderMode(v)
		case "_CullMode":
			def.rendering.cullMode = mtoonCullMode(v)
		case "_Cutoff":
			def.color.cutoutThresholdValue = v

		case "_BumpScale":
			def.lighting.normal.normalScaleValue = v
		case "_LightColorAttenuation":
			def.lighting.lightingInfluence.lightColorAttenuationValue = v
		case "_ShadeShift":
			def.lighting.litAndShadeMixing.shadingShiftValue = float32(v)
		case "_ShadeToony":
			def.lighting.litAndShadeMixing.shadingToonyValue = float32(v)
		case "_ShadingGradeRate", "_ReceiveShadowRate":
			// Not supported

		case "_IndirectLightIntensity":
			def.lighting.lightingInfluence.giIntensityValue = v
		case "_RimFresnelPower":
			def.rim.rimFresnelPowerValue = v
		case "_RimLift":
			def.rim.rimLiftValue = v
		case "_RimLightingMix":
			def.rim.rimLightingMixValue = v

		case "_OutlineColorMode":
			def.outline.outlineColorMode = mtoonOutlineColorMode(v)
		case "_OutlineLightingMix":
			def.outline.outlineLightingMixValue = v
		case "_OutlineScaledMaxDistance":
			def.outline.outlineScaledMaxDistanceValue = v
		case "_OutlineWidth":
			def.outline.outlineWidthValue = v
		case "_OutlineWidthMode":
			if 2 < v {
				v = 0
			}
			def.outline.outlineWidthMode = mtoonOutlineWidthMode(v)

		case "_UvAnimRotation":
			def.textureOption.uvAnimationRotationSpeedValue = v
		case "_UvAnimScrollX":
			def.textureOption.uvAnimationScrollXSpeedValue = v
		case "_UvAnimScrollY":
			def.textureOption.uvAnimationScrollYSpeedValue = v
		case "_OutlineCullMode",
			"_ZWrite",
			"_DstBlend",
			"_SrcBlend",
			"_MToonVersion",
			"_DebugMode":
			// Auto generated
		}
	}

	var value vrm0XMToonValue

	for k, v := range material.TextureProperties {
		index := v
		switch k {
		case "_MainTex":
			value.textureIndexMap.mainTex = &index
		case "_ShadeTexture":
			value.textureIndexMap.shadeTexture = &index
		case "_Bumpvalue.textureIndexMap":
			value.textureIndexMap.bumpMap = &index
		case "_ReceiveShadowTexture":
			value.textureIndexMap.receiveShadowTexture = &index
		case "_ShadingGradeTexture":
			value.textureIndexMap.shadingGradeTexture = &index
		case "_Emissionvalue.textureIndexMap":
			value.textureIndexMap.emissionMap = &index
		case "_RimTexture":
			value.textureIndexMap.rimTexture = &index
		case "_SphereAdd":
			value.textureIndexMap.sphereAdd = &index
		case "_OutlineWidthTexture":
			value.textureIndexMap.outlineWidthTexture = &index
		case "_UvAnimMaskTexture":
			value.textureIndexMap.uvAnimMaskTexture = &index
		}
	}

	def.rendering.renderQueueOffsetNumber = *material.RenderQueue - getMToonRenderQueueRequirementDefault(def.rendering.renderMode)

	return &value
}

type colorSpace int

const (
	colorSpaceSRGB colorSpace = iota
	colorSpaceLinear
)

func convertColorSpace(src []float64, srcColorSpace, dstColorSpace colorSpace) color.RGBA {
	if len(src) != 4 {
		return colornames.Magenta
	}

	srcColor := color.RGBA{
		R: uint8(src[0] * 255),
		G: uint8(src[1] * 255),
		B: uint8(src[2] * 255),
		A: 0xFF,
	}

	if srcColorSpace == colorSpaceSRGB {
		if dstColorSpace == colorSpaceSRGB {
			return srcColor
		} else {
			r, g, b, a := srgb.LineariseColor(srcColor).RGBA()
			return color.RGBA{
				R: uint8(r),
				G: uint8(g),
				B: uint8(b),
				A: uint8(a),
			}
		}
	} else {
		if dstColorSpace == colorSpaceSRGB {
			rgb, a := linear.RGBFromLinear(srcColor)
			return color.RGBA{
				R: uint8(rgb.R * 255),
				G: uint8(rgb.G * 255),
				B: uint8(rgb.B * 255),
				A: uint8(a * 255),
			}
		} else {
			return srcColor
		}
	}
}

// https://github.com/Santarh/MToon/blob/master/MToon/Scripts/MToonDefinition.cs
type mtoonDefinition struct {
	meta          metaDefinition
	rendering     renderingDefinition
	color         mtoonColorDefinition
	lighting      mtoonLightingDefinition
	emission      mtoonEmissionDefinition
	matcap        mtoonMatCapDefinition
	rim           mtoonRimDefinition
	outline       mtoonOutlineDefinition
	textureOption mtoonTextureUvCoordsDefinition
}

type metaDefinition struct {
	implementation string
	versionNumber  int
}

type renderingDefinition struct {
	renderMode              mtoonRenderMode
	cullMode                mtoonCullMode
	renderQueueOffsetNumber int
}

// https://github.com/Santarh/MToon/blob/cbf459cc899b00fe0d2e2c05d91ff83f8befaa23/MToon/Scripts/Enums.cs#L23
type mtoonRenderMode int

const (
	mtoonRenderModeOpaque mtoonRenderMode = iota
	mtoonRenderModeCutout
	mtoonRenderModeTransparent
	mtoonRenderModeTransparentWithZWrite
)

// https://github.com/Santarh/MToon/blob/cbf459cc899b00fe0d2e2c05d91ff83f8befaa23/MToon/Scripts/Enums.cs#L31
type mtoonCullMode int

const (
	mtoonCullModeOff mtoonCullMode = iota
	mtoonCullModeFront
	mtoonCullModeBack
)

type mtoonColorDefinition struct {
	litColor             color.RGBA
	shadeColor           color.RGBA
	cutoutThresholdValue float64
}

type mtoonLightingDefinition struct {
	litAndShadeMixing struct {
		shadingShiftValue float32
		shadingToonyValue float32
	}
	lightingInfluence struct {
		lightColorAttenuationValue float64
		giIntensityValue           float64
	}
	normal struct {
		normalScaleValue float64
	}
}

type mtoonEmissionDefinition struct {
	emissionColor color.RGBA
}

type mtoonMatCapDefinition struct{}

type mtoonRimDefinition struct {
	rimColor             color.RGBA
	rimLightingMixValue  float64
	rimFresnelPowerValue float64
	rimLiftValue         float64
}

type mtoonOutlineDefinition struct {
	outlineWidthMode              mtoonOutlineWidthMode
	outlineWidthValue             float64
	outlineColorMode              mtoonOutlineColorMode
	outlineColor                  color.RGBA
	outlineScaledMaxDistanceValue float64
	outlineLightingMixValue       float64
}

type mtoonOutlineColorMode int

const (
	mtoonOutlineColorModeFixedColor mtoonOutlineColorMode = iota
	mtoonOutlineColorModeMixedLighting
)

type mtoonOutlineWidthMode int

const (
	mtoonOutlineWidthModeNone mtoonOutlineWidthMode = iota
	mtoonOutlineWidthModeWorldCoordinates
	mtoonOutlineWidthModeScreenCoordinates
)

type mtoonTextureUvCoordsDefinition struct {
	uvAnimationScrollXSpeedValue  float64
	uvAnimationScrollYSpeedValue  float64
	uvAnimationRotationSpeedValue float64
}

func getMToonRenderQueueRequirementDefault(renderMode mtoonRenderMode) int {
	const (
		shaderDefaultQueue    = -1
		firstTransparentQueue = 2501
	)

	switch renderMode {
	case mtoonRenderModeOpaque:
		return shaderDefaultQueue
	case mtoonRenderModeCutout:
		//TODO unity RenderQueue.AlphaTest
	case mtoonRenderModeTransparent:
		//TODO unity RenderQueue.Transparent
	case mtoonRenderModeTransparentWithZWrite:
		return firstTransparentQueue
	}
	log.Printf("unrecognized mtoonRenderMode: %d", renderMode)
	return shaderDefaultQueue
}
