// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickstable


type DataDatabricksTableTableInfo struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/data-sources/table#access_point DataDatabricksTable#access_point}.
	AccessPoint *string `field:"optional" json:"accessPoint" yaml:"accessPoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/data-sources/table#browse_only DataDatabricksTable#browse_only}.
	BrowseOnly interface{} `field:"optional" json:"browseOnly" yaml:"browseOnly"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/data-sources/table#catalog_name DataDatabricksTable#catalog_name}.
	CatalogName *string `field:"optional" json:"catalogName" yaml:"catalogName"`
	// columns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/data-sources/table#columns DataDatabricksTable#columns}
	Columns interface{} `field:"optional" json:"columns" yaml:"columns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/data-sources/table#comment DataDatabricksTable#comment}.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/data-sources/table#created_at DataDatabricksTable#created_at}.
	CreatedAt *float64 `field:"optional" json:"createdAt" yaml:"createdAt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/data-sources/table#created_by DataDatabricksTable#created_by}.
	CreatedBy *string `field:"optional" json:"createdBy" yaml:"createdBy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/data-sources/table#data_access_configuration_id DataDatabricksTable#data_access_configuration_id}.
	DataAccessConfigurationId *string `field:"optional" json:"dataAccessConfigurationId" yaml:"dataAccessConfigurationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/data-sources/table#data_source_format DataDatabricksTable#data_source_format}.
	DataSourceFormat *string `field:"optional" json:"dataSourceFormat" yaml:"dataSourceFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/data-sources/table#deleted_at DataDatabricksTable#deleted_at}.
	DeletedAt *float64 `field:"optional" json:"deletedAt" yaml:"deletedAt"`
	// delta_runtime_properties_kvpairs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/data-sources/table#delta_runtime_properties_kvpairs DataDatabricksTable#delta_runtime_properties_kvpairs}
	DeltaRuntimePropertiesKvpairs *DataDatabricksTableTableInfoDeltaRuntimePropertiesKvpairs `field:"optional" json:"deltaRuntimePropertiesKvpairs" yaml:"deltaRuntimePropertiesKvpairs"`
	// effective_predictive_optimization_flag block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/data-sources/table#effective_predictive_optimization_flag DataDatabricksTable#effective_predictive_optimization_flag}
	EffectivePredictiveOptimizationFlag *DataDatabricksTableTableInfoEffectivePredictiveOptimizationFlag `field:"optional" json:"effectivePredictiveOptimizationFlag" yaml:"effectivePredictiveOptimizationFlag"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/data-sources/table#enable_predictive_optimization DataDatabricksTable#enable_predictive_optimization}.
	EnablePredictiveOptimization *string `field:"optional" json:"enablePredictiveOptimization" yaml:"enablePredictiveOptimization"`
	// encryption_details block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/data-sources/table#encryption_details DataDatabricksTable#encryption_details}
	EncryptionDetails *DataDatabricksTableTableInfoEncryptionDetails `field:"optional" json:"encryptionDetails" yaml:"encryptionDetails"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/data-sources/table#full_name DataDatabricksTable#full_name}.
	FullName *string `field:"optional" json:"fullName" yaml:"fullName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/data-sources/table#metastore_id DataDatabricksTable#metastore_id}.
	MetastoreId *string `field:"optional" json:"metastoreId" yaml:"metastoreId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/data-sources/table#name DataDatabricksTable#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/data-sources/table#owner DataDatabricksTable#owner}.
	Owner *string `field:"optional" json:"owner" yaml:"owner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/data-sources/table#pipeline_id DataDatabricksTable#pipeline_id}.
	PipelineId *string `field:"optional" json:"pipelineId" yaml:"pipelineId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/data-sources/table#properties DataDatabricksTable#properties}.
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
	// row_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/data-sources/table#row_filter DataDatabricksTable#row_filter}
	RowFilter *DataDatabricksTableTableInfoRowFilter `field:"optional" json:"rowFilter" yaml:"rowFilter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/data-sources/table#schema_name DataDatabricksTable#schema_name}.
	SchemaName *string `field:"optional" json:"schemaName" yaml:"schemaName"`
	// securable_kind_manifest block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/data-sources/table#securable_kind_manifest DataDatabricksTable#securable_kind_manifest}
	SecurableKindManifest *DataDatabricksTableTableInfoSecurableKindManifest `field:"optional" json:"securableKindManifest" yaml:"securableKindManifest"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/data-sources/table#sql_path DataDatabricksTable#sql_path}.
	SqlPath *string `field:"optional" json:"sqlPath" yaml:"sqlPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/data-sources/table#storage_credential_name DataDatabricksTable#storage_credential_name}.
	StorageCredentialName *string `field:"optional" json:"storageCredentialName" yaml:"storageCredentialName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/data-sources/table#storage_location DataDatabricksTable#storage_location}.
	StorageLocation *string `field:"optional" json:"storageLocation" yaml:"storageLocation"`
	// table_constraints block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/data-sources/table#table_constraints DataDatabricksTable#table_constraints}
	TableConstraints interface{} `field:"optional" json:"tableConstraints" yaml:"tableConstraints"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/data-sources/table#table_id DataDatabricksTable#table_id}.
	TableId *string `field:"optional" json:"tableId" yaml:"tableId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/data-sources/table#table_type DataDatabricksTable#table_type}.
	TableType *string `field:"optional" json:"tableType" yaml:"tableType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/data-sources/table#updated_at DataDatabricksTable#updated_at}.
	UpdatedAt *float64 `field:"optional" json:"updatedAt" yaml:"updatedAt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/data-sources/table#updated_by DataDatabricksTable#updated_by}.
	UpdatedBy *string `field:"optional" json:"updatedBy" yaml:"updatedBy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/data-sources/table#view_definition DataDatabricksTable#view_definition}.
	ViewDefinition *string `field:"optional" json:"viewDefinition" yaml:"viewDefinition"`
	// view_dependencies block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/data-sources/table#view_dependencies DataDatabricksTable#view_dependencies}
	ViewDependencies *DataDatabricksTableTableInfoViewDependencies `field:"optional" json:"viewDependencies" yaml:"viewDependencies"`
}

