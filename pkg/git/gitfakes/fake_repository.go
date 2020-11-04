/*
Copyright 2020 The Kubernetes Authors.

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

// Code generated by counterfeiter. DO NOT EDIT.
package gitfakes

import (
	"sync"

	gita "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/storer"
	"k8s.io/release/pkg/git"
)

type FakeRepository struct {
	BranchesStub        func() (storer.ReferenceIter, error)
	branchesMutex       sync.RWMutex
	branchesArgsForCall []struct {
	}
	branchesReturns struct {
		result1 storer.ReferenceIter
		result2 error
	}
	branchesReturnsOnCall map[int]struct {
		result1 storer.ReferenceIter
		result2 error
	}
	CommitObjectStub        func(plumbing.Hash) (*object.Commit, error)
	commitObjectMutex       sync.RWMutex
	commitObjectArgsForCall []struct {
		arg1 plumbing.Hash
	}
	commitObjectReturns struct {
		result1 *object.Commit
		result2 error
	}
	commitObjectReturnsOnCall map[int]struct {
		result1 *object.Commit
		result2 error
	}
	CreateRemoteStub        func(*config.RemoteConfig) (*gita.Remote, error)
	createRemoteMutex       sync.RWMutex
	createRemoteArgsForCall []struct {
		arg1 *config.RemoteConfig
	}
	createRemoteReturns struct {
		result1 *gita.Remote
		result2 error
	}
	createRemoteReturnsOnCall map[int]struct {
		result1 *gita.Remote
		result2 error
	}
	DeleteRemoteStub        func(string) error
	deleteRemoteMutex       sync.RWMutex
	deleteRemoteArgsForCall []struct {
		arg1 string
	}
	deleteRemoteReturns struct {
		result1 error
	}
	deleteRemoteReturnsOnCall map[int]struct {
		result1 error
	}
	HeadStub        func() (*plumbing.Reference, error)
	headMutex       sync.RWMutex
	headArgsForCall []struct {
	}
	headReturns struct {
		result1 *plumbing.Reference
		result2 error
	}
	headReturnsOnCall map[int]struct {
		result1 *plumbing.Reference
		result2 error
	}
	RemoteStub        func(string) (*gita.Remote, error)
	remoteMutex       sync.RWMutex
	remoteArgsForCall []struct {
		arg1 string
	}
	remoteReturns struct {
		result1 *gita.Remote
		result2 error
	}
	remoteReturnsOnCall map[int]struct {
		result1 *gita.Remote
		result2 error
	}
	RemotesStub        func() ([]*gita.Remote, error)
	remotesMutex       sync.RWMutex
	remotesArgsForCall []struct {
	}
	remotesReturns struct {
		result1 []*gita.Remote
		result2 error
	}
	remotesReturnsOnCall map[int]struct {
		result1 []*gita.Remote
		result2 error
	}
	ResolveRevisionStub        func(plumbing.Revision) (*plumbing.Hash, error)
	resolveRevisionMutex       sync.RWMutex
	resolveRevisionArgsForCall []struct {
		arg1 plumbing.Revision
	}
	resolveRevisionReturns struct {
		result1 *plumbing.Hash
		result2 error
	}
	resolveRevisionReturnsOnCall map[int]struct {
		result1 *plumbing.Hash
		result2 error
	}
	TagsStub        func() (storer.ReferenceIter, error)
	tagsMutex       sync.RWMutex
	tagsArgsForCall []struct {
	}
	tagsReturns struct {
		result1 storer.ReferenceIter
		result2 error
	}
	tagsReturnsOnCall map[int]struct {
		result1 storer.ReferenceIter
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRepository) Branches() (storer.ReferenceIter, error) {
	fake.branchesMutex.Lock()
	ret, specificReturn := fake.branchesReturnsOnCall[len(fake.branchesArgsForCall)]
	fake.branchesArgsForCall = append(fake.branchesArgsForCall, struct {
	}{})
	fake.recordInvocation("Branches", []interface{}{})
	fake.branchesMutex.Unlock()
	if fake.BranchesStub != nil {
		return fake.BranchesStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.branchesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRepository) BranchesCallCount() int {
	fake.branchesMutex.RLock()
	defer fake.branchesMutex.RUnlock()
	return len(fake.branchesArgsForCall)
}

func (fake *FakeRepository) BranchesCalls(stub func() (storer.ReferenceIter, error)) {
	fake.branchesMutex.Lock()
	defer fake.branchesMutex.Unlock()
	fake.BranchesStub = stub
}

func (fake *FakeRepository) BranchesReturns(result1 storer.ReferenceIter, result2 error) {
	fake.branchesMutex.Lock()
	defer fake.branchesMutex.Unlock()
	fake.BranchesStub = nil
	fake.branchesReturns = struct {
		result1 storer.ReferenceIter
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) BranchesReturnsOnCall(i int, result1 storer.ReferenceIter, result2 error) {
	fake.branchesMutex.Lock()
	defer fake.branchesMutex.Unlock()
	fake.BranchesStub = nil
	if fake.branchesReturnsOnCall == nil {
		fake.branchesReturnsOnCall = make(map[int]struct {
			result1 storer.ReferenceIter
			result2 error
		})
	}
	fake.branchesReturnsOnCall[i] = struct {
		result1 storer.ReferenceIter
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) CommitObject(arg1 plumbing.Hash) (*object.Commit, error) {
	fake.commitObjectMutex.Lock()
	ret, specificReturn := fake.commitObjectReturnsOnCall[len(fake.commitObjectArgsForCall)]
	fake.commitObjectArgsForCall = append(fake.commitObjectArgsForCall, struct {
		arg1 plumbing.Hash
	}{arg1})
	fake.recordInvocation("CommitObject", []interface{}{arg1})
	fake.commitObjectMutex.Unlock()
	if fake.CommitObjectStub != nil {
		return fake.CommitObjectStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.commitObjectReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRepository) CommitObjectCallCount() int {
	fake.commitObjectMutex.RLock()
	defer fake.commitObjectMutex.RUnlock()
	return len(fake.commitObjectArgsForCall)
}

func (fake *FakeRepository) CommitObjectCalls(stub func(plumbing.Hash) (*object.Commit, error)) {
	fake.commitObjectMutex.Lock()
	defer fake.commitObjectMutex.Unlock()
	fake.CommitObjectStub = stub
}

func (fake *FakeRepository) CommitObjectArgsForCall(i int) plumbing.Hash {
	fake.commitObjectMutex.RLock()
	defer fake.commitObjectMutex.RUnlock()
	argsForCall := fake.commitObjectArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRepository) CommitObjectReturns(result1 *object.Commit, result2 error) {
	fake.commitObjectMutex.Lock()
	defer fake.commitObjectMutex.Unlock()
	fake.CommitObjectStub = nil
	fake.commitObjectReturns = struct {
		result1 *object.Commit
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) CommitObjectReturnsOnCall(i int, result1 *object.Commit, result2 error) {
	fake.commitObjectMutex.Lock()
	defer fake.commitObjectMutex.Unlock()
	fake.CommitObjectStub = nil
	if fake.commitObjectReturnsOnCall == nil {
		fake.commitObjectReturnsOnCall = make(map[int]struct {
			result1 *object.Commit
			result2 error
		})
	}
	fake.commitObjectReturnsOnCall[i] = struct {
		result1 *object.Commit
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) CreateRemote(arg1 *config.RemoteConfig) (*gita.Remote, error) {
	fake.createRemoteMutex.Lock()
	ret, specificReturn := fake.createRemoteReturnsOnCall[len(fake.createRemoteArgsForCall)]
	fake.createRemoteArgsForCall = append(fake.createRemoteArgsForCall, struct {
		arg1 *config.RemoteConfig
	}{arg1})
	fake.recordInvocation("CreateRemote", []interface{}{arg1})
	fake.createRemoteMutex.Unlock()
	if fake.CreateRemoteStub != nil {
		return fake.CreateRemoteStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.createRemoteReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRepository) CreateRemoteCallCount() int {
	fake.createRemoteMutex.RLock()
	defer fake.createRemoteMutex.RUnlock()
	return len(fake.createRemoteArgsForCall)
}

func (fake *FakeRepository) CreateRemoteCalls(stub func(*config.RemoteConfig) (*gita.Remote, error)) {
	fake.createRemoteMutex.Lock()
	defer fake.createRemoteMutex.Unlock()
	fake.CreateRemoteStub = stub
}

func (fake *FakeRepository) CreateRemoteArgsForCall(i int) *config.RemoteConfig {
	fake.createRemoteMutex.RLock()
	defer fake.createRemoteMutex.RUnlock()
	argsForCall := fake.createRemoteArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRepository) CreateRemoteReturns(result1 *gita.Remote, result2 error) {
	fake.createRemoteMutex.Lock()
	defer fake.createRemoteMutex.Unlock()
	fake.CreateRemoteStub = nil
	fake.createRemoteReturns = struct {
		result1 *gita.Remote
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) CreateRemoteReturnsOnCall(i int, result1 *gita.Remote, result2 error) {
	fake.createRemoteMutex.Lock()
	defer fake.createRemoteMutex.Unlock()
	fake.CreateRemoteStub = nil
	if fake.createRemoteReturnsOnCall == nil {
		fake.createRemoteReturnsOnCall = make(map[int]struct {
			result1 *gita.Remote
			result2 error
		})
	}
	fake.createRemoteReturnsOnCall[i] = struct {
		result1 *gita.Remote
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) DeleteRemote(arg1 string) error {
	fake.deleteRemoteMutex.Lock()
	ret, specificReturn := fake.deleteRemoteReturnsOnCall[len(fake.deleteRemoteArgsForCall)]
	fake.deleteRemoteArgsForCall = append(fake.deleteRemoteArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("DeleteRemote", []interface{}{arg1})
	fake.deleteRemoteMutex.Unlock()
	if fake.DeleteRemoteStub != nil {
		return fake.DeleteRemoteStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.deleteRemoteReturns
	return fakeReturns.result1
}

func (fake *FakeRepository) DeleteRemoteCallCount() int {
	fake.deleteRemoteMutex.RLock()
	defer fake.deleteRemoteMutex.RUnlock()
	return len(fake.deleteRemoteArgsForCall)
}

func (fake *FakeRepository) DeleteRemoteCalls(stub func(string) error) {
	fake.deleteRemoteMutex.Lock()
	defer fake.deleteRemoteMutex.Unlock()
	fake.DeleteRemoteStub = stub
}

func (fake *FakeRepository) DeleteRemoteArgsForCall(i int) string {
	fake.deleteRemoteMutex.RLock()
	defer fake.deleteRemoteMutex.RUnlock()
	argsForCall := fake.deleteRemoteArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRepository) DeleteRemoteReturns(result1 error) {
	fake.deleteRemoteMutex.Lock()
	defer fake.deleteRemoteMutex.Unlock()
	fake.DeleteRemoteStub = nil
	fake.deleteRemoteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRepository) DeleteRemoteReturnsOnCall(i int, result1 error) {
	fake.deleteRemoteMutex.Lock()
	defer fake.deleteRemoteMutex.Unlock()
	fake.DeleteRemoteStub = nil
	if fake.deleteRemoteReturnsOnCall == nil {
		fake.deleteRemoteReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteRemoteReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRepository) Head() (*plumbing.Reference, error) {
	fake.headMutex.Lock()
	ret, specificReturn := fake.headReturnsOnCall[len(fake.headArgsForCall)]
	fake.headArgsForCall = append(fake.headArgsForCall, struct {
	}{})
	fake.recordInvocation("Head", []interface{}{})
	fake.headMutex.Unlock()
	if fake.HeadStub != nil {
		return fake.HeadStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.headReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRepository) HeadCallCount() int {
	fake.headMutex.RLock()
	defer fake.headMutex.RUnlock()
	return len(fake.headArgsForCall)
}

func (fake *FakeRepository) HeadCalls(stub func() (*plumbing.Reference, error)) {
	fake.headMutex.Lock()
	defer fake.headMutex.Unlock()
	fake.HeadStub = stub
}

func (fake *FakeRepository) HeadReturns(result1 *plumbing.Reference, result2 error) {
	fake.headMutex.Lock()
	defer fake.headMutex.Unlock()
	fake.HeadStub = nil
	fake.headReturns = struct {
		result1 *plumbing.Reference
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) HeadReturnsOnCall(i int, result1 *plumbing.Reference, result2 error) {
	fake.headMutex.Lock()
	defer fake.headMutex.Unlock()
	fake.HeadStub = nil
	if fake.headReturnsOnCall == nil {
		fake.headReturnsOnCall = make(map[int]struct {
			result1 *plumbing.Reference
			result2 error
		})
	}
	fake.headReturnsOnCall[i] = struct {
		result1 *plumbing.Reference
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) Remote(arg1 string) (*gita.Remote, error) {
	fake.remoteMutex.Lock()
	ret, specificReturn := fake.remoteReturnsOnCall[len(fake.remoteArgsForCall)]
	fake.remoteArgsForCall = append(fake.remoteArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Remote", []interface{}{arg1})
	fake.remoteMutex.Unlock()
	if fake.RemoteStub != nil {
		return fake.RemoteStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.remoteReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRepository) RemoteCallCount() int {
	fake.remoteMutex.RLock()
	defer fake.remoteMutex.RUnlock()
	return len(fake.remoteArgsForCall)
}

func (fake *FakeRepository) RemoteCalls(stub func(string) (*gita.Remote, error)) {
	fake.remoteMutex.Lock()
	defer fake.remoteMutex.Unlock()
	fake.RemoteStub = stub
}

func (fake *FakeRepository) RemoteArgsForCall(i int) string {
	fake.remoteMutex.RLock()
	defer fake.remoteMutex.RUnlock()
	argsForCall := fake.remoteArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRepository) RemoteReturns(result1 *gita.Remote, result2 error) {
	fake.remoteMutex.Lock()
	defer fake.remoteMutex.Unlock()
	fake.RemoteStub = nil
	fake.remoteReturns = struct {
		result1 *gita.Remote
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) RemoteReturnsOnCall(i int, result1 *gita.Remote, result2 error) {
	fake.remoteMutex.Lock()
	defer fake.remoteMutex.Unlock()
	fake.RemoteStub = nil
	if fake.remoteReturnsOnCall == nil {
		fake.remoteReturnsOnCall = make(map[int]struct {
			result1 *gita.Remote
			result2 error
		})
	}
	fake.remoteReturnsOnCall[i] = struct {
		result1 *gita.Remote
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) Remotes() ([]*gita.Remote, error) {
	fake.remotesMutex.Lock()
	ret, specificReturn := fake.remotesReturnsOnCall[len(fake.remotesArgsForCall)]
	fake.remotesArgsForCall = append(fake.remotesArgsForCall, struct {
	}{})
	fake.recordInvocation("Remotes", []interface{}{})
	fake.remotesMutex.Unlock()
	if fake.RemotesStub != nil {
		return fake.RemotesStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.remotesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRepository) RemotesCallCount() int {
	fake.remotesMutex.RLock()
	defer fake.remotesMutex.RUnlock()
	return len(fake.remotesArgsForCall)
}

func (fake *FakeRepository) RemotesCalls(stub func() ([]*gita.Remote, error)) {
	fake.remotesMutex.Lock()
	defer fake.remotesMutex.Unlock()
	fake.RemotesStub = stub
}

func (fake *FakeRepository) RemotesReturns(result1 []*gita.Remote, result2 error) {
	fake.remotesMutex.Lock()
	defer fake.remotesMutex.Unlock()
	fake.RemotesStub = nil
	fake.remotesReturns = struct {
		result1 []*gita.Remote
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) RemotesReturnsOnCall(i int, result1 []*gita.Remote, result2 error) {
	fake.remotesMutex.Lock()
	defer fake.remotesMutex.Unlock()
	fake.RemotesStub = nil
	if fake.remotesReturnsOnCall == nil {
		fake.remotesReturnsOnCall = make(map[int]struct {
			result1 []*gita.Remote
			result2 error
		})
	}
	fake.remotesReturnsOnCall[i] = struct {
		result1 []*gita.Remote
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) ResolveRevision(arg1 plumbing.Revision) (*plumbing.Hash, error) {
	fake.resolveRevisionMutex.Lock()
	ret, specificReturn := fake.resolveRevisionReturnsOnCall[len(fake.resolveRevisionArgsForCall)]
	fake.resolveRevisionArgsForCall = append(fake.resolveRevisionArgsForCall, struct {
		arg1 plumbing.Revision
	}{arg1})
	fake.recordInvocation("ResolveRevision", []interface{}{arg1})
	fake.resolveRevisionMutex.Unlock()
	if fake.ResolveRevisionStub != nil {
		return fake.ResolveRevisionStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.resolveRevisionReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRepository) ResolveRevisionCallCount() int {
	fake.resolveRevisionMutex.RLock()
	defer fake.resolveRevisionMutex.RUnlock()
	return len(fake.resolveRevisionArgsForCall)
}

func (fake *FakeRepository) ResolveRevisionCalls(stub func(plumbing.Revision) (*plumbing.Hash, error)) {
	fake.resolveRevisionMutex.Lock()
	defer fake.resolveRevisionMutex.Unlock()
	fake.ResolveRevisionStub = stub
}

func (fake *FakeRepository) ResolveRevisionArgsForCall(i int) plumbing.Revision {
	fake.resolveRevisionMutex.RLock()
	defer fake.resolveRevisionMutex.RUnlock()
	argsForCall := fake.resolveRevisionArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRepository) ResolveRevisionReturns(result1 *plumbing.Hash, result2 error) {
	fake.resolveRevisionMutex.Lock()
	defer fake.resolveRevisionMutex.Unlock()
	fake.ResolveRevisionStub = nil
	fake.resolveRevisionReturns = struct {
		result1 *plumbing.Hash
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) ResolveRevisionReturnsOnCall(i int, result1 *plumbing.Hash, result2 error) {
	fake.resolveRevisionMutex.Lock()
	defer fake.resolveRevisionMutex.Unlock()
	fake.ResolveRevisionStub = nil
	if fake.resolveRevisionReturnsOnCall == nil {
		fake.resolveRevisionReturnsOnCall = make(map[int]struct {
			result1 *plumbing.Hash
			result2 error
		})
	}
	fake.resolveRevisionReturnsOnCall[i] = struct {
		result1 *plumbing.Hash
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) Tags() (storer.ReferenceIter, error) {
	fake.tagsMutex.Lock()
	ret, specificReturn := fake.tagsReturnsOnCall[len(fake.tagsArgsForCall)]
	fake.tagsArgsForCall = append(fake.tagsArgsForCall, struct {
	}{})
	fake.recordInvocation("Tags", []interface{}{})
	fake.tagsMutex.Unlock()
	if fake.TagsStub != nil {
		return fake.TagsStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.tagsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRepository) TagsCallCount() int {
	fake.tagsMutex.RLock()
	defer fake.tagsMutex.RUnlock()
	return len(fake.tagsArgsForCall)
}

func (fake *FakeRepository) TagsCalls(stub func() (storer.ReferenceIter, error)) {
	fake.tagsMutex.Lock()
	defer fake.tagsMutex.Unlock()
	fake.TagsStub = stub
}

func (fake *FakeRepository) TagsReturns(result1 storer.ReferenceIter, result2 error) {
	fake.tagsMutex.Lock()
	defer fake.tagsMutex.Unlock()
	fake.TagsStub = nil
	fake.tagsReturns = struct {
		result1 storer.ReferenceIter
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) TagsReturnsOnCall(i int, result1 storer.ReferenceIter, result2 error) {
	fake.tagsMutex.Lock()
	defer fake.tagsMutex.Unlock()
	fake.TagsStub = nil
	if fake.tagsReturnsOnCall == nil {
		fake.tagsReturnsOnCall = make(map[int]struct {
			result1 storer.ReferenceIter
			result2 error
		})
	}
	fake.tagsReturnsOnCall[i] = struct {
		result1 storer.ReferenceIter
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.branchesMutex.RLock()
	defer fake.branchesMutex.RUnlock()
	fake.commitObjectMutex.RLock()
	defer fake.commitObjectMutex.RUnlock()
	fake.createRemoteMutex.RLock()
	defer fake.createRemoteMutex.RUnlock()
	fake.deleteRemoteMutex.RLock()
	defer fake.deleteRemoteMutex.RUnlock()
	fake.headMutex.RLock()
	defer fake.headMutex.RUnlock()
	fake.remoteMutex.RLock()
	defer fake.remoteMutex.RUnlock()
	fake.remotesMutex.RLock()
	defer fake.remotesMutex.RUnlock()
	fake.resolveRevisionMutex.RLock()
	defer fake.resolveRevisionMutex.RUnlock()
	fake.tagsMutex.RLock()
	defer fake.tagsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRepository) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ git.Repository = new(FakeRepository)
