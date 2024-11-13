// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1 "github.com/openshift/api/security/v1"
	securityv1 "github.com/openshift/client-go/security/applyconfigurations/security/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRangeAllocations implements RangeAllocationInterface
type FakeRangeAllocations struct {
	Fake *FakeSecurityV1
}

var rangeallocationsResource = v1.SchemeGroupVersion.WithResource("rangeallocations")

var rangeallocationsKind = v1.SchemeGroupVersion.WithKind("RangeAllocation")

// Get takes name of the rangeAllocation, and returns the corresponding rangeAllocation object, and an error if there is any.
func (c *FakeRangeAllocations) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.RangeAllocation, err error) {
	emptyResult := &v1.RangeAllocation{}
	obj, err := c.Fake.
		Invokes(testing.NewRootGetActionWithOptions(rangeallocationsResource, name, options), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.RangeAllocation), err
}

// List takes label and field selectors, and returns the list of RangeAllocations that match those selectors.
func (c *FakeRangeAllocations) List(ctx context.Context, opts metav1.ListOptions) (result *v1.RangeAllocationList, err error) {
	emptyResult := &v1.RangeAllocationList{}
	obj, err := c.Fake.
		Invokes(testing.NewRootListActionWithOptions(rangeallocationsResource, rangeallocationsKind, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.RangeAllocationList{ListMeta: obj.(*v1.RangeAllocationList).ListMeta}
	for _, item := range obj.(*v1.RangeAllocationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested rangeAllocations.
func (c *FakeRangeAllocations) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchActionWithOptions(rangeallocationsResource, opts))
}

// Create takes the representation of a rangeAllocation and creates it.  Returns the server's representation of the rangeAllocation, and an error, if there is any.
func (c *FakeRangeAllocations) Create(ctx context.Context, rangeAllocation *v1.RangeAllocation, opts metav1.CreateOptions) (result *v1.RangeAllocation, err error) {
	emptyResult := &v1.RangeAllocation{}
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateActionWithOptions(rangeallocationsResource, rangeAllocation, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.RangeAllocation), err
}

// Update takes the representation of a rangeAllocation and updates it. Returns the server's representation of the rangeAllocation, and an error, if there is any.
func (c *FakeRangeAllocations) Update(ctx context.Context, rangeAllocation *v1.RangeAllocation, opts metav1.UpdateOptions) (result *v1.RangeAllocation, err error) {
	emptyResult := &v1.RangeAllocation{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateActionWithOptions(rangeallocationsResource, rangeAllocation, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.RangeAllocation), err
}

// Delete takes name of the rangeAllocation and deletes it. Returns an error if one occurs.
func (c *FakeRangeAllocations) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(rangeallocationsResource, name, opts), &v1.RangeAllocation{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRangeAllocations) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewRootDeleteCollectionActionWithOptions(rangeallocationsResource, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1.RangeAllocationList{})
	return err
}

// Patch applies the patch and returns the patched rangeAllocation.
func (c *FakeRangeAllocations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.RangeAllocation, err error) {
	emptyResult := &v1.RangeAllocation{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(rangeallocationsResource, name, pt, data, opts, subresources...), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.RangeAllocation), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied rangeAllocation.
func (c *FakeRangeAllocations) Apply(ctx context.Context, rangeAllocation *securityv1.RangeAllocationApplyConfiguration, opts metav1.ApplyOptions) (result *v1.RangeAllocation, err error) {
	if rangeAllocation == nil {
		return nil, fmt.Errorf("rangeAllocation provided to Apply must not be nil")
	}
	data, err := json.Marshal(rangeAllocation)
	if err != nil {
		return nil, err
	}
	name := rangeAllocation.Name
	if name == nil {
		return nil, fmt.Errorf("rangeAllocation.Name must be provided to Apply")
	}
	emptyResult := &v1.RangeAllocation{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(rangeallocationsResource, *name, types.ApplyPatchType, data, opts.ToPatchOptions()), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.RangeAllocation), err
}
