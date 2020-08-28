// Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package client

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// Finalizer checks and removes the finalizers of given resource.
type Finalizer interface {
	// Finalize removes the resource finalizers (so it can be garbage collected).
	Finalize(ctx context.Context, c client.Client, obj runtime.Object) error

	// HasFinalizers checks whether the given resource has finalizers.
	HasFinalizers(obj runtime.Object) (bool, error)
}

// Cleaner is capable of deleting and finalizing resources.
type Cleaner interface {
	// Clean cleans the given resource(s). It first tries to delete them. If they are 'hanging'
	// in deletion state and `FinalizeGracePeriodSeconds` is specified, then they are finalized
	// once the `deletionTimestamp` is beyond that amount in the past.
	Clean(ctx context.Context, c client.Client, obj runtime.Object, opts ...CleanOption) error
}

// GoneEnsurer ensures that resource(s) are gone.
type GoneEnsurer interface {
	// EnsureGone ensures that the given resource is gone. If the resource is not gone, it will throw
	// a NewObjectsRemaining error.
	EnsureGone(ctx context.Context, c client.Client, obj runtime.Object, opts ...client.ListOption) error
}

// GoneEnsurerFunc is a function that implements GoneEnsurer.
type GoneEnsurerFunc func(ctx context.Context, c client.Client, obj runtime.Object, opts ...client.ListOption) error

// CleanOps are ops to clean.
type CleanOps interface {
	// CleanAndEnsureGone cleans the resource(s) and ensures that it/they are gone afterwards.
	CleanAndEnsureGone(ctx context.Context, c client.Client, obj runtime.Object, opts ...CleanOption) error
}
