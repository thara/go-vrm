package vrm1migrate

import (
	"fmt"

	"github.com/thara/go-vrm-migrate/vrm0"
	"github.com/thara/go-vrm-migrate/vrm1"
)

func migrateHumanoid(ext0 *vrm0.VRMExtension) (vrm1.Humanoid, error) {
	humanoid := vrm1.Humanoid{}

	for _, bone := range ext0.Humanoid.HumanBones {
		if bone.Bone == nil {
			continue
		}

		switch *bone.Bone {
		case vrm0.HumanoidBoneTypeChest:
			humanoid.HumanBones.Chest = &vrm1.HumanoidBone{Node: bone.Node}
		case vrm0.HumanoidBoneTypeHead:
			humanoid.HumanBones.Head = vrm1.HumanoidBone{Node: bone.Node}
		case vrm0.HumanoidBoneTypeHips:
			humanoid.HumanBones.Hips = vrm1.HumanoidBone{Node: bone.Node}
		case vrm0.HumanoidBoneTypeJaw:
			humanoid.HumanBones.Jaw = &vrm1.HumanoidBone{Node: bone.Node}
		case vrm0.HumanoidBoneTypeLeftEye:
			humanoid.HumanBones.LeftEye = &vrm1.HumanoidBone{Node: bone.Node}
		case vrm0.HumanoidBoneTypeLeftFoot:
			humanoid.HumanBones.LeftFoot = vrm1.HumanoidBone{Node: bone.Node}
		case vrm0.HumanoidBoneTypeLeftHand:
			humanoid.HumanBones.LeftHand = vrm1.HumanoidBone{Node: bone.Node}
		case vrm0.HumanoidBoneTypeLeftIndexDistal:
			humanoid.HumanBones.LeftIndexDistal = &vrm1.HumanoidBone{Node: bone.Node}
		case vrm0.HumanoidBoneTypeLeftIndexIntermediate:
			humanoid.HumanBones.LeftIndexIntermediate = &vrm1.HumanoidBone{Node: bone.Node}
		case vrm0.HumanoidBoneTypeLeftIndexProximal:
			humanoid.HumanBones.LeftIndexProximal = &vrm1.HumanoidBone{Node: bone.Node}
		case vrm0.HumanoidBoneTypeLeftLittleDistal:
			humanoid.HumanBones.LeftLittleDistal = &vrm1.HumanoidBone{Node: bone.Node}
		case vrm0.HumanoidBoneTypeLeftLittleIntermediate:
			humanoid.HumanBones.LeftLittleIntermediate = &vrm1.HumanoidBone{Node: bone.Node}
		case vrm0.HumanoidBoneTypeLeftLittleProximal:
			humanoid.HumanBones.LeftLittleProximal = &vrm1.HumanoidBone{Node: bone.Node}
		case vrm0.HumanoidBoneTypeLeftLowerArm:
			humanoid.HumanBones.LeftLowerArm = vrm1.HumanoidBone{Node: bone.Node}
		case vrm0.HumanoidBoneTypeLeftLowerLeg:
			humanoid.HumanBones.LeftLowerLeg = vrm1.HumanoidBone{Node: bone.Node}
		case vrm0.HumanoidBoneTypeLeftMiddleDistal:
			humanoid.HumanBones.LeftMiddleDistal = &vrm1.HumanoidBone{Node: bone.Node}
		case vrm0.HumanoidBoneTypeLeftMiddleIntermediate:
			humanoid.HumanBones.LeftMiddleIntermediate = &vrm1.HumanoidBone{Node: bone.Node}
		case vrm0.HumanoidBoneTypeLeftMiddleProximal:
			humanoid.HumanBones.LeftMiddleProximal = &vrm1.HumanoidBone{Node: bone.Node}
		case vrm0.HumanoidBoneTypeLeftRingDistal:
			humanoid.HumanBones.LeftRingDistal = &vrm1.HumanoidBone{Node: bone.Node}
		case vrm0.HumanoidBoneTypeLeftRingIntermediate:
			humanoid.HumanBones.LeftRingIntermediate = &vrm1.HumanoidBone{Node: bone.Node}
		case vrm0.HumanoidBoneTypeLeftRingProximal:
			humanoid.HumanBones.LeftRingProximal = &vrm1.HumanoidBone{Node: bone.Node}
		case vrm0.HumanoidBoneTypeLeftShoulder:
			humanoid.HumanBones.LeftShoulder = &vrm1.HumanoidBone{Node: bone.Node}
		case vrm0.HumanoidBoneTypeLeftThumbDistal:
			humanoid.HumanBones.LeftThumbDistal = &vrm1.HumanoidBone{Node: bone.Node}
		case vrm0.HumanoidBoneTypeLeftThumbIntermediate:
			//NOTE https://github.com/vrm-c/UniVRM/blob/6ba7499a2f18101ac1445e8468e527aa40a24d67/Assets/VRM10/Runtime/Migration/MigrationVrmHumanoid.cs#L74
			humanoid.HumanBones.LeftThumbProximal = &vrm1.HumanoidBone{Node: bone.Node}
		case vrm0.HumanoidBoneTypeLeftThumbProximal:
			//NOTE https://github.com/vrm-c/UniVRM/blob/6ba7499a2f18101ac1445e8468e527aa40a24d67/Assets/VRM10/Runtime/Migration/MigrationVrmHumanoid.cs#L73
			humanoid.HumanBones.LeftThumbMetacarpal = &vrm1.HumanoidBone{Node: bone.Node}
		case vrm0.HumanoidBoneTypeLeftToes:
			humanoid.HumanBones.LeftToes = &vrm1.HumanoidBone{Node: bone.Node}
		case vrm0.HumanoidBoneTypeLeftUpperArm:
			humanoid.HumanBones.LeftUpperArm = vrm1.HumanoidBone{Node: bone.Node}
		case vrm0.HumanoidBoneTypeLeftUpperLeg:
			humanoid.HumanBones.LeftUpperLeg = vrm1.HumanoidBone{Node: bone.Node}
		case vrm0.HumanoidBoneTypeNeck:
			humanoid.HumanBones.Neck = &vrm1.HumanoidBone{Node: bone.Node}
		case vrm0.HumanoidBoneTypeRightEye:
			humanoid.HumanBones.RightEye = &vrm1.HumanoidBone{Node: bone.Node}
		case vrm0.HumanoidBoneTypeRightFoot:
			humanoid.HumanBones.RightFoot = vrm1.HumanoidBone{Node: bone.Node}
		case vrm0.HumanoidBoneTypeRightHand:
			humanoid.HumanBones.RightHand = vrm1.HumanoidBone{Node: bone.Node}
		case vrm0.HumanoidBoneTypeRightIndexDistal:
			humanoid.HumanBones.RightIndexDistal = &vrm1.HumanoidBone{Node: bone.Node}
		case vrm0.HumanoidBoneTypeRightIndexIntermediate:
			humanoid.HumanBones.RightIndexIntermediate = &vrm1.HumanoidBone{Node: bone.Node}
		case vrm0.HumanoidBoneTypeRightIndexProximal:
			humanoid.HumanBones.RightIndexProximal = &vrm1.HumanoidBone{Node: bone.Node}
		case vrm0.HumanoidBoneTypeRightLittleDistal:
			humanoid.HumanBones.RightLittleDistal = &vrm1.HumanoidBone{Node: bone.Node}
		case vrm0.HumanoidBoneTypeRightLittleIntermediate:
			humanoid.HumanBones.RightLittleIntermediate = &vrm1.HumanoidBone{Node: bone.Node}
		case vrm0.HumanoidBoneTypeRightLittleProximal:
			humanoid.HumanBones.RightLittleProximal = &vrm1.HumanoidBone{Node: bone.Node}
		case vrm0.HumanoidBoneTypeRightLowerArm:
			humanoid.HumanBones.RightLowerArm = vrm1.HumanoidBone{Node: bone.Node}
		case vrm0.HumanoidBoneTypeRightLowerLeg:
			humanoid.HumanBones.RightLowerLeg = vrm1.HumanoidBone{Node: bone.Node}
		case vrm0.HumanoidBoneTypeRightMiddleDistal:
			humanoid.HumanBones.RightMiddleDistal = &vrm1.HumanoidBone{Node: bone.Node}
		case vrm0.HumanoidBoneTypeRightMiddleIntermediate:
			humanoid.HumanBones.RightMiddleIntermediate = &vrm1.HumanoidBone{Node: bone.Node}
		case vrm0.HumanoidBoneTypeRightMiddleProximal:
			humanoid.HumanBones.RightMiddleProximal = &vrm1.HumanoidBone{Node: bone.Node}
		case vrm0.HumanoidBoneTypeRightRingDistal:
			humanoid.HumanBones.RightRingDistal = &vrm1.HumanoidBone{Node: bone.Node}
		case vrm0.HumanoidBoneTypeRightRingIntermediate:
			//NOTE https://github.com/vrm-c/UniVRM/blob/6ba7499a2f18101ac1445e8468e527aa40a24d67/Assets/VRM10/Runtime/Migration/MigrationVrmHumanoid.cs#L89
			humanoid.HumanBones.RightRingIntermediate = &vrm1.HumanoidBone{Node: bone.Node}
		case vrm0.HumanoidBoneTypeRightRingProximal:
			//NOTE https://github.com/vrm-c/UniVRM/blob/6ba7499a2f18101ac1445e8468e527aa40a24d67/Assets/VRM10/Runtime/Migration/MigrationVrmHumanoid.cs#L88
			humanoid.HumanBones.RightThumbMetacarpal = &vrm1.HumanoidBone{Node: bone.Node}
		case vrm0.HumanoidBoneTypeRightShoulder:
			humanoid.HumanBones.RightShoulder = &vrm1.HumanoidBone{Node: bone.Node}
		case vrm0.HumanoidBoneTypeRightThumbDistal:
			humanoid.HumanBones.RightThumbDistal = &vrm1.HumanoidBone{Node: bone.Node}
		case vrm0.HumanoidBoneTypeRightThumbIntermediate:
			humanoid.HumanBones.RightThumbProximal = &vrm1.HumanoidBone{Node: bone.Node}
		case vrm0.HumanoidBoneTypeRightThumbProximal:
			humanoid.HumanBones.RightThumbProximal = &vrm1.HumanoidBone{Node: bone.Node}
		case vrm0.HumanoidBoneTypeRightToes:
			humanoid.HumanBones.RightToes = &vrm1.HumanoidBone{Node: bone.Node}
		case vrm0.HumanoidBoneTypeRightUpperArm:
			humanoid.HumanBones.RightUpperArm = vrm1.HumanoidBone{Node: bone.Node}
		case vrm0.HumanoidBoneTypeRightUpperLeg:
			humanoid.HumanBones.RightUpperLeg = vrm1.HumanoidBone{Node: bone.Node}
		case vrm0.HumanoidBoneTypeSpine:
			humanoid.HumanBones.Spine = vrm1.HumanoidBone{Node: bone.Node}
		case vrm0.HumanoidBoneTypeUpperChest:
			humanoid.HumanBones.UpperChest = &vrm1.HumanoidBone{Node: bone.Node}
		default:
			return vrm1.Humanoid{}, fmt.Errorf("unknown bone: %s", *bone.Bone)
		}
	}

	return humanoid, nil
}
