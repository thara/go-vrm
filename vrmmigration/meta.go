package vrmmigrate

import (
	"fmt"

	"github.com/qmuntal/gltf"
	"github.com/thara/go-vrm-migrate/vrm0"
	"github.com/thara/go-vrm-migrate/vrm1x"
)

func migrateMeta(doc *gltf.Document, meta0 *vrm0.Meta) (vrm1x.Meta, error) {
	meta := vrm1x.Meta{
		LicenseUrl:       "https://vrm.dev/licenses/1.0/",
		AvatarPermission: vrm1x.MetaAvatarPermissionOnlyAuthor,
		CommercialUsage:  vrm1x.MetaCommercialUsagePersonalNonProfit,
		CreditNotation:   vrm1x.MetaCreditNotationRequired,
		Modification:     vrm1x.MetaModificationProhibited,
	}

	if meta0.Title != nil {
		meta.Name = *meta0.Title
	}
	meta.Version = meta0.Version
	if meta0.Author != nil {
		meta.Authors = append(meta.Authors, *meta0.Author)
	}
	meta.ContactInformation = meta0.ContactInformation
	if meta0.Reference != nil {
		meta.References = append(meta.References, *meta0.Reference)
	}
	if meta0.Texture != nil {
		i := *meta0.Texture
		if i == -1 {
			meta.ThumbnailImage = -1
		} else {
			tex := doc.Textures[i]
			meta.ThumbnailImage = tex.Source
		}
	}
	if meta0.AllowedUserName != nil {
		switch *meta0.AllowedUserName {
		case vrm0.MetaAllowedUserNameOnlyAuthor:
			meta.AvatarPermission = vrm1x.MetaAvatarPermissionOnlyAuthor
		case vrm0.MetaAllowedUserNameExplicitlyLicensedPerson:
			meta.AvatarPermission = vrm1x.MetaAvatarPermissionOnlySeparatelyLicensedPerson
		case vrm0.MetaAllowedUserNameEveryone:
			meta.AvatarPermission = vrm1x.MetaAvatarPermissionEveryone
		default:
			return meta, fmt.Errorf("can not migrate allowedUserName(%s) from vrm 0.x", *meta0.AllowedUserName)
		}
	}
	if meta0.ViolentUssageName != nil {
		switch *meta0.ViolentUssageName {
		case vrm0.MetaViolentUssageNameAllow:
			meta.AllowExcessivelyViolentUsage = true
		case vrm0.MetaViolentUssageNameDisallow:
			meta.AllowExcessivelyViolentUsage = false
		default:
			return meta, fmt.Errorf("can not migrate violentUssageName(%s) from vrm 0.x", *meta0.ViolentUssageName)
		}
	}
	if meta0.SexualUssageName != nil {
		switch *meta0.SexualUssageName {
		case vrm0.MetaSexualUssageNameAllow:
			meta.AllowExcessivelySexualUsage = true
		case vrm0.MetaSexualUssageNameDisallow:
			meta.AllowExcessivelySexualUsage = false
		default:
			return meta, fmt.Errorf("can not migrate sexualUssageName(%s) from vrm 0.x", *meta0.SexualUssageName)
		}
	}
	if meta0.CommercialUssageName != nil {
		switch *meta0.CommercialUssageName {
		case vrm0.MetaCommercialUssageNameAllow:
			meta.CommercialUsage = vrm1x.MetaCommercialUsagePersonalProfit
		case vrm0.MetaCommercialUssageNameDisallow:
			meta.CommercialUsage = vrm1x.MetaCommercialUsagePersonalNonProfit
		default:
			return meta, fmt.Errorf("can not migrate commercialUssageName(%s) from vrm 0.x", *meta0.CommercialUssageName)
		}
	}

	otherLicenseUrl := meta0.OtherLicenseUrl
	otherPermissionUrl := meta0.OtherPermissionUrl

	hasLicense := otherLicenseUrl != nil && 1 < len(*otherLicenseUrl)
	hasPermission := otherPermissionUrl != nil && 1 < len(*otherPermissionUrl)

	if hasLicense && hasPermission {
		if *otherLicenseUrl == *otherPermissionUrl {
			meta.OtherLicenseUrl = otherLicenseUrl
		} else {
			url := fmt.Sprintf("%s,%s", *otherLicenseUrl, *otherPermissionUrl)
			meta.OtherLicenseUrl = &url
		}
	} else if hasLicense {
		meta.OtherLicenseUrl = otherLicenseUrl
	} else if hasPermission {
		meta.OtherLicenseUrl = otherPermissionUrl
	}

	return meta, nil
}
