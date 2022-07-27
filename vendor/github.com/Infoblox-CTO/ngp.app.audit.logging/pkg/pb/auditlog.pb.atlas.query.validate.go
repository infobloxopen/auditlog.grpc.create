// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/Infoblox-CTO/ngp.app.audit.logging/pkg/pb/auditlog.proto

package pb // import "github.com/Infoblox-CTO/ngp.app.audit.logging/pkg/pb"

import options "github.com/infobloxopen/protoc-gen-atlas-query-validate/options"
import query "github.com/infobloxopen/atlas-app-toolkit/query"
import _ "github.com/envoyproxy/protoc-gen-validate/validate"
import _ "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.

var AuditlogMethodsRequireFilteringValidation = map[string]map[string]options.FilteringOption{
	"/service.AuditLogging/ListAuditLogs": map[string]options.FilteringOption{
		"id":               options.FilteringOption{ValueType: options.QueryValidate_NUMBER},
		"created_at":       options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"action":           options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"result":           options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"app_id":           options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"resource_id":      options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"resource_type":    options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"user_name":        options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"client_ip":        options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"resource_desc":    options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"message":          options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"request_id":       options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"event_version":    options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"event_metadata.*": options.FilteringOption{Deny: []options.QueryValidate_FilterOperator{options.QueryValidate_ALL}, ValueType: options.QueryValidate_STRING},
		"subject_type":     options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"session_type":     options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"subject_groups":   options.FilteringOption{Deny: []options.QueryValidate_FilterOperator{options.QueryValidate_ALL}, ValueType: options.QueryValidate_DEFAULT},
		"session_id":       options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"http_url":         options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"http_method":      options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"http_req_body":    options.FilteringOption{Deny: []options.QueryValidate_FilterOperator{options.QueryValidate_ALL}, ValueType: options.QueryValidate_STRING},
		"http_code":        options.FilteringOption{ValueType: options.QueryValidate_NUMBER},
		"http_resp_body":   options.FilteringOption{Deny: []options.QueryValidate_FilterOperator{options.QueryValidate_ALL}, ValueType: options.QueryValidate_STRING},
	},
}
var AuditlogMethodsRequireSortingValidation = map[string][]string{
	"/service.AuditLogging/ListAuditLogs": []string{
		"id",
		"created_at",
		"action",
		"result",
		"app_id",
		"resource_id",
		"resource_type",
		"user_name",
		"client_ip",
		"resource_desc",
		"message",
		"request_id",
		"event_version",
		"event_metadata",
		"subject_type",
		"session_type",
		"session_id",
		"http_url",
		"http_method",
		"http_req_body",
		"http_code",
		"http_resp_body",
	},
}
var AuditlogMethodsRequireFieldSelectionValidation = map[string][]string{
	"/service.AuditLogging/ListAuditLogs": {
		"id",
		"created_at",
		"action",
		"result",
		"app_id",
		"resource_id",
		"resource_type",
		"user_name",
		"client_ip",
		"resource_desc",
		"message",
		"request_id",
		"event_version",
		"event_metadata.value",
		"event_metadata",
		"subject_type",
		"session_type",
		"subject_groups",
		"session_id",
		"http_url",
		"http_method",
		"http_req_body",
		"http_code",
		"http_resp_body",
	},
}

func AuditlogValidateFiltering(methodName string, f *query.Filtering) error {
	info, ok := AuditlogMethodsRequireFilteringValidation[methodName]
	if !ok {
		return nil
	}
	return options.ValidateFiltering(f, info)
}
func AuditlogValidateSorting(methodName string, s *query.Sorting) error {
	info, ok := AuditlogMethodsRequireSortingValidation[methodName]
	if !ok {
		return nil
	}
	return options.ValidateSorting(s, info)
}
func AuditlogValidateFieldSelection(methodName string, s *query.FieldSelection) error {
	info, ok := AuditlogMethodsRequireFieldSelectionValidation[methodName]
	if !ok {
		return nil
	}
	return options.ValidateFieldSelection(s, info)
}