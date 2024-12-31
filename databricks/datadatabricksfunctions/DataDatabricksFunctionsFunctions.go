// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksfunctions


type DataDatabricksFunctionsFunctions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.62.1/docs/data-sources/functions#browse_only DataDatabricksFunctions#browse_only}.
	BrowseOnly interface{} `field:"optional" json:"browseOnly" yaml:"browseOnly"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.62.1/docs/data-sources/functions#catalog_name DataDatabricksFunctions#catalog_name}.
	CatalogName *string `field:"optional" json:"catalogName" yaml:"catalogName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.62.1/docs/data-sources/functions#comment DataDatabricksFunctions#comment}.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.62.1/docs/data-sources/functions#created_at DataDatabricksFunctions#created_at}.
	CreatedAt *float64 `field:"optional" json:"createdAt" yaml:"createdAt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.62.1/docs/data-sources/functions#created_by DataDatabricksFunctions#created_by}.
	CreatedBy *string `field:"optional" json:"createdBy" yaml:"createdBy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.62.1/docs/data-sources/functions#data_type DataDatabricksFunctions#data_type}.
	DataType *string `field:"optional" json:"dataType" yaml:"dataType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.62.1/docs/data-sources/functions#external_language DataDatabricksFunctions#external_language}.
	ExternalLanguage *string `field:"optional" json:"externalLanguage" yaml:"externalLanguage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.62.1/docs/data-sources/functions#external_name DataDatabricksFunctions#external_name}.
	ExternalName *string `field:"optional" json:"externalName" yaml:"externalName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.62.1/docs/data-sources/functions#full_data_type DataDatabricksFunctions#full_data_type}.
	FullDataType *string `field:"optional" json:"fullDataType" yaml:"fullDataType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.62.1/docs/data-sources/functions#full_name DataDatabricksFunctions#full_name}.
	FullName *string `field:"optional" json:"fullName" yaml:"fullName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.62.1/docs/data-sources/functions#function_id DataDatabricksFunctions#function_id}.
	FunctionId *string `field:"optional" json:"functionId" yaml:"functionId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.62.1/docs/data-sources/functions#input_params DataDatabricksFunctions#input_params}.
	InputParams interface{} `field:"optional" json:"inputParams" yaml:"inputParams"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.62.1/docs/data-sources/functions#is_deterministic DataDatabricksFunctions#is_deterministic}.
	IsDeterministic interface{} `field:"optional" json:"isDeterministic" yaml:"isDeterministic"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.62.1/docs/data-sources/functions#is_null_call DataDatabricksFunctions#is_null_call}.
	IsNullCall interface{} `field:"optional" json:"isNullCall" yaml:"isNullCall"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.62.1/docs/data-sources/functions#metastore_id DataDatabricksFunctions#metastore_id}.
	MetastoreId *string `field:"optional" json:"metastoreId" yaml:"metastoreId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.62.1/docs/data-sources/functions#name DataDatabricksFunctions#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.62.1/docs/data-sources/functions#owner DataDatabricksFunctions#owner}.
	Owner *string `field:"optional" json:"owner" yaml:"owner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.62.1/docs/data-sources/functions#parameter_style DataDatabricksFunctions#parameter_style}.
	ParameterStyle *string `field:"optional" json:"parameterStyle" yaml:"parameterStyle"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.62.1/docs/data-sources/functions#properties DataDatabricksFunctions#properties}.
	Properties *string `field:"optional" json:"properties" yaml:"properties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.62.1/docs/data-sources/functions#return_params DataDatabricksFunctions#return_params}.
	ReturnParams interface{} `field:"optional" json:"returnParams" yaml:"returnParams"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.62.1/docs/data-sources/functions#routine_body DataDatabricksFunctions#routine_body}.
	RoutineBody *string `field:"optional" json:"routineBody" yaml:"routineBody"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.62.1/docs/data-sources/functions#routine_definition DataDatabricksFunctions#routine_definition}.
	RoutineDefinition *string `field:"optional" json:"routineDefinition" yaml:"routineDefinition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.62.1/docs/data-sources/functions#routine_dependencies DataDatabricksFunctions#routine_dependencies}.
	RoutineDependencies interface{} `field:"optional" json:"routineDependencies" yaml:"routineDependencies"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.62.1/docs/data-sources/functions#schema_name DataDatabricksFunctions#schema_name}.
	SchemaName *string `field:"optional" json:"schemaName" yaml:"schemaName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.62.1/docs/data-sources/functions#security_type DataDatabricksFunctions#security_type}.
	SecurityType *string `field:"optional" json:"securityType" yaml:"securityType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.62.1/docs/data-sources/functions#specific_name DataDatabricksFunctions#specific_name}.
	SpecificName *string `field:"optional" json:"specificName" yaml:"specificName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.62.1/docs/data-sources/functions#sql_data_access DataDatabricksFunctions#sql_data_access}.
	SqlDataAccess *string `field:"optional" json:"sqlDataAccess" yaml:"sqlDataAccess"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.62.1/docs/data-sources/functions#sql_path DataDatabricksFunctions#sql_path}.
	SqlPath *string `field:"optional" json:"sqlPath" yaml:"sqlPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.62.1/docs/data-sources/functions#updated_at DataDatabricksFunctions#updated_at}.
	UpdatedAt *float64 `field:"optional" json:"updatedAt" yaml:"updatedAt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.62.1/docs/data-sources/functions#updated_by DataDatabricksFunctions#updated_by}.
	UpdatedBy *string `field:"optional" json:"updatedBy" yaml:"updatedBy"`
}

