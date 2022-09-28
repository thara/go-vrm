package vrm1migrate

import (
	"errors"
	"fmt"

	"github.com/qmuntal/gltf"
	"github.com/thara/go-vrm-migrate/vrm0"
	"github.com/thara/go-vrm-migrate/vrm1"
)

func Migrate(doc *gltf.Document) error {
	// porting https://github.com/vrm-c/UniVRM/blob/8864846a7f13ffcd6f516a7cdd304b50bf30d71a/Assets/VRM10/Runtime/Migration/MigrationVrm.cs#L71
	ext0, ok := vrm0.GetVRMExtension(doc)
	if !ok {
		return errors.New("not found VRM 0.0 extension in gltf.Document")
	}

	meshToNode := make(map[uint32]int)
	for i, n := range doc.Nodes {
		if n.Mesh != nil {
			meshToNode[*n.Mesh] = i
		}
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

	ext1.Humanoid, err = migrateHumanoid(ext0)
	if err != nil {
		return fmt.Errorf("failed to migrate humanoid: %w", err)
	}

	if ext0.BlendShapeMaster != nil {
		ext1.Expressions = migrateExpression(doc, ext0, meshToNode)
	}

	//TODO springBone & collider (optional)

	//TODO Material

	return nil
}
