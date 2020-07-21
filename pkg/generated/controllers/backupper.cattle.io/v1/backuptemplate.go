/*
Copyright 2020 Rancher Labs, Inc.

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

// Code generated by main. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/mrajashree/backup/pkg/apis/backupper.cattle.io/v1"
	"github.com/rancher/lasso/pkg/client"
	"github.com/rancher/lasso/pkg/controller"
	"github.com/rancher/wrangler/pkg/generic"
	"k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

type BackupTemplateHandler func(string, *v1.BackupTemplate) (*v1.BackupTemplate, error)

type BackupTemplateController interface {
	generic.ControllerMeta
	BackupTemplateClient

	OnChange(ctx context.Context, name string, sync BackupTemplateHandler)
	OnRemove(ctx context.Context, name string, sync BackupTemplateHandler)
	Enqueue(namespace, name string)
	EnqueueAfter(namespace, name string, duration time.Duration)

	Cache() BackupTemplateCache
}

type BackupTemplateClient interface {
	Create(*v1.BackupTemplate) (*v1.BackupTemplate, error)
	Update(*v1.BackupTemplate) (*v1.BackupTemplate, error)

	Delete(namespace, name string, options *metav1.DeleteOptions) error
	Get(namespace, name string, options metav1.GetOptions) (*v1.BackupTemplate, error)
	List(namespace string, opts metav1.ListOptions) (*v1.BackupTemplateList, error)
	Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error)
	Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.BackupTemplate, err error)
}

type BackupTemplateCache interface {
	Get(namespace, name string) (*v1.BackupTemplate, error)
	List(namespace string, selector labels.Selector) ([]*v1.BackupTemplate, error)

	AddIndexer(indexName string, indexer BackupTemplateIndexer)
	GetByIndex(indexName, key string) ([]*v1.BackupTemplate, error)
}

type BackupTemplateIndexer func(obj *v1.BackupTemplate) ([]string, error)

type backupTemplateController struct {
	controller    controller.SharedController
	client        *client.Client
	gvk           schema.GroupVersionKind
	groupResource schema.GroupResource
}

func NewBackupTemplateController(gvk schema.GroupVersionKind, resource string, namespaced bool, controller controller.SharedControllerFactory) BackupTemplateController {
	c := controller.ForResourceKind(gvk.GroupVersion().WithResource(resource), gvk.Kind, namespaced)
	return &backupTemplateController{
		controller: c,
		client:     c.Client(),
		gvk:        gvk,
		groupResource: schema.GroupResource{
			Group:    gvk.Group,
			Resource: resource,
		},
	}
}

func FromBackupTemplateHandlerToHandler(sync BackupTemplateHandler) generic.Handler {
	return func(key string, obj runtime.Object) (ret runtime.Object, err error) {
		var v *v1.BackupTemplate
		if obj == nil {
			v, err = sync(key, nil)
		} else {
			v, err = sync(key, obj.(*v1.BackupTemplate))
		}
		if v == nil {
			return nil, err
		}
		return v, err
	}
}

func (c *backupTemplateController) Updater() generic.Updater {
	return func(obj runtime.Object) (runtime.Object, error) {
		newObj, err := c.Update(obj.(*v1.BackupTemplate))
		if newObj == nil {
			return nil, err
		}
		return newObj, err
	}
}

func UpdateBackupTemplateDeepCopyOnChange(client BackupTemplateClient, obj *v1.BackupTemplate, handler func(obj *v1.BackupTemplate) (*v1.BackupTemplate, error)) (*v1.BackupTemplate, error) {
	if obj == nil {
		return obj, nil
	}

	copyObj := obj.DeepCopy()
	newObj, err := handler(copyObj)
	if newObj != nil {
		copyObj = newObj
	}
	if obj.ResourceVersion == copyObj.ResourceVersion && !equality.Semantic.DeepEqual(obj, copyObj) {
		return client.Update(copyObj)
	}

	return copyObj, err
}

func (c *backupTemplateController) AddGenericHandler(ctx context.Context, name string, handler generic.Handler) {
	c.controller.RegisterHandler(ctx, name, controller.SharedControllerHandlerFunc(handler))
}

func (c *backupTemplateController) AddGenericRemoveHandler(ctx context.Context, name string, handler generic.Handler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), handler))
}

func (c *backupTemplateController) OnChange(ctx context.Context, name string, sync BackupTemplateHandler) {
	c.AddGenericHandler(ctx, name, FromBackupTemplateHandlerToHandler(sync))
}

func (c *backupTemplateController) OnRemove(ctx context.Context, name string, sync BackupTemplateHandler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), FromBackupTemplateHandlerToHandler(sync)))
}

func (c *backupTemplateController) Enqueue(namespace, name string) {
	c.controller.Enqueue(namespace, name)
}

func (c *backupTemplateController) EnqueueAfter(namespace, name string, duration time.Duration) {
	c.controller.EnqueueAfter(namespace, name, duration)
}

func (c *backupTemplateController) Informer() cache.SharedIndexInformer {
	return c.controller.Informer()
}

func (c *backupTemplateController) GroupVersionKind() schema.GroupVersionKind {
	return c.gvk
}

func (c *backupTemplateController) Cache() BackupTemplateCache {
	return &backupTemplateCache{
		indexer:  c.Informer().GetIndexer(),
		resource: c.groupResource,
	}
}

func (c *backupTemplateController) Create(obj *v1.BackupTemplate) (*v1.BackupTemplate, error) {
	result := &v1.BackupTemplate{}
	return result, c.client.Create(context.TODO(), obj.Namespace, obj, result, metav1.CreateOptions{})
}

func (c *backupTemplateController) Update(obj *v1.BackupTemplate) (*v1.BackupTemplate, error) {
	result := &v1.BackupTemplate{}
	return result, c.client.Update(context.TODO(), obj.Namespace, obj, result, metav1.UpdateOptions{})
}

func (c *backupTemplateController) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	if options == nil {
		options = &metav1.DeleteOptions{}
	}
	return c.client.Delete(context.TODO(), namespace, name, *options)
}

func (c *backupTemplateController) Get(namespace, name string, options metav1.GetOptions) (*v1.BackupTemplate, error) {
	result := &v1.BackupTemplate{}
	return result, c.client.Get(context.TODO(), namespace, name, result, options)
}

func (c *backupTemplateController) List(namespace string, opts metav1.ListOptions) (*v1.BackupTemplateList, error) {
	result := &v1.BackupTemplateList{}
	return result, c.client.List(context.TODO(), namespace, result, opts)
}

func (c *backupTemplateController) Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error) {
	return c.client.Watch(context.TODO(), namespace, opts)
}

func (c *backupTemplateController) Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (*v1.BackupTemplate, error) {
	result := &v1.BackupTemplate{}
	return result, c.client.Patch(context.TODO(), namespace, name, pt, data, result, metav1.PatchOptions{}, subresources...)
}

type backupTemplateCache struct {
	indexer  cache.Indexer
	resource schema.GroupResource
}

func (c *backupTemplateCache) Get(namespace, name string) (*v1.BackupTemplate, error) {
	obj, exists, err := c.indexer.GetByKey(namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(c.resource, name)
	}
	return obj.(*v1.BackupTemplate), nil
}

func (c *backupTemplateCache) List(namespace string, selector labels.Selector) (ret []*v1.BackupTemplate, err error) {

	err = cache.ListAllByNamespace(c.indexer, namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.BackupTemplate))
	})

	return ret, err
}

func (c *backupTemplateCache) AddIndexer(indexName string, indexer BackupTemplateIndexer) {
	utilruntime.Must(c.indexer.AddIndexers(map[string]cache.IndexFunc{
		indexName: func(obj interface{}) (strings []string, e error) {
			return indexer(obj.(*v1.BackupTemplate))
		},
	}))
}

func (c *backupTemplateCache) GetByIndex(indexName, key string) (result []*v1.BackupTemplate, err error) {
	objs, err := c.indexer.ByIndex(indexName, key)
	if err != nil {
		return nil, err
	}
	result = make([]*v1.BackupTemplate, 0, len(objs))
	for _, obj := range objs {
		result = append(result, obj.(*v1.BackupTemplate))
	}
	return result, nil
}