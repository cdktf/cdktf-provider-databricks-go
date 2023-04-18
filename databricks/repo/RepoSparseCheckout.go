package repo


type RepoSparseCheckout struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.14.3/docs/resources/repo#patterns Repo#patterns}.
	Patterns *[]*string `field:"required" json:"patterns" yaml:"patterns"`
}

