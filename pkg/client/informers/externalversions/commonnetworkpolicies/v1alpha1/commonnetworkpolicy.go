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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	commonnetworkpoliciesv1alpha1 "github.com/bells17/common-network-policy-operator/pkg/apis/commonnetworkpolicies/v1alpha1"
	versioned "github.com/bells17/common-network-policy-operator/pkg/client/clientset/versioned"
	internalinterfaces "github.com/bells17/common-network-policy-operator/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/bells17/common-network-policy-operator/pkg/client/listers/commonnetworkpolicies/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// CommonNetworkPolicyInformer provides access to a shared informer and lister for
// CommonNetworkPolicies.
type CommonNetworkPolicyInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.CommonNetworkPolicyLister
}

type commonNetworkPolicyInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewCommonNetworkPolicyInformer constructs a new informer for CommonNetworkPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCommonNetworkPolicyInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCommonNetworkPolicyInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredCommonNetworkPolicyInformer constructs a new informer for CommonNetworkPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCommonNetworkPolicyInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CommonnetworkpoliciesV1alpha1().CommonNetworkPolicies().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CommonnetworkpoliciesV1alpha1().CommonNetworkPolicies().Watch(options)
			},
		},
		&commonnetworkpoliciesv1alpha1.CommonNetworkPolicy{},
		resyncPeriod,
		indexers,
	)
}

func (f *commonNetworkPolicyInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCommonNetworkPolicyInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *commonNetworkPolicyInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&commonnetworkpoliciesv1alpha1.CommonNetworkPolicy{}, f.defaultInformer)
}

func (f *commonNetworkPolicyInformer) Lister() v1alpha1.CommonNetworkPolicyLister {
	return v1alpha1.NewCommonNetworkPolicyLister(f.Informer().GetIndexer())
}
