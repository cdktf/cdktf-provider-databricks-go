// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sqlquery


type SqlQueryParameter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.34.0/docs/resources/sql_query#name SqlQuery#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// date block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.34.0/docs/resources/sql_query#date SqlQuery#date}
	Date *SqlQueryParameterDate `field:"optional" json:"date" yaml:"date"`
	// date_range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.34.0/docs/resources/sql_query#date_range SqlQuery#date_range}
	DateRange *SqlQueryParameterDateRange `field:"optional" json:"dateRange" yaml:"dateRange"`
	// datetime block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.34.0/docs/resources/sql_query#datetime SqlQuery#datetime}
	Datetime *SqlQueryParameterDatetime `field:"optional" json:"datetime" yaml:"datetime"`
	// datetime_range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.34.0/docs/resources/sql_query#datetime_range SqlQuery#datetime_range}
	DatetimeRange *SqlQueryParameterDatetimeRange `field:"optional" json:"datetimeRange" yaml:"datetimeRange"`
	// datetimesec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.34.0/docs/resources/sql_query#datetimesec SqlQuery#datetimesec}
	Datetimesec *SqlQueryParameterDatetimesec `field:"optional" json:"datetimesec" yaml:"datetimesec"`
	// datetimesec_range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.34.0/docs/resources/sql_query#datetimesec_range SqlQuery#datetimesec_range}
	DatetimesecRange *SqlQueryParameterDatetimesecRange `field:"optional" json:"datetimesecRange" yaml:"datetimesecRange"`
	// enum block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.34.0/docs/resources/sql_query#enum SqlQuery#enum}
	Enum *SqlQueryParameterEnum `field:"optional" json:"enum" yaml:"enum"`
	// number block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.34.0/docs/resources/sql_query#number SqlQuery#number}
	Number *SqlQueryParameterNumber `field:"optional" json:"number" yaml:"number"`
	// query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.34.0/docs/resources/sql_query#query SqlQuery#query}
	Query *SqlQueryParameterQuery `field:"optional" json:"query" yaml:"query"`
	// text block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.34.0/docs/resources/sql_query#text SqlQuery#text}
	Text *SqlQueryParameterText `field:"optional" json:"text" yaml:"text"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.34.0/docs/resources/sql_query#title SqlQuery#title}.
	Title *string `field:"optional" json:"title" yaml:"title"`
}

