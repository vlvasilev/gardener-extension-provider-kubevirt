// Copyright (c) 2020 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
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

package validation_test

import (
	api "github.com/gardener/gardener-extension-provider-kubevirt/pkg/apis/kubevirt"
	. "github.com/gardener/gardener-extension-provider-kubevirt/pkg/apis/kubevirt/validation"

	gardencorev1beta1 "github.com/gardener/gardener/pkg/apis/core/v1beta1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

var _ = Describe("ControlPlaneConfig validation", func() {
	var (
		nilPath *field.Path

		controlPlane *api.ControlPlaneConfig
	)

	BeforeEach(func() {
		controlPlane = &api.ControlPlaneConfig{}
	})

	Describe("#ValidateControlPlaneConfig", func() {
		It("should return no errors for a valid configuration", func() {
			Expect(ValidateControlPlaneConfig(controlPlane, nilPath)).To(BeEmpty())
		})
	})

	Describe("#ValidateControlPlaneConfigUpdate", func() {
		It("should return no errors for an unchanged config", func() {
			Expect(ValidateControlPlaneConfigUpdate(controlPlane, controlPlane, nilPath)).To(BeEmpty())
		})
	})

	Describe("#ValidateControlPlaneConfigAgainstCloudProfile", func() {
		var (
			cloudProfileConfig *api.CloudProfileConfig
			cloudProfile       *gardencorev1beta1.CloudProfile
		)

		BeforeEach(func() {
			cloudProfile = &gardencorev1beta1.CloudProfile{}
			cloudProfileConfig = &api.CloudProfileConfig{}
		})

		It("should return no errors if the name is not defined in the constraints", func() {
			errorList := ValidateControlPlaneConfigAgainstCloudProfile(controlPlane, "testRegion", cloudProfile, cloudProfileConfig, nilPath)
			Expect(errorList).To(BeEmpty())
		})

	})
})
