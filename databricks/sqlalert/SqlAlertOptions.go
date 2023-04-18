package sqlalert


type SqlAlertOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.14.3/docs/resources/sql_alert#column SqlAlert#column}.
	Column *string `field:"required" json:"column" yaml:"column"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.14.3/docs/resources/sql_alert#op SqlAlert#op}.
	Op *string `field:"required" json:"op" yaml:"op"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.14.3/docs/resources/sql_alert#value SqlAlert#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.14.3/docs/resources/sql_alert#custom_body SqlAlert#custom_body}.
	CustomBody *string `field:"optional" json:"customBody" yaml:"customBody"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.14.3/docs/resources/sql_alert#custom_subject SqlAlert#custom_subject}.
	CustomSubject *string `field:"optional" json:"customSubject" yaml:"customSubject"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.14.3/docs/resources/sql_alert#muted SqlAlert#muted}.
	Muted interface{} `field:"optional" json:"muted" yaml:"muted"`
}

