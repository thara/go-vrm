package vrm1

import (
	"encoding/json"
	"fmt"

	"github.com/qmuntal/gltf"
)

const extensionName = "VRMC_vrm"

func init() {
	gltf.RegisterExtension(extensionName, unmarshal)
}

func unmarshal(data []byte) (interface{}, error) {
	var vrm VRMExtension
	err := json.Unmarshal(data, &vrm)
	return &vrm, fmt.Errorf("failed to unmarshal json: %w", err)
}

func GetVRMExtension(doc *gltf.Document) (*VRMExtension, bool) {
	ext, ok := doc.Extensions[extensionName]
	if !ok {
		return nil, false
	}
	v, ok := ext.(*VRMExtension)
	return v, ok
}

func AddVRMExtension(doc *gltf.Document, vrm *VRMExtension) {
	doc.ExtensionsUsed = append(doc.ExtensionsUsed, extensionName)
	if doc.Extensions == nil {
		doc.Extensions = gltf.Extensions{}
	}
	doc.Extensions[extensionName] = vrm
}
