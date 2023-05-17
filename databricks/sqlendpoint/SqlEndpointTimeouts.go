package sqlendpoint


type SqlEndpointTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.16.1/docs/resources/sql_endpoint#create SqlEndpoint#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
}

