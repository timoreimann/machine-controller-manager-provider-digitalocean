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

package provider

import (
	"encoding/json"

	"github.com/gardener/machine-controller-manager/pkg/apis/machine/v1alpha1"
	"github.com/gardener/machine-controller-manager/pkg/util/provider/machinecodes/codes"
	"github.com/gardener/machine-controller-manager/pkg/util/provider/machinecodes/status"
	corev1 "k8s.io/api/core/v1"

	api "github.com/gardener/machine-controller-manager-provider-digitalocean/pkg/provider/apis"
	"github.com/gardener/machine-controller-manager-provider-digitalocean/pkg/provider/apis/validation"
	"github.com/gardener/machine-controller-manager-provider-digitalocean/pkg/provider/internal/util"
)

// decodeProviderSpecAndSecret converts request parameters to api.ProviderSpec & api.Secrets
func decodeProviderSpecAndSecret(machineClass *v1alpha1.MachineClass, secret *corev1.Secret) (*api.ProviderSpec, string, error) {
	var (
		providerSpec *api.ProviderSpec
	)

	// Extract providerSpec
	err := json.Unmarshal(machineClass.ProviderSpec.Raw, &providerSpec)
	if err != nil {
		return nil, "", status.Error(codes.Internal, err.Error())
	}

	//Validate the Spec and Secrets
	err = validation.ValidateProviderSpecNSecret(providerSpec, secret)
	if err != nil {
		return nil, "", status.Error(codes.InvalidArgument, err.Error())
	}

	return providerSpec, util.ExtractToken(secret), nil
}
