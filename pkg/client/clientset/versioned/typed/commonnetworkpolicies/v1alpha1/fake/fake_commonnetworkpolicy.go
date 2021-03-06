/*
Copyright The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/bells17/common-network-policy-operator/pkg/apis/commonnetworkpolicies/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCommonNetworkPolicies implements CommonNetworkPolicyInterface
type FakeCommonNetworkPolicies struct {
	Fake *FakeCommonnetworkpoliciesV1alpha1
}

var commonnetworkpoliciesResource = schema.GroupVersionResource{Group: "commonnetworkpolicies.bells17.io", Version: "v1alpha1", Resource: "commonnetworkpolicies"}

var commonnetworkpoliciesKind = schema.GroupVersionKind{Group: "commonnetworkpolicies.bells17.io", Version: "v1alpha1", Kind: "CommonNetworkPolicy"}

// Get takes name of the commonNetworkPolicy, and returns the corresponding commonNetworkPolicy object, and an error if there is any.
func (c *FakeCommonNetworkPolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.CommonNetworkPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(commonnetworkpoliciesResource, name), &v1alpha1.CommonNetworkPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CommonNetworkPolicy), err
}

// List takes label and field selectors, and returns the list of CommonNetworkPolicies that match those selectors.
func (c *FakeCommonNetworkPolicies) List(opts v1.ListOptions) (result *v1alpha1.CommonNetworkPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(commonnetworkpoliciesResource, commonnetworkpoliciesKind, opts), &v1alpha1.CommonNetworkPolicyList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CommonNetworkPolicyList{ListMeta: obj.(*v1alpha1.CommonNetworkPolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.CommonNetworkPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested commonNetworkPolicies.
func (c *FakeCommonNetworkPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(commonnetworkpoliciesResource, opts))
}

// Create takes the representation of a commonNetworkPolicy and creates it.  Returns the server's representation of the commonNetworkPolicy, and an error, if there is any.
func (c *FakeCommonNetworkPolicies) Create(commonNetworkPolicy *v1alpha1.CommonNetworkPolicy) (result *v1alpha1.CommonNetworkPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(commonnetworkpoliciesResource, commonNetworkPolicy), &v1alpha1.CommonNetworkPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CommonNetworkPolicy), err
}

// Update takes the representation of a commonNetworkPolicy and updates it. Returns the server's representation of the commonNetworkPolicy, and an error, if there is any.
func (c *FakeCommonNetworkPolicies) Update(commonNetworkPolicy *v1alpha1.CommonNetworkPolicy) (result *v1alpha1.CommonNetworkPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(commonnetworkpoliciesResource, commonNetworkPolicy), &v1alpha1.CommonNetworkPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CommonNetworkPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCommonNetworkPolicies) UpdateStatus(commonNetworkPolicy *v1alpha1.CommonNetworkPolicy) (*v1alpha1.CommonNetworkPolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(commonnetworkpoliciesResource, "status", commonNetworkPolicy), &v1alpha1.CommonNetworkPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CommonNetworkPolicy), err
}

// Delete takes name of the commonNetworkPolicy and deletes it. Returns an error if one occurs.
func (c *FakeCommonNetworkPolicies) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(commonnetworkpoliciesResource, name), &v1alpha1.CommonNetworkPolicy{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCommonNetworkPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(commonnetworkpoliciesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.CommonNetworkPolicyList{})
	return err
}

// Patch applies the patch and returns the patched commonNetworkPolicy.
func (c *FakeCommonNetworkPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CommonNetworkPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(commonnetworkpoliciesResource, name, data, subresources...), &v1alpha1.CommonNetworkPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CommonNetworkPolicy), err
}
