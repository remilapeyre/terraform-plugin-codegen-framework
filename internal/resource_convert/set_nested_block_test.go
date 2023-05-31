package resource_convert

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-codegen-spec/resource"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github/hashicorp/terraform-provider-code-generator/internal/resource_generate"
)

func TestConvertSetNestedBlock(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         *resource.SetNestedBlock
		expected      resource_generate.GeneratorSetNestedBlock
		expectedError error
	}{
		"nil": {
			expectedError: fmt.Errorf("*resource.SetNestedBlock is nil"),
		},
		"attributes-nil": {
			input: &resource.SetNestedBlock{
				NestedObject: resource.NestedBlockObject{
					Attributes: []resource.Attribute{
						{
							Name: "empty",
						},
					},
				},
			},
			expectedError: fmt.Errorf("attribute type is not defined: %+v", resource.Attribute{
				Name: "empty",
			}),
		},
		"attributes-bool": {
			input: &resource.SetNestedBlock{
				NestedObject: resource.NestedBlockObject{
					Attributes: []resource.Attribute{
						{
							Name: "bool_attribute",
							Bool: &resource.BoolAttribute{
								ComputedOptionalRequired: "optional",
							},
						},
					},
				},
			},
			expected: resource_generate.GeneratorSetNestedBlock{
				NestedObject: resource_generate.GeneratorNestedBlockObject{
					Attributes: map[string]resource_generate.GeneratorAttribute{
						"bool_attribute": resource_generate.GeneratorBoolAttribute{
							BoolAttribute: schema.BoolAttribute{
								Optional: true,
							},
						},
					},
				},
			},
		},
		"attributes-list-bool": {
			input: &resource.SetNestedBlock{
				NestedObject: resource.NestedBlockObject{
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
			},
			expected: resource_generate.GeneratorSetNestedBlock{
				NestedObject: resource_generate.GeneratorNestedBlockObject{
					Attributes: map[string]resource_generate.GeneratorAttribute{
						"list_attribute": resource_generate.GeneratorListAttribute{
							ListAttribute: schema.ListAttribute{
								ElementType: types.BoolType,
								Optional:    true,
							},
						},
					},
				},
			},
		},
		"attributes-list-nested-bool": {
			input: &resource.SetNestedBlock{
				NestedObject: resource.NestedBlockObject{
					Attributes: []resource.Attribute{
						{
							Name: "nested_attribute",
							SetNested: &resource.SetNestedAttribute{
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
			},
			expected: resource_generate.GeneratorSetNestedBlock{
				NestedObject: resource_generate.GeneratorNestedBlockObject{
					Attributes: map[string]resource_generate.GeneratorAttribute{
						"nested_attribute": resource_generate.GeneratorSetNestedAttribute{
							NestedObject: resource_generate.GeneratorNestedAttributeObject{
								Attributes: map[string]resource_generate.GeneratorAttribute{
									"nested_bool": resource_generate.GeneratorBoolAttribute{
										BoolAttribute: schema.BoolAttribute{
											Computed: true,
										},
									},
								},
							},
							SetNestedAttribute: schema.SetNestedAttribute{
								Optional: true,
							},
						},
					},
				},
			},
		},
		"attributes-object-bool": {
			input: &resource.SetNestedBlock{
				NestedObject: resource.NestedBlockObject{
					Attributes: []resource.Attribute{
						{
							Name: "object_attribute",
							Object: &resource.ObjectAttribute{
								AttributeTypes: []specschema.ObjectAttributeType{
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
			},
			expected: resource_generate.GeneratorSetNestedBlock{
				NestedObject: resource_generate.GeneratorNestedBlockObject{
					Attributes: map[string]resource_generate.GeneratorAttribute{
						"object_attribute": resource_generate.GeneratorObjectAttribute{
							ObjectAttribute: schema.ObjectAttribute{
								AttributeTypes: map[string]attr.Type{
									"obj_bool": types.BoolType,
								},
								Optional: true,
							},
						},
					},
				},
			},
		},
		"attributes-single-nested-bool": {
			input: &resource.SetNestedBlock{
				NestedObject: resource.NestedBlockObject{
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
			},
			expected: resource_generate.GeneratorSetNestedBlock{
				NestedObject: resource_generate.GeneratorNestedBlockObject{
					Attributes: map[string]resource_generate.GeneratorAttribute{
						"nested_attribute": resource_generate.GeneratorSingleNestedAttribute{
							Attributes: map[string]resource_generate.GeneratorAttribute{
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
		},

		"blocks-nil": {
			input: &resource.SetNestedBlock{
				NestedObject: resource.NestedBlockObject{
					Blocks: []resource.Block{
						{
							Name: "empty",
						},
					},
				},
			},
			expectedError: fmt.Errorf("block type is not defined: %+v", resource.Block{
				Name: "empty",
			}),
		},

		"blocks-list-nested-bool": {
			input: &resource.SetNestedBlock{
				NestedObject: resource.NestedBlockObject{
					Blocks: []resource.Block{
						{
							Name: "nested_block",
							SetNested: &resource.SetNestedBlock{
								NestedObject: resource.NestedBlockObject{
									Attributes: []resource.Attribute{
										{
											Name: "nested_bool",
											Bool: &resource.BoolAttribute{
												ComputedOptionalRequired: "computed",
											},
										},
									},
								},
							},
						},
					},
				},
			},
			expected: resource_generate.GeneratorSetNestedBlock{
				NestedObject: resource_generate.GeneratorNestedBlockObject{
					Blocks: map[string]resource_generate.GeneratorBlock{
						"nested_block": resource_generate.GeneratorSetNestedBlock{
							NestedObject: resource_generate.GeneratorNestedBlockObject{
								Attributes: map[string]resource_generate.GeneratorAttribute{
									"bool_attribute": resource_generate.GeneratorBoolAttribute{
										BoolAttribute: schema.BoolAttribute{
											Optional: true,
										},
									},
								},
							},
						},
					},
				},
			},
		},

		"blocks-single-nested-bool": {
			input: &resource.SetNestedBlock{
				NestedObject: resource.NestedBlockObject{
					Blocks: []resource.Block{
						{
							Name: "nested_block",
							SingleNested: &resource.SingleNestedBlock{
								Attributes: []resource.Attribute{
									{
										Name: "nested_bool",
										Bool: &resource.BoolAttribute{
											ComputedOptionalRequired: "computed",
										},
									},
								},
							},
						},
					},
				},
			},
			expected: resource_generate.GeneratorSetNestedBlock{
				NestedObject: resource_generate.GeneratorNestedBlockObject{
					Blocks: map[string]resource_generate.GeneratorBlock{
						"nested_block": resource_generate.GeneratorSingleNestedBlock{
							Attributes: map[string]resource_generate.GeneratorAttribute{
								"bool_attribute": resource_generate.GeneratorBoolAttribute{
									BoolAttribute: schema.BoolAttribute{
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
		},

		"custom_type": {
			input: &resource.SetNestedBlock{
				CustomType: &specschema.CustomType{
					Import:    pointer("github.com/"),
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
			},
			expected: resource_generate.GeneratorSetNestedBlock{
				CustomType: &specschema.CustomType{
					Import:    pointer("github.com/"),
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
			},
		},
		"deprecation_message": {
			input: &resource.SetNestedBlock{
				DeprecationMessage: pointer("deprecation message"),
			},
			expected: resource_generate.GeneratorSetNestedBlock{
				SetNestedBlock: schema.SetNestedBlock{
					DeprecationMessage: "deprecation message",
				},
			},
		},
		"description": {
			input: &resource.SetNestedBlock{
				Description: pointer("description"),
			},
			expected: resource_generate.GeneratorSetNestedBlock{
				SetNestedBlock: schema.SetNestedBlock{
					Description:         "description",
					MarkdownDescription: "description",
				},
			},
		},
		"validators": {
			input: &resource.SetNestedBlock{
				Validators: []specschema.SetValidator{
					{
						Custom: &specschema.CustomValidator{
							Import:           pointer("github.com/.../myvalidator"),
							SchemaDefinition: "myvalidator.Validate()",
						},
					},
				},
			},
			expected: resource_generate.GeneratorSetNestedBlock{
				Validators: []specschema.SetValidator{
					{
						Custom: &specschema.CustomValidator{
							Import:           pointer("github.com/.../myvalidator"),
							SchemaDefinition: "myvalidator.Validate()",
						},
					},
				},
			},
		},
		"plan-modifiers": {
			input: &resource.SetNestedBlock{
				PlanModifiers: []specschema.SetPlanModifier{
					{
						Custom: &specschema.CustomPlanModifier{
							Import:           pointer("github.com/.../my_planmodifier"),
							SchemaDefinition: "my_planmodifier.Modify()",
						},
					},
				},
			},
			expected: resource_generate.GeneratorSetNestedBlock{
				PlanModifiers: []specschema.SetPlanModifier{
					{
						Custom: &specschema.CustomPlanModifier{
							Import:           pointer("github.com/.../my_planmodifier"),
							SchemaDefinition: "my_planmodifier.Modify()",
						},
					},
				},
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := convertSetNestedBlock(testCase.input)

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}