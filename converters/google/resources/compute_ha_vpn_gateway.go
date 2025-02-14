// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"reflect"
)

const ComputeHaVpnGatewayAssetType string = "compute.googleapis.com/HaVpnGateway"

func resourceConverterComputeHaVpnGateway() ResourceConverter {
	return ResourceConverter{
		AssetType: ComputeHaVpnGatewayAssetType,
		Convert:   GetComputeHaVpnGatewayCaiObject,
	}
}

func GetComputeHaVpnGatewayCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	name, err := assetName(d, config, "//compute.googleapis.com/projects/{{project}}/regions/{{region}}/vpnGateways/{{name}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetComputeHaVpnGatewayApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: ComputeHaVpnGatewayAssetType,
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/compute/v1/rest",
				DiscoveryName:        "HaVpnGateway",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetComputeHaVpnGatewayApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	descriptionProp, err := expandComputeHaVpnGatewayDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	nameProp, err := expandComputeHaVpnGatewayName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	networkProp, err := expandComputeHaVpnGatewayNetwork(d.Get("network"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("network"); !isEmptyValue(reflect.ValueOf(networkProp)) && (ok || !reflect.DeepEqual(v, networkProp)) {
		obj["network"] = networkProp
	}
	vpnInterfacesProp, err := expandComputeHaVpnGatewayVpnInterfaces(d.Get("vpn_interfaces"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("vpn_interfaces"); !isEmptyValue(reflect.ValueOf(vpnInterfacesProp)) && (ok || !reflect.DeepEqual(v, vpnInterfacesProp)) {
		obj["vpnInterfaces"] = vpnInterfacesProp
	}
	regionProp, err := expandComputeHaVpnGatewayRegion(d.Get("region"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("region"); !isEmptyValue(reflect.ValueOf(regionProp)) && (ok || !reflect.DeepEqual(v, regionProp)) {
		obj["region"] = regionProp
	}

	return obj, nil
}

func expandComputeHaVpnGatewayDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHaVpnGatewayName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHaVpnGatewayNetwork(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("networks", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for network: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeHaVpnGatewayVpnInterfaces(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedId, err := expandComputeHaVpnGatewayVpnInterfacesId(original["id"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedId); val.IsValid() && !isEmptyValue(val) {
			transformed["id"] = transformedId
		}

		transformedIpAddress, err := expandComputeHaVpnGatewayVpnInterfacesIpAddress(original["ip_address"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedIpAddress); val.IsValid() && !isEmptyValue(val) {
			transformed["ipAddress"] = transformedIpAddress
		}

		transformedInterconnectAttachment, err := expandComputeHaVpnGatewayVpnInterfacesInterconnectAttachment(original["interconnect_attachment"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedInterconnectAttachment); val.IsValid() && !isEmptyValue(val) {
			transformed["interconnectAttachment"] = transformedInterconnectAttachment
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeHaVpnGatewayVpnInterfacesId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHaVpnGatewayVpnInterfacesIpAddress(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHaVpnGatewayVpnInterfacesInterconnectAttachment(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseRegionalFieldValue("interconnectAttachments", v.(string), "project", "region", "zone", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for interconnect_attachment: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeHaVpnGatewayRegion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("regions", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for region: %s", err)
	}
	return f.RelativeLink(), nil
}
