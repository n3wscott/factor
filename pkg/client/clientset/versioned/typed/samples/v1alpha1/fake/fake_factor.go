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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "knative.dev/sample-controller/pkg/apis/samples/v1alpha1"
)

// FakeFactors implements FactorInterface
type FakeFactors struct {
	Fake *FakeSamplesV1alpha1
	ns   string
}

var factorsResource = schema.GroupVersionResource{Group: "samples.knative.dev", Version: "v1alpha1", Resource: "factors"}

var factorsKind = schema.GroupVersionKind{Group: "samples.knative.dev", Version: "v1alpha1", Kind: "Factor"}

// Get takes name of the factor, and returns the corresponding factor object, and an error if there is any.
func (c *FakeFactors) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Factor, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(factorsResource, c.ns, name), &v1alpha1.Factor{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Factor), err
}

// List takes label and field selectors, and returns the list of Factors that match those selectors.
func (c *FakeFactors) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.FactorList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(factorsResource, factorsKind, c.ns, opts), &v1alpha1.FactorList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.FactorList{ListMeta: obj.(*v1alpha1.FactorList).ListMeta}
	for _, item := range obj.(*v1alpha1.FactorList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested factors.
func (c *FakeFactors) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(factorsResource, c.ns, opts))

}

// Create takes the representation of a factor and creates it.  Returns the server's representation of the factor, and an error, if there is any.
func (c *FakeFactors) Create(ctx context.Context, factor *v1alpha1.Factor, opts v1.CreateOptions) (result *v1alpha1.Factor, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(factorsResource, c.ns, factor), &v1alpha1.Factor{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Factor), err
}

// Update takes the representation of a factor and updates it. Returns the server's representation of the factor, and an error, if there is any.
func (c *FakeFactors) Update(ctx context.Context, factor *v1alpha1.Factor, opts v1.UpdateOptions) (result *v1alpha1.Factor, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(factorsResource, c.ns, factor), &v1alpha1.Factor{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Factor), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFactors) UpdateStatus(ctx context.Context, factor *v1alpha1.Factor, opts v1.UpdateOptions) (*v1alpha1.Factor, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(factorsResource, "status", c.ns, factor), &v1alpha1.Factor{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Factor), err
}

// Delete takes name of the factor and deletes it. Returns an error if one occurs.
func (c *FakeFactors) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(factorsResource, c.ns, name), &v1alpha1.Factor{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFactors) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(factorsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.FactorList{})
	return err
}

// Patch applies the patch and returns the patched factor.
func (c *FakeFactors) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Factor, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(factorsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Factor{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Factor), err
}
