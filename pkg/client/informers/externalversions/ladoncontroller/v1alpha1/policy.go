/*
Copyright 2018 The Kubernetes Authors.

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

	ladoncontroller_v1alpha1 "github.com/kminehart/ladon-resource-manager/pkg/apis/ladoncontroller/v1alpha1"
	versioned "github.com/kminehart/ladon-resource-manager/pkg/client/clientset/versioned"
	internalinterfaces "github.com/kminehart/ladon-resource-manager/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/kminehart/ladon-resource-manager/pkg/client/listers/ladoncontroller/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// LadonPolicyInformer provides access to a shared informer and lister for
// Policies.
type LadonPolicyInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.LadonPolicyLister
}

type policyInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewLadonPolicyInformer constructs a new informer for LadonPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewLadonPolicyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredLadonPolicyInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredLadonPolicyInformer constructs a new informer for LadonPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredLadonPolicyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.LadoncontrollerV1alpha1().Policies(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.LadoncontrollerV1alpha1().Policies(namespace).Watch(options)
			},
		},
		&ladoncontroller_v1alpha1.LadonPolicy{},
		resyncPeriod,
		indexers,
	)
}

func (f *policyInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredLadonPolicyInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *policyInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&ladoncontroller_v1alpha1.LadonPolicy{}, f.defaultInformer)
}

func (f *policyInformer) Lister() v1alpha1.LadonPolicyLister {
	return v1alpha1.NewLadonPolicyLister(f.Informer().GetIndexer())
}
