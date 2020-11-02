/*
Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved.

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

// Package validation - validation is used to validate cloud specific ProviderSpec
package validation

import (
	"errors"

	utilerror "k8s.io/apimachinery/pkg/util/errors"

	api "github.com/gardener/machine-controller-manager-provider-digitalocean/pkg/provider/apis"
	"github.com/gardener/machine-controller-manager-provider-digitalocean/pkg/provider/internal/util"

	corev1 "k8s.io/api/core/v1"
)

// ValidateProviderSpecNSecret validates provider spec and secret to check if all fields are present and valid
func ValidateProviderSpecNSecret(spec *api.ProviderSpec, secret *corev1.Secret) error {
	var errs []error

	if spec.Name == "" {
		errs = append(errs, errors.New("name is missing"))
	}
	if spec.Image == "" {
		errs = append(errs, errors.New("image is missing"))
	}
	if spec.Size == "" {
		errs = append(errs, errors.New("size is missing"))
	}
	if spec.Region == "" {
		errs = append(errs, errors.New("region is missing"))
	}
	if spec.VPCUUID == "" {
		errs = append(errs, errors.New("VPC UUID is missing"))
	}
	if len(spec.Tags) == 0 {
		errs = append(errs, errors.New("tags are missing"))
	}

	token := util.ExtractToken(secret)
	if token == "" {
		errs = append(errs, errors.New("token is missing from secret"))
	}

	return utilerror.NewAggregate(errs)
}
