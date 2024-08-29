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

package v1

import (
	"context"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	corev1 "k8s.io/client-go/applyconfigurations/core/v1"
	gentype "k8s.io/client-go/gentype"
	scheme "k8s.io/client-go/kubernetes/scheme"
)

// PersistentVolumesGetter has a method to return a PersistentVolumeInterface.
// A group's client should implement this interface.
type PersistentVolumesGetter interface {
	PersistentVolumes() PersistentVolumeInterface
}

// PersistentVolumeInterface has methods to work with PersistentVolume resources.
type PersistentVolumeInterface interface {
	Create(ctx context.Context, persistentVolume *v1.PersistentVolume, opts metav1.CreateOptions) (*v1.PersistentVolume, error)
	Update(ctx context.Context, persistentVolume *v1.PersistentVolume, opts metav1.UpdateOptions) (*v1.PersistentVolume, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, persistentVolume *v1.PersistentVolume, opts metav1.UpdateOptions) (*v1.PersistentVolume, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.PersistentVolume, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.PersistentVolumeList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.PersistentVolume, err error)
	Apply(ctx context.Context, persistentVolume *corev1.PersistentVolumeApplyConfiguration, opts metav1.ApplyOptions) (result *v1.PersistentVolume, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, persistentVolume *corev1.PersistentVolumeApplyConfiguration, opts metav1.ApplyOptions) (result *v1.PersistentVolume, err error)
	PersistentVolumeExpansion
}

// persistentVolumes implements PersistentVolumeInterface
type persistentVolumes struct {
	*gentype.ClientWithListAndApply[*v1.PersistentVolume, *v1.PersistentVolumeList, *corev1.PersistentVolumeApplyConfiguration]
}

// newPersistentVolumes returns a PersistentVolumes
func newPersistentVolumes(c *CoreV1Client) *persistentVolumes {
	return &persistentVolumes{
		gentype.NewClientWithListAndApply[*v1.PersistentVolume, *v1.PersistentVolumeList, *corev1.PersistentVolumeApplyConfiguration](
			"persistentvolumes",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *v1.PersistentVolume { return &v1.PersistentVolume{} },
			func() *v1.PersistentVolumeList { return &v1.PersistentVolumeList{} }),
	}
}
