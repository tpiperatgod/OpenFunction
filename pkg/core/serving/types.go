/*
Copyright 2022 The OpenFunction Authors.

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

package serving

import (
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/apiutil"

	"github.com/openfunction/pkg/core/serving/knative"
	"github.com/openfunction/pkg/core/serving/openfuncasync"
)

func Registry(mgr ctrl.Manager) []client.Object {
	rm, err := apiutil.NewDiscoveryRESTMapper(mgr.GetConfig())
	if err != nil {
		return nil
	}

	var objs []client.Object
	objs = append(objs, knative.Registry(rm)...)
	objs = append(objs, openfuncasync.Registry(rm)...)

	return objs
}