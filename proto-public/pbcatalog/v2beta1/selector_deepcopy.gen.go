// Code generated by protoc-gen-deepcopy. DO NOT EDIT.
package catalogv2beta1

import (
	proto "google.golang.org/protobuf/proto"
)

// DeepCopyInto supports using WorkloadSelector within kubernetes types, where deepcopy-gen is used.
func (in *WorkloadSelector) DeepCopyInto(out *WorkloadSelector) {
	proto.Reset(out)
	proto.Merge(out, proto.Clone(in))
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadSelector. Required by controller-gen.
func (in *WorkloadSelector) DeepCopy() *WorkloadSelector {
	if in == nil {
		return nil
	}
	out := new(WorkloadSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadSelector. Required by controller-gen.
func (in *WorkloadSelector) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}
