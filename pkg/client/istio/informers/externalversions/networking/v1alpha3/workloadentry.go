/*
Copyright 2020 The Knative Authors

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

package v1alpha3

import (
	"context"
	time "time"

	networkingv1alpha3 "istio.io/client-go/pkg/apis/networking/v1alpha3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	versioned "knative.dev/networking/pkg/client/istio/clientset/versioned"
	internalinterfaces "knative.dev/networking/pkg/client/istio/informers/externalversions/internalinterfaces"
	v1alpha3 "knative.dev/networking/pkg/client/istio/listers/networking/v1alpha3"
)

// WorkloadEntryInformer provides access to a shared informer and lister for
// WorkloadEntries.
type WorkloadEntryInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha3.WorkloadEntryLister
}

type workloadEntryInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewWorkloadEntryInformer constructs a new informer for WorkloadEntry type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewWorkloadEntryInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredWorkloadEntryInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredWorkloadEntryInformer constructs a new informer for WorkloadEntry type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredWorkloadEntryInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NetworkingV1alpha3().WorkloadEntries(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NetworkingV1alpha3().WorkloadEntries(namespace).Watch(context.TODO(), options)
			},
		},
		&networkingv1alpha3.WorkloadEntry{},
		resyncPeriod,
		indexers,
	)
}

func (f *workloadEntryInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredWorkloadEntryInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *workloadEntryInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&networkingv1alpha3.WorkloadEntry{}, f.defaultInformer)
}

func (f *workloadEntryInformer) Lister() v1alpha3.WorkloadEntryLister {
	return v1alpha3.NewWorkloadEntryLister(f.Informer().GetIndexer())
}
