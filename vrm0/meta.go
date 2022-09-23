package vrm0

import (
	"github.com/thara/go-vrm-migrate/vrm0/internal"
)

type Meta = internal.VrmMetaSchemaJson

type MetaAllowedUserName = internal.VrmMetaSchemaJsonAllowedUserName

const (
	MetaAllowedUserNameEveryone                 = internal.VrmMetaSchemaJsonAllowedUserNameEveryone
	MetaAllowedUserNameExplicitlyLicensedPerson = internal.VrmMetaSchemaJsonAllowedUserNameExplicitlyLicensedPerson
	MetaAllowedUserNameOnlyAuthor               = internal.VrmMetaSchemaJsonAllowedUserNameOnlyAuthor
)

type MetaCommercialUssageName = internal.VrmMetaSchemaJsonCommercialUssageName

const (
	MetaCommercialUssageNameAllow    = internal.VrmMetaSchemaJsonCommercialUssageNameAllow
	MetaCommercialUssageNameDisallow = internal.VrmMetaSchemaJsonCommercialUssageNameDisallow
)

type MetaLicenseName = internal.VrmMetaSchemaJsonLicenseName

const (
	MetaLicenseNameCC0                      = internal.VrmMetaSchemaJsonLicenseNameCC0
	MetaLicenseNameCCBY                     = internal.VrmMetaSchemaJsonLicenseNameCCBY
	MetaLicenseNameCCBYNC                   = internal.VrmMetaSchemaJsonLicenseNameCCBYNC
	MetaLicenseNameCCBYNCND                 = internal.VrmMetaSchemaJsonLicenseNameCCBYNCND
	MetaLicenseNameCCBYNCSA                 = internal.VrmMetaSchemaJsonLicenseNameCCBYNCSA
	MetaLicenseNameCCBYND                   = internal.VrmMetaSchemaJsonLicenseNameCCBYND
	MetaLicenseNameCCBYSA                   = internal.VrmMetaSchemaJsonLicenseNameCCBYSA
	MetaLicenseNameOther                    = internal.VrmMetaSchemaJsonLicenseNameOther
	MetaLicenseNameRedistributionProhibited = internal.VrmMetaSchemaJsonLicenseNameRedistributionProhibited
)

type MetaSexualUssageName = internal.VrmMetaSchemaJsonSexualUssageName

const (
	MetaSexualUssageNameAllow    = internal.VrmMetaSchemaJsonSexualUssageNameAllow
	MetaSexualUssageNameDisallow = internal.VrmMetaSchemaJsonSexualUssageNameDisallow
)

type MetaViolentUssageName = internal.VrmMetaSchemaJsonViolentUssageName

const (
	MetaViolentUssageNameAllow    = internal.VrmMetaSchemaJsonViolentUssageNameAllow
	MetaViolentUssageNameDisallow = internal.VrmMetaSchemaJsonViolentUssageNameDisallow
)
