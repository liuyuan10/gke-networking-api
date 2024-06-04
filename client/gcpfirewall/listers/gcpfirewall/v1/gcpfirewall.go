/*
Copyright 2024 Google LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/GoogleCloudPlatform/gke-networking-api/apis/gcpfirewall/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// GCPFirewallLister helps list GCPFirewalls.
// All objects returned here must be treated as read-only.
type GCPFirewallLister interface {
	// List lists all GCPFirewalls in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.GCPFirewall, err error)
	// Get retrieves the GCPFirewall from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.GCPFirewall, error)
	GCPFirewallListerExpansion
}

// gCPFirewallLister implements the GCPFirewallLister interface.
type gCPFirewallLister struct {
	indexer cache.Indexer
}

// NewGCPFirewallLister returns a new GCPFirewallLister.
func NewGCPFirewallLister(indexer cache.Indexer) GCPFirewallLister {
	return &gCPFirewallLister{indexer: indexer}
}

// List lists all GCPFirewalls in the indexer.
func (s *gCPFirewallLister) List(selector labels.Selector) (ret []*v1.GCPFirewall, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.GCPFirewall))
	})
	return ret, err
}

// Get retrieves the GCPFirewall from the index for a given name.
func (s *gCPFirewallLister) Get(name string) (*v1.GCPFirewall, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("gcpfirewall"), name)
	}
	return obj.(*v1.GCPFirewall), nil
}
