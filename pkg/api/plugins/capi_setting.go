// Code generated by go-swagger; DO NOT EDIT.

//
//   Copyright (c) 2022 Intel Corporation.
//
//   SPDX-License-Identifier: Apache-2.0
//
//
//

package plugins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CapiSetting capi setting
//
// swagger:model Capi-setting
type CapiSetting struct {

	// byoh config
	ByohConfig *CapiSettingByohConfig `json:"Byoh_config,omitempty"`

	// c r i
	CRI *CapiSettingCRI `json:"CRI,omitempty"`

	// infra provider
	InfraProvider *CapiSettingInfraProvider `json:"Infra_provider,omitempty"`

	// ironic config
	IronicConfig *CapiSettingIronicConfig `json:"Ironic_config,omitempty"`

	// provider
	Provider string `json:"Provider,omitempty"`

	// registry
	Registry *CapiSettingRegistry `json:"Registry,omitempty"`
}

// Validate validates this capi setting
func (m *CapiSetting) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateByohConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCRI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInfraProvider(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIronicConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegistry(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CapiSetting) validateByohConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.ByohConfig) { // not required
		return nil
	}

	if m.ByohConfig != nil {
		if err := m.ByohConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Byoh_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Byoh_config")
			}
			return err
		}
	}

	return nil
}

func (m *CapiSetting) validateCRI(formats strfmt.Registry) error {
	if swag.IsZero(m.CRI) { // not required
		return nil
	}

	if m.CRI != nil {
		if err := m.CRI.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("CRI")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("CRI")
			}
			return err
		}
	}

	return nil
}

func (m *CapiSetting) validateInfraProvider(formats strfmt.Registry) error {
	if swag.IsZero(m.InfraProvider) { // not required
		return nil
	}

	if m.InfraProvider != nil {
		if err := m.InfraProvider.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Infra_provider")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Infra_provider")
			}
			return err
		}
	}

	return nil
}

func (m *CapiSetting) validateIronicConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.IronicConfig) { // not required
		return nil
	}

	if m.IronicConfig != nil {
		if err := m.IronicConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Ironic_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Ironic_config")
			}
			return err
		}
	}

	return nil
}

func (m *CapiSetting) validateRegistry(formats strfmt.Registry) error {
	if swag.IsZero(m.Registry) { // not required
		return nil
	}

	if m.Registry != nil {
		if err := m.Registry.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Registry")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Registry")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this capi setting based on the context it is used
func (m *CapiSetting) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateByohConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCRI(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInfraProvider(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIronicConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRegistry(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CapiSetting) contextValidateByohConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.ByohConfig != nil {
		if err := m.ByohConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Byoh_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Byoh_config")
			}
			return err
		}
	}

	return nil
}

func (m *CapiSetting) contextValidateCRI(ctx context.Context, formats strfmt.Registry) error {

	if m.CRI != nil {
		if err := m.CRI.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("CRI")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("CRI")
			}
			return err
		}
	}

	return nil
}

func (m *CapiSetting) contextValidateInfraProvider(ctx context.Context, formats strfmt.Registry) error {

	if m.InfraProvider != nil {
		if err := m.InfraProvider.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Infra_provider")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Infra_provider")
			}
			return err
		}
	}

	return nil
}

func (m *CapiSetting) contextValidateIronicConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.IronicConfig != nil {
		if err := m.IronicConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Ironic_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Ironic_config")
			}
			return err
		}
	}

	return nil
}

