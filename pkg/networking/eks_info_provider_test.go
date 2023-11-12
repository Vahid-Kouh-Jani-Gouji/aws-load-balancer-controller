package networking

import (
	"context"
	"net/netip"
	"testing"

	awssdk "github.com/aws/aws-sdk-go/aws"
	ec2sdk "github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"sigs.k8s.io/aws-load-balancer-controller/pkg/aws/services"
)

func Test_defaultEKSInfoResolver_ListCIDRs(t *testing.T) {
	type describeClusterWithContextCall struct {
		req  *eks.DescribeClusterInput
		resp *eks.DescribeClusterOutput
		err  error
	}
	type describeSubnetsCall struct {
		req  *ec2sdk.DescribeSubnetsInput
		resp *ec2sdk.DescribeSubnetsOutput
		err  error
	}
	tests := []struct {
		name        string
		clusterName string
		clusterCall describeClusterWithContextCall
		subnetCall  describeSubnetsCall
		want        []netip.Prefix
		wantErr     error
	}{
		{
			name:        "list CIDRs",
			clusterName: "cluster_1",
			clusterCall: describeClusterWithContextCall{
				req: &eks.DescribeClusterInput{
					Name: awssdk.String("cluster_1"),
				},
				resp: &eks.DescribeClusterOutput{
					Cluster: &eks.Cluster{
						ResourcesVpcConfig: &eks.VpcConfigResponse{
							SubnetIds: awssdk.StringSlice([]string{"subnet-01234567890abcdef", "subnet-0bb1c79de3abcdef"}),
						},
					},
				},
			},
			subnetCall: describeSubnetsCall{
				req: &ec2sdk.DescribeSubnetsInput{
					SubnetIds: awssdk.StringSlice([]string{"subnet-01234567890abcdef", "subnet-0bb1c79de3abcdef"}),
				},
				resp: &ec2sdk.DescribeSubnetsOutput{
					Subnets: []*ec2sdk.Subnet{
						{
							CidrBlock: awssdk.String("10.0.0.0/24"),
						},
						{
							CidrBlock: awssdk.String("10.0.1.0/24"),
						},
					},
				},
			},
			want: []netip.Prefix{
				netip.MustParsePrefix("10.0.0.0/24"),
				netip.MustParsePrefix("10.0.1.0/24"),
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			eksClient := services.NewMockEKS(ctrl)
			eksClient.EXPECT().DescribeClusterWithContext(gomock.Any(), tt.clusterCall.req).Return(tt.clusterCall.resp, tt.clusterCall.err)
			ec2Client := services.NewMockEC2(ctrl)
			ec2Client.EXPECT().DescribeSubnets(tt.subnetCall.req).Return(tt.subnetCall.resp, tt.subnetCall.err)

			c := &defaultEKSInfoResolver{
				eksClient:   eksClient,
				ec2Client:   ec2Client,
				clusterName: tt.clusterName,
			}
			got, err := c.ListCIDRs(context.Background())
			if tt.wantErr != nil {
				assert.EqualError(t, err, tt.wantErr.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func Test_defaultEKSInfoResolver_IsInstanceInCluster(t *testing.T) {
	type describeInstancesAsListCall struct {
		req  *ec2sdk.DescribeInstancesInput
		resp []*ec2sdk.Instance
		err  error
	}
	type args struct {
		instanceID string
	}
	tests := []struct {
		name        string
		clusterName string
		call        describeInstancesAsListCall
		args        args
		want        bool
		wantErr     error
	}{
		{
			name:        "instance is a node",
			clusterName: "cluster_1",
			call: describeInstancesAsListCall{
				req: &ec2sdk.DescribeInstancesInput{
					InstanceIds: []*string{
						awssdk.String("i-0fa2d0064e848c69e"),
					},
				},
				resp: []*ec2sdk.Instance{
					{
						Tags: []*ec2sdk.Tag{
							{
								Key:   awssdk.String("kubernetes.io/cluster/cluster_1"),
								Value: awssdk.String("owned"),
							},
						},
					},
				},
			},
			args:    args{instanceID: "i-0fa2d0064e848c69e"},
			want:    true,
			wantErr: nil,
		},
		{
			name:        "instance is not node",
			clusterName: "cluster_1",
			call: describeInstancesAsListCall{
				req: &ec2sdk.DescribeInstancesInput{
					InstanceIds: []*string{
						awssdk.String("i-0fa2d0064e848c69e"),
					},
				},
				resp: []*ec2sdk.Instance{
					{
						Tags: []*ec2sdk.Tag{
							{
								Key:   awssdk.String("terraform"),
								Value: awssdk.String("true"),
							},
						},
					},
				},
			},
			args:    args{instanceID: "i-0fa2d0064e848c69e"},
			want:    false,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			eksClient := services.NewMockEKS(ctrl)
			ec2Client := services.NewMockEC2(ctrl)
			ec2Client.EXPECT().DescribeInstancesAsList(gomock.Any(), tt.call.req).Return(tt.call.resp, tt.call.err)

			c := &defaultEKSInfoResolver{
				eksClient:   eksClient,
				ec2Client:   ec2Client,
				clusterName: tt.clusterName,
			}
			got, err := c.IsInstanceInCluster(context.Background(), tt.args.instanceID)
			if tt.wantErr != nil {
				assert.EqualError(t, err, tt.wantErr.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}