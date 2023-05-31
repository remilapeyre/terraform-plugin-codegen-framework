package provider_convert

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-codegen-spec/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"

	"github/hashicorp/terraform-provider-code-generator/internal/provider_generate"
)

func convertInt64Attribute(a *provider.Int64Attribute) (provider_generate.GeneratorInt64Attribute, error) {
	if a == nil {
		return provider_generate.GeneratorInt64Attribute{}, fmt.Errorf("*provider.Int64Attribute is nil")
	}

	return provider_generate.GeneratorInt64Attribute{
		Int64Attribute: schema.Int64Attribute{
			Required:            isRequired(a.OptionalRequired),
			Optional:            isOptional(a.OptionalRequired),
			Sensitive:           isSensitive(a.Sensitive),
			Description:         description(a.Description),
			MarkdownDescription: description(a.Description),
			DeprecationMessage:  deprecationMessage(a.DeprecationMessage),
		},
		CustomType: a.CustomType,
		Validators: a.Validators,
	}, nil
}