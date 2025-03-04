// Copyright The prometheus-operator Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	monitoringv1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	applyconfigurationmonitoringv1 "github.com/prometheus-operator/prometheus-operator/pkg/client/applyconfiguration/monitoring/v1"
	v1 "k8s.io/api/core/v1"
)

// ConsulSDConfigApplyConfiguration represents a declarative configuration of the ConsulSDConfig type for use
// with apply.
type ConsulSDConfigApplyConfiguration struct {
	Server                                                       *string                                                             `json:"server,omitempty"`
	PathPrefix                                                   *string                                                             `json:"pathPrefix,omitempty"`
	TokenRef                                                     *v1.SecretKeySelector                                               `json:"tokenRef,omitempty"`
	Datacenter                                                   *string                                                             `json:"datacenter,omitempty"`
	Namespace                                                    *string                                                             `json:"namespace,omitempty"`
	Partition                                                    *string                                                             `json:"partition,omitempty"`
	Scheme                                                       *string                                                             `json:"scheme,omitempty"`
	Services                                                     []string                                                            `json:"services,omitempty"`
	Tags                                                         []string                                                            `json:"tags,omitempty"`
	TagSeparator                                                 *string                                                             `json:"tagSeparator,omitempty"`
	NodeMeta                                                     map[string]string                                                   `json:"nodeMeta,omitempty"`
	Filter                                                       *string                                                             `json:"filter,omitempty"`
	AllowStale                                                   *bool                                                               `json:"allowStale,omitempty"`
	RefreshInterval                                              *monitoringv1.Duration                                              `json:"refreshInterval,omitempty"`
	BasicAuth                                                    *applyconfigurationmonitoringv1.BasicAuthApplyConfiguration         `json:"basicAuth,omitempty"`
	Authorization                                                *applyconfigurationmonitoringv1.SafeAuthorizationApplyConfiguration `json:"authorization,omitempty"`
	OAuth2                                                       *applyconfigurationmonitoringv1.OAuth2ApplyConfiguration            `json:"oauth2,omitempty"`
	applyconfigurationmonitoringv1.ProxyConfigApplyConfiguration `json:",inline"`
	FollowRedirects                                              *bool                                                           `json:"followRedirects,omitempty"`
	EnableHttp2                                                  *bool                                                           `json:"enableHTTP2,omitempty"`
	TLSConfig                                                    *applyconfigurationmonitoringv1.SafeTLSConfigApplyConfiguration `json:"tlsConfig,omitempty"`
}

// ConsulSDConfigApplyConfiguration constructs a declarative configuration of the ConsulSDConfig type for use with
// apply.
func ConsulSDConfig() *ConsulSDConfigApplyConfiguration {
	return &ConsulSDConfigApplyConfiguration{}
}

// WithServer sets the Server field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Server field is set to the value of the last call.
func (b *ConsulSDConfigApplyConfiguration) WithServer(value string) *ConsulSDConfigApplyConfiguration {
	b.Server = &value
	return b
}

// WithPathPrefix sets the PathPrefix field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PathPrefix field is set to the value of the last call.
func (b *ConsulSDConfigApplyConfiguration) WithPathPrefix(value string) *ConsulSDConfigApplyConfiguration {
	b.PathPrefix = &value
	return b
}

// WithTokenRef sets the TokenRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TokenRef field is set to the value of the last call.
func (b *ConsulSDConfigApplyConfiguration) WithTokenRef(value v1.SecretKeySelector) *ConsulSDConfigApplyConfiguration {
	b.TokenRef = &value
	return b
}

// WithDatacenter sets the Datacenter field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Datacenter field is set to the value of the last call.
func (b *ConsulSDConfigApplyConfiguration) WithDatacenter(value string) *ConsulSDConfigApplyConfiguration {
	b.Datacenter = &value
	return b
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *ConsulSDConfigApplyConfiguration) WithNamespace(value string) *ConsulSDConfigApplyConfiguration {
	b.Namespace = &value
	return b
}

// WithPartition sets the Partition field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Partition field is set to the value of the last call.
func (b *ConsulSDConfigApplyConfiguration) WithPartition(value string) *ConsulSDConfigApplyConfiguration {
	b.Partition = &value
	return b
}

// WithScheme sets the Scheme field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Scheme field is set to the value of the last call.
func (b *ConsulSDConfigApplyConfiguration) WithScheme(value string) *ConsulSDConfigApplyConfiguration {
	b.Scheme = &value
	return b
}

// WithServices adds the given value to the Services field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Services field.
func (b *ConsulSDConfigApplyConfiguration) WithServices(values ...string) *ConsulSDConfigApplyConfiguration {
	for i := range values {
		b.Services = append(b.Services, values[i])
	}
	return b
}

// WithTags adds the given value to the Tags field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Tags field.
func (b *ConsulSDConfigApplyConfiguration) WithTags(values ...string) *ConsulSDConfigApplyConfiguration {
	for i := range values {
		b.Tags = append(b.Tags, values[i])
	}
	return b
}

// WithTagSeparator sets the TagSeparator field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TagSeparator field is set to the value of the last call.
func (b *ConsulSDConfigApplyConfiguration) WithTagSeparator(value string) *ConsulSDConfigApplyConfiguration {
	b.TagSeparator = &value
	return b
}

