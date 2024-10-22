//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// SPDX-License-Identifier: AGPL-3.0-only

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v0alpha1

import (
	common "k8s.io/kube-openapi/pkg/common"
	spec "k8s.io/kube-openapi/pkg/validation/spec"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/grafana/grafana/pkg/apis/dashboardsnapshot/v0alpha1.DashboardCreateCommand":         schema_pkg_apis_dashboardsnapshot_v0alpha1_DashboardCreateCommand(ref),
		"github.com/grafana/grafana/pkg/apis/dashboardsnapshot/v0alpha1.DashboardCreateResponse":        schema_pkg_apis_dashboardsnapshot_v0alpha1_DashboardCreateResponse(ref),
		"github.com/grafana/grafana/pkg/apis/dashboardsnapshot/v0alpha1.DashboardSnapshot":              schema_pkg_apis_dashboardsnapshot_v0alpha1_DashboardSnapshot(ref),
		"github.com/grafana/grafana/pkg/apis/dashboardsnapshot/v0alpha1.DashboardSnapshotList":          schema_pkg_apis_dashboardsnapshot_v0alpha1_DashboardSnapshotList(ref),
		"github.com/grafana/grafana/pkg/apis/dashboardsnapshot/v0alpha1.DashboardSnapshotWithDeleteKey": schema_pkg_apis_dashboardsnapshot_v0alpha1_DashboardSnapshotWithDeleteKey(ref),
		"github.com/grafana/grafana/pkg/apis/dashboardsnapshot/v0alpha1.FullDashboardSnapshot":          schema_pkg_apis_dashboardsnapshot_v0alpha1_FullDashboardSnapshot(ref),
		"github.com/grafana/grafana/pkg/apis/dashboardsnapshot/v0alpha1.SharingOptions":                 schema_pkg_apis_dashboardsnapshot_v0alpha1_SharingOptions(ref),
		"github.com/grafana/grafana/pkg/apis/dashboardsnapshot/v0alpha1.SharingOptionsList":             schema_pkg_apis_dashboardsnapshot_v0alpha1_SharingOptionsList(ref),
		"github.com/grafana/grafana/pkg/apis/dashboardsnapshot/v0alpha1.SnapshotInfo":                   schema_pkg_apis_dashboardsnapshot_v0alpha1_SnapshotInfo(ref),
		"github.com/grafana/grafana/pkg/apis/dashboardsnapshot/v0alpha1.SnapshotSharingOptions":         schema_pkg_apis_dashboardsnapshot_v0alpha1_SnapshotSharingOptions(ref),
	}
}

func schema_pkg_apis_dashboardsnapshot_v0alpha1_DashboardCreateCommand(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "These are the values expected to be sent from an end user",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"name": {
						SchemaProps: spec.SchemaProps{
							Description: "Snapshot name required:false",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"dashboard": {
						SchemaProps: spec.SchemaProps{
							Description: "The complete dashboard model. required:true",
							Ref:         ref("github.com/grafana/grafana/pkg/apis/common/v0alpha1.Unstructured"),
						},
					},
					"expires": {
						SchemaProps: spec.SchemaProps{
							Description: "When the snapshot should expire in seconds in seconds. Default is never to expire. required:false default:0",
							Default:     0,
							Type:        []string{"integer"},
							Format:      "int64",
						},
					},
					"external": {
						SchemaProps: spec.SchemaProps{
							Description: "these are passed when storing an external snapshot ref Save the snapshot on an external server rather than locally. required:false default: false",
							Default:     false,
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
				},
				Required: []string{"name", "dashboard", "expires", "external"},
			},
		},
		Dependencies: []string{
			"github.com/grafana/grafana/pkg/apis/common/v0alpha1.Unstructured"},
	}
}

func schema_pkg_apis_dashboardsnapshot_v0alpha1_DashboardCreateResponse(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "The create response",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"key": {
						SchemaProps: spec.SchemaProps{
							Description: "The unique key",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"deleteKey": {
						SchemaProps: spec.SchemaProps{
							Description: "A unique key that will allow delete",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"url": {
						SchemaProps: spec.SchemaProps{
							Description: "Absolute URL to show the dashboard",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"deleteUrl": {
						SchemaProps: spec.SchemaProps{
							Description: "URL that will delete the response",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"key", "deleteKey", "url", "deleteUrl"},
			},
		},
	}
}

func schema_pkg_apis_dashboardsnapshot_v0alpha1_DashboardSnapshot(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Description: "Snapshot summary info",
							Default:     map[string]interface{}{},
							Ref:         ref("github.com/grafana/grafana/pkg/apis/dashboardsnapshot/v0alpha1.SnapshotInfo"),
						},
					},
				},
				Required: []string{"spec"},
			},
		},
		Dependencies: []string{
			"github.com/grafana/grafana/pkg/apis/dashboardsnapshot/v0alpha1.SnapshotInfo", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_dashboardsnapshot_v0alpha1_DashboardSnapshotList(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"),
						},
					},
					"items": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("github.com/grafana/grafana/pkg/apis/dashboardsnapshot/v0alpha1.DashboardSnapshot"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/grafana/grafana/pkg/apis/dashboardsnapshot/v0alpha1.DashboardSnapshot", "k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"},
	}
}

