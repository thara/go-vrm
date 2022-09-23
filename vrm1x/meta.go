package vrm1x

import "github.com/thara/go-vrm-migrate/vrm1x/internal"

type Meta = internal.VRMCVrmMetaSchemaJson

type MetaAvatarPermission = internal.VRMCVrmMetaSchemaJsonAvatarPermission

const (
	MetaAvatarPermissionOnlyAuthor                   = internal.VRMCVrmMetaSchemaJsonAvatarPermissionOnlyAuthor
	MetaAvatarPermissionOnlySeparatelyLicensedPerson = internal.VRMCVrmMetaSchemaJsonAvatarPermissionOnlySeparatelyLicensedPerson
	MetaAvatarPermissionEveryone                     = internal.VRMCVrmMetaSchemaJsonAvatarPermissionEveryone
)

type MetaCommercialUsage = internal.VRMCVrmMetaSchemaJsonCommercialUsage

const (
	MetaCommercialUsagePersonalNonProfit = internal.VRMCVrmMetaSchemaJsonCommercialUsagePersonalNonProfit
	MetaCommercialUsagePersonalProfit    = internal.VRMCVrmMetaSchemaJsonCommercialUsagePersonalProfit
	MetaCommercialUsageCorporation       = internal.VRMCVrmMetaSchemaJsonCommercialUsageCorporation
)

type MetaCreditNotation = internal.VRMCVrmMetaSchemaJsonCreditNotation

const (
	MetaCreditNotationRequired    = internal.VRMCVrmMetaSchemaJsonCreditNotationRequired
	MetaCreditNotationUnnecessary = internal.VRMCVrmMetaSchemaJsonCreditNotationUnnecessary
)

type MetaModification = internal.VRMCVrmMetaSchemaJsonModification

const (
	MetaModificationAllowModification               = internal.VRMCVrmMetaSchemaJsonModificationAllowModification
	MetaModificationAllowModificationRedistribution = internal.VRMCVrmMetaSchemaJsonModificationAllowModificationRedistribution
	MetaModificationProhibited                      = internal.VRMCVrmMetaSchemaJsonModificationProhibited
)
