// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ClusterWorkflowTemplateLister helps list ClusterWorkflowTemplates.
// All objects returned here must be treated as read-only.
type ClusterWorkflowTemplateLister interface {
	// List lists all ClusterWorkflowTemplates in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ClusterWorkflowTemplate, err error)
	// Get retrieves the ClusterWorkflowTemplate from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ClusterWorkflowTemplate, error)
	ClusterWorkflowTemplateListerExpansion
}

// clusterWorkflowTemplateLister implements the ClusterWorkflowTemplateLister interface.
type clusterWorkflowTemplateLister struct {
	indexer cache.Indexer
}

// NewClusterWorkflowTemplateLister returns a new ClusterWorkflowTemplateLister.
func NewClusterWorkflowTemplateLister(indexer cache.Indexer) ClusterWorkflowTemplateLister {
	return &clusterWorkflowTemplateLister{indexer: indexer}
}

// List lists all ClusterWorkflowTemplates in the indexer.
func (s *clusterWorkflowTemplateLister) List(selector labels.Selector) (ret []*v1alpha1.ClusterWorkflowTemplate, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ClusterWorkflowTemplate))
	})
	return ret, err
}

// Get retrieves the ClusterWorkflowTemplate from the index for a given name.
func (s *clusterWorkflowTemplateLister) Get(name string) (*v1alpha1.ClusterWorkflowTemplate, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("clusterworkflowtemplate"), name)
	}
	return obj.(*v1alpha1.ClusterWorkflowTemplate), nil
}
