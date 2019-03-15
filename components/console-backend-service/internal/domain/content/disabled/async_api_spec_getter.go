// Code generated by failery v1.0.0. DO NOT EDIT.

package disabled

import storage "github.com/kyma-project/kyma/components/console-backend-service/internal/domain/content/storage"

// asyncApiSpecGetter is an autogenerated failing mock type for the asyncApiSpecGetter type
type asyncApiSpecGetter struct {
	err error
}

// NewAsyncApiSpecGetter creates a new asyncApiSpecGetter type instance
func NewAsyncApiSpecGetter(err error) *asyncApiSpecGetter {
	return &asyncApiSpecGetter{err: err}
}

// Find provides a failing mock function with given fields: kind, id
func (_m *asyncApiSpecGetter) Find(kind string, id string) (*storage.AsyncApiSpec, error) {
	var r0 *storage.AsyncApiSpec
	var r1 error
	r1 = _m.err

	return r0, r1
}