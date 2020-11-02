/*
Copyright (c) 2020 SAP SE or an SAP affiliate company. All rights reserved.

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

package spi

import (
	"context"

	"github.com/digitalocean/godo"
)

// SessionProviderInterface provides an interface to deal with cloud provider session
// Example interfaces are listed below.
type SessionProviderInterface interface {
	CreateDroplet(ctx context.Context, req *godo.DropletCreateRequest) (*godo.Droplet, error)
	DeleteDroplet(ctx context.Context, id int) error
	// NewSession(*corev1.Secret, string) (*session.Session, error)
	// NewEC2API(*session.Session) ec2iface.EC2API
}

func (p *PluginSPIImpl) CreateDroplet(ctx context.Context, req *godo.DropletCreateRequest) (*godo.Droplet, error) {
	d, _, err := p.client.Droplets.Create(ctx, req)
	return d, err
}

func (p *PluginSPIImpl) DeleteDroplet(ctx context.Context, id int) error {
	_, err := p.client.Droplets.Delete(ctx, id)
	return err
}

// PluginSPIImpl is the real implementation of SPI interface that makes the calls to the provider SDK.
type PluginSPIImpl struct{
	client *godo.Client
}

func NewPluginSPIImpl(token string) *PluginSPIImpl {
	return &PluginSPIImpl{
		client: godo.NewFromToken(token),
	}
}