func schema_pkg_apis_dashboardsnapshot_v0alpha1_DashboardSnapshotWithDeleteKey(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "This is returned from the POST command",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Description: "Snapshot summary info",
							Default:     map[string]interface{}{},
							Ref:         ref("github.com/grafana/grafana/pkg/apis/dashboardsnapshot/v0alpha1.SnapshotInfo"),
						},
					},
					"deleteKey": {
						SchemaProps: spec.SchemaProps{
							Description: "The delete key is only returned when the item is created.  It is not returned from a get request",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"spec"},
			},
		},
		Dependencies: []string{
			"github.com/grafana/grafana/pkg/apis/dashboardsnapshot/v0alpha1.SnapshotInfo", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_dashboardsnapshot_v0alpha1_FullDashboardSnapshot(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "This is the snapshot returned from the subresource",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"info": {
						SchemaProps: spec.SchemaProps{
							Description: "Snapshot summary info",
							Default:     map[string]interface{}{},
							Ref:         ref("github.com/grafana/grafana/pkg/apis/dashboardsnapshot/v0alpha1.SnapshotInfo"),
						},
					},
					"dashboard": {
						SchemaProps: spec.SchemaProps{
							Description: "The raw dashboard (unstructured for now)",
							Ref:         ref("github.com/grafana/grafana/pkg/apis/common/v0alpha1.Unstructured"),
						},
					},
				},
				Required: []string{"info", "dashboard"},
			},
		},
		Dependencies: []string{
			"github.com/grafana/grafana/pkg/apis/common/v0alpha1.Unstructured", "github.com/grafana/grafana/pkg/apis/dashboardsnapshot/v0alpha1.SnapshotInfo", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_dashboardsnapshot_v0alpha1_SharingOptions(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Represents an options object that must be named for each namespace/team/user",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Description: "Show the options inline",
							Default:     map[string]interface{}{},
							Ref:         ref("github.com/grafana/grafana/pkg/apis/dashboardsnapshot/v0alpha1.SnapshotSharingOptions"),
						},
					},
				},
				Required: []string{"spec"},
			},
		},
		Dependencies: []string{
			"github.com/grafana/grafana/pkg/apis/dashboardsnapshot/v0alpha1.SnapshotSharingOptions", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_dashboardsnapshot_v0alpha1_SharingOptionsList(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Represents a list of namespaced options",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"),
						},
					},
					"items": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("github.com/grafana/grafana/pkg/apis/dashboardsnapshot/v0alpha1.SharingOptions"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/grafana/grafana/pkg/apis/dashboardsnapshot/v0alpha1.SharingOptions", "k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"},
	}
}

func schema_pkg_apis_dashboardsnapshot_v0alpha1_SnapshotInfo(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"title": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"expires": {
						SchemaProps: spec.SchemaProps{
							Description: "Optionally auto-remove the snapshot at a future date",
							Type:        []string{"integer"},
							Format:      "int64",
						},
					},
					"external": {
						SchemaProps: spec.SchemaProps{
							Description: "When set to true, the snapshot exists in a remote server",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"externalUrl": {
						SchemaProps: spec.SchemaProps{
							Description: "The external URL where the snapshot can be seen",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"originalUrl": {
						SchemaProps: spec.SchemaProps{
							Description: "The URL that created the dashboard originally",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"timestamp": {
						SchemaProps: spec.SchemaProps{
							Description: "Snapshot creation timestamp",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
	}
}

func schema_pkg_apis_dashboardsnapshot_v0alpha1_SnapshotSharingOptions(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Each tenant, may have different sharing options This is currently set using custom.ini, but multi-tenant support will need to be managed differently",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"snapshotEnabled": {
						SchemaProps: spec.SchemaProps{
							Default: false,
							Type:    []string{"boolean"},
							Format:  "",
						},
					},
					"externalSnapshotURL": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"externalSnapshotName": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"externalEnabled": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
				},
				Required: []string{"snapshotEnabled"},
			},
		},
	}
}
