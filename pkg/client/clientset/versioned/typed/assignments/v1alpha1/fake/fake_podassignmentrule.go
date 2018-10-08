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

package fake

import (
	v1alpha1 "github.com/domoinc/kube-valet/pkg/apis/assignments/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePodAssignmentRules implements PodAssignmentRuleInterface
type FakePodAssignmentRules struct {
	Fake *FakeAssignmentsV1alpha1
	ns   string
}

var podassignmentrulesResource = schema.GroupVersionResource{Group: "assignments.kube-valet.io", Version: "v1alpha1", Resource: "podassignmentrules"}

var podassignmentrulesKind = schema.GroupVersionKind{Group: "assignments.kube-valet.io", Version: "v1alpha1", Kind: "PodAssignmentRule"}

// Get takes name of the podAssignmentRule, and returns the corresponding podAssignmentRule object, and an error if there is any.
func (c *FakePodAssignmentRules) Get(name string, options v1.GetOptions) (result *v1alpha1.PodAssignmentRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(podassignmentrulesResource, c.ns, name), &v1alpha1.PodAssignmentRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PodAssignmentRule), err
}

// List takes label and field selectors, and returns the list of PodAssignmentRules that match those selectors.
func (c *FakePodAssignmentRules) List(opts v1.ListOptions) (result *v1alpha1.PodAssignmentRuleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(podassignmentrulesResource, podassignmentrulesKind, c.ns, opts), &v1alpha1.PodAssignmentRuleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PodAssignmentRuleList{}
	for _, item := range obj.(*v1alpha1.PodAssignmentRuleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested podAssignmentRules.
func (c *FakePodAssignmentRules) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(podassignmentrulesResource, c.ns, opts))

}

// Create takes the representation of a podAssignmentRule and creates it.  Returns the server's representation of the podAssignmentRule, and an error, if there is any.
func (c *FakePodAssignmentRules) Create(podAssignmentRule *v1alpha1.PodAssignmentRule) (result *v1alpha1.PodAssignmentRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(podassignmentrulesResource, c.ns, podAssignmentRule), &v1alpha1.PodAssignmentRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PodAssignmentRule), err
}

// Update takes the representation of a podAssignmentRule and updates it. Returns the server's representation of the podAssignmentRule, and an error, if there is any.
func (c *FakePodAssignmentRules) Update(podAssignmentRule *v1alpha1.PodAssignmentRule) (result *v1alpha1.PodAssignmentRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(podassignmentrulesResource, c.ns, podAssignmentRule), &v1alpha1.PodAssignmentRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PodAssignmentRule), err
}

// Delete takes name of the podAssignmentRule and deletes it. Returns an error if one occurs.
func (c *FakePodAssignmentRules) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(podassignmentrulesResource, c.ns, name), &v1alpha1.PodAssignmentRule{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePodAssignmentRules) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(podassignmentrulesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.PodAssignmentRuleList{})
	return err
}

// Patch applies the patch and returns the patched podAssignmentRule.
func (c *FakePodAssignmentRules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PodAssignmentRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(podassignmentrulesResource, c.ns, name, data, subresources...), &v1alpha1.PodAssignmentRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PodAssignmentRule), err
}
