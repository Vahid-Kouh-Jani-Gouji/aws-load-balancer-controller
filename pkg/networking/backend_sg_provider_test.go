package networking

import (
	"context"
	networking "k8s.io/api/networking/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	mock_client "sigs.k8s.io/aws-load-balancer-controller/mocks/controller-runtime/client"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"testing"

	awssdk "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	ec2sdk "github.com/aws/aws-sdk-go/service/ec2"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"sigs.k8s.io/aws-load-balancer-controller/pkg/aws/services"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

const (
	defaultVPCID       = "vpc-xxxyyy"
	defaultClusterName = "testCluster"
)

func Test_defaultBackendSGProvider_Get(t *testing.T) {
	type describeSecurityGroupsAsListCall struct {
		req  *ec2sdk.DescribeSecurityGroupsInput
		resp []*ec2sdk.SecurityGroup
		err  error
	}
	type createSecurityGroupWithContexCall struct {
		req  *ec2sdk.CreateSecurityGroupInput
		resp *ec2sdk.CreateSecurityGroupOutput
		err  error
	}
	type fields struct {
		enabled         bool
		backendSGs      []string
		describeSGCalls []describeSecurityGroupsAsListCall
		createSGCalls   []createSecurityGroupWithContexCall
	}
	defaultEC2Filters := []*ec2sdk.Filter{
		{
			Name:   awssdk.String("vpc-id"),
			Values: awssdk.StringSlice([]string{defaultVPCID}),
		},
		{
			Name:   awssdk.String("group-name"),
			Values: awssdk.StringSlice([]string{"k8s-testCluster-traffic-a0c9fe55ad53c12"}),
		},
	}
	tests := []struct {
		name    string
		want    []string
		fields  fields
		wantErr error
	}{
		{
			name: "backend sg disabled",
		},
		{
			name: "backend sg enabled",
			fields: fields{
				enabled:    true,
				backendSGs: []string{"sg-xxx", "sg-yyy"},
			},
			want: []string{"sg-xxx", "sg-yyy"},
		},
		{
			name: "backend sg enabled, auto-gen, SG exists",
			fields: fields{
				enabled: true,
				describeSGCalls: []describeSecurityGroupsAsListCall{
					{
						req: &ec2sdk.DescribeSecurityGroupsInput{
							Filters: defaultEC2Filters,
						},
						resp: []*ec2sdk.SecurityGroup{
							{
								GroupId: awssdk.String("sg-autogen"),
							},
						},
					},
				},
			},
			want: []string{"sg-autogen"},
		},
		{
			name: "backend sg enabled, auto-gen new SG",
			fields: fields{
				enabled: true,
				describeSGCalls: []describeSecurityGroupsAsListCall{
					{
						req: &ec2sdk.DescribeSecurityGroupsInput{
							Filters: defaultEC2Filters,
						},
						err: awserr.New("InvalidGroup.NotFound", "", nil),
					},
				},
				createSGCalls: []createSecurityGroupWithContexCall{
					{
						req: &ec2sdk.CreateSecurityGroupInput{
							Description: awssdk.String(sgDescription),
							GroupName:   awssdk.String("k8s-testCluster-traffic-a0c9fe55ad53c12"),
							TagSpecifications: []*ec2sdk.TagSpecification{
								{
									ResourceType: awssdk.String("security-group"),
									Tags: []*ec2sdk.Tag{
										{
											Key:   awssdk.String("elbv2.k8s.aws/cluster"),
											Value: awssdk.String(defaultClusterName),
										},
										{
											Key:   awssdk.String("elbv2.k8s.aws/type"),
											Value: awssdk.String("backend"),
										},
									},
								},
							},
							VpcId: awssdk.String(defaultVPCID),
						},
						resp: &ec2sdk.CreateSecurityGroupOutput{
							GroupId: awssdk.String("sg-newauto"),
						},
					},
				},
			},
			want: []string{"sg-newauto"},
		},
		{
			name: "describe SG call returns error",
			fields: fields{
				enabled: true,
				describeSGCalls: []describeSecurityGroupsAsListCall{
					{
						req: &ec2sdk.DescribeSecurityGroupsInput{
							Filters: defaultEC2Filters,
						},
						err: awserr.New("Some.Other.Error", "describe security group as list error", nil),
					},
				},
			},
			wantErr: errors.New("Some.Other.Error: describe security group as list error"),
		},
		{
			name: "create SG call returns error",
			fields: fields{
				enabled: true,
				describeSGCalls: []describeSecurityGroupsAsListCall{
					{
						req: &ec2sdk.DescribeSecurityGroupsInput{
							Filters: defaultEC2Filters,
						},
						err: awserr.New("InvalidGroup.NotFound", "", nil),
					},
				},
				createSGCalls: []createSecurityGroupWithContexCall{
					{
						req: &ec2sdk.CreateSecurityGroupInput{
							Description: awssdk.String(sgDescription),
							GroupName:   awssdk.String("k8s-testCluster-traffic-a0c9fe55ad53c12"),
							TagSpecifications: []*ec2sdk.TagSpecification{
								{
									ResourceType: awssdk.String("security-group"),
									Tags: []*ec2sdk.Tag{
										{
											Key:   awssdk.String("elbv2.k8s.aws/cluster"),
											Value: awssdk.String(defaultClusterName),
										},
										{
											Key:   awssdk.String("elbv2.k8s.aws/type"),
											Value: awssdk.String("backend"),
										},
									},
								},
							},
							VpcId: awssdk.String(defaultVPCID),
						},
						err: awserr.New("Create.Error", "unable to create security group", nil),
					},
				},
			},
			wantErr: errors.New("Create.Error: unable to create security group"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ec2Client := services.NewMockEC2(ctrl)
			for _, call := range tt.fields.describeSGCalls {
				ec2Client.EXPECT().DescribeSecurityGroupsAsList(context.Background(), call.req).Return(call.resp, call.err)
			}
			for _, call := range tt.fields.createSGCalls {
				ec2Client.EXPECT().CreateSecurityGroupWithContext(context.Background(), call.req).Return(call.resp, call.err)
			}
			k8sClient := mock_client.NewMockClient(ctrl)
			sgProvider := NewBackendSGProvider(defaultClusterName, tt.fields.enabled, tt.fields.backendSGs,
				defaultVPCID, ec2Client, k8sClient, &log.NullLogger{})

			got, err := sgProvider.Get(context.Background())
			if tt.wantErr != nil {
				assert.EqualError(t, err, tt.wantErr.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func Test_defaultBackendSGProvider_Release(t *testing.T) {
	type env struct {
		ingresses []*networking.Ingress
	}
	type listIngressCall struct {
		ingresses []*networking.Ingress
		err       error
	}
	type deleteSecurityGroupWithContextCall struct {
		req  *ec2sdk.DeleteSecurityGroupInput
		resp *ec2sdk.DeleteSecurityGroupOutput
		err  error
	}
	type fields struct {
		enabled          bool
		autogenSG        string
		backendSGs       []string
		listIngressCalls []listIngressCall
		deleteSGCalls    []deleteSecurityGroupWithContextCall
	}
	tests := []struct {
		name    string
		env     env
		fields  fields
		wantErr error
	}{
		{
			name: "backend sg not enabled",
		},
		{
			name: "backend sg specified via flags",
			fields: fields{
				enabled:    true,
				backendSGs: []string{"sg-first", "sg-second"},
			},
		},
		{
			name: "backend sg autogenerated",
			fields: fields{
				enabled:   true,
				autogenSG: "sg-autogen",
				listIngressCalls: []listIngressCall{
					{
						ingresses: []*networking.Ingress{},
					},
				},
				deleteSGCalls: []deleteSecurityGroupWithContextCall{
					{
						req: &ec2sdk.DeleteSecurityGroupInput{
							GroupId: awssdk.String("sg-autogen"),
						},
						resp: &ec2sdk.DeleteSecurityGroupOutput{},
					},
				},
			},
		},
		{
			name: "backend sg required due to standalone ingress",
			fields: fields{
				enabled:   true,
				autogenSG: "sg-autogen",
				listIngressCalls: []listIngressCall{
					{
						ingresses: []*networking.Ingress{
							{
								ObjectMeta: metav1.ObjectMeta{
									Namespace: "regular-ns",
									Name:      "ing-nofinalizer",
								},
							},
							{
								ObjectMeta: metav1.ObjectMeta{
									Namespace:  "awesome-ns",
									Name:       "ing-1",
									Finalizers: []string{"ingress.k8s.aws/resources"},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "backend sg required for ingress group",
			fields: fields{
				enabled:   true,
				autogenSG: "sg-autogen",
				listIngressCalls: []listIngressCall{
					{
						ingresses: []*networking.Ingress{
							{
								ObjectMeta: metav1.ObjectMeta{
									Namespace:  "awesome-ns",
									Name:       "ing-1",
									Finalizers: []string{"group.ingress.k8s.aws/awesome-group"},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "First SG delete attempt fails",
			fields: fields{
				enabled:   true,
				autogenSG: "sg-autogen",
				listIngressCalls: []listIngressCall{
					{
						ingresses: []*networking.Ingress{},
					},
				},
				deleteSGCalls: []deleteSecurityGroupWithContextCall{
					{
						req: &ec2sdk.DeleteSecurityGroupInput{
							GroupId: awssdk.String("sg-autogen"),
						},
						err: awserr.New("DependencyViolation", "", nil),
					},
					{
						req: &ec2sdk.DeleteSecurityGroupInput{
							GroupId: awssdk.String("sg-autogen"),
						},
						resp: &ec2sdk.DeleteSecurityGroupOutput{},
					},
				},
			},
		},
		{
			name: "SG delete attempt fails return non-dependency violation error",
			fields: fields{
				enabled:   true,
				autogenSG: "sg-autogen",
				listIngressCalls: []listIngressCall{
					{},
					{},
				},
				deleteSGCalls: []deleteSecurityGroupWithContextCall{
					{
						req: &ec2sdk.DeleteSecurityGroupInput{
							GroupId: awssdk.String("sg-autogen"),
						},
						err: awserr.New("Something.Else", "unable to delete SG", nil),
					},
				},
			},
			wantErr: errors.New("failed to delete securityGroup: Something.Else: unable to delete SG"),
		},
		{
			name: "k8s list returns error",
			fields: fields{
				enabled:   true,
				autogenSG: "sg-autogen",
				listIngressCalls: []listIngressCall{
					{
						err: errors.New("failed"),
					},
				},
			},
			wantErr: errors.New("unable to list ingresses: failed"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ec2Client := services.NewMockEC2(ctrl)
			k8sClient := mock_client.NewMockClient(ctrl)
			sgProvider := NewBackendSGProvider(defaultClusterName, tt.fields.enabled, tt.fields.backendSGs,
				defaultVPCID, ec2Client, k8sClient, &log.NullLogger{})
			if len(tt.fields.autogenSG) > 0 {
				sgProvider.backendSGs = nil
				sgProvider.autoGeneratedSG = tt.fields.autogenSG
			}
			var deleteCalls []*gomock.Call
			for _, call := range tt.fields.deleteSGCalls {
				deleteCalls = append(deleteCalls, ec2Client.EXPECT().DeleteSecurityGroupWithContext(context.Background(), call.req).Return(call.resp, call.err))
			}
			if len(deleteCalls) > 0 {
				gomock.InAnyOrder(deleteCalls)
			}
			for _, call := range tt.fields.listIngressCalls {
				k8sClient.EXPECT().List(gomock.Any(), &networking.IngressList{}, gomock.Any()).DoAndReturn(
					func(ctx context.Context, ingList *networking.IngressList, opts ...client.ListOption) error {
						for _, ing := range call.ingresses {
							ingList.Items = append(ingList.Items, *(ing.DeepCopy()))
						}
						return call.err
					},
				).AnyTimes()
			}
			for _, ing := range tt.env.ingresses {
				assert.NoError(t, k8sClient.Create(context.Background(), ing.DeepCopy()))
			}
			gotErr := sgProvider.Release(context.Background())
			if tt.wantErr != nil {
				assert.EqualError(t, gotErr, tt.wantErr.Error())
			} else {
				assert.NoError(t, gotErr)
			}
		})
	}
}
