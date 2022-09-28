package vrmmigrate

import (
	"errors"
	"fmt"

	"github.com/qmuntal/gltf"
	"github.com/thara/go-vrm-migrate/vrm0"
	"github.com/thara/go-vrm-migrate/vrm1x"
)

func Migrate(doc *gltf.Document) error {
	vrm0, ok := vrm0.GetVRMExtension(doc)
	if !ok {
		return errors.New("not found VRM 0.0 extension in gltf.Document")
	}

	vrm1 := vrm1x.VRMExtension{
		SpecVersion: "1.0",
	}
	vrm1x.AddVRMExtension(doc, &vrm1)

	var err error

	vrm1.Meta, err = migrateMeta(doc, vrm0.Meta)
	if err != nil {
		return fmt.Errorf("failed to migrate meta: %w", err)
	}

	//TODO humanoid

	//TODO blendshape

	//TODO springBone & collider (optional)

	//TODO Material

	return nil
}
