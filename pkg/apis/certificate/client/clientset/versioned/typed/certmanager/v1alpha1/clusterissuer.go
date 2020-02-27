/*
Copyright 2018 Jetstack Ltd.

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
package v1alpha1

import (
	v1alpha1 "github.com/choerodon/choerodon-cluster-agent/pkg/apis/certificate/apis/certmanager/v1alpha1"
	scheme "github.com/choerodon/choerodon-cluster-agent/pkg/apis/certificate/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ClusterIssuersGetter has a method to return a ClusterIssuerInterface.
// A group's client should implement this interface.
type ClusterIssuersGetter interface {
	ClusterIssuers() ClusterIssuerInterface
}

// ClusterIssuerInterface has methods to work with ClusterIssuer resources.
type ClusterIssuerInterface interface {
	Create(*v1alpha1.ClusterIssuer) (*v1alpha1.ClusterIssuer, error)
	Update(*v1alpha1.ClusterIssuer) (*v1alpha1.ClusterIssuer, error)
	UpdateStatus(*v1alpha1.ClusterIssuer) (*v1alpha1.ClusterIssuer, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ClusterIssuer, error)
	List(opts v1.ListOptions) (*v1alpha1.ClusterIssuerList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ClusterIssuer, err error)
	ClusterIssuerExpansion
}

// clusterIssuers implements ClusterIssuerInterface
type clusterIssuers struct {
	client rest.Interface
}

// newClusterIssuers returns a ClusterIssuers
func newClusterIssuers(c *CertmanagerV1alpha1Client) *clusterIssuers {
	return &clusterIssuers{
		client: c.RESTClient(),
	}
}

// Get takes name of the clusterIssuer, and returns the corresponding clusterIssuer object, and an error if there is any.
func (c *clusterIssuers) Get(name string, options v1.GetOptions) (result *v1alpha1.ClusterIssuer, err error) {
	result = &v1alpha1.ClusterIssuer{}
	err = c.client.Get().
		Resource("clusterissuers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ClusterIssuers that match those selectors.
func (c *clusterIssuers) List(opts v1.ListOptions) (result *v1alpha1.ClusterIssuerList, err error) {
	result = &v1alpha1.ClusterIssuerList{}
	err = c.client.Get().
		Resource("clusterissuers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested clusterIssuers.
func (c *clusterIssuers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Resource("clusterissuers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a clusterIssuer and creates it.  Returns the server's representation of the clusterIssuer, and an error, if there is any.
func (c *clusterIssuers) Create(clusterIssuer *v1alpha1.ClusterIssuer) (result *v1alpha1.ClusterIssuer, err error) {
	result = &v1alpha1.ClusterIssuer{}
	err = c.client.Post().
		Resource("clusterissuers").
		Body(clusterIssuer).
		Do().
		Into(result)
	return
}

// Update takes the representation of a clusterIssuer and updates it. Returns the server's representation of the clusterIssuer, and an error, if there is any.
func (c *clusterIssuers) Update(clusterIssuer *v1alpha1.ClusterIssuer) (result *v1alpha1.ClusterIssuer, err error) {
	result = &v1alpha1.ClusterIssuer{}
	err = c.client.Put().
		Resource("clusterissuers").
		Name(clusterIssuer.Name).
		Body(clusterIssuer).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *clusterIssuers) UpdateStatus(clusterIssuer *v1alpha1.ClusterIssuer) (result *v1alpha1.ClusterIssuer, err error) {
	result = &v1alpha1.ClusterIssuer{}
	err = c.client.Put().
		Resource("clusterissuers").
		Name(clusterIssuer.Name).
		SubResource("status").
		Body(clusterIssuer).
		Do().
		Into(result)
	return
}

// Delete takes name of the clusterIssuer and deletes it. Returns an error if one occurs.
func (c *clusterIssuers) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("clusterissuers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *clusterIssuers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Resource("clusterissuers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched clusterIssuer.
func (c *clusterIssuers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ClusterIssuer, err error) {
	result = &v1alpha1.ClusterIssuer{}
	err = c.client.Patch(pt).
		Resource("clusterissuers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}