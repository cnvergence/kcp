/*
Copyright The KCP Authors.

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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

// ObjectReferenceApplyConfiguration represents a declarative configuration of the ObjectReference type for use
// with apply.
type ObjectReferenceApplyConfiguration struct {
	APIVersion *string `json:"apiVersion,omitempty"`
	Kind       *string `json:"kind,omitempty"`
	Name       *string `json:"name,omitempty"`
	Namespace  *string `json:"namespace,omitempty"`
}

// ObjectReferenceApplyConfiguration constructs a declarative configuration of the ObjectReference type for use with
// apply.
func ObjectReference() *ObjectReferenceApplyConfiguration {
	return &ObjectReferenceApplyConfiguration{}
}

// WithAPIVersion sets the APIVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIVersion field is set to the value of the last call.
func (b *ObjectReferenceApplyConfiguration) WithAPIVersion(value string) *ObjectReferenceApplyConfiguration {
	b.APIVersion = &value
	return b
}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *ObjectReferenceApplyConfiguration) WithKind(value string) *ObjectReferenceApplyConfiguration {
	b.Kind = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *ObjectReferenceApplyConfiguration) WithName(value string) *ObjectReferenceApplyConfiguration {
	b.Name = &value
	return b
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *ObjectReferenceApplyConfiguration) WithNamespace(value string) *ObjectReferenceApplyConfiguration {
	b.Namespace = &value
	return b
}
