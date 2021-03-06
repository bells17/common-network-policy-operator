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

package v1alpha1

import (
	v1alpha1 "github.com/bells17/common-network-policy-operator/pkg/apis/commonnetworkpolicies/v1alpha1"
	scheme "github.com/bells17/common-network-policy-operator/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CommonNetworkPoliciesGetter has a method to return a CommonNetworkPolicyInterface.
// A group's client should implement this interface.
type CommonNetworkPoliciesGetter interface {
	CommonNetworkPolicies() CommonNetworkPolicyInterface
}

// CommonNetworkPolicyInterface has methods to work with CommonNetworkPolicy resources.
type CommonNetworkPolicyInterface interface {
	Create(*v1alpha1.CommonNetworkPolicy) (*v1alpha1.CommonNetworkPolicy, error)
	Update(*v1alpha1.CommonNetworkPolicy) (*v1alpha1.CommonNetworkPolicy, error)
	UpdateStatus(*v1alpha1.CommonNetworkPolicy) (*v1alpha1.CommonNetworkPolicy, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.CommonNetworkPolicy, error)
	List(opts v1.ListOptions) (*v1alpha1.CommonNetworkPolicyList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CommonNetworkPolicy, err error)
	CommonNetworkPolicyExpansion
}

// commonNetworkPolicies implements CommonNetworkPolicyInterface
type commonNetworkPolicies struct {
	client rest.Interface
}

// newCommonNetworkPolicies returns a CommonNetworkPolicies
func newCommonNetworkPolicies(c *CommonnetworkpoliciesV1alpha1Client) *commonNetworkPolicies {
	return &commonNetworkPolicies{
		client: c.RESTClient(),
	}
}

// Get takes name of the commonNetworkPolicy, and returns the corresponding commonNetworkPolicy object, and an error if there is any.
func (c *commonNetworkPolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.CommonNetworkPolicy, err error) {
	result = &v1alpha1.CommonNetworkPolicy{}
	err = c.client.Get().
		Resource("commonnetworkpolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CommonNetworkPolicies that match those selectors.
func (c *commonNetworkPolicies) List(opts v1.ListOptions) (result *v1alpha1.CommonNetworkPolicyList, err error) {
	result = &v1alpha1.CommonNetworkPolicyList{}
	err = c.client.Get().
		Resource("commonnetworkpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested commonNetworkPolicies.
func (c *commonNetworkPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Resource("commonnetworkpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a commonNetworkPolicy and creates it.  Returns the server's representation of the commonNetworkPolicy, and an error, if there is any.
func (c *commonNetworkPolicies) Create(commonNetworkPolicy *v1alpha1.CommonNetworkPolicy) (result *v1alpha1.CommonNetworkPolicy, err error) {
	result = &v1alpha1.CommonNetworkPolicy{}
	err = c.client.Post().
		Resource("commonnetworkpolicies").
		Body(commonNetworkPolicy).
		Do().
		Into(result)
	return
}

// Update takes the representation of a commonNetworkPolicy and updates it. Returns the server's representation of the commonNetworkPolicy, and an error, if there is any.
func (c *commonNetworkPolicies) Update(commonNetworkPolicy *v1alpha1.CommonNetworkPolicy) (result *v1alpha1.CommonNetworkPolicy, err error) {
	result = &v1alpha1.CommonNetworkPolicy{}
	err = c.client.Put().
		Resource("commonnetworkpolicies").
		Name(commonNetworkPolicy.Name).
		Body(commonNetworkPolicy).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *commonNetworkPolicies) UpdateStatus(commonNetworkPolicy *v1alpha1.CommonNetworkPolicy) (result *v1alpha1.CommonNetworkPolicy, err error) {
	result = &v1alpha1.CommonNetworkPolicy{}
	err = c.client.Put().
		Resource("commonnetworkpolicies").
		Name(commonNetworkPolicy.Name).
		SubResource("status").
		Body(commonNetworkPolicy).
		Do().
		Into(result)
	return
}

// Delete takes name of the commonNetworkPolicy and deletes it. Returns an error if one occurs.
func (c *commonNetworkPolicies) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("commonnetworkpolicies").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *commonNetworkPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Resource("commonnetworkpolicies").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched commonNetworkPolicy.
func (c *commonNetworkPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CommonNetworkPolicy, err error) {
	result = &v1alpha1.CommonNetworkPolicy{}
	err = c.client.Patch(pt).
		Resource("commonnetworkpolicies").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
