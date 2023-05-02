//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlockDeviceObservation) DeepCopyInto(out *BlockDeviceObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlockDeviceObservation.
func (in *BlockDeviceObservation) DeepCopy() *BlockDeviceObservation {
	if in == nil {
		return nil
	}
	out := new(BlockDeviceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlockDeviceParameters) DeepCopyInto(out *BlockDeviceParameters) {
	*out = *in
	if in.BootIndex != nil {
		in, out := &in.BootIndex, &out.BootIndex
		*out = new(float64)
		**out = **in
	}
	if in.DeleteOnTermination != nil {
		in, out := &in.DeleteOnTermination, &out.DeleteOnTermination
		*out = new(bool)
		**out = **in
	}
	if in.DestinationType != nil {
		in, out := &in.DestinationType, &out.DestinationType
		*out = new(string)
		**out = **in
	}
	if in.DeviceType != nil {
		in, out := &in.DeviceType, &out.DeviceType
		*out = new(string)
		**out = **in
	}
	if in.DiskBus != nil {
		in, out := &in.DiskBus, &out.DiskBus
		*out = new(string)
		**out = **in
	}
	if in.GuestFormat != nil {
		in, out := &in.GuestFormat, &out.GuestFormat
		*out = new(string)
		**out = **in
	}
	if in.SourceType != nil {
		in, out := &in.SourceType, &out.SourceType
		*out = new(string)
		**out = **in
	}
	if in.UUID != nil {
		in, out := &in.UUID, &out.UUID
		*out = new(string)
		**out = **in
	}
	if in.VolumeSize != nil {
		in, out := &in.VolumeSize, &out.VolumeSize
		*out = new(float64)
		**out = **in
	}
	if in.VolumeType != nil {
		in, out := &in.VolumeType, &out.VolumeType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlockDeviceParameters.
func (in *BlockDeviceParameters) DeepCopy() *BlockDeviceParameters {
	if in == nil {
		return nil
	}
	out := new(BlockDeviceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceV2) DeepCopyInto(out *InstanceV2) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceV2.
func (in *InstanceV2) DeepCopy() *InstanceV2 {
	if in == nil {
		return nil
	}
	out := new(InstanceV2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InstanceV2) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceV2List) DeepCopyInto(out *InstanceV2List) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]InstanceV2, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceV2List.
func (in *InstanceV2List) DeepCopy() *InstanceV2List {
	if in == nil {
		return nil
	}
	out := new(InstanceV2List)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InstanceV2List) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceV2Observation) DeepCopyInto(out *InstanceV2Observation) {
	*out = *in
	if in.AllMetadata != nil {
		in, out := &in.AllMetadata, &out.AllMetadata
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.AllTags != nil {
		in, out := &in.AllTags, &out.AllTags
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Created != nil {
		in, out := &in.Created, &out.Created
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Network != nil {
		in, out := &in.Network, &out.Network
		*out = make([]NetworkObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Updated != nil {
		in, out := &in.Updated, &out.Updated
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceV2Observation.
func (in *InstanceV2Observation) DeepCopy() *InstanceV2Observation {
	if in == nil {
		return nil
	}
	out := new(InstanceV2Observation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceV2Parameters) DeepCopyInto(out *InstanceV2Parameters) {
	*out = *in
	if in.AccessIPV4 != nil {
		in, out := &in.AccessIPV4, &out.AccessIPV4
		*out = new(string)
		**out = **in
	}
	if in.AccessIPV6 != nil {
		in, out := &in.AccessIPV6, &out.AccessIPV6
		*out = new(string)
		**out = **in
	}
	if in.AdminPassSecretRef != nil {
		in, out := &in.AdminPassSecretRef, &out.AdminPassSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.AvailabilityZone != nil {
		in, out := &in.AvailabilityZone, &out.AvailabilityZone
		*out = new(string)
		**out = **in
	}
	if in.AvailabilityZoneHints != nil {
		in, out := &in.AvailabilityZoneHints, &out.AvailabilityZoneHints
		*out = new(string)
		**out = **in
	}
	if in.BlockDevice != nil {
		in, out := &in.BlockDevice, &out.BlockDevice
		*out = make([]BlockDeviceParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ConfigDrive != nil {
		in, out := &in.ConfigDrive, &out.ConfigDrive
		*out = new(bool)
		**out = **in
	}
	if in.FlavorID != nil {
		in, out := &in.FlavorID, &out.FlavorID
		*out = new(string)
		**out = **in
	}
	if in.FlavorName != nil {
		in, out := &in.FlavorName, &out.FlavorName
		*out = new(string)
		**out = **in
	}
	if in.FloatingIP != nil {
		in, out := &in.FloatingIP, &out.FloatingIP
		*out = new(string)
		**out = **in
	}
	if in.ForceDelete != nil {
		in, out := &in.ForceDelete, &out.ForceDelete
		*out = new(bool)
		**out = **in
	}
	if in.ImageID != nil {
		in, out := &in.ImageID, &out.ImageID
		*out = new(string)
		**out = **in
	}
	if in.ImageName != nil {
		in, out := &in.ImageName, &out.ImageName
		*out = new(string)
		**out = **in
	}
	if in.KeyPair != nil {
		in, out := &in.KeyPair, &out.KeyPair
		*out = new(string)
		**out = **in
	}
	if in.KeyPairRef != nil {
		in, out := &in.KeyPairRef, &out.KeyPairRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.KeyPairSelector != nil {
		in, out := &in.KeyPairSelector, &out.KeyPairSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Network != nil {
		in, out := &in.Network, &out.Network
		*out = make([]NetworkParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NetworkMode != nil {
		in, out := &in.NetworkMode, &out.NetworkMode
		*out = new(string)
		**out = **in
	}
	if in.Personality != nil {
		in, out := &in.Personality, &out.Personality
		*out = make([]PersonalityParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PowerState != nil {
		in, out := &in.PowerState, &out.PowerState
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.SchedulerHints != nil {
		in, out := &in.SchedulerHints, &out.SchedulerHints
		*out = make([]SchedulerHintsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SecurityGroups != nil {
		in, out := &in.SecurityGroups, &out.SecurityGroups
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.StopBeforeDestroy != nil {
		in, out := &in.StopBeforeDestroy, &out.StopBeforeDestroy
		*out = new(bool)
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
	if in.UserData != nil {
		in, out := &in.UserData, &out.UserData
		*out = new(string)
		**out = **in
	}
	if in.VendorOptions != nil {
		in, out := &in.VendorOptions, &out.VendorOptions
		*out = make([]VendorOptionsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Volume != nil {
		in, out := &in.Volume, &out.Volume
		*out = make([]VolumeParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceV2Parameters.
func (in *InstanceV2Parameters) DeepCopy() *InstanceV2Parameters {
	if in == nil {
		return nil
	}
	out := new(InstanceV2Parameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceV2Spec) DeepCopyInto(out *InstanceV2Spec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceV2Spec.
func (in *InstanceV2Spec) DeepCopy() *InstanceV2Spec {
	if in == nil {
		return nil
	}
	out := new(InstanceV2Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceV2Status) DeepCopyInto(out *InstanceV2Status) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceV2Status.
func (in *InstanceV2Status) DeepCopy() *InstanceV2Status {
	if in == nil {
		return nil
	}
	out := new(InstanceV2Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeypairV2) DeepCopyInto(out *KeypairV2) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeypairV2.
func (in *KeypairV2) DeepCopy() *KeypairV2 {
	if in == nil {
		return nil
	}
	out := new(KeypairV2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeypairV2) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeypairV2List) DeepCopyInto(out *KeypairV2List) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KeypairV2, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeypairV2List.
func (in *KeypairV2List) DeepCopy() *KeypairV2List {
	if in == nil {
		return nil
	}
	out := new(KeypairV2List)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeypairV2List) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeypairV2Observation) DeepCopyInto(out *KeypairV2Observation) {
	*out = *in
	if in.Fingerprint != nil {
		in, out := &in.Fingerprint, &out.Fingerprint
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.PrivateKey != nil {
		in, out := &in.PrivateKey, &out.PrivateKey
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeypairV2Observation.
func (in *KeypairV2Observation) DeepCopy() *KeypairV2Observation {
	if in == nil {
		return nil
	}
	out := new(KeypairV2Observation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeypairV2Parameters) DeepCopyInto(out *KeypairV2Parameters) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PublicKey != nil {
		in, out := &in.PublicKey, &out.PublicKey
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.UserID != nil {
		in, out := &in.UserID, &out.UserID
		*out = new(string)
		**out = **in
	}
	if in.ValueSpecs != nil {
		in, out := &in.ValueSpecs, &out.ValueSpecs
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeypairV2Parameters.
func (in *KeypairV2Parameters) DeepCopy() *KeypairV2Parameters {
	if in == nil {
		return nil
	}
	out := new(KeypairV2Parameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeypairV2Spec) DeepCopyInto(out *KeypairV2Spec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeypairV2Spec.
func (in *KeypairV2Spec) DeepCopy() *KeypairV2Spec {
	if in == nil {
		return nil
	}
	out := new(KeypairV2Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeypairV2Status) DeepCopyInto(out *KeypairV2Status) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeypairV2Status.
func (in *KeypairV2Status) DeepCopy() *KeypairV2Status {
	if in == nil {
		return nil
	}
	out := new(KeypairV2Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkObservation) DeepCopyInto(out *NetworkObservation) {
	*out = *in
	if in.Mac != nil {
		in, out := &in.Mac, &out.Mac
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkObservation.
func (in *NetworkObservation) DeepCopy() *NetworkObservation {
	if in == nil {
		return nil
	}
	out := new(NetworkObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkParameters) DeepCopyInto(out *NetworkParameters) {
	*out = *in
	if in.AccessNetwork != nil {
		in, out := &in.AccessNetwork, &out.AccessNetwork
		*out = new(bool)
		**out = **in
	}
	if in.FixedIPV4 != nil {
		in, out := &in.FixedIPV4, &out.FixedIPV4
		*out = new(string)
		**out = **in
	}
	if in.FixedIPV6 != nil {
		in, out := &in.FixedIPV6, &out.FixedIPV6
		*out = new(string)
		**out = **in
	}
	if in.FloatingIP != nil {
		in, out := &in.FloatingIP, &out.FloatingIP
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(string)
		**out = **in
	}
	if in.UUID != nil {
		in, out := &in.UUID, &out.UUID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkParameters.
func (in *NetworkParameters) DeepCopy() *NetworkParameters {
	if in == nil {
		return nil
	}
	out := new(NetworkParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PersonalityObservation) DeepCopyInto(out *PersonalityObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PersonalityObservation.
func (in *PersonalityObservation) DeepCopy() *PersonalityObservation {
	if in == nil {
		return nil
	}
	out := new(PersonalityObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PersonalityParameters) DeepCopyInto(out *PersonalityParameters) {
	*out = *in
	if in.Content != nil {
		in, out := &in.Content, &out.Content
		*out = new(string)
		**out = **in
	}
	if in.File != nil {
		in, out := &in.File, &out.File
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PersonalityParameters.
func (in *PersonalityParameters) DeepCopy() *PersonalityParameters {
	if in == nil {
		return nil
	}
	out := new(PersonalityParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchedulerHintsObservation) DeepCopyInto(out *SchedulerHintsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchedulerHintsObservation.
func (in *SchedulerHintsObservation) DeepCopy() *SchedulerHintsObservation {
	if in == nil {
		return nil
	}
	out := new(SchedulerHintsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchedulerHintsParameters) DeepCopyInto(out *SchedulerHintsParameters) {
	*out = *in
	if in.AdditionalProperties != nil {
		in, out := &in.AdditionalProperties, &out.AdditionalProperties
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.BuildNearHostIP != nil {
		in, out := &in.BuildNearHostIP, &out.BuildNearHostIP
		*out = new(string)
		**out = **in
	}
	if in.DifferentCell != nil {
		in, out := &in.DifferentCell, &out.DifferentCell
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.DifferentHost != nil {
		in, out := &in.DifferentHost, &out.DifferentHost
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Group != nil {
		in, out := &in.Group, &out.Group
		*out = new(string)
		**out = **in
	}
	if in.Query != nil {
		in, out := &in.Query, &out.Query
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.SameHost != nil {
		in, out := &in.SameHost, &out.SameHost
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.TargetCell != nil {
		in, out := &in.TargetCell, &out.TargetCell
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchedulerHintsParameters.
func (in *SchedulerHintsParameters) DeepCopy() *SchedulerHintsParameters {
	if in == nil {
		return nil
	}
	out := new(SchedulerHintsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VendorOptionsObservation) DeepCopyInto(out *VendorOptionsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VendorOptionsObservation.
func (in *VendorOptionsObservation) DeepCopy() *VendorOptionsObservation {
	if in == nil {
		return nil
	}
	out := new(VendorOptionsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VendorOptionsParameters) DeepCopyInto(out *VendorOptionsParameters) {
	*out = *in
	if in.DetachPortsBeforeDestroy != nil {
		in, out := &in.DetachPortsBeforeDestroy, &out.DetachPortsBeforeDestroy
		*out = new(bool)
		**out = **in
	}
	if in.IgnoreResizeConfirmation != nil {
		in, out := &in.IgnoreResizeConfirmation, &out.IgnoreResizeConfirmation
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VendorOptionsParameters.
func (in *VendorOptionsParameters) DeepCopy() *VendorOptionsParameters {
	if in == nil {
		return nil
	}
	out := new(VendorOptionsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeObservation) DeepCopyInto(out *VolumeObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeObservation.
func (in *VolumeObservation) DeepCopy() *VolumeObservation {
	if in == nil {
		return nil
	}
	out := new(VolumeObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeParameters) DeepCopyInto(out *VolumeParameters) {
	*out = *in
	if in.Device != nil {
		in, out := &in.Device, &out.Device
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.VolumeID != nil {
		in, out := &in.VolumeID, &out.VolumeID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeParameters.
func (in *VolumeParameters) DeepCopy() *VolumeParameters {
	if in == nil {
		return nil
	}
	out := new(VolumeParameters)
	in.DeepCopyInto(out)
	return out
}