func (m *CapiSetting) contextValidateRegistry(ctx context.Context, formats strfmt.Registry) error {

	if m.Registry != nil {
		if err := m.Registry.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Registry")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Registry")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CapiSetting) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CapiSetting) UnmarshalBinary(b []byte) error {
	var res CapiSetting
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CapiSettingByohConfig capi setting byoh config
//
// swagger:model CapiSettingByohConfig
type CapiSettingByohConfig struct {

	// bundle image
	// Pattern: ^[a-zA-Z_$][a-zA-Z_.\-$0-9]*$
	BundleImage string `json:"Bundle_image,omitempty"`

	// bundle registry
	// Pattern: ^[a-zA-Z_$][a-zA-Z_.\-$0-9]*$
	BundleRegistry string `json:"Bundle_registry,omitempty"`

	// bundle tag
	// Pattern: ^[a-zA-Z_$][a-zA-Z_.\-$0-9]*$
	BundleTag string `json:"Bundle_tag,omitempty"`

	// download bin url
	// Pattern: (?:(?:https?|http|ftp|file|oci)://|www.|ftp.)(?:([-A-Z0-9+&@#/%=~_|$?!:,.]*)|[-A-Z0-9+&@#/%=~_|$?!:,.])*(?:([-A-Z0-9+&@#/%=~_|$?!:,.]*)|[A-Z0-9+&@#/%=~_|$])
	DownloadBinURL string `json:"Download_bin_url,omitempty"`

	// host agent bin url
	// Pattern: (?:(?:https?|http|ftp|file|oci)://|www.|ftp.)(?:([-A-Z0-9+&@#/%=~_|$?!:,.]*)|[-A-Z0-9+&@#/%=~_|$?!:,.])*(?:([-A-Z0-9+&@#/%=~_|$?!:,.]*)|[A-Z0-9+&@#/%=~_|$])
	HostAgentBinURL string `json:"Host_agent_bin_url,omitempty"`
}

// Validate validates this capi setting byoh config
func (m *CapiSettingByohConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBundleImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBundleRegistry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBundleTag(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDownloadBinURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostAgentBinURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CapiSettingByohConfig) validateBundleImage(formats strfmt.Registry) error {
	if swag.IsZero(m.BundleImage) { // not required
		return nil
	}

	if err := validate.Pattern("Byoh_config"+"."+"Bundle_image", "body", m.BundleImage, `^[a-zA-Z_$][a-zA-Z_.\-$0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *CapiSettingByohConfig) validateBundleRegistry(formats strfmt.Registry) error {
	if swag.IsZero(m.BundleRegistry) { // not required
		return nil
	}

	if err := validate.Pattern("Byoh_config"+"."+"Bundle_registry", "body", m.BundleRegistry, `^[a-zA-Z_$][a-zA-Z_.\-$0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *CapiSettingByohConfig) validateBundleTag(formats strfmt.Registry) error {
	if swag.IsZero(m.BundleTag) { // not required
		return nil
	}

	if err := validate.Pattern("Byoh_config"+"."+"Bundle_tag", "body", m.BundleTag, `^[a-zA-Z_$][a-zA-Z_.\-$0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *CapiSettingByohConfig) validateDownloadBinURL(formats strfmt.Registry) error {
	if swag.IsZero(m.DownloadBinURL) { // not required
		return nil
	}

	if err := validate.Pattern("Byoh_config"+"."+"Download_bin_url", "body", m.DownloadBinURL, `(?:(?:https?|http|ftp|file|oci)://|www.|ftp.)(?:([-A-Z0-9+&@#/%=~_|$?!:,.]*)|[-A-Z0-9+&@#/%=~_|$?!:,.])*(?:([-A-Z0-9+&@#/%=~_|$?!:,.]*)|[A-Z0-9+&@#/%=~_|$])`); err != nil {
		return err
	}

	return nil
}

func (m *CapiSettingByohConfig) validateHostAgentBinURL(formats strfmt.Registry) error {
	if swag.IsZero(m.HostAgentBinURL) { // not required
		return nil
	}

	if err := validate.Pattern("Byoh_config"+"."+"Host_agent_bin_url", "body", m.HostAgentBinURL, `(?:(?:https?|http|ftp|file|oci)://|www.|ftp.)(?:([-A-Z0-9+&@#/%=~_|$?!:,.]*)|[-A-Z0-9+&@#/%=~_|$?!:,.])*(?:([-A-Z0-9+&@#/%=~_|$?!:,.]*)|[A-Z0-9+&@#/%=~_|$])`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this capi setting byoh config based on context it is used
func (m *CapiSettingByohConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CapiSettingByohConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CapiSettingByohConfig) UnmarshalBinary(b []byte) error {
	var res CapiSettingByohConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CapiSettingCRI capi setting c r i
//
// swagger:model CapiSettingCRI
type CapiSettingCRI struct {

	// bin URL
	BinURL string `json:"BinURL,omitempty"`

	// endpoint
	// Pattern: ^[a-zA-Z.\/][a-zA-Z0-9-_.\/]*$
	Endpoint string `json:"Endpoint,omitempty"`

	// name
	Name string `json:"Name,omitempty"`

	// version
	Version string `json:"Version,omitempty"`
}

// Validate validates this capi setting c r i
func (m *CapiSettingCRI) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndpoint(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CapiSettingCRI) validateEndpoint(formats strfmt.Registry) error {
	if swag.IsZero(m.Endpoint) { // not required
		return nil
	}

	if err := validate.Pattern("CRI"+"."+"Endpoint", "body", m.Endpoint, `^[a-zA-Z.\/][a-zA-Z0-9-_.\/]*$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this capi setting c r i based on context it is used
func (m *CapiSettingCRI) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CapiSettingCRI) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CapiSettingCRI) UnmarshalBinary(b []byte) error {
	var res CapiSettingCRI
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CapiSettingInfraProvider capi setting infra provider
//
// swagger:model CapiSettingInfraProvider
type CapiSettingInfraProvider struct {

	// authorized ssh public key
	// Pattern: ^[a-zA-Z_$][a-zA-Z_.\-$0-9]*$
	AuthorizedSSHPublicKey string `json:"Authorized_ssh_public_key,omitempty"`

	// management cluster kubeconfig
	// Pattern: ^[a-zA-Z_$][a-zA-Z_.\-$0-9]*$
	ManagementClusterKubeconfig string `json:"Management_cluster_kubeconfig,omitempty"`

	// workload cluster control plane num
	WorkloadClusterControlPlaneNum int64 `json:"Workload_cluster_control_plane_num,omitempty"`

	// workload cluster controlplane endpoint
	// Pattern: ^[a-zA-Z_$][a-zA-Z_.\-$0-9]*$
	WorkloadClusterControlplaneEndpoint string `json:"Workload_cluster_controlplane_endpoint,omitempty"`

	// workload cluster name
	// Pattern: ^[a-zA-Z_$][a-zA-Z_.\-$0-9]*$
	WorkloadClusterName string `json:"Workload_cluster_name,omitempty"`

	// workload cluster namespace
	// Pattern: ^[a-zA-Z_$][a-zA-Z_.\-$0-9]*$
	WorkloadClusterNamespace string `json:"Workload_cluster_namespace,omitempty"`

	// workload cluster network
	// Pattern: ^[a-zA-Z_$][a-zA-Z_.\-$0-9]*$
	WorkloadClusterNetwork string `json:"Workload_cluster_network,omitempty"`

	// workload cluster network gateway
	// Pattern: ^[a-zA-Z_$][a-zA-Z_.\-$0-9]*$
	WorkloadClusterNetworkGateway string `json:"Workload_cluster_network_gateway,omitempty"`

	// workload cluster nic name
	// Pattern: ^[a-zA-Z_$][a-zA-Z_.\-$0-9]*$
	WorkloadClusterNicName string `json:"Workload_cluster_nic_name,omitempty"`

	// workload cluster node address end
	// Pattern: ^[a-zA-Z_$][a-zA-Z_.\-$0-9]*$
	WorkloadClusterNodeAddressEnd string `json:"Workload_cluster_node_address_end,omitempty"`

	// workload cluster node address prefix
	// Pattern: ^[a-zA-Z_$][a-zA-Z_.\-$0-9]*$
	WorkloadClusterNodeAddressPrefix string `json:"Workload_cluster_node_address_prefix,omitempty"`

	// workload cluster node address start
	// Pattern: ^[a-zA-Z_$][a-zA-Z_.\-$0-9]*$
	WorkloadClusterNodeAddressStart string `json:"Workload_cluster_node_address_start,omitempty"`

	// workload cluster node username
	// Pattern: ^[a-zA-Z_$][a-zA-Z_.\-$0-9]*$
	WorkloadClusterNodeUsername string `json:"Workload_cluster_node_username,omitempty"`

	// workload cluster worker node num
	WorkloadClusterWorkerNodeNum int64 `json:"Workload_cluster_worker_node_num,omitempty"`
}

// Validate validates this capi setting infra provider
func (m *CapiSettingInfraProvider) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthorizedSSHPublicKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateManagementClusterKubeconfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkloadClusterControlplaneEndpoint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkloadClusterName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkloadClusterNamespace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkloadClusterNetwork(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkloadClusterNetworkGateway(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkloadClusterNicName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkloadClusterNodeAddressEnd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkloadClusterNodeAddressPrefix(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkloadClusterNodeAddressStart(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkloadClusterNodeUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CapiSettingInfraProvider) validateAuthorizedSSHPublicKey(formats strfmt.Registry) error {
	if swag.IsZero(m.AuthorizedSSHPublicKey) { // not required
		return nil
	}

	if err := validate.Pattern("Infra_provider"+"."+"Authorized_ssh_public_key", "body", m.AuthorizedSSHPublicKey, `^[a-zA-Z_$][a-zA-Z_.\-$0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *CapiSettingInfraProvider) validateManagementClusterKubeconfig(formats strfmt.Registry) error {
	if swag.IsZero(m.ManagementClusterKubeconfig) { // not required
		return nil
	}

	if err := validate.Pattern("Infra_provider"+"."+"Management_cluster_kubeconfig", "body", m.ManagementClusterKubeconfig, `^[a-zA-Z_$][a-zA-Z_.\-$0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *CapiSettingInfraProvider) validateWorkloadClusterControlplaneEndpoint(formats strfmt.Registry) error {
	if swag.IsZero(m.WorkloadClusterControlplaneEndpoint) { // not required
		return nil
	}

	if err := validate.Pattern("Infra_provider"+"."+"Workload_cluster_controlplane_endpoint", "body", m.WorkloadClusterControlplaneEndpoint, `^[a-zA-Z_$][a-zA-Z_.\-$0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *CapiSettingInfraProvider) validateWorkloadClusterName(formats strfmt.Registry) error {
	if swag.IsZero(m.WorkloadClusterName) { // not required
		return nil
	}

	if err := validate.Pattern("Infra_provider"+"."+"Workload_cluster_name", "body", m.WorkloadClusterName, `^[a-zA-Z_$][a-zA-Z_.\-$0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *CapiSettingInfraProvider) validateWorkloadClusterNamespace(formats strfmt.Registry) error {
	if swag.IsZero(m.WorkloadClusterNamespace) { // not required
		return nil
	}

	if err := validate.Pattern("Infra_provider"+"."+"Workload_cluster_namespace", "body", m.WorkloadClusterNamespace, `^[a-zA-Z_$][a-zA-Z_.\-$0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *CapiSettingInfraProvider) validateWorkloadClusterNetwork(formats strfmt.Registry) error {
	if swag.IsZero(m.WorkloadClusterNetwork) { // not required
		return nil
	}

	if err := validate.Pattern("Infra_provider"+"."+"Workload_cluster_network", "body", m.WorkloadClusterNetwork, `^[a-zA-Z_$][a-zA-Z_.\-$0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *CapiSettingInfraProvider) validateWorkloadClusterNetworkGateway(formats strfmt.Registry) error {
	if swag.IsZero(m.WorkloadClusterNetworkGateway) { // not required
		return nil
	}

	if err := validate.Pattern("Infra_provider"+"."+"Workload_cluster_network_gateway", "body", m.WorkloadClusterNetworkGateway, `^[a-zA-Z_$][a-zA-Z_.\-$0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *CapiSettingInfraProvider) validateWorkloadClusterNicName(formats strfmt.Registry) error {
	if swag.IsZero(m.WorkloadClusterNicName) { // not required
		return nil
	}

	if err := validate.Pattern("Infra_provider"+"."+"Workload_cluster_nic_name", "body", m.WorkloadClusterNicName, `^[a-zA-Z_$][a-zA-Z_.\-$0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *CapiSettingInfraProvider) validateWorkloadClusterNodeAddressEnd(formats strfmt.Registry) error {
	if swag.IsZero(m.WorkloadClusterNodeAddressEnd) { // not required
		return nil
	}

	if err := validate.Pattern("Infra_provider"+"."+"Workload_cluster_node_address_end", "body", m.WorkloadClusterNodeAddressEnd, `^[a-zA-Z_$][a-zA-Z_.\-$0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *CapiSettingInfraProvider) validateWorkloadClusterNodeAddressPrefix(formats strfmt.Registry) error {
	if swag.IsZero(m.WorkloadClusterNodeAddressPrefix) { // not required
		return nil
	}

	if err := validate.Pattern("Infra_provider"+"."+"Workload_cluster_node_address_prefix", "body", m.WorkloadClusterNodeAddressPrefix, `^[a-zA-Z_$][a-zA-Z_.\-$0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *CapiSettingInfraProvider) validateWorkloadClusterNodeAddressStart(formats strfmt.Registry) error {
	if swag.IsZero(m.WorkloadClusterNodeAddressStart) { // not required
		return nil
	}

	if err := validate.Pattern("Infra_provider"+"."+"Workload_cluster_node_address_start", "body", m.WorkloadClusterNodeAddressStart, `^[a-zA-Z_$][a-zA-Z_.\-$0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *CapiSettingInfraProvider) validateWorkloadClusterNodeUsername(formats strfmt.Registry) error {
	if swag.IsZero(m.WorkloadClusterNodeUsername) { // not required
		return nil
	}

	if err := validate.Pattern("Infra_provider"+"."+"Workload_cluster_node_username", "body", m.WorkloadClusterNodeUsername, `^[a-zA-Z_$][a-zA-Z_.\-$0-9]*$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this capi setting infra provider based on context it is used
func (m *CapiSettingInfraProvider) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CapiSettingInfraProvider) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CapiSettingInfraProvider) UnmarshalBinary(b []byte) error {
	var res CapiSettingInfraProvider
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CapiSettingIronicConfig capi setting ironic config
//
// swagger:model CapiSettingIronicConfig
type CapiSettingIronicConfig struct {

	// ironic dhcp range
	// Pattern: ^[a-zA-Z_$][a-zA-Z_.\-$0-9]*$
	IronicDhcpRange string `json:"Ironic_dhcp_range,omitempty"`

	// ironic http port
	// Pattern: ^[a-zA-Z_$][a-zA-Z_.\-$0-9]*$
	IronicHTTPPort string `json:"Ironic_http_port,omitempty"`

	// ironic os image
	// Pattern: ^[a-zA-Z_$][a-zA-Z_.\-$0-9]*$
	IronicOsImage string `json:"Ironic_os_image,omitempty"`

	// ironic provision ip
	// Pattern: ^[a-zA-Z_$][a-zA-Z_.\-$0-9]*$
	IronicProvisionIP string `json:"Ironic_provision_ip,omitempty"`

	// ironic provision nic
	// Pattern: ^[a-zA-Z_$][a-zA-Z_.\-$0-9]*$
	IronicProvisionNic string `json:"Ironic_provision_nic,omitempty"`
}

// Validate validates this capi setting ironic config
func (m *CapiSettingIronicConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIronicDhcpRange(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIronicHTTPPort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIronicOsImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIronicProvisionIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIronicProvisionNic(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CapiSettingIronicConfig) validateIronicDhcpRange(formats strfmt.Registry) error {
	if swag.IsZero(m.IronicDhcpRange) { // not required
		return nil
	}

	if err := validate.Pattern("Ironic_config"+"."+"Ironic_dhcp_range", "body", m.IronicDhcpRange, `^[a-zA-Z_$][a-zA-Z_.\-$0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *CapiSettingIronicConfig) validateIronicHTTPPort(formats strfmt.Registry) error {
	if swag.IsZero(m.IronicHTTPPort) { // not required
		return nil
	}

	if err := validate.Pattern("Ironic_config"+"."+"Ironic_http_port", "body", m.IronicHTTPPort, `^[a-zA-Z_$][a-zA-Z_.\-$0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *CapiSettingIronicConfig) validateIronicOsImage(formats strfmt.Registry) error {
	if swag.IsZero(m.IronicOsImage) { // not required
		return nil
	}

	if err := validate.Pattern("Ironic_config"+"."+"Ironic_os_image", "body", m.IronicOsImage, `^[a-zA-Z_$][a-zA-Z_.\-$0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *CapiSettingIronicConfig) validateIronicProvisionIP(formats strfmt.Registry) error {
	if swag.IsZero(m.IronicProvisionIP) { // not required
		return nil
	}

	if err := validate.Pattern("Ironic_config"+"."+"Ironic_provision_ip", "body", m.IronicProvisionIP, `^[a-zA-Z_$][a-zA-Z_.\-$0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *CapiSettingIronicConfig) validateIronicProvisionNic(formats strfmt.Registry) error {
	if swag.IsZero(m.IronicProvisionNic) { // not required
		return nil
	}

	if err := validate.Pattern("Ironic_config"+"."+"Ironic_provision_nic", "body", m.IronicProvisionNic, `^[a-zA-Z_$][a-zA-Z_.\-$0-9]*$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this capi setting ironic config based on context it is used
func (m *CapiSettingIronicConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CapiSettingIronicConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CapiSettingIronicConfig) UnmarshalBinary(b []byte) error {
	var res CapiSettingIronicConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CapiSettingRegistry capi setting registry
//
// swagger:model CapiSettingRegistry
type CapiSettingRegistry struct {

	// auth
	Auth string `json:"Auth,omitempty"`
}

// Validate validates this capi setting registry
func (m *CapiSettingRegistry) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this capi setting registry based on context it is used
func (m *CapiSettingRegistry) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CapiSettingRegistry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CapiSettingRegistry) UnmarshalBinary(b []byte) error {
	var res CapiSettingRegistry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
