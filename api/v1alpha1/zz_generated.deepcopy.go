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

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GraphSpec) DeepCopyInto(out *GraphSpec) {
	*out = *in
	in.Pod.DeepCopyInto(&out.Pod)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GraphSpec.
func (in *GraphSpec) DeepCopy() *GraphSpec {
	if in == nil {
		return nil
	}
	out := new(GraphSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Repro) DeepCopyInto(out *Repro) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Repro.
func (in *Repro) DeepCopy() *Repro {
	if in == nil {
		return nil
	}
	out := new(Repro)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Repro) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReproList) DeepCopyInto(out *ReproList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Repro, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReproList.
func (in *ReproList) DeepCopy() *ReproList {
	if in == nil {
		return nil
	}
	out := new(ReproList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReproList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReproSpec) DeepCopyInto(out *ReproSpec) {
	*out = *in
	in.Graph.DeepCopyInto(&out.Graph)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReproSpec.
func (in *ReproSpec) DeepCopy() *ReproSpec {
	if in == nil {
		return nil
	}
	out := new(ReproSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReproStatus) DeepCopyInto(out *ReproStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReproStatus.
func (in *ReproStatus) DeepCopy() *ReproStatus {
	if in == nil {
		return nil
	}
	out := new(ReproStatus)
	in.DeepCopyInto(out)
	return out
}
