// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickscatalog


type DataDatabricksCatalogCatalogInfo struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.47.0/docs/data-sources/catalog#browse_only DataDatabricksCatalog#browse_only}.
	BrowseOnly interface{} `field:"optional" json:"browseOnly" yaml:"browseOnly"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.47.0/docs/data-sources/catalog#catalog_type DataDatabricksCatalog#catalog_type}.
	CatalogType *string `field:"optional" json:"catalogType" yaml:"catalogType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.47.0/docs/data-sources/catalog#comment DataDatabricksCatalog#comment}.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.47.0/docs/data-sources/catalog#connection_name DataDatabricksCatalog#connection_name}.
	ConnectionName *string `field:"optional" json:"connectionName" yaml:"connectionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.47.0/docs/data-sources/catalog#created_at DataDatabricksCatalog#created_at}.
	CreatedAt *float64 `field:"optional" json:"createdAt" yaml:"createdAt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.47.0/docs/data-sources/catalog#created_by DataDatabricksCatalog#created_by}.
	CreatedBy *string `field:"optional" json:"createdBy" yaml:"createdBy"`
	// effective_predictive_optimization_flag block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.47.0/docs/data-sources/catalog#effective_predictive_optimization_flag DataDatabricksCatalog#effective_predictive_optimization_flag}
	EffectivePredictiveOptimizationFlag *DataDatabricksCatalogCatalogInfoEffectivePredictiveOptimizationFlag `field:"optional" json:"effectivePredictiveOptimizationFlag" yaml:"effectivePredictiveOptimizationFlag"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.47.0/docs/data-sources/catalog#enable_predictive_optimization DataDatabricksCatalog#enable_predictive_optimization}.
	EnablePredictiveOptimization *string `field:"optional" json:"enablePredictiveOptimization" yaml:"enablePredictiveOptimization"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.47.0/docs/data-sources/catalog#full_name DataDatabricksCatalog#full_name}.
	FullName *string `field:"optional" json:"fullName" yaml:"fullName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.47.0/docs/data-sources/catalog#isolation_mode DataDatabricksCatalog#isolation_mode}.
	IsolationMode *string `field:"optional" json:"isolationMode" yaml:"isolationMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.47.0/docs/data-sources/catalog#metastore_id DataDatabricksCatalog#metastore_id}.
	MetastoreId *string `field:"optional" json:"metastoreId" yaml:"metastoreId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.47.0/docs/data-sources/catalog#name DataDatabricksCatalog#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.47.0/docs/data-sources/catalog#options DataDatabricksCatalog#options}.
	Options *map[string]*string `field:"optional" json:"options" yaml:"options"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.47.0/docs/data-sources/catalog#owner DataDatabricksCatalog#owner}.
	Owner *string `field:"optional" json:"owner" yaml:"owner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.47.0/docs/data-sources/catalog#properties DataDatabricksCatalog#properties}.
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.47.0/docs/data-sources/catalog#provider_name DataDatabricksCatalog#provider_name}.
	ProviderName *string `field:"optional" json:"providerName" yaml:"providerName"`
	// provisioning_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.47.0/docs/data-sources/catalog#provisioning_info DataDatabricksCatalog#provisioning_info}
	ProvisioningInfo *DataDatabricksCatalogCatalogInfoProvisioningInfo `field:"optional" json:"provisioningInfo" yaml:"provisioningInfo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.47.0/docs/data-sources/catalog#securable_kind DataDatabricksCatalog#securable_kind}.
	SecurableKind *string `field:"optional" json:"securableKind" yaml:"securableKind"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.47.0/docs/data-sources/catalog#securable_type DataDatabricksCatalog#securable_type}.
	SecurableType *string `field:"optional" json:"securableType" yaml:"securableType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.47.0/docs/data-sources/catalog#share_name DataDatabricksCatalog#share_name}.
	ShareName *string `field:"optional" json:"shareName" yaml:"shareName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.47.0/docs/data-sources/catalog#storage_location DataDatabricksCatalog#storage_location}.
	StorageLocation *string `field:"optional" json:"storageLocation" yaml:"storageLocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.47.0/docs/data-sources/catalog#storage_root DataDatabricksCatalog#storage_root}.
	StorageRoot *string `field:"optional" json:"storageRoot" yaml:"storageRoot"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.47.0/docs/data-sources/catalog#updated_at DataDatabricksCatalog#updated_at}.
	UpdatedAt *float64 `field:"optional" json:"updatedAt" yaml:"updatedAt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.47.0/docs/data-sources/catalog#updated_by DataDatabricksCatalog#updated_by}.
	UpdatedBy *string `field:"optional" json:"updatedBy" yaml:"updatedBy"`
}

