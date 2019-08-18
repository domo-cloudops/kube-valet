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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/domoinc/kube-valet/pkg/apis/assignments/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// PodAssignmentRuleLister helps list PodAssignmentRules.
type PodAssignmentRuleLister interface {
	// List lists all PodAssignmentRules in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.PodAssignmentRule, err error)
	// PodAssignmentRules returns an object that can list and get PodAssignmentRules.
	PodAssignmentRules(namespace string) PodAssignmentRuleNamespaceLister
	PodAssignmentRuleListerExpansion
}

// podAssignmentRuleLister implements the PodAssignmentRuleLister interface.
type podAssignmentRuleLister struct {
	indexer cache.Indexer
}

// NewPodAssignmentRuleLister returns a new PodAssignmentRuleLister.
func NewPodAssignmentRuleLister(indexer cache.Indexer) PodAssignmentRuleLister {
	return &podAssignmentRuleLister{indexer: indexer}
}

// List lists all PodAssignmentRules in the indexer.
func (s *podAssignmentRuleLister) List(selector labels.Selector) (ret []*v1alpha1.PodAssignmentRule, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PodAssignmentRule))
	})
	return ret, err
}

// PodAssignmentRules returns an object that can list and get PodAssignmentRules.
func (s *podAssignmentRuleLister) PodAssignmentRules(namespace string) PodAssignmentRuleNamespaceLister {
	return podAssignmentRuleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PodAssignmentRuleNamespaceLister helps list and get PodAssignmentRules.
type PodAssignmentRuleNamespaceLister interface {
	// List lists all PodAssignmentRules in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.PodAssignmentRule, err error)
	// Get retrieves the PodAssignmentRule from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.PodAssignmentRule, error)
	PodAssignmentRuleNamespaceListerExpansion
}

// podAssignmentRuleNamespaceLister implements the PodAssignmentRuleNamespaceLister
// interface.
type podAssignmentRuleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all PodAssignmentRules in the indexer for a given namespace.
func (s podAssignmentRuleNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.PodAssignmentRule, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PodAssignmentRule))
	})
	return ret, err
}

// Get retrieves the PodAssignmentRule from the indexer for a given namespace and name.
func (s podAssignmentRuleNamespaceLister) Get(name string) (*v1alpha1.PodAssignmentRule, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("podassignmentrule"), name)
	}
	return obj.(*v1alpha1.PodAssignmentRule), nil
}
