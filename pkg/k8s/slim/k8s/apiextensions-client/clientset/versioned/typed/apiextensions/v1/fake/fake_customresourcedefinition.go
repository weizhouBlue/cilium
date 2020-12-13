// Copyright 2017-2020 Authors of Cilium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	apiextensionsv1 "github.com/cilium/cilium/pkg/k8s/slim/k8s/apis/apiextensions/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCustomResourceDefinitions implements CustomResourceDefinitionInterface
type FakeCustomResourceDefinitions struct {
	Fake *FakeApiextensionsV1
}

var customresourcedefinitionsResource = schema.GroupVersionResource{Group: "apiextensions.k8s.io", Version: "v1", Resource: "customresourcedefinitions"}

var customresourcedefinitionsKind = schema.GroupVersionKind{Group: "apiextensions.k8s.io", Version: "v1", Kind: "CustomResourceDefinition"}

// Get takes name of the customResourceDefinition, and returns the corresponding customResourceDefinition object, and an error if there is any.
func (c *FakeCustomResourceDefinitions) Get(ctx context.Context, name string, options v1.GetOptions) (result *apiextensionsv1.CustomResourceDefinition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(customresourcedefinitionsResource, name), &apiextensionsv1.CustomResourceDefinition{})
	if obj == nil {
		return nil, err
	}
	return obj.(*apiextensionsv1.CustomResourceDefinition), err
}

// List takes label and field selectors, and returns the list of CustomResourceDefinitions that match those selectors.
func (c *FakeCustomResourceDefinitions) List(ctx context.Context, opts v1.ListOptions) (result *apiextensionsv1.CustomResourceDefinitionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(customresourcedefinitionsResource, customresourcedefinitionsKind, opts), &apiextensionsv1.CustomResourceDefinitionList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &apiextensionsv1.CustomResourceDefinitionList{ListMeta: obj.(*apiextensionsv1.CustomResourceDefinitionList).ListMeta}
	for _, item := range obj.(*apiextensionsv1.CustomResourceDefinitionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested customResourceDefinitions.
func (c *FakeCustomResourceDefinitions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(customresourcedefinitionsResource, opts))
}

// Create takes the representation of a customResourceDefinition and creates it.  Returns the server's representation of the customResourceDefinition, and an error, if there is any.
func (c *FakeCustomResourceDefinitions) Create(ctx context.Context, customResourceDefinition *apiextensionsv1.CustomResourceDefinition, opts v1.CreateOptions) (result *apiextensionsv1.CustomResourceDefinition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(customresourcedefinitionsResource, customResourceDefinition), &apiextensionsv1.CustomResourceDefinition{})
	if obj == nil {
		return nil, err
	}
	return obj.(*apiextensionsv1.CustomResourceDefinition), err
}

// Update takes the representation of a customResourceDefinition and updates it. Returns the server's representation of the customResourceDefinition, and an error, if there is any.
func (c *FakeCustomResourceDefinitions) Update(ctx context.Context, customResourceDefinition *apiextensionsv1.CustomResourceDefinition, opts v1.UpdateOptions) (result *apiextensionsv1.CustomResourceDefinition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(customresourcedefinitionsResource, customResourceDefinition), &apiextensionsv1.CustomResourceDefinition{})
	if obj == nil {
		return nil, err
	}
	return obj.(*apiextensionsv1.CustomResourceDefinition), err
}

// Delete takes name of the customResourceDefinition and deletes it. Returns an error if one occurs.
func (c *FakeCustomResourceDefinitions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(customresourcedefinitionsResource, name), &apiextensionsv1.CustomResourceDefinition{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCustomResourceDefinitions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(customresourcedefinitionsResource, listOpts)

	_, err := c.Fake.Invokes(action, &apiextensionsv1.CustomResourceDefinitionList{})
	return err
}

// Patch applies the patch and returns the patched customResourceDefinition.
func (c *FakeCustomResourceDefinitions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *apiextensionsv1.CustomResourceDefinition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(customresourcedefinitionsResource, name, pt, data, subresources...), &apiextensionsv1.CustomResourceDefinition{})
	if obj == nil {
		return nil, err
	}
	return obj.(*apiextensionsv1.CustomResourceDefinition), err
}
