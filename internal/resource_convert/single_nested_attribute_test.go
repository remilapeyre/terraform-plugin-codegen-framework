// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource_convert

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-codegen-spec/code"
	"github.com/hashicorp/terraform-plugin-codegen-spec/resource"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/resource_generate"
	generatorschema "github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

func TestConvertSingleNestedAttribute(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         *resource.SingleNestedAttribute
		expected      resource_generate.GeneratorSingleNestedAttribute
		expectedError error
	}{
		"nil": {
			expectedError: fmt.Errorf("*resource.SingleNestedAttribute is nil"),
		},
		"attributes-nil": {
			input: &resource.SingleNestedAttribute{
				Attributes: []resource.Attribute{
					{
						Name: "empty",
					},
				},
			},
			expectedError: fmt.Errorf("attribute type not defined: %+v", resource.Attribute{
				Name: "empty",
			}),
		},
		"attributes-bool": {
			input: &resource.SingleNestedAttribute{
				Attributes: []resource.Attribute{
					{
						Name: "bool_attribute",
						Bool: &resource.BoolAttribute{
							ComputedOptionalRequired: "optional",
						},
					},
				},
			},
			expected: resource_generate.GeneratorSingleNestedAttribute{
				Attributes: generatorschema.GeneratorAttributes{
					"bool_attribute": resource_generate.GeneratorBoolAttribute{
						BoolAttribute: schema.BoolAttribute{
							Optional: true,
						},
					},
				},
			},
		},
		"attributes-list-bool": {
			input: &resource.SingleNestedAttribute{
				Attributes: []resource.Attribute{
					{
						Name: "list_attribute",
						List: &resource.ListAttribute{
							ComputedOptionalRequired: "optional",
							ElementType: specschema.ElementType{
								Bool: &specschema.BoolType{},
							},
						},
					},
				},
			},
			expected: resource_generate.GeneratorSingleNestedAttribute{
				Attributes: generatorschema.GeneratorAttributes{
					"list_attribute": resource_generate.GeneratorListAttribute{
						ListAttribute: schema.ListAttribute{
							Optional: true,
						},
						ElementType: specschema.ElementType{
							Bool: &specschema.BoolType{},
						},
					},
				},
			},
		},
		"attributes-list-nested-bool": {
			input: &resource.SingleNestedAttribute{
				Attributes: []resource.Attribute{
					{
						Name: "nested_attribute",
						ListNested: &resource.ListNestedAttribute{
							NestedObject: resource.NestedAttributeObject{
								Attributes: []resource.Attribute{
									{
										Name: "nested_bool",
										Bool: &resource.BoolAttribute{
											ComputedOptionalRequired: "computed",
										},
									},
								},
							},
							ComputedOptionalRequired: "optional",
						},
					},
				},
			},
			expected: resource_generate.GeneratorSingleNestedAttribute{
				Attributes: generatorschema.GeneratorAttributes{
					"nested_attribute": resource_generate.GeneratorListNestedAttribute{
						NestedObject: resource_generate.GeneratorNestedAttributeObject{
							Attributes: generatorschema.GeneratorAttributes{
								"nested_bool": resource_generate.GeneratorBoolAttribute{
									BoolAttribute: schema.BoolAttribute{
										Computed: true,
									},
								},
							},
						},
						ListNestedAttribute: schema.ListNestedAttribute{
							Optional: true,
						},
					},
				},
			},
		},
		"attributes-object-bool": {
			input: &resource.SingleNestedAttribute{
				Attributes: []resource.Attribute{
					{
						Name: "object_attribute",
						Object: &resource.ObjectAttribute{
							AttributeTypes: specschema.ObjectAttributeTypes{
								{
									Name: "obj_bool",
									Bool: &specschema.BoolType{},
								},
							},
							ComputedOptionalRequired: "optional",
						},
					},
				},
			},
			expected: resource_generate.GeneratorSingleNestedAttribute{
				Attributes: generatorschema.GeneratorAttributes{
					"object_attribute": resource_generate.GeneratorObjectAttribute{
						ObjectAttribute: schema.ObjectAttribute{
							Optional: true,
						},
						AttributeTypes: specschema.ObjectAttributeTypes{
							{
								Name: "obj_bool",
								Bool: &specschema.BoolType{},
							},
						},
					},
				},
			},
		},
		"attributes-single-nested-bool": {
			input: &resource.SingleNestedAttribute{
				Attributes: []resource.Attribute{
					{
						Name: "nested_attribute",
						SingleNested: &resource.SingleNestedAttribute{
							Attributes: []resource.Attribute{
								{
									Name: "nested_bool",
									Bool: &resource.BoolAttribute{
										ComputedOptionalRequired: "computed",
									},
								},
							},
							ComputedOptionalRequired: "optional",
						},
					},
				},
			},
			expected: resource_generate.GeneratorSingleNestedAttribute{
				Attributes: generatorschema.GeneratorAttributes{
					"nested_attribute": resource_generate.GeneratorSingleNestedAttribute{
						Attributes: generatorschema.GeneratorAttributes{
							"nested_bool": resource_generate.GeneratorBoolAttribute{
								BoolAttribute: schema.BoolAttribute{
									Computed: true,
								},
							},
						},
						SingleNestedAttribute: schema.SingleNestedAttribute{
							Optional: true,
						},
					},
				},
			},
		},
		"computed": {
			input: &resource.SingleNestedAttribute{
				ComputedOptionalRequired: "computed",
			},
			expected: resource_generate.GeneratorSingleNestedAttribute{
				SingleNestedAttribute: schema.SingleNestedAttribute{
					Computed: true,
				},
			},
		},
		"computed_optional": {
			input: &resource.SingleNestedAttribute{
				ComputedOptionalRequired: "computed_optional",
			},
			expected: resource_generate.GeneratorSingleNestedAttribute{
				SingleNestedAttribute: schema.SingleNestedAttribute{
					Computed: true,
					Optional: true,
				},
			},
		},
		"optional": {
			input: &resource.SingleNestedAttribute{
				ComputedOptionalRequired: "optional",
			},
			expected: resource_generate.GeneratorSingleNestedAttribute{
				SingleNestedAttribute: schema.SingleNestedAttribute{
					Optional: true,
				},
			},
		},
		"required": {
			input: &resource.SingleNestedAttribute{
				ComputedOptionalRequired: "required",
			},
			expected: resource_generate.GeneratorSingleNestedAttribute{
				SingleNestedAttribute: schema.SingleNestedAttribute{
					Required: true,
				},
			},
		},
		"custom_type": {
			input: &resource.SingleNestedAttribute{
				CustomType: &specschema.CustomType{
					Import: &code.Import{
						Path: "github.com/",
					},
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
			},
			expected: resource_generate.GeneratorSingleNestedAttribute{
				CustomType: &specschema.CustomType{
					Import: &code.Import{
						Path: "github.com/",
					},
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
			},
		},
		"deprecation_message": {
			input: &resource.SingleNestedAttribute{
				DeprecationMessage: pointer("deprecation message"),
			},
			expected: resource_generate.GeneratorSingleNestedAttribute{
				SingleNestedAttribute: schema.SingleNestedAttribute{
					DeprecationMessage: "deprecation message",
				},
			},
		},
		"description": {
			input: &resource.SingleNestedAttribute{
				Description: pointer("description"),
			},
			expected: resource_generate.GeneratorSingleNestedAttribute{
				SingleNestedAttribute: schema.SingleNestedAttribute{
					Description:         "description",
					MarkdownDescription: "description",
				},
			},
		},
		"sensitive": {
			input: &resource.SingleNestedAttribute{
				Sensitive: pointer(true),
			},
			expected: resource_generate.GeneratorSingleNestedAttribute{
				SingleNestedAttribute: schema.SingleNestedAttribute{
					Sensitive: true,
				},
			},
		},
		"validators": {
			input: &resource.SingleNestedAttribute{
				Validators: specschema.ObjectValidators{
					{
						Custom: &specschema.CustomValidator{
							Imports: []code.Import{
								{
									Path: "github.com/.../myvalidator",
								},
							},
							SchemaDefinition: "myvalidator.Validate()",
						},
					},
				},
			},
			expected: resource_generate.GeneratorSingleNestedAttribute{
				Validators: specschema.ObjectValidators{
					{
						Custom: &specschema.CustomValidator{
							Imports: []code.Import{
								{
									Path: "github.com/.../myvalidator",
								},
							},
							SchemaDefinition: "myvalidator.Validate()",
						},
					},
				},
			},
		},
		"plan-modifiers": {
			input: &resource.SingleNestedAttribute{
				PlanModifiers: specschema.ObjectPlanModifiers{
					{
						Custom: &specschema.CustomPlanModifier{
							Imports: []code.Import{
								{
									Path: "github.com/.../my_planmodifier",
								},
							},
							SchemaDefinition: "my_planmodifier.Modify()",
						},
					},
				},
			},
			expected: resource_generate.GeneratorSingleNestedAttribute{
				PlanModifiers: specschema.ObjectPlanModifiers{
					{
						Custom: &specschema.CustomPlanModifier{
							Imports: []code.Import{
								{
									Path: "github.com/.../my_planmodifier",
								},
							},
							SchemaDefinition: "my_planmodifier.Modify()",
						},
					},
				},
			},
		},
		"default": {
			input: &resource.SingleNestedAttribute{
				Default: &specschema.ObjectDefault{
					Custom: &specschema.CustomDefault{
						Imports: []code.Import{
							{
								Path: "github.com/.../my_default",
							},
						},
						SchemaDefinition: "my_default.Default()",
					},
				},
			},
			expected: resource_generate.GeneratorSingleNestedAttribute{
				Default: &specschema.ObjectDefault{
					Custom: &specschema.CustomDefault{
						Imports: []code.Import{
							{
								Path: "github.com/.../my_default",
							},
						},
						SchemaDefinition: "my_default.Default()",
					},
				},
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := convertSingleNestedAttribute(testCase.input)

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
