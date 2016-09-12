/*
Copyright 2016 The Kubernetes Authors All rights reserved.

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

package storage // import "k8s.io/helm/pkg/storage"

import rspb "k8s.io/helm/pkg/proto/hapi/release"

// FilterFunc returns true if the release object satisfies
// the predicate of the underlying func.
type FilterFunc func(*rspb.Release) bool

// Check applies the FilterFunc to the release object.
func (fn FilterFunc) Check(rls *rspb.Release) bool {
	if rls == nil {
		return false
	}
	return fn(rls)
}

// Any returns a FilterFunc that filters a list of releases
// determined by the predicate 'f0 || f1 || ... || fn'.
func Any(filters ...FilterFunc) FilterFunc {
	return func(rls *rspb.Release) bool {
		for _, filter := range filters {
			if filter(rls) {
				return true
			}
		}
		return false
	}
}

// All returns a FilterFunc that filters a list of releases
// determined by the predicate 'f0 && f1 && ... && fn'.
func All(filters ...FilterFunc) FilterFunc {
	return func(rls *rspb.Release) bool {
		for _, filter := range filters {
			if !filter(rls) {
				return false
			}
		}
		return true
	}
}

// StatusFilter filters a set of releases by status code.
func StatusFilter(status rspb.Status_Code) FilterFunc {
	return FilterFunc(func(rls *rspb.Release) bool {
		if rls == nil {
			return true
		}
		return rls.GetInfo().GetStatus().Code == status
	})
}
