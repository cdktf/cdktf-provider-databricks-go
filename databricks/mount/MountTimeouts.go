package mount


type MountTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.23.0/docs/resources/mount#default Mount#default}.
	Default *string `field:"optional" json:"default" yaml:"default"`
}

