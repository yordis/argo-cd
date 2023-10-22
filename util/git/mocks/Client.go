// Code generated by mockery v2.40.2. DO NOT EDIT.

package mocks

import (
	git "github.com/argoproj/argo-cd/v2/util/git"
	mock "github.com/stretchr/testify/mock"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// ChangedFiles provides a mock function with given fields: revision, targetRevision
func (_m *Client) ChangedFiles(revision string, targetRevision string) ([]string, error) {
	ret := _m.Called(revision, targetRevision)

	if len(ret) == 0 {
		panic("no return value specified for ChangedFiles")
	}

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) ([]string, error)); ok {
		return rf(revision, targetRevision)
	}
	if rf, ok := ret.Get(0).(func(string, string) []string); ok {
		r0 = rf(revision, targetRevision)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(revision, targetRevision)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Checkout provides a mock function with given fields: revision, submoduleEnabled
func (_m *Client) Checkout(revision string, submoduleEnabled bool) error {
	ret := _m.Called(revision, submoduleEnabled)

	if len(ret) == 0 {
		panic("no return value specified for Checkout")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, bool) error); ok {
		r0 = rf(revision, submoduleEnabled)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CommitSHA provides a mock function with given fields:
func (_m *Client) CommitSHA() (string, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for CommitSHA")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func() (string, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Fetch provides a mock function with given fields: revision, depth
func (_m *Client) Fetch(revision string, depth uint64) error {
	ret := _m.Called(revision, depth)

	if len(ret) == 0 {
		panic("no return value specified for Fetch")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, uint64) error); ok {
		r0 = rf(revision, depth)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Init provides a mock function with given fields:
func (_m *Client) Init() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Init")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IsAnnotatedTag provides a mock function with given fields: _a0
func (_m *Client) IsAnnotatedTag(_a0 string) bool {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for IsAnnotatedTag")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// LsFiles provides a mock function with given fields: path, enableNewGitFileGlobbing
func (_m *Client) LsFiles(path string, enableNewGitFileGlobbing bool) ([]string, error) {
	ret := _m.Called(path, enableNewGitFileGlobbing)

	if len(ret) == 0 {
		panic("no return value specified for LsFiles")
	}

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, bool) ([]string, error)); ok {
		return rf(path, enableNewGitFileGlobbing)
	}
	if rf, ok := ret.Get(0).(func(string, bool) []string); ok {
		r0 = rf(path, enableNewGitFileGlobbing)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(string, bool) error); ok {
		r1 = rf(path, enableNewGitFileGlobbing)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LsLargeFiles provides a mock function with given fields:
func (_m *Client) LsLargeFiles() ([]string, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for LsLargeFiles")
	}

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]string, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LsRefs provides a mock function with given fields:
func (_m *Client) LsRefs() (*git.Refs, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for LsRefs")
	}

	var r0 *git.Refs
	var r1 error
	if rf, ok := ret.Get(0).(func() (*git.Refs, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *git.Refs); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*git.Refs)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LsRemote provides a mock function with given fields: revision
func (_m *Client) LsRemote(revision string) (string, error) {
	ret := _m.Called(revision)

	if len(ret) == 0 {
		panic("no return value specified for LsRemote")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(revision)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(revision)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(revision)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RevisionMetadata provides a mock function with given fields: revision
func (_m *Client) RevisionMetadata(revision string) (*git.RevisionMetadata, error) {
	ret := _m.Called(revision)

	if len(ret) == 0 {
		panic("no return value specified for RevisionMetadata")
	}

	var r0 *git.RevisionMetadata
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*git.RevisionMetadata, error)); ok {
		return rf(revision)
	}
	if rf, ok := ret.Get(0).(func(string) *git.RevisionMetadata); ok {
		r0 = rf(revision)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*git.RevisionMetadata)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(revision)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Root provides a mock function with given fields:
func (_m *Client) Root() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Root")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Submodule provides a mock function with given fields:
func (_m *Client) Submodule() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Submodule")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// VerifyCommitSignature provides a mock function with given fields: _a0
func (_m *Client) VerifyCommitSignature(_a0 string) (string, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for VerifyCommitSignature")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewClient creates a new instance of Client. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *Client {
	mock := &Client{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
