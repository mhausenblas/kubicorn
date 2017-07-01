package profiles

import "github.com/kris-nova/kubicorn/apis/cluster"

func NewSimpleAmazonCluster(name string) *cluster.Cluster {
	return &cluster.Cluster{
		Name:     name,
		Cloud:    cluster.Cloud_Amazon,
		Location: "us-west-2",
		Network: &cluster.Network{
			Type: cluster.NetworkType_Public,
			CIDR: "10.0.0.0/16",
		},
		ServerPools: []*cluster.ServerPool{
			{
				Type: cluster.ServerPoolType_Master,
				Name: "amazon-master",
				Subnets: []*cluster.Subnet{
					{
						CIDR:     "10.0.0.0/24",
						Location: "us-west-2a",
					},
				},
			},
			{
				Type: cluster.ServerPoolType_Node,
				Name: "amazon-node",
				Subnets: []*cluster.Subnet{
					{
						CIDR:     "10.0.100.0/24",
						Location: "us-west-2b",
					},
				},
			},
		},
	}
}
