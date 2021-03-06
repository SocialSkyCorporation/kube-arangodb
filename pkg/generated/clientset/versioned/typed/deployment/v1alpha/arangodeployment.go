//
// DISCLAIMER
//
// Copyright 2018 ArangoDB GmbH, Cologne, Germany
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Copyright holder is ArangoDB GmbH, Cologne, Germany
//

// Code generated by client-gen. DO NOT EDIT.

package v1alpha

import (
	v1alpha "github.com/arangodb/kube-arangodb/pkg/apis/deployment/v1alpha"
	scheme "github.com/arangodb/kube-arangodb/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ArangoDeploymentsGetter has a method to return a ArangoDeploymentInterface.
// A group's client should implement this interface.
type ArangoDeploymentsGetter interface {
	ArangoDeployments(namespace string) ArangoDeploymentInterface
}

// ArangoDeploymentInterface has methods to work with ArangoDeployment resources.
type ArangoDeploymentInterface interface {
	Create(*v1alpha.ArangoDeployment) (*v1alpha.ArangoDeployment, error)
	Update(*v1alpha.ArangoDeployment) (*v1alpha.ArangoDeployment, error)
	UpdateStatus(*v1alpha.ArangoDeployment) (*v1alpha.ArangoDeployment, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha.ArangoDeployment, error)
	List(opts v1.ListOptions) (*v1alpha.ArangoDeploymentList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha.ArangoDeployment, err error)
	ArangoDeploymentExpansion
}

// arangoDeployments implements ArangoDeploymentInterface
type arangoDeployments struct {
	client rest.Interface
	ns     string
}

// newArangoDeployments returns a ArangoDeployments
func newArangoDeployments(c *DatabaseV1alphaClient, namespace string) *arangoDeployments {
	return &arangoDeployments{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the arangoDeployment, and returns the corresponding arangoDeployment object, and an error if there is any.
func (c *arangoDeployments) Get(name string, options v1.GetOptions) (result *v1alpha.ArangoDeployment, err error) {
	result = &v1alpha.ArangoDeployment{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("arangodeployments").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ArangoDeployments that match those selectors.
func (c *arangoDeployments) List(opts v1.ListOptions) (result *v1alpha.ArangoDeploymentList, err error) {
	result = &v1alpha.ArangoDeploymentList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("arangodeployments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested arangoDeployments.
func (c *arangoDeployments) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("arangodeployments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a arangoDeployment and creates it.  Returns the server's representation of the arangoDeployment, and an error, if there is any.
func (c *arangoDeployments) Create(arangoDeployment *v1alpha.ArangoDeployment) (result *v1alpha.ArangoDeployment, err error) {
	result = &v1alpha.ArangoDeployment{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("arangodeployments").
		Body(arangoDeployment).
		Do().
		Into(result)
	return
}

// Update takes the representation of a arangoDeployment and updates it. Returns the server's representation of the arangoDeployment, and an error, if there is any.
func (c *arangoDeployments) Update(arangoDeployment *v1alpha.ArangoDeployment) (result *v1alpha.ArangoDeployment, err error) {
	result = &v1alpha.ArangoDeployment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("arangodeployments").
		Name(arangoDeployment.Name).
		Body(arangoDeployment).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *arangoDeployments) UpdateStatus(arangoDeployment *v1alpha.ArangoDeployment) (result *v1alpha.ArangoDeployment, err error) {
	result = &v1alpha.ArangoDeployment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("arangodeployments").
		Name(arangoDeployment.Name).
		SubResource("status").
		Body(arangoDeployment).
		Do().
		Into(result)
	return
}

// Delete takes name of the arangoDeployment and deletes it. Returns an error if one occurs.
func (c *arangoDeployments) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("arangodeployments").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *arangoDeployments) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("arangodeployments").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched arangoDeployment.
func (c *arangoDeployments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha.ArangoDeployment, err error) {
	result = &v1alpha.ArangoDeployment{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("arangodeployments").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
