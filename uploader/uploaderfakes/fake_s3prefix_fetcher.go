// Code generated by counterfeiter. DO NOT EDIT.
package uploaderfakes

import (
	"sync"

	"github.com/pivotal-cf/pivnet-resource/uploader"
)

type FakeS3PrefixFetcher struct {
	S3PrefixForProductSlugStub        func(productSlug string) (string, error)
	s3PrefixForProductSlugMutex       sync.RWMutex
	s3PrefixForProductSlugArgsForCall []struct {
		productSlug string
	}
	s3PrefixForProductSlugReturns struct {
		result1 string
		result2 error
	}
	s3PrefixForProductSlugReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeS3PrefixFetcher) S3PrefixForProductSlug(productSlug string) (string, error) {
	fake.s3PrefixForProductSlugMutex.Lock()
	ret, specificReturn := fake.s3PrefixForProductSlugReturnsOnCall[len(fake.s3PrefixForProductSlugArgsForCall)]
	fake.s3PrefixForProductSlugArgsForCall = append(fake.s3PrefixForProductSlugArgsForCall, struct {
		productSlug string
	}{productSlug})
	fake.recordInvocation("S3PrefixForProductSlug", []interface{}{productSlug})
	fake.s3PrefixForProductSlugMutex.Unlock()
	if fake.S3PrefixForProductSlugStub != nil {
		return fake.S3PrefixForProductSlugStub(productSlug)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.s3PrefixForProductSlugReturns.result1, fake.s3PrefixForProductSlugReturns.result2
}

func (fake *FakeS3PrefixFetcher) S3PrefixForProductSlugCallCount() int {
	fake.s3PrefixForProductSlugMutex.RLock()
	defer fake.s3PrefixForProductSlugMutex.RUnlock()
	return len(fake.s3PrefixForProductSlugArgsForCall)
}

func (fake *FakeS3PrefixFetcher) S3PrefixForProductSlugArgsForCall(i int) string {
	fake.s3PrefixForProductSlugMutex.RLock()
	defer fake.s3PrefixForProductSlugMutex.RUnlock()
	return fake.s3PrefixForProductSlugArgsForCall[i].productSlug
}

func (fake *FakeS3PrefixFetcher) S3PrefixForProductSlugReturns(result1 string, result2 error) {
	fake.S3PrefixForProductSlugStub = nil
	fake.s3PrefixForProductSlugReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeS3PrefixFetcher) S3PrefixForProductSlugReturnsOnCall(i int, result1 string, result2 error) {
	fake.S3PrefixForProductSlugStub = nil
	if fake.s3PrefixForProductSlugReturnsOnCall == nil {
		fake.s3PrefixForProductSlugReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.s3PrefixForProductSlugReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeS3PrefixFetcher) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.s3PrefixForProductSlugMutex.RLock()
	defer fake.s3PrefixForProductSlugMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeS3PrefixFetcher) recordInvocation(key string, args []interface{}) {
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

var _ uploader.S3PrefixFetcher = new(FakeS3PrefixFetcher)
