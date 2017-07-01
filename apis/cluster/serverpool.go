package cluster

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	ServerPoolType_Master = "master"
	ServerPoolType_Node   = "node"
	ServerPoolType_Hybrid = "hybrid"
)

type ServerPool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Count             int
	Type              string
	Name              string
	Subnets           []*Subnet
}
