// +build !ignore_autogenerated

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/mattermost/mattermost-operator/pkg/apis/mattermost/v1alpha1.ClusterInstallation":       schema_pkg_apis_mattermost_v1alpha1_ClusterInstallation(ref),
		"github.com/mattermost/mattermost-operator/pkg/apis/mattermost/v1alpha1.ClusterInstallationSpec":   schema_pkg_apis_mattermost_v1alpha1_ClusterInstallationSpec(ref),
		"github.com/mattermost/mattermost-operator/pkg/apis/mattermost/v1alpha1.ClusterInstallationStatus": schema_pkg_apis_mattermost_v1alpha1_ClusterInstallationStatus(ref),
	}
}

func schema_pkg_apis_mattermost_v1alpha1_ClusterInstallation(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ClusterInstallation is the Schema for the clusterinstallations API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Description: "Specification of the desired behavior of the Mattermost cluster. More info: https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#spec-and-status",
							Ref:         ref("github.com/mattermost/mattermost-operator/pkg/apis/mattermost/v1alpha1.ClusterInstallationSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Description: "Most recent observed status of the Mattermost cluster. Read-only. Not included when requesting from the apiserver, only from the Mattermost Operator API itself. More info: https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#spec-and-status",
							Ref:         ref("github.com/mattermost/mattermost-operator/pkg/apis/mattermost/v1alpha1.ClusterInstallationStatus"),
						},
					},
				},
				Required: []string{"spec"},
			},
		},
		Dependencies: []string{
			"github.com/mattermost/mattermost-operator/pkg/apis/mattermost/v1alpha1.ClusterInstallationSpec", "github.com/mattermost/mattermost-operator/pkg/apis/mattermost/v1alpha1.ClusterInstallationStatus"},
	}
}

func schema_pkg_apis_mattermost_v1alpha1_ClusterInstallationSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ClusterInstallationSpec defines the desired state of ClusterInstallation",
				Properties: map[string]spec.Schema{
					"image": {
						SchemaProps: spec.SchemaProps{
							Description: "Image defines the ClusterInstallation Docker image.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"version": {
						SchemaProps: spec.SchemaProps{
							Description: "Version defines the ClusterInstallation Docker image version.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"replicas": {
						SchemaProps: spec.SchemaProps{
							Description: "Replicas defines the number of Mattermost instances in a ClusterInstallation resource",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"ingressName": {
						SchemaProps: spec.SchemaProps{
							Description: "IngressName defines the name to be used when creating the ingress rules",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"nodeSelector": {
						SchemaProps: spec.SchemaProps{
							Description: "NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must match a node's labels for the pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"affinity": {
						SchemaProps: spec.SchemaProps{
							Description: "If specified, affinity will define the pod's scheduling constraints",
							Ref:         ref("k8s.io/api/core/v1.Affinity"),
						},
					},
					"minioStorageSize": {
						SchemaProps: spec.SchemaProps{
							Description: "MinioStorageSize defines the storage size for minio. ie 50Gi",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"databaseType": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/mattermost/mattermost-operator/pkg/apis/mattermost/v1alpha1.DatabaseType"),
						},
					},
				},
				Required: []string{"ingressName"},
			},
		},
		Dependencies: []string{
			"github.com/mattermost/mattermost-operator/pkg/apis/mattermost/v1alpha1.DatabaseType", "k8s.io/api/core/v1.Affinity"},
	}
}

func schema_pkg_apis_mattermost_v1alpha1_ClusterInstallationStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ClusterInstallationStatus defines the observed state of ClusterInstallation",
				Properties: map[string]spec.Schema{
					"paused": {
						SchemaProps: spec.SchemaProps{
							Description: "Represents whether any actions on the underlying managed objects are being performed. Only delete actions will be performed.",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"replicas": {
						SchemaProps: spec.SchemaProps{
							Description: "Total number of non-terminated pods targeted by this Mattermost deployment (their labels match the selector).",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
				},
				Required: []string{"paused", "replicas"},
			},
		},
		Dependencies: []string{},
	}
}