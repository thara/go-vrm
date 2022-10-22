package vrm1migrate

import (
	"log"

	"github.com/qmuntal/gltf"
	"github.com/qmuntal/gltf/ext/texturetransform"
	"github.com/qmuntal/gltf/ext/unlit"
	"github.com/thara/go-vrm-migrate/vrm0"
	"github.com/thara/go-vrm-migrate/vrm1"
)

func migrateMaterial(doc *gltf.Document, ext0 *vrm0.VRMExtension) {
	// https://github.com/vrm-c/UniVRM/blob/8864846a7f13ffcd6f516a7cdd304b50bf30d71a/Assets/VRM10/Runtime/Migration/MigrationMaterials.cs#L18

	var disabledVertexColor bool
	if migrateLegacyUnlitMaterial(doc, ext0.MaterialProperties) {
		disabledVertexColor = true
	}
	if migrateUnlitTransparentZWriteMaterial(doc, ext0.MaterialProperties) {
		disabledVertexColor = true
	}
	migrateMToonMaterial(doc, ext0)

	if disabledVertexColor {
		for _, m := range doc.Meshes {
			for _, p := range m.Primitives {
				delete(p.Attributes, "COLOR_0")
			}
		}
	}
}

func migrateLegacyUnlitMaterial(doc *gltf.Document, vrm0Materials []vrm0.Material) (migrated bool) {
	for i, m := range vrm0Materials {
		newMaterial := migrateUnlitMaterial(&m)
		if newMaterial != nil {
			doc.Materials[i] = newMaterial
			migrated = true
		}
	}
	return
}

func migrateUnlitMaterial(m *vrm0.Material) *gltf.Material {
	unlitMaterial := gltf.Material{
		Name: *m.Name,
		PBRMetallicRoughness: &gltf.PBRMetallicRoughness{
			MetallicFactor:  ptr(float32(0.0)),
			RoughnessFactor: ptr(float32(1.0)),
		},
		Extensions: gltf.Extensions{
			unlit.ExtensionName: unlit.Unlit{},
		},
	}

	shaderName := ""
	if m.Shader != nil {
		shaderName = *m.Shader
	}
	switch shaderName {
	case "Unlit/Color":
		unlitMaterial.PBRMetallicRoughness.BaseColorFactor = getBaseColorFactor(m)
		unlitMaterial.PBRMetallicRoughness.BaseColorTexture = nil
	case "Unlit/Texture":
		unlitMaterial.PBRMetallicRoughness.BaseColorFactor = &baseColorDefault
		unlitMaterial.PBRMetallicRoughness.BaseColorTexture = getBaseColorTexture(m)
	case "Unlit/Transparent":
		unlitMaterial.PBRMetallicRoughness.BaseColorFactor = &baseColorDefault
		unlitMaterial.PBRMetallicRoughness.BaseColorTexture = getBaseColorTexture(m)
		unlitMaterial.AlphaMode = gltf.AlphaBlend
	case "Unlit/Transparent Cutout":
		unlitMaterial.PBRMetallicRoughness.BaseColorFactor = &baseColorDefault
		unlitMaterial.PBRMetallicRoughness.BaseColorTexture = getBaseColorTexture(m)
		unlitMaterial.AlphaMode = gltf.AlphaMask
		unlitMaterial.AlphaCutoff = getCutOff(m)
	case "VRM/UnlitTexture":
		unlitMaterial.PBRMetallicRoughness.BaseColorFactor = &baseColorDefault
		unlitMaterial.PBRMetallicRoughness.BaseColorTexture = getBaseColorTexture(m)
	case "VRM/UnlitTransparent":
		unlitMaterial.PBRMetallicRoughness.BaseColorFactor = &baseColorDefault
		unlitMaterial.PBRMetallicRoughness.BaseColorTexture = getBaseColorTexture(m)
		unlitMaterial.AlphaMode = gltf.AlphaBlend
	case "VRM/UnlitCutout":
		unlitMaterial.PBRMetallicRoughness.BaseColorFactor = &baseColorDefault
		unlitMaterial.PBRMetallicRoughness.BaseColorTexture = getBaseColorTexture(m)
		unlitMaterial.AlphaMode = gltf.AlphaMask
		unlitMaterial.AlphaCutoff = getCutOff(m)
	case "VRM/UnlitTransparentZWrite":
		return nil
	default:
		return nil
	}
	return &unlitMaterial
}

var baseColorDefault = [4]float32{1, 1, 1, 1}

func getBaseColorFactor(material *vrm0.Material) *[4]float32 {
	factor, ok := material.VectorProperties["_Color"]
	if !ok || len(factor) != 4 {
		log.Print("not found _Color in vectorProperties")
		return &baseColorDefault
	}
	f := [4]float32{float32(factor[0]), float32(factor[1]), float32(factor[2]), float32(factor[3])}
	return &f
}

