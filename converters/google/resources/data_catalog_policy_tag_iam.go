// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import "fmt"

// Provide a separate asset type constant so we don't have to worry about name conflicts between IAM and non-IAM converter files
const DataCatalogPolicyTagIAMAssetType string = "datacatalog.googleapis.com/PolicyTag"

func resourceConverterDataCatalogPolicyTagIamPolicy() ResourceConverter {
	return ResourceConverter{
		AssetType:         DataCatalogPolicyTagIAMAssetType,
		Convert:           GetDataCatalogPolicyTagIamPolicyCaiObject,
		MergeCreateUpdate: MergeDataCatalogPolicyTagIamPolicy,
	}
}

func resourceConverterDataCatalogPolicyTagIamBinding() ResourceConverter {
	return ResourceConverter{
		AssetType:         DataCatalogPolicyTagIAMAssetType,
		Convert:           GetDataCatalogPolicyTagIamBindingCaiObject,
		FetchFullResource: FetchDataCatalogPolicyTagIamPolicy,
		MergeCreateUpdate: MergeDataCatalogPolicyTagIamBinding,
		MergeDelete:       MergeDataCatalogPolicyTagIamBindingDelete,
	}
}

func resourceConverterDataCatalogPolicyTagIamMember() ResourceConverter {
	return ResourceConverter{
		AssetType:         DataCatalogPolicyTagIAMAssetType,
		Convert:           GetDataCatalogPolicyTagIamMemberCaiObject,
		FetchFullResource: FetchDataCatalogPolicyTagIamPolicy,
		MergeCreateUpdate: MergeDataCatalogPolicyTagIamMember,
		MergeDelete:       MergeDataCatalogPolicyTagIamMemberDelete,
	}
}

func GetDataCatalogPolicyTagIamPolicyCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	return newDataCatalogPolicyTagIamAsset(d, config, expandIamPolicyBindings)
}

func GetDataCatalogPolicyTagIamBindingCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	return newDataCatalogPolicyTagIamAsset(d, config, expandIamRoleBindings)
}

func GetDataCatalogPolicyTagIamMemberCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	return newDataCatalogPolicyTagIamAsset(d, config, expandIamMemberBindings)
}

func MergeDataCatalogPolicyTagIamPolicy(existing, incoming Asset) Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeDataCatalogPolicyTagIamBinding(existing, incoming Asset) Asset {
	return mergeIamAssets(existing, incoming, mergeAuthoritativeBindings)
}

func MergeDataCatalogPolicyTagIamBindingDelete(existing, incoming Asset) Asset {
	return mergeDeleteIamAssets(existing, incoming, mergeDeleteAuthoritativeBindings)
}

func MergeDataCatalogPolicyTagIamMember(existing, incoming Asset) Asset {
	return mergeIamAssets(existing, incoming, mergeAdditiveBindings)
}

func MergeDataCatalogPolicyTagIamMemberDelete(existing, incoming Asset) Asset {
	return mergeDeleteIamAssets(existing, incoming, mergeDeleteAdditiveBindings)
}

func newDataCatalogPolicyTagIamAsset(
	d TerraformResourceData,
	config *Config,
	expandBindings func(d TerraformResourceData) ([]IAMBinding, error),
) ([]Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := assetName(d, config, "//datacatalog.googleapis.com/{{policy_tag}}")
	if err != nil {
		return []Asset{}, err
	}

	return []Asset{{
		Name: name,
		Type: DataCatalogPolicyTagIAMAssetType,
		IAMPolicy: &IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchDataCatalogPolicyTagIamPolicy(d TerraformResourceData, config *Config) (Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("policy_tag"); !ok {
		return Asset{}, ErrEmptyIdentityField
	}

	return fetchIamPolicy(
		DataCatalogPolicyTagIamUpdaterProducer,
		d,
		config,
		"//datacatalog.googleapis.com/{{policy_tag}}",
		DataCatalogPolicyTagIAMAssetType,
	)
}