// WithNodeMeta puts the entries into the NodeMeta field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the NodeMeta field,
// overwriting an existing map entries in NodeMeta field with the same key.
func (b *ConsulSDConfigApplyConfiguration) WithNodeMeta(entries map[string]string) *ConsulSDConfigApplyConfiguration {
	if b.NodeMeta == nil && len(entries) > 0 {
		b.NodeMeta = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.NodeMeta[k] = v
	}
	return b
}

// WithFilter sets the Filter field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Filter field is set to the value of the last call.
func (b *ConsulSDConfigApplyConfiguration) WithFilter(value string) *ConsulSDConfigApplyConfiguration {
	b.Filter = &value
	return b
}

// WithAllowStale sets the AllowStale field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AllowStale field is set to the value of the last call.
func (b *ConsulSDConfigApplyConfiguration) WithAllowStale(value bool) *ConsulSDConfigApplyConfiguration {
	b.AllowStale = &value
	return b
}

// WithRefreshInterval sets the RefreshInterval field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RefreshInterval field is set to the value of the last call.
func (b *ConsulSDConfigApplyConfiguration) WithRefreshInterval(value monitoringv1.Duration) *ConsulSDConfigApplyConfiguration {
	b.RefreshInterval = &value
	return b
}

// WithBasicAuth sets the BasicAuth field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the BasicAuth field is set to the value of the last call.
func (b *ConsulSDConfigApplyConfiguration) WithBasicAuth(value *applyconfigurationmonitoringv1.BasicAuthApplyConfiguration) *ConsulSDConfigApplyConfiguration {
	b.BasicAuth = value
	return b
}

// WithAuthorization sets the Authorization field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Authorization field is set to the value of the last call.
func (b *ConsulSDConfigApplyConfiguration) WithAuthorization(value *applyconfigurationmonitoringv1.SafeAuthorizationApplyConfiguration) *ConsulSDConfigApplyConfiguration {
	b.Authorization = value
	return b
}

// WithOAuth2 sets the OAuth2 field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the OAuth2 field is set to the value of the last call.
func (b *ConsulSDConfigApplyConfiguration) WithOAuth2(value *applyconfigurationmonitoringv1.OAuth2ApplyConfiguration) *ConsulSDConfigApplyConfiguration {
	b.OAuth2 = value
	return b
}

// WithProxyURL sets the ProxyURL field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ProxyURL field is set to the value of the last call.
func (b *ConsulSDConfigApplyConfiguration) WithProxyURL(value string) *ConsulSDConfigApplyConfiguration {
	b.ProxyConfigApplyConfiguration.ProxyURL = &value
	return b
}

// WithNoProxy sets the NoProxy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NoProxy field is set to the value of the last call.
func (b *ConsulSDConfigApplyConfiguration) WithNoProxy(value string) *ConsulSDConfigApplyConfiguration {
	b.ProxyConfigApplyConfiguration.NoProxy = &value
	return b
}

// WithProxyFromEnvironment sets the ProxyFromEnvironment field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ProxyFromEnvironment field is set to the value of the last call.
func (b *ConsulSDConfigApplyConfiguration) WithProxyFromEnvironment(value bool) *ConsulSDConfigApplyConfiguration {
	b.ProxyConfigApplyConfiguration.ProxyFromEnvironment = &value
	return b
}

// WithProxyConnectHeader puts the entries into the ProxyConnectHeader field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the ProxyConnectHeader field,
// overwriting an existing map entries in ProxyConnectHeader field with the same key.
func (b *ConsulSDConfigApplyConfiguration) WithProxyConnectHeader(entries map[string][]v1.SecretKeySelector) *ConsulSDConfigApplyConfiguration {
	if b.ProxyConfigApplyConfiguration.ProxyConnectHeader == nil && len(entries) > 0 {
		b.ProxyConfigApplyConfiguration.ProxyConnectHeader = make(map[string][]v1.SecretKeySelector, len(entries))
	}
	for k, v := range entries {
		b.ProxyConfigApplyConfiguration.ProxyConnectHeader[k] = v
	}
	return b
}

// WithFollowRedirects sets the FollowRedirects field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the FollowRedirects field is set to the value of the last call.
func (b *ConsulSDConfigApplyConfiguration) WithFollowRedirects(value bool) *ConsulSDConfigApplyConfiguration {
	b.FollowRedirects = &value
	return b
}

// WithEnableHttp2 sets the EnableHttp2 field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the EnableHttp2 field is set to the value of the last call.
func (b *ConsulSDConfigApplyConfiguration) WithEnableHttp2(value bool) *ConsulSDConfigApplyConfiguration {
	b.EnableHttp2 = &value
	return b
}

// WithTLSConfig sets the TLSConfig field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TLSConfig field is set to the value of the last call.
func (b *ConsulSDConfigApplyConfiguration) WithTLSConfig(value *applyconfigurationmonitoringv1.SafeTLSConfigApplyConfiguration) *ConsulSDConfigApplyConfiguration {
	b.TLSConfig = value
	return b
}
