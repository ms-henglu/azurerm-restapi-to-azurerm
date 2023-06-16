//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armappplatform

import "encoding/json"

func unmarshalCertificatePropertiesClassification(rawMsg json.RawMessage) (CertificatePropertiesClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b CertificatePropertiesClassification
	switch m["type"] {
	case "ContentCertificate":
		b = &ContentCertificateProperties{}
	case "KeyVaultCertificate":
		b = &KeyVaultCertificateProperties{}
	default:
		b = &CertificateProperties{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalUserSourceInfoClassification(rawMsg json.RawMessage) (UserSourceInfoClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b UserSourceInfoClassification
	switch m["type"] {
	case "BuildResult":
		b = &BuildResultUserSourceInfo{}
	case "Jar":
		b = &JarUploadedUserSourceInfo{}
	case "NetCoreZip":
		b = &NetCoreZipUploadedUserSourceInfo{}
	case "Source":
		b = &SourceUploadedUserSourceInfo{}
	case "UploadedUserSourceInfo":
		b = &UploadedUserSourceInfo{}
	default:
		b = &UserSourceInfo{}
	}
	return b, json.Unmarshal(rawMsg, b)
}