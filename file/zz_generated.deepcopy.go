//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021 Kong Inc.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package file

import (
	kong "github.com/kong/go-kong/kong"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Content) DeepCopyInto(out *Content) {
	*out = *in
	if in.Transform != nil {
		in, out := &in.Transform, &out.Transform
		*out = new(bool)
		**out = **in
	}
	if in.Info != nil {
		in, out := &in.Info, &out.Info
		*out = new(Info)
		(*in).DeepCopyInto(*out)
	}
	if in.Konnect != nil {
		in, out := &in.Konnect, &out.Konnect
		*out = new(Konnect)
		**out = **in
	}
	if in.Services != nil {
		in, out := &in.Services, &out.Services
		*out = make([]FService, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Routes != nil {
		in, out := &in.Routes, &out.Routes
		*out = make([]FRoute, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Consumers != nil {
		in, out := &in.Consumers, &out.Consumers
		*out = make([]FConsumer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ConsumerGroups != nil {
		in, out := &in.ConsumerGroups, &out.ConsumerGroups
		*out = make([]FConsumerGroupObject, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Plugins != nil {
		in, out := &in.Plugins, &out.Plugins
		*out = make([]FPlugin, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Upstreams != nil {
		in, out := &in.Upstreams, &out.Upstreams
		*out = make([]FUpstream, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Certificates != nil {
		in, out := &in.Certificates, &out.Certificates
		*out = make([]FCertificate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CACertificates != nil {
		in, out := &in.CACertificates, &out.CACertificates
		*out = make([]FCACertificate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RBACRoles != nil {
		in, out := &in.RBACRoles, &out.RBACRoles
		*out = make([]FRBACRole, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PluginConfigs != nil {
		in, out := &in.PluginConfigs, &out.PluginConfigs
		*out = make(map[string]kong.Configuration, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	if in.ServicePackages != nil {
		in, out := &in.ServicePackages, &out.ServicePackages
		*out = make([]FServicePackage, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Content.
func (in *Content) DeepCopy() *Content {
	if in == nil {
		return nil
	}
	out := new(Content)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FCACertificate) DeepCopyInto(out *FCACertificate) {
	*out = *in
	in.CACertificate.DeepCopyInto(&out.CACertificate)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FCACertificate.
func (in *FCACertificate) DeepCopy() *FCACertificate {
	if in == nil {
		return nil
	}
	out := new(FCACertificate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FCertificate) DeepCopyInto(out *FCertificate) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Cert != nil {
		in, out := &in.Cert, &out.Cert
		*out = new(string)
		**out = **in
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = new(int64)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.SNIs != nil {
		in, out := &in.SNIs, &out.SNIs
		*out = make([]kong.SNI, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FCertificate.
func (in *FCertificate) DeepCopy() *FCertificate {
	if in == nil {
		return nil
	}
	out := new(FCertificate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FConsumer) DeepCopyInto(out *FConsumer) {
	*out = *in
	in.Consumer.DeepCopyInto(&out.Consumer)
	if in.Plugins != nil {
		in, out := &in.Plugins, &out.Plugins
		*out = make([]*FPlugin, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(FPlugin)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.KeyAuths != nil {
		in, out := &in.KeyAuths, &out.KeyAuths
		*out = make([]*kong.KeyAuth, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(kong.KeyAuth)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.HMACAuths != nil {
		in, out := &in.HMACAuths, &out.HMACAuths
		*out = make([]*kong.HMACAuth, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(kong.HMACAuth)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.JWTAuths != nil {
		in, out := &in.JWTAuths, &out.JWTAuths
		*out = make([]*kong.JWTAuth, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(kong.JWTAuth)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.BasicAuths != nil {
		in, out := &in.BasicAuths, &out.BasicAuths
		*out = make([]*kong.BasicAuth, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(kong.BasicAuth)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Oauth2Creds != nil {
		in, out := &in.Oauth2Creds, &out.Oauth2Creds
		*out = make([]*kong.Oauth2Credential, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(kong.Oauth2Credential)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.ACLGroups != nil {
		in, out := &in.ACLGroups, &out.ACLGroups
		*out = make([]*kong.ACLGroup, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(kong.ACLGroup)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.MTLSAuths != nil {
		in, out := &in.MTLSAuths, &out.MTLSAuths
		*out = make([]*kong.MTLSAuth, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(kong.MTLSAuth)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FConsumer.
func (in *FConsumer) DeepCopy() *FConsumer {
	if in == nil {
		return nil
	}
	out := new(FConsumer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FConsumerGroupObject) DeepCopyInto(out *FConsumerGroupObject) {
	*out = *in
	in.ConsumerGroup.DeepCopyInto(&out.ConsumerGroup)
	if in.Consumers != nil {
		in, out := &in.Consumers, &out.Consumers
		*out = make([]*kong.Consumer, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(kong.Consumer)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FConsumerGroupObject.
func (in *FConsumerGroupObject) DeepCopy() *FConsumerGroupObject {
	if in == nil {
		return nil
	}
	out := new(FConsumerGroupObject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FDocument) DeepCopyInto(out *FDocument) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
	if in.Published != nil {
		in, out := &in.Published, &out.Published
		*out = new(bool)
		**out = **in
	}
	if in.Content != nil {
		in, out := &in.Content, &out.Content
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FDocument.
func (in *FDocument) DeepCopy() *FDocument {
	if in == nil {
		return nil
	}
	out := new(FDocument)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FPlugin) DeepCopyInto(out *FPlugin) {
	*out = *in
	in.Plugin.DeepCopyInto(&out.Plugin)
	if in.ConfigSource != nil {
		in, out := &in.ConfigSource, &out.ConfigSource
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FPlugin.
func (in *FPlugin) DeepCopy() *FPlugin {
	if in == nil {
		return nil
	}
	out := new(FPlugin)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FRBACEndpointPermission) DeepCopyInto(out *FRBACEndpointPermission) {
	*out = *in
	in.RBACEndpointPermission.DeepCopyInto(&out.RBACEndpointPermission)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FRBACEndpointPermission.
func (in *FRBACEndpointPermission) DeepCopy() *FRBACEndpointPermission {
	if in == nil {
		return nil
	}
	out := new(FRBACEndpointPermission)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FRBACRole) DeepCopyInto(out *FRBACRole) {
	*out = *in
	in.RBACRole.DeepCopyInto(&out.RBACRole)
	if in.EndpointPermissions != nil {
		in, out := &in.EndpointPermissions, &out.EndpointPermissions
		*out = make([]*FRBACEndpointPermission, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(FRBACEndpointPermission)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FRBACRole.
func (in *FRBACRole) DeepCopy() *FRBACRole {
	if in == nil {
		return nil
	}
	out := new(FRBACRole)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FRoute) DeepCopyInto(out *FRoute) {
	*out = *in
	in.Route.DeepCopyInto(&out.Route)
	if in.Plugins != nil {
		in, out := &in.Plugins, &out.Plugins
		*out = make([]*FPlugin, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(FPlugin)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FRoute.
func (in *FRoute) DeepCopy() *FRoute {
	if in == nil {
		return nil
	}
	out := new(FRoute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FService) DeepCopyInto(out *FService) {
	*out = *in
	in.Service.DeepCopyInto(&out.Service)
	if in.Routes != nil {
		in, out := &in.Routes, &out.Routes
		*out = make([]*FRoute, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(FRoute)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Plugins != nil {
		in, out := &in.Plugins, &out.Plugins
		*out = make([]*FPlugin, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(FPlugin)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.URL != nil {
		in, out := &in.URL, &out.URL
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FService.
func (in *FService) DeepCopy() *FService {
	if in == nil {
		return nil
	}
	out := new(FService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FServicePackage) DeepCopyInto(out *FServicePackage) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Versions != nil {
		in, out := &in.Versions, &out.Versions
		*out = make([]FServiceVersion, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Document != nil {
		in, out := &in.Document, &out.Document
		*out = new(FDocument)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FServicePackage.
func (in *FServicePackage) DeepCopy() *FServicePackage {
	if in == nil {
		return nil
	}
	out := new(FServicePackage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FServiceVersion) DeepCopyInto(out *FServiceVersion) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
	if in.Implementation != nil {
		in, out := &in.Implementation, &out.Implementation
		*out = new(Implementation)
		(*in).DeepCopyInto(*out)
	}
	if in.Document != nil {
		in, out := &in.Document, &out.Document
		*out = new(FDocument)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FServiceVersion.
func (in *FServiceVersion) DeepCopy() *FServiceVersion {
	if in == nil {
		return nil
	}
	out := new(FServiceVersion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FTarget) DeepCopyInto(out *FTarget) {
	*out = *in
	in.Target.DeepCopyInto(&out.Target)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FTarget.
func (in *FTarget) DeepCopy() *FTarget {
	if in == nil {
		return nil
	}
	out := new(FTarget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FUpstream) DeepCopyInto(out *FUpstream) {
	*out = *in
	in.Upstream.DeepCopyInto(&out.Upstream)
	if in.Targets != nil {
		in, out := &in.Targets, &out.Targets
		*out = make([]*FTarget, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(FTarget)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FUpstream.
func (in *FUpstream) DeepCopy() *FUpstream {
	if in == nil {
		return nil
	}
	out := new(FUpstream)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Implementation) DeepCopyInto(out *Implementation) {
	*out = *in
	if in.Kong != nil {
		in, out := &in.Kong, &out.Kong
		*out = new(Kong)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Implementation.
func (in *Implementation) DeepCopy() *Implementation {
	if in == nil {
		return nil
	}
	out := new(Implementation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Info) DeepCopyInto(out *Info) {
	*out = *in
	if in.SelectorTags != nil {
		in, out := &in.SelectorTags, &out.SelectorTags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.Defaults.DeepCopyInto(&out.Defaults)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Info.
func (in *Info) DeepCopy() *Info {
	if in == nil {
		return nil
	}
	out := new(Info)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Kong) DeepCopyInto(out *Kong) {
	*out = *in
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = new(FService)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Kong.
func (in *Kong) DeepCopy() *Kong {
	if in == nil {
		return nil
	}
	out := new(Kong)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KongDefaults) DeepCopyInto(out *KongDefaults) {
	*out = *in
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = new(kong.Service)
		(*in).DeepCopyInto(*out)
	}
	if in.Route != nil {
		in, out := &in.Route, &out.Route
		*out = new(kong.Route)
		(*in).DeepCopyInto(*out)
	}
	if in.Upstream != nil {
		in, out := &in.Upstream, &out.Upstream
		*out = new(kong.Upstream)
		(*in).DeepCopyInto(*out)
	}
	if in.Target != nil {
		in, out := &in.Target, &out.Target
		*out = new(kong.Target)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KongDefaults.
func (in *KongDefaults) DeepCopy() *KongDefaults {
	if in == nil {
		return nil
	}
	out := new(KongDefaults)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Konnect) DeepCopyInto(out *Konnect) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Konnect.
func (in *Konnect) DeepCopy() *Konnect {
	if in == nil {
		return nil
	}
	out := new(Konnect)
	in.DeepCopyInto(out)
	return out
}
