// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pipeline


type PipelineNotification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.92.0/docs/resources/pipeline#alerts Pipeline#alerts}.
	Alerts *[]*string `field:"optional" json:"alerts" yaml:"alerts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.92.0/docs/resources/pipeline#email_recipients Pipeline#email_recipients}.
	EmailRecipients *[]*string `field:"optional" json:"emailRecipients" yaml:"emailRecipients"`
}

