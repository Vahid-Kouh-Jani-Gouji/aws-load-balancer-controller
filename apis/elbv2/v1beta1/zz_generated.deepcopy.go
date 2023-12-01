//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Attribute) DeepCopyInto(out *Attribute) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Attribute.
func (in *Attribute) DeepCopy() *Attribute {
	if in == nil {
		return nil
	}
	out := new(Attribute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPBlock) DeepCopyInto(out *IPBlock) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPBlock.
func (in *IPBlock) DeepCopy() *IPBlock {
	if in == nil {
		return nil
	}
	out := new(IPBlock)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressClassParams) DeepCopyInto(out *IngressClassParams) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressClassParams.
func (in *IngressClassParams) DeepCopy() *IngressClassParams {
	if in == nil {
		return nil
	}
	out := new(IngressClassParams)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IngressClassParams) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressClassParamsList) DeepCopyInto(out *IngressClassParamsList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IngressClassParams, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressClassParamsList.
func (in *IngressClassParamsList) DeepCopy() *IngressClassParamsList {
	if in == nil {
		return nil
	}
	out := new(IngressClassParamsList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IngressClassParamsList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressClassParamsSpec) DeepCopyInto(out *IngressClassParamsSpec) {
	*out = *in
	if in.NamespaceSelector != nil {
		in, out := &in.NamespaceSelector, &out.NamespaceSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.Group != nil {
		in, out := &in.Group, &out.Group
		*out = new(IngressGroup)
		**out = **in
	}
	if in.Scheme != nil {
		in, out := &in.Scheme, &out.Scheme
		*out = new(LoadBalancerScheme)
		**out = **in
	}
	if in.InboundCIDRs != nil {
		in, out := &in.InboundCIDRs, &out.InboundCIDRs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Subnets != nil {
		in, out := &in.Subnets, &out.Subnets
		*out = new(SubnetSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.IPAddressType != nil {
		in, out := &in.IPAddressType, &out.IPAddressType
		*out = new(IPAddressType)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]Tag, len(*in))
		copy(*out, *in)
	}
	if in.LoadBalancerAttributes != nil {
		in, out := &in.LoadBalancerAttributes, &out.LoadBalancerAttributes
		*out = make([]Attribute, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressClassParamsSpec.
func (in *IngressClassParamsSpec) DeepCopy() *IngressClassParamsSpec {
	if in == nil {
		return nil
	}
	out := new(IngressClassParamsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressGroup) DeepCopyInto(out *IngressGroup) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressGroup.
func (in *IngressGroup) DeepCopy() *IngressGroup {
	if in == nil {
		return nil
	}
	out := new(IngressGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkingIngressRule) DeepCopyInto(out *NetworkingIngressRule) {
	*out = *in
	if in.From != nil {
		in, out := &in.From, &out.From
		*out = make([]NetworkingPeer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]NetworkingPort, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkingIngressRule.
func (in *NetworkingIngressRule) DeepCopy() *NetworkingIngressRule {
	if in == nil {
		return nil
	}
	out := new(NetworkingIngressRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkingPeer) DeepCopyInto(out *NetworkingPeer) {
	*out = *in
	if in.IPBlock != nil {
		in, out := &in.IPBlock, &out.IPBlock
		*out = new(IPBlock)
		**out = **in
	}
	if in.SecurityGroup != nil {
		in, out := &in.SecurityGroup, &out.SecurityGroup
		*out = new(SecurityGroup)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkingPeer.
func (in *NetworkingPeer) DeepCopy() *NetworkingPeer {
	if in == nil {
		return nil
	}
	out := new(NetworkingPeer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkingPort) DeepCopyInto(out *NetworkingPort) {
	*out = *in
	if in.Protocol != nil {
		in, out := &in.Protocol, &out.Protocol
		*out = new(NetworkingProtocol)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(intstr.IntOrString)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkingPort.
func (in *NetworkingPort) DeepCopy() *NetworkingPort {
	if in == nil {
		return nil
	}
	out := new(NetworkingPort)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityGroup) DeepCopyInto(out *SecurityGroup) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityGroup.
func (in *SecurityGroup) DeepCopy() *SecurityGroup {
	if in == nil {
		return nil
	}
	out := new(SecurityGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceReference) DeepCopyInto(out *ServiceReference) {
	*out = *in
	out.Port = in.Port
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceReference.
func (in *ServiceReference) DeepCopy() *ServiceReference {
	if in == nil {
		return nil
	}
	out := new(ServiceReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetSelector) DeepCopyInto(out *SubnetSelector) {
	*out = *in
	if in.IDs != nil {
		in, out := &in.IDs, &out.IDs
		*out = make([]SubnetID, len(*in))
		copy(*out, *in)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string][]string, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetSelector.
func (in *SubnetSelector) DeepCopy() *SubnetSelector {
	if in == nil {
		return nil
	}
	out := new(SubnetSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Tag) DeepCopyInto(out *Tag) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Tag.
func (in *Tag) DeepCopy() *Tag {
	if in == nil {
		return nil
	}
	out := new(Tag)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetGroupBinding) DeepCopyInto(out *TargetGroupBinding) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetGroupBinding.
func (in *TargetGroupBinding) DeepCopy() *TargetGroupBinding {
	if in == nil {
		return nil
	}
	out := new(TargetGroupBinding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TargetGroupBinding) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetGroupBindingList) DeepCopyInto(out *TargetGroupBindingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TargetGroupBinding, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetGroupBindingList.
func (in *TargetGroupBindingList) DeepCopy() *TargetGroupBindingList {
	if in == nil {
		return nil
	}
	out := new(TargetGroupBindingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TargetGroupBindingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetGroupBindingNetworking) DeepCopyInto(out *TargetGroupBindingNetworking) {
	*out = *in
	if in.Ingress != nil {
		in, out := &in.Ingress, &out.Ingress
		*out = make([]NetworkingIngressRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetGroupBindingNetworking.
func (in *TargetGroupBindingNetworking) DeepCopy() *TargetGroupBindingNetworking {
	if in == nil {
		return nil
	}
	out := new(TargetGroupBindingNetworking)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetGroupBindingSpec) DeepCopyInto(out *TargetGroupBindingSpec) {
	*out = *in
	if in.TargetType != nil {
		in, out := &in.TargetType, &out.TargetType
		*out = new(TargetType)
		**out = **in
	}
	out.ServiceRef = in.ServiceRef
	if in.Networking != nil {
		in, out := &in.Networking, &out.Networking
		*out = new(TargetGroupBindingNetworking)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.IPAddressType != nil {
		in, out := &in.IPAddressType, &out.IPAddressType
		*out = new(TargetGroupIPAddressType)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetGroupBindingSpec.
func (in *TargetGroupBindingSpec) DeepCopy() *TargetGroupBindingSpec {
	if in == nil {
		return nil
	}
	out := new(TargetGroupBindingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetGroupBindingStatus) DeepCopyInto(out *TargetGroupBindingStatus) {
	*out = *in
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetGroupBindingStatus.
func (in *TargetGroupBindingStatus) DeepCopy() *TargetGroupBindingStatus {
	if in == nil {
		return nil
	}
	out := new(TargetGroupBindingStatus)
	in.DeepCopyInto(out)
	return out
}
