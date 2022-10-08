package vrm1

import (
	"encoding/json"
	"fmt"

	"github.com/qmuntal/gltf"
)

const (
	extensionNameVRM        = "VRMC_vrm"
	extensionNameSpringBone = "VRMC_springBone"
)

func init() {
	gltf.RegisterExtension(extensionNameVRM, unmarshalVRM)
	gltf.RegisterExtension(extensionNameSpringBone, unmarshalSpringBone)
}

func unmarshalVRM(data []byte) (interface{}, error) {
	var vrm VRMExtension
	err := json.Unmarshal(data, &vrm)
	return &vrm, fmt.Errorf("failed to unmarshal json: %w", err)
}

func GetVRMExtension(doc *gltf.Document) (*VRMExtension, bool) {
	ext, ok := doc.Extensions[extensionNameVRM]
	if !ok {
		return nil, false
	}
	v, ok := ext.(*VRMExtension)
	return v, ok
}

func AddVRMExtension(doc *gltf.Document, ext *VRMExtension) {
	doc.ExtensionsUsed = append(doc.ExtensionsUsed, extensionNameVRM)
	if doc.Extensions == nil {
		doc.Extensions = gltf.Extensions{}
	}
	doc.Extensions[extensionNameVRM] = ext
}

func unmarshalSpringBone(data []byte) (interface{}, error) {
	var ext SpringBone
	err := json.Unmarshal(data, &ext)
	return &ext, fmt.Errorf("failed to unmarshal json: %w", err)
}

func GetSpringBoneExtension(doc *gltf.Document) (*SpringBone, bool) {
	ext, ok := doc.Extensions[extensionNameVRM]
	if !ok {
		return nil, false
	}
	v, ok := ext.(*SpringBone)
	return v, ok
}

func AddSpringBoneExtension(doc *gltf.Document, ext *SpringBone) {
	doc.ExtensionsUsed = append(doc.ExtensionsUsed, extensionNameSpringBone)
	if doc.Extensions == nil {
		doc.Extensions = gltf.Extensions{}
	}
	doc.Extensions[extensionNameSpringBone] = ext
}
