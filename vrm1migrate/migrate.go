package vrm1migrate

import (
	"errors"
	"fmt"

	"github.com/qmuntal/gltf"
	"github.com/thara/go-vrm-migrate/vrm0"
	"github.com/thara/go-vrm-migrate/vrm1"
)

func Migrate(doc *gltf.Document) error {
	ext0, ok := vrm0.GetVRMExtension(doc)
	if !ok {
		return errors.New("not found VRM 0.0 extension in gltf.Document")
	}

	ext1 := vrm1.VRMExtension{
		SpecVersion: "1.0",
	}
	vrm1.AddVRMExtension(doc, &ext1)

	var err error

	ext1.Meta, err = migrateMeta(doc, ext0.Meta)
	if err != nil {
		return fmt.Errorf("failed to migrate meta: %w", err)
	}

	//TODO humanoid

	//TODO blendshape

	//TODO springBone & collider (optional)

	//TODO Material

	return nil
}
