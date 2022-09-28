package vrm0

import "github.com/thara/go-vrm-migrate/vrm0/internal"

type Humanoid = internal.VrmHumanoidSchemaJson
type HumanoidBone = internal.VrmHumanoidBoneSchemaJson

type HumanoidBoneType = internal.VrmHumanoidBoneSchemaJsonBone

const (
	HumanoidBoneTypeChest                   = internal.VrmHumanoidBoneSchemaJsonBoneChest
	HumanoidBoneTypeHead                    = internal.VrmHumanoidBoneSchemaJsonBoneHead
	HumanoidBoneTypeHips                    = internal.VrmHumanoidBoneSchemaJsonBoneHips
	HumanoidBoneTypeJaw                     = internal.VrmHumanoidBoneSchemaJsonBoneJaw
	HumanoidBoneTypeLeftEye                 = internal.VrmHumanoidBoneSchemaJsonBoneLeftEye
	HumanoidBoneTypeLeftFoot                = internal.VrmHumanoidBoneSchemaJsonBoneLeftFoot
	HumanoidBoneTypeLeftHand                = internal.VrmHumanoidBoneSchemaJsonBoneLeftHand
	HumanoidBoneTypeLeftIndexDistal         = internal.VrmHumanoidBoneSchemaJsonBoneLeftIndexDistal
	HumanoidBoneTypeLeftIndexIntermediate   = internal.VrmHumanoidBoneSchemaJsonBoneLeftIndexIntermediate
	HumanoidBoneTypeLeftIndexProximal       = internal.VrmHumanoidBoneSchemaJsonBoneLeftIndexProximal
	HumanoidBoneTypeLeftLittleDistal        = internal.VrmHumanoidBoneSchemaJsonBoneLeftLittleDistal
	HumanoidBoneTypeLeftLittleIntermediate  = internal.VrmHumanoidBoneSchemaJsonBoneLeftLittleIntermediate
	HumanoidBoneTypeLeftLittleProximal      = internal.VrmHumanoidBoneSchemaJsonBoneLeftLittleProximal
	HumanoidBoneTypeLeftLowerArm            = internal.VrmHumanoidBoneSchemaJsonBoneLeftLowerArm
	HumanoidBoneTypeLeftLowerLeg            = internal.VrmHumanoidBoneSchemaJsonBoneLeftLowerLeg
	HumanoidBoneTypeLeftMiddleDistal        = internal.VrmHumanoidBoneSchemaJsonBoneLeftMiddleDistal
	HumanoidBoneTypeLeftMiddleIntermediate  = internal.VrmHumanoidBoneSchemaJsonBoneLeftMiddleIntermediate
	HumanoidBoneTypeLeftMiddleProximal      = internal.VrmHumanoidBoneSchemaJsonBoneLeftMiddleProximal
	HumanoidBoneTypeLeftRingDistal          = internal.VrmHumanoidBoneSchemaJsonBoneLeftRingDistal
	HumanoidBoneTypeLeftRingIntermediate    = internal.VrmHumanoidBoneSchemaJsonBoneLeftRingIntermediate
	HumanoidBoneTypeLeftRingProximal        = internal.VrmHumanoidBoneSchemaJsonBoneLeftRingProximal
	HumanoidBoneTypeLeftShoulder            = internal.VrmHumanoidBoneSchemaJsonBoneLeftShoulder
	HumanoidBoneTypeLeftThumbDistal         = internal.VrmHumanoidBoneSchemaJsonBoneLeftThumbDistal
	HumanoidBoneTypeLeftThumbIntermediate   = internal.VrmHumanoidBoneSchemaJsonBoneLeftThumbIntermediate
	HumanoidBoneTypeLeftThumbProximal       = internal.VrmHumanoidBoneSchemaJsonBoneLeftThumbProximal
	HumanoidBoneTypeLeftToes                = internal.VrmHumanoidBoneSchemaJsonBoneLeftToes
	HumanoidBoneTypeLeftUpperArm            = internal.VrmHumanoidBoneSchemaJsonBoneLeftUpperArm
	HumanoidBoneTypeLeftUpperLeg            = internal.VrmHumanoidBoneSchemaJsonBoneLeftUpperLeg
	HumanoidBoneTypeNeck                    = internal.VrmHumanoidBoneSchemaJsonBoneNeck
	HumanoidBoneTypeRightEye                = internal.VrmHumanoidBoneSchemaJsonBoneRightEye
	HumanoidBoneTypeRightFoot               = internal.VrmHumanoidBoneSchemaJsonBoneRightFoot
	HumanoidBoneTypeRightHand               = internal.VrmHumanoidBoneSchemaJsonBoneRightHand
	HumanoidBoneTypeRightIndexDistal        = internal.VrmHumanoidBoneSchemaJsonBoneRightIndexDistal
	HumanoidBoneTypeRightIndexIntermediate  = internal.VrmHumanoidBoneSchemaJsonBoneRightIndexIntermediate
	HumanoidBoneTypeRightIndexProximal      = internal.VrmHumanoidBoneSchemaJsonBoneRightIndexProximal
	HumanoidBoneTypeRightLittleDistal       = internal.VrmHumanoidBoneSchemaJsonBoneRightLittleDistal
	HumanoidBoneTypeRightLittleIntermediate = internal.VrmHumanoidBoneSchemaJsonBoneRightLittleIntermediate
	HumanoidBoneTypeRightLittleProximal     = internal.VrmHumanoidBoneSchemaJsonBoneRightLittleProximal
	HumanoidBoneTypeRightLowerArm           = internal.VrmHumanoidBoneSchemaJsonBoneRightLowerArm
	HumanoidBoneTypeRightLowerLeg           = internal.VrmHumanoidBoneSchemaJsonBoneRightLowerLeg
	HumanoidBoneTypeRightMiddleDistal       = internal.VrmHumanoidBoneSchemaJsonBoneRightMiddleDistal
	HumanoidBoneTypeRightMiddleIntermediate = internal.VrmHumanoidBoneSchemaJsonBoneRightMiddleIntermediate
	HumanoidBoneTypeRightMiddleProximal     = internal.VrmHumanoidBoneSchemaJsonBoneRightMiddleProximal
	HumanoidBoneTypeRightRingDistal         = internal.VrmHumanoidBoneSchemaJsonBoneRightRingDistal
	HumanoidBoneTypeRightRingIntermediate   = internal.VrmHumanoidBoneSchemaJsonBoneRightRingIntermediate
	HumanoidBoneTypeRightRingProximal       = internal.VrmHumanoidBoneSchemaJsonBoneRightRingProximal
	HumanoidBoneTypeRightShoulder           = internal.VrmHumanoidBoneSchemaJsonBoneRightShoulder
	HumanoidBoneTypeRightThumbDistal        = internal.VrmHumanoidBoneSchemaJsonBoneRightThumbDistal
	HumanoidBoneTypeRightThumbIntermediate  = internal.VrmHumanoidBoneSchemaJsonBoneRightThumbIntermediate
	HumanoidBoneTypeRightThumbProximal      = internal.VrmHumanoidBoneSchemaJsonBoneRightThumbProximal
	HumanoidBoneTypeRightToes               = internal.VrmHumanoidBoneSchemaJsonBoneRightToes
	HumanoidBoneTypeRightUpperArm           = internal.VrmHumanoidBoneSchemaJsonBoneRightUpperArm
	HumanoidBoneTypeRightUpperLeg           = internal.VrmHumanoidBoneSchemaJsonBoneRightUpperLeg
	HumanoidBoneTypeSpine                   = internal.VrmHumanoidBoneSchemaJsonBoneSpine
	HumanoidBoneTypeUpperChest              = internal.VrmHumanoidBoneSchemaJsonBoneUpperChest
)
