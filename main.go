package main

import (
	"context"
	"fmt"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	rest "k8s.io/client-go/rest"
	v1alpha1 "sigs.k8s.io/multi-tenancy/poc/tenant-controller/pkg/apis/tenants/v1alpha1"
	scheme "sigs.k8s.io/multi-tenancy/poc/tenant-controller/pkg/clients/tenants/clientset/v1alpha1/scheme"
	tenclient "sigs.k8s.io/multi-tenancy/poc/tenant-controller/pkg/clients/tenants/clientset/v1alpha1/typed/tenants/v1alpha1/"
)

var demo rest.Interface

type tenants struct {
	client rest.Interface
}

func (c *tenants) Get(name string, options v1.GetOptions) (result *v1alpha1.Tenant, err error) {
	result = &v1alpha1.Tenant{}
	err = c.client.Get().
		Resource("tenants").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(context.TODO()).
		Into(result)
	return
}

func (c *tenants) Create(string name, tenant *v1alpha1.Tenant) (result *v1alpha1.Tenant, err error) {
	result = &v1alpha1.Tenant{}
	err = c.client.Post().
		Resource("tenants").
		Name(name).
		Body(tenant).
		Do(context.TODO()).
		Into(result)
	return
}

func main() {

	c := tenclient.New(demo)
	demoClient := c.RESTCLIENT()

	t := &tenants{
		client: demoClient,
	}

	tenant * v1alpha1.Tenant

	result, err := t.Create("do-something", tenant)
	fmt.Println(result, err.Error())

	tenants, err := t.Get("do-something", v1.GetOptions{})
	fmt.Println(tenants, err.Error())
}