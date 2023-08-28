// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"dashed-import-schema/plant-provider/internal"
	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

type Nursery struct {
	pulumi.CustomResourceState
}

// NewNursery registers a new resource with the given unique name, arguments, and options.
func NewNursery(ctx *pulumi.Context,
	name string, args *NurseryArgs, opts ...pulumi.ResourceOption) (*Nursery, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Varieties == nil {
		return nil, errors.New("invalid value for required argument 'Varieties'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Nursery
	err := ctx.RegisterResource("plant:tree/v1:Nursery", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNursery gets an existing Nursery resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNursery(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NurseryState, opts ...pulumi.ResourceOption) (*Nursery, error) {
	var resource Nursery
	err := ctx.ReadResource("plant:tree/v1:Nursery", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Nursery resources.
type nurseryState struct {
}

type NurseryState struct {
}

func (NurseryState) ElementType() reflect.Type {
	return reflect.TypeOf((*nurseryState)(nil)).Elem()
}

type nurseryArgs struct {
	// The sizes of trees available
	Sizes map[string]TreeSize `pulumi:"sizes"`
	// The varieties available
	Varieties []RubberTreeVariety `pulumi:"varieties"`
}

// The set of arguments for constructing a Nursery resource.
type NurseryArgs struct {
	// The sizes of trees available
	Sizes TreeSizeMapInput
	// The varieties available
	Varieties RubberTreeVarietyArrayInput
}

func (NurseryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nurseryArgs)(nil)).Elem()
}

type NurseryInput interface {
	pulumi.Input

	ToNurseryOutput() NurseryOutput
	ToNurseryOutputWithContext(ctx context.Context) NurseryOutput
}

func (*Nursery) ElementType() reflect.Type {
	return reflect.TypeOf((**Nursery)(nil)).Elem()
}

func (i *Nursery) ToNurseryOutput() NurseryOutput {
	return i.ToNurseryOutputWithContext(context.Background())
}

func (i *Nursery) ToNurseryOutputWithContext(ctx context.Context) NurseryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NurseryOutput)
}

func (i *Nursery) ToOutput(ctx context.Context) pulumix.Output[*Nursery] {
	return pulumix.Output[*Nursery]{
		OutputState: i.ToNurseryOutputWithContext(ctx).OutputState,
	}
}

type NurseryOutput struct{ *pulumi.OutputState }

func (NurseryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Nursery)(nil)).Elem()
}

func (o NurseryOutput) ToNurseryOutput() NurseryOutput {
	return o
}

func (o NurseryOutput) ToNurseryOutputWithContext(ctx context.Context) NurseryOutput {
	return o
}

func (o NurseryOutput) ToOutput(ctx context.Context) pulumix.Output[*Nursery] {
	return pulumix.Output[*Nursery]{
		OutputState: o.OutputState,
	}
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NurseryInput)(nil)).Elem(), &Nursery{})
	pulumi.RegisterOutputType(NurseryOutput{})
}
