// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package example

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
	"go-overridden-internal-module-name/example/utilities"
)

type Resource struct {
	pulumi.CustomResourceState

	Bar pulumi.StringPtrOutput `pulumi:"bar"`
	Baz pulumi.StringPtrOutput `pulumi:"baz"`
}

// NewResource registers a new resource with the given unique name, arguments, and options.
func NewResource(ctx *pulumi.Context,
	name string, args *ResourceArgs, opts ...pulumi.ResourceOption) (*Resource, error) {
	if args == nil {
		args = &ResourceArgs{}
	}

	if args.Bar != nil {
		args.Bar = pulumi.ToSecret(args.Bar).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"bar",
		"baz",
	})
	opts = append(opts, secrets)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Resource
	err := ctx.RegisterResource("example::Resource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResource gets an existing Resource resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourceState, opts ...pulumi.ResourceOption) (*Resource, error) {
	var resource Resource
	err := ctx.ReadResource("example::Resource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Resource resources.
type resourceState struct {
}

type ResourceState struct {
}

func (ResourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceState)(nil)).Elem()
}

type resourceArgs struct {
	Bar *string `pulumi:"bar"`
}

// The set of arguments for constructing a Resource resource.
type ResourceArgs struct {
	Bar pulumi.StringPtrInput
}

func (ResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceArgs)(nil)).Elem()
}

type ResourceInput interface {
	pulumi.Input

	ToResourceOutput() ResourceOutput
	ToResourceOutputWithContext(ctx context.Context) ResourceOutput
}

func (*Resource) ElementType() reflect.Type {
	return reflect.TypeOf((**Resource)(nil)).Elem()
}

func (i *Resource) ToResourceOutput() ResourceOutput {
	return i.ToResourceOutputWithContext(context.Background())
}

func (i *Resource) ToResourceOutputWithContext(ctx context.Context) ResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceOutput)
}

func (i *Resource) ToOutput(ctx context.Context) pulumix.Output[*Resource] {
	return pulumix.Output[*Resource]{
		OutputState: i.ToResourceOutputWithContext(ctx).OutputState,
	}
}

type ResourceOutput struct{ *pulumi.OutputState }

func (ResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Resource)(nil)).Elem()
}

func (o ResourceOutput) ToResourceOutput() ResourceOutput {
	return o
}

func (o ResourceOutput) ToResourceOutputWithContext(ctx context.Context) ResourceOutput {
	return o
}

func (o ResourceOutput) ToOutput(ctx context.Context) pulumix.Output[*Resource] {
	return pulumix.Output[*Resource]{
		OutputState: o.OutputState,
	}
}

func (o ResourceOutput) Bar() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Resource) pulumi.StringPtrOutput { return v.Bar }).(pulumi.StringPtrOutput)
}

func (o ResourceOutput) Baz() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Resource) pulumi.StringPtrOutput { return v.Baz }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceInput)(nil)).Elem(), &Resource{})
	pulumi.RegisterOutputType(ResourceOutput{})
}
