// +build !ignore_autogenerated

/*
Copyright 2021.

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
func (in *ConsoleApplication) DeepCopyInto(out *ConsoleApplication) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsoleApplication.
func (in *ConsoleApplication) DeepCopy() *ConsoleApplication {
	if in == nil {
		return nil
	}
	out := new(ConsoleApplication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConsoleApplication) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsoleApplicationList) DeepCopyInto(out *ConsoleApplicationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ConsoleApplication, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsoleApplicationList.
func (in *ConsoleApplicationList) DeepCopy() *ConsoleApplicationList {
	if in == nil {
		return nil
	}
	out := new(ConsoleApplicationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConsoleApplicationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsoleApplicationSpec) DeepCopyInto(out *ConsoleApplicationSpec) {
	*out = *in
	if in.IsEnabled != nil {
		in, out := &in.IsEnabled, &out.IsEnabled
		*out = new(bool)
		**out = **in
	}
	if in.KfdefApplications != nil {
		in, out := &in.KfdefApplications, &out.KfdefApplications
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.Enable.DeepCopyInto(&out.Enable)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsoleApplicationSpec.
func (in *ConsoleApplicationSpec) DeepCopy() *ConsoleApplicationSpec {
	if in == nil {
		return nil
	}
	out := new(ConsoleApplicationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsoleApplicationStatus) DeepCopyInto(out *ConsoleApplicationStatus) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsoleApplicationStatus.
func (in *ConsoleApplicationStatus) DeepCopy() *ConsoleApplicationStatus {
	if in == nil {
		return nil
	}
	out := new(ConsoleApplicationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Enable) DeepCopyInto(out *Enable) {
	*out = *in
	if in.Variables != nil {
		in, out := &in.Variables, &out.Variables
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.VariableDisplayText != nil {
		in, out := &in.VariableDisplayText, &out.VariableDisplayText
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.VariableHelpText != nil {
		in, out := &in.VariableHelpText, &out.VariableHelpText
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Enable.
func (in *Enable) DeepCopy() *Enable {
	if in == nil {
		return nil
	}
	out := new(Enable)
	in.DeepCopyInto(out)
	return out
}
