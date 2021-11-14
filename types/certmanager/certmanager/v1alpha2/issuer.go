// *** WARNING: this file was generated by crd2pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha2

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An Issuer represents a certificate issuing authority which can be referenced as part of `issuerRef` fields. It is scoped to a single namespace and can therefore only be referenced by resources within the same namespace.
type Issuer struct {
	pulumi.CustomResourceState

	ApiVersion pulumi.StringPtrOutput     `pulumi:"apiVersion"`
	Kind       pulumi.StringPtrOutput     `pulumi:"kind"`
	Metadata   metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// Desired state of the Issuer resource.
	Spec IssuerSpecPtrOutput `pulumi:"spec"`
	// Status of the Issuer. This is set and managed automatically.
	Status IssuerStatusPtrOutput `pulumi:"status"`
}

// NewIssuer registers a new resource with the given unique name, arguments, and options.
func NewIssuer(ctx *pulumi.Context,
	name string, args *IssuerArgs, opts ...pulumi.ResourceOption) (*Issuer, error) {
	if args == nil {
		args = &IssuerArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("cert-manager.io/v1alpha2")
	args.Kind = pulumi.StringPtr("Issuer")
	var resource Issuer
	err := ctx.RegisterResource("kubernetes:cert-manager.io/v1alpha2:Issuer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIssuer gets an existing Issuer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIssuer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IssuerState, opts ...pulumi.ResourceOption) (*Issuer, error) {
	var resource Issuer
	err := ctx.ReadResource("kubernetes:cert-manager.io/v1alpha2:Issuer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Issuer resources.
type issuerState struct {
	ApiVersion *string            `pulumi:"apiVersion"`
	Kind       *string            `pulumi:"kind"`
	Metadata   *metav1.ObjectMeta `pulumi:"metadata"`
	// Desired state of the Issuer resource.
	Spec *IssuerSpec `pulumi:"spec"`
	// Status of the Issuer. This is set and managed automatically.
	Status *IssuerStatus `pulumi:"status"`
}

type IssuerState struct {
	ApiVersion pulumi.StringPtrInput
	Kind       pulumi.StringPtrInput
	Metadata   metav1.ObjectMetaPtrInput
	// Desired state of the Issuer resource.
	Spec IssuerSpecPtrInput
	// Status of the Issuer. This is set and managed automatically.
	Status IssuerStatusPtrInput
}

func (IssuerState) ElementType() reflect.Type {
	return reflect.TypeOf((*issuerState)(nil)).Elem()
}

type issuerArgs struct {
	ApiVersion *string            `pulumi:"apiVersion"`
	Kind       *string            `pulumi:"kind"`
	Metadata   *metav1.ObjectMeta `pulumi:"metadata"`
	// Desired state of the Issuer resource.
	Spec *IssuerSpec `pulumi:"spec"`
	// Status of the Issuer. This is set and managed automatically.
	Status *IssuerStatus `pulumi:"status"`
}

// The set of arguments for constructing a Issuer resource.
type IssuerArgs struct {
	ApiVersion pulumi.StringPtrInput
	Kind       pulumi.StringPtrInput
	Metadata   metav1.ObjectMetaPtrInput
	// Desired state of the Issuer resource.
	Spec IssuerSpecPtrInput
	// Status of the Issuer. This is set and managed automatically.
	Status IssuerStatusPtrInput
}

func (IssuerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*issuerArgs)(nil)).Elem()
}

type IssuerInput interface {
	pulumi.Input

	ToIssuerOutput() IssuerOutput
	ToIssuerOutputWithContext(ctx context.Context) IssuerOutput
}

func (*Issuer) ElementType() reflect.Type {
	return reflect.TypeOf((*Issuer)(nil))
}

func (i *Issuer) ToIssuerOutput() IssuerOutput {
	return i.ToIssuerOutputWithContext(context.Background())
}

func (i *Issuer) ToIssuerOutputWithContext(ctx context.Context) IssuerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IssuerOutput)
}

type IssuerOutput struct {
	*pulumi.OutputState
}

func (IssuerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Issuer)(nil))
}

func (o IssuerOutput) ToIssuerOutput() IssuerOutput {
	return o
}

func (o IssuerOutput) ToIssuerOutputWithContext(ctx context.Context) IssuerOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(IssuerOutput{})
}