func getBaseColorTexture(material *vrm0.Material) *gltf.TextureInfo {
	tex := gltf.TextureInfo{
		Index: uint32(material.TextureProperties["_MainTex"]),
	}
	offsetX, offsetY, scaleX, scaleY := getBaseColorTextureOffsetScale(material)
	tex.Extensions = gltf.Extensions{
		texturetransform.ExtensionName: texturetransform.TextureTranform{
			Offset: [2]float32{offsetX, offsetY},
			Scale:  [2]float32{scaleX, scaleY},
		},
	}
	return &tex
}

func getBaseColorTextureOffsetScale(material *vrm0.Material) (offsetX, offsetY, scaleX, scaleY float32) {
	v, ok := material.VectorProperties["_MainTex"]
	if !ok || len(v) != 4 {
		log.Print("not found _MainTex in vectorProperties")
		return 0, 0, 1, 1
	}
	scaleX = float32(v[2])
	scaleY = float32(v[3])
	offsetX = float32(v[0])
	offsetY = 1 - float32(v[1]) - scaleY
	return
}

func getCutOff(material *vrm0.Material) *float32 {
	v, ok := material.FloatProperties["_Cutoff"]
	if !ok {
		log.Print("not found _Cutoff in floatProperties")
		return ptr(float32(0.5))
	}
	return ptr(float32(v))
}

func migrateUnlitTransparentZWriteMaterial(doc *gltf.Document, vrm0Materials []vrm0.Material) (migrated bool) {

	mapper := getRenderQueueMapper(vrm0Materials)

	for i, m := range vrm0Materials {
		newMaterial, ok := migrateUnlitTransparentZWriteMaterialShader(&m, doc.Materials[i].Name, mapper)
		if ok {
			doc.Materials[i] = newMaterial
			migrated = true
		}
	}
	return
}

const unity0XDefaultRenderQueue = 2501
const maxRenderQueueOffset = 9

func getRenderQueueMapper(vrm0Materials []vrm0.Material) map[int]int {
	var renderQueues []int
	renderQueueDup := map[int]bool{}
	for _, m := range vrm0Materials {
		var renderQueue int
		if m.RenderQueue == nil || *m.RenderQueue == -1 {
			renderQueue = unity0XDefaultRenderQueue
		} else {
			renderQueue = *m.RenderQueue
		}
		if !renderQueueDup[renderQueue] {
			renderQueues = append(renderQueues, renderQueue)
		}
	}

	mapper := map[int]int{}
	var currentQueueOffset int
	for _, q := range renderQueues {
		mapper[q] = currentQueueOffset

		if maxRenderQueueOffset < currentQueueOffset+1 {
			currentQueueOffset = currentQueueOffset + 1
		} else {
			currentQueueOffset = maxRenderQueueOffset
		}
	}
	return mapper
}

func migrateUnlitTransparentZWriteMaterialShader(
	material *vrm0.Material, materialName string, renderQueueMapper map[int]int) (*gltf.Material, bool) {

	shaderName := ""
	if material.Shader != nil {
		shaderName = *material.Shader
	}
	if shaderName != "VRM/UnlitTransparentZWrite" {
		return nil, false
	}

	baseColorFactor := getBaseColorFactor(material)
	baseColorTex := getBaseColorTexture(material)
	emissiveTex := gltf.TextureInfo{
		Index:      baseColorTex.Index,
		Extensions: baseColorTex.Extensions,
	}

	var renderQueue int
	if material.RenderQueue == nil || *material.RenderQueue == -1 {
		renderQueue = unity0XDefaultRenderQueue
	}
	renderQueueOffset := renderQueueMapper[renderQueue]

	mToonMaterial := gltf.Material{
		Name: materialName,
		PBRMetallicRoughness: &gltf.PBRMetallicRoughness{
			BaseColorFactor:  ptr([4]float32{0, 0, 0, baseColorFactor[3]}),
			BaseColorTexture: baseColorTex,
			RoughnessFactor:  ptr(float32(1.0)),
		},
		AlphaMode:       gltf.AlphaBlend,
		AlphaCutoff:     ptr(float32(0.5)),
		EmissiveFactor:  [3]float32{baseColorFactor[0], baseColorFactor[1], baseColorFactor[2]},
		EmissiveTexture: &emissiveTex,
	}

	vrm1MToon := vrm1.MaterialsMToon{
		SpecVersion:             "1.0",
		TransparentWithZWrite:   true,
		RenderQueueOffsetNumber: renderQueueOffset,
		ShadeColorFactor:        []float64{0, 0, 0},
		OutlineWidthMode:        vrm1.MaterialsMtoonOutlineWidthModeNone,
	}

	mToonMaterial.Extensions = gltf.Extensions{
		vrm1.ExtensionNameMaterialsMToon: &vrm1MToon,
	}
	return &mToonMaterial, true
}
