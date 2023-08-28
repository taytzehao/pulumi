// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package example

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
	"replace-on-change/example/internal"
)

type Cat struct {
	pulumi.CustomResourceState

	Foes    ToyMapOutput           `pulumi:"foes"`
	Friends ToyArrayOutput         `pulumi:"friends"`
	Name    pulumi.StringPtrOutput `pulumi:"name"`
	Other   GodOutput              `pulumi:"other"`
	Toy     ToyPtrOutput           `pulumi:"toy"`
}

// NewCat registers a new resource with the given unique name, arguments, and options.
func NewCat(ctx *pulumi.Context,
	name string, args *CatArgs, opts ...pulumi.ResourceOption) (*Cat, error) {
	if args == nil {
		args = &CatArgs{}
	}

	secrets := pulumi.AdditionalSecretOutputs([]string{
		"name",
	})
	opts = append(opts, secrets)
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"foes.*.associated.color",
		"foes.*.color",
		"friends[*].associated.color",
		"friends[*].color",
		"name",
		"toy.color",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Cat
	err := ctx.RegisterResource("example::Cat", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCat gets an existing Cat resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCat(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CatState, opts ...pulumi.ResourceOption) (*Cat, error) {
	var resource Cat
	err := ctx.ReadResource("example::Cat", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Cat resources.
type catState struct {
}

type CatState struct {
}

func (CatState) ElementType() reflect.Type {
	return reflect.TypeOf((*catState)(nil)).Elem()
}

type catArgs struct {
}

// The set of arguments for constructing a Cat resource.
type CatArgs struct {
}

func (CatArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*catArgs)(nil)).Elem()
}

type CatInput interface {
	pulumi.Input

	ToCatOutput() CatOutput
	ToCatOutputWithContext(ctx context.Context) CatOutput
}

func (*Cat) ElementType() reflect.Type {
	return reflect.TypeOf((**Cat)(nil)).Elem()
}

func (i *Cat) ToCatOutput() CatOutput {
	return i.ToCatOutputWithContext(context.Background())
}

func (i *Cat) ToCatOutputWithContext(ctx context.Context) CatOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CatOutput)
}

func (i *Cat) ToOutput(ctx context.Context) pulumix.Output[*Cat] {
	return pulumix.Output[*Cat]{
		OutputState: i.ToCatOutputWithContext(ctx).OutputState,
	}
}

// CatArrayInput is an input type that accepts CatArray and CatArrayOutput values.
// You can construct a concrete instance of `CatArrayInput` via:
//
//	CatArray{ CatArgs{...} }
type CatArrayInput interface {
	pulumi.Input

	ToCatArrayOutput() CatArrayOutput
	ToCatArrayOutputWithContext(context.Context) CatArrayOutput
}

type CatArray []CatInput

func (CatArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Cat)(nil)).Elem()
}

func (i CatArray) ToCatArrayOutput() CatArrayOutput {
	return i.ToCatArrayOutputWithContext(context.Background())
}

func (i CatArray) ToCatArrayOutputWithContext(ctx context.Context) CatArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CatArrayOutput)
}

func (i CatArray) ToOutput(ctx context.Context) pulumix.Output[[]*Cat] {
	return pulumix.Output[[]*Cat]{
		OutputState: i.ToCatArrayOutputWithContext(ctx).OutputState,
	}
}

// CatMapInput is an input type that accepts CatMap and CatMapOutput values.
// You can construct a concrete instance of `CatMapInput` via:
//
//	CatMap{ "key": CatArgs{...} }
type CatMapInput interface {
	pulumi.Input

	ToCatMapOutput() CatMapOutput
	ToCatMapOutputWithContext(context.Context) CatMapOutput
}

type CatMap map[string]CatInput

func (CatMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Cat)(nil)).Elem()
}

func (i CatMap) ToCatMapOutput() CatMapOutput {
	return i.ToCatMapOutputWithContext(context.Background())
}

func (i CatMap) ToCatMapOutputWithContext(ctx context.Context) CatMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CatMapOutput)
}

func (i CatMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Cat] {
	return pulumix.Output[map[string]*Cat]{
		OutputState: i.ToCatMapOutputWithContext(ctx).OutputState,
	}
}

type CatOutput struct{ *pulumi.OutputState }

func (CatOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Cat)(nil)).Elem()
}

func (o CatOutput) ToCatOutput() CatOutput {
	return o
}

func (o CatOutput) ToCatOutputWithContext(ctx context.Context) CatOutput {
	return o
}

func (o CatOutput) ToOutput(ctx context.Context) pulumix.Output[*Cat] {
	return pulumix.Output[*Cat]{
		OutputState: o.OutputState,
	}
}

func (o CatOutput) Foes() ToyMapOutput {
	return o.ApplyT(func(v *Cat) ToyMapOutput { return v.Foes }).(ToyMapOutput)
}

func (o CatOutput) Friends() ToyArrayOutput {
	return o.ApplyT(func(v *Cat) ToyArrayOutput { return v.Friends }).(ToyArrayOutput)
}

func (o CatOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cat) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o CatOutput) Other() GodOutput {
	return o.ApplyT(func(v *Cat) GodOutput { return v.Other }).(GodOutput)
}

func (o CatOutput) Toy() ToyPtrOutput {
	return o.ApplyT(func(v *Cat) ToyPtrOutput { return v.Toy }).(ToyPtrOutput)
}

type CatArrayOutput struct{ *pulumi.OutputState }

func (CatArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Cat)(nil)).Elem()
}

func (o CatArrayOutput) ToCatArrayOutput() CatArrayOutput {
	return o
}

func (o CatArrayOutput) ToCatArrayOutputWithContext(ctx context.Context) CatArrayOutput {
	return o
}

func (o CatArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Cat] {
	return pulumix.Output[[]*Cat]{
		OutputState: o.OutputState,
	}
}

func (o CatArrayOutput) Index(i pulumi.IntInput) CatOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Cat {
		return vs[0].([]*Cat)[vs[1].(int)]
	}).(CatOutput)
}

type CatMapOutput struct{ *pulumi.OutputState }

func (CatMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Cat)(nil)).Elem()
}

func (o CatMapOutput) ToCatMapOutput() CatMapOutput {
	return o
}

func (o CatMapOutput) ToCatMapOutputWithContext(ctx context.Context) CatMapOutput {
	return o
}

func (o CatMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Cat] {
	return pulumix.Output[map[string]*Cat]{
		OutputState: o.OutputState,
	}
}

func (o CatMapOutput) MapIndex(k pulumi.StringInput) CatOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Cat {
		return vs[0].(map[string]*Cat)[vs[1].(string)]
	}).(CatOutput)
}

func init() {
	pulumi.RegisterOutputType(CatOutput{})
	pulumi.RegisterOutputType(CatArrayOutput{})
	pulumi.RegisterOutputType(CatMapOutput{})
}
