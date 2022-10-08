package vrm1migrate

import (
	"errors"
	"fmt"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/qmuntal/gltf"

	"github.com/thara/go-vrm-migrate/vrm0"
	"github.com/thara/go-vrm-migrate/vrm1"
)

func migrateSpringBone(doc *gltf.Document, ext0 *vrm0.VRMExtension) (*vrm1.SpringBone, error) {
	// https://github.com/vrm-c/UniVRM/blob/8864846a7f13ffcd6f516a7cdd304b50bf30d71a/Assets/VRM10/Runtime/Migration/MigrationVrmSpringBone.cs#L183
	springBone := vrm1.SpringBone{
		SpecVersion: "1.0",
	}
	for _, colGrp := range ext0.SecondaryAnimation.ColliderGroups {
		var colliders []int
		if colGrp.Node == nil {
			continue
		}
		node := *colGrp.Node
		for _, col := range colGrp.Colliders {
			if col.Offset == nil || col.Radius == nil {
				continue
			}
			if col.Offset.X == nil || col.Offset.Y == nil || col.Offset.Z == nil {
				continue
			}
			colliders = append(colliders, len(springBone.Colliders))
			offset := migrateVec(*col.Offset.X, *col.Offset.Y, *col.Offset.Z)
			springBone.Colliders = append(springBone.Colliders, vrm1.SpringBoneCollider{
				Node: node,
				Shape: vrm1.SpringBoneColliderShape{
					Sphere: &vrm1.SpringBoneColliderShapeSphere{
						Offset: offset,
						Radius: *col.Radius,
					},
				},
			})
		}
		springBone.ColliderGroups = append(springBone.ColliderGroups, vrm1.SpringBoneColliderGroup{
			Colliders: colliders,
		})
	}

	for i, boneGroup := range ext0.SecondaryAnimation.BoneGroups {
		springs, err := migrateSpringBoneGroup(doc, &boneGroup)
		if err != nil {
			return nil, fmt.Errorf("secondaryAnimation.BoneGroups[%d]: %w", i, err)
		}
		for _, s := range springs {
			springBone.Springs = append(springBone.Springs, *s)
		}
	}

	return &springBone, nil
}

func migrateSpringBoneGroup(
	doc *gltf.Document,
	boneGroup *vrm0.SecondaryAnimationSpring,
) (springs []*vrm1.SpringBoneSpring, err error) {
	// https://github.com/vrm-c/UniVRM/blob/8864846a7f13ffcd6f516a7cdd304b50bf30d71a/Assets/VRM10/Runtime/Migration/MigrationVrmSpringBone.cs#L10

	if boneGroup.DragForce == nil {
		return nil, errors.New("dragForce is nil")
	}
	if boneGroup.GravityDir == nil ||
		boneGroup.GravityDir.X == nil || boneGroup.GravityDir.Y == nil || boneGroup.GravityDir.Z == nil {
		return nil, errors.New("gravityDir is empty")
	}
	if boneGroup.GravityPower == nil {
		return nil, errors.New("gravityPower is nil")
	}
	if boneGroup.HitRadius == nil {
		return nil, errors.New("hitRadius is nil")
	}
	if boneGroup.Stiffiness == nil {
		return nil, errors.New("stiffiness is nil")
	}

	for _, rootBone := range boneGroup.Bones {
		if rootBone < 0 || len(doc.Nodes) <= rootBone {
			continue
		}
		springs, err = createJointRecursive(doc, boneGroup, uint32(rootBone), 1, nil, springs)
		if err != nil {
			return nil, fmt.Errorf("createJointRecursive(%d): %w", rootBone, err)
		}
	}
	return
}

func createJointRecursive(
	doc *gltf.Document, boneGroup *vrm0.SecondaryAnimationSpring,
	nodeIndex uint32, level int, spring *vrm1.SpringBoneSpring,
	springs []*vrm1.SpringBoneSpring) ([]*vrm1.SpringBoneSpring, error) {
	var err error

	if spring == nil && 0 < level {
		spring = &vrm1.SpringBoneSpring{
			Name:           boneGroup.Comment,
			ColliderGroups: boneGroup.ColliderGroups,
		}
		springs = append(springs, spring)
	}
	if spring != nil {
		spring.Joints = append(spring.Joints, vrm1.SpringBoneJoint{
			Node:         nodeIndex,
			DragForce:    *boneGroup.DragForce,
			GravityDir:   migrateVec(*boneGroup.GravityDir.X, *boneGroup.GravityDir.Y, *boneGroup.GravityDir.Z),
			GravityPower: *boneGroup.GravityPower,
			HitRadius:    *boneGroup.HitRadius,
			Stiffness:    *boneGroup.Stiffiness,
		})
	}

	if 0 < len(doc.Nodes[nodeIndex].Children) {
		for i, childIndex := range doc.Nodes[nodeIndex].Children {
			if childIndex < 0 || len(doc.Nodes) < int(childIndex) {
				continue
			}
			if i == 0 {
				springs, err = createJointRecursive(doc, boneGroup, childIndex, level+1, spring, springs)
			} else {
				springs, err = createJointRecursive(doc, boneGroup, childIndex, 0, nil, springs)
			}
		}
	} else {
		if spring != nil && 0 < len(spring.Joints) {
			last := spring.Joints[len(spring.Joints)-1]
			if last.Node == nil {
				return springs, nil
			}

			// https://github.com/vrm-c/UniVRM/blob/8864846a7f13ffcd6f516a7cdd304b50bf30d71a/Assets/VRM10/Runtime/Migration/MigrationVrmSpringBone.cs#L139
			lastNode := doc.Nodes[last.Node.(int)]
			name := ""
			if 0 < len(lastNode.Name) {
				name = lastNode.Name
			}
			v1 := mgl32.Vec3{lastNode.Translation[0], lastNode.Translation[1], lastNode.Translation[2]}
			delta := v1.Normalize().Mul(0.07)
			tail := &gltf.Node{
				Name:        name + "_end",
				Translation: [3]float32{delta.X(), delta.Y(), delta.Z()},
			}
			tailIndex := len(doc.Nodes)
			doc.Nodes = append(doc.Nodes, tail)
			if 0 < len(lastNode.Children) {
				return nil, errors.New("invalid last node")
			}
			lastNode.Children = append(lastNode.Children, uint32(tailIndex))

			spring.Joints = append(spring.Joints, vrm1.SpringBoneJoint{
				Node: tailIndex,
			})
		}
	}
	return springs, err
}
