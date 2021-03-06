package k8s

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Namespace represents a Kubernetes namespace.
type Namespace struct{}

// NewNamespace returns a new Namespace.
func NewNamespace() Res {
	return &Namespace{}
}

// Get a namespace.
func (*Namespace) Get(_, n string) (interface{}, error) {
	opts := metav1.GetOptions{}
	return conn.dialOrDie().CoreV1().Namespaces().Get(n, opts)
}

// List all namespaces on the cluster.
func (*Namespace) List(_ string) (Collection, error) {
	opts := metav1.ListOptions{}

	rr, err := conn.dialOrDie().CoreV1().Namespaces().List(opts)
	if err != nil {
		return Collection{}, err
	}

	cc := make(Collection, len(rr.Items))
	for i, r := range rr.Items {
		cc[i] = r
	}

	return cc, nil
}

// Delete a namespace.
func (*Namespace) Delete(_, n string) error {
	opts := metav1.DeleteOptions{}
	return conn.dialOrDie().CoreV1().Namespaces().Delete(n, &opts)
}
