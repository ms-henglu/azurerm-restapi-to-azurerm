//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmonitor

// ActionGroupsClientCreateNotificationsAtActionGroupResourceLevelResponse contains the response from method ActionGroupsClient.BeginCreateNotificationsAtActionGroupResourceLevel.
type ActionGroupsClientCreateNotificationsAtActionGroupResourceLevelResponse struct {
	TestNotificationDetailsResponse
}

// ActionGroupsClientCreateOrUpdateResponse contains the response from method ActionGroupsClient.CreateOrUpdate.
type ActionGroupsClientCreateOrUpdateResponse struct {
	ActionGroupResource
}

// ActionGroupsClientDeleteResponse contains the response from method ActionGroupsClient.Delete.
type ActionGroupsClientDeleteResponse struct {
	// placeholder for future response values
}

// ActionGroupsClientEnableReceiverResponse contains the response from method ActionGroupsClient.EnableReceiver.
type ActionGroupsClientEnableReceiverResponse struct {
	// placeholder for future response values
}

// ActionGroupsClientGetResponse contains the response from method ActionGroupsClient.Get.
type ActionGroupsClientGetResponse struct {
	ActionGroupResource
}

// ActionGroupsClientGetTestNotificationsAtActionGroupResourceLevelResponse contains the response from method ActionGroupsClient.GetTestNotificationsAtActionGroupResourceLevel.
type ActionGroupsClientGetTestNotificationsAtActionGroupResourceLevelResponse struct {
	TestNotificationDetailsResponse
}

// ActionGroupsClientListByResourceGroupResponse contains the response from method ActionGroupsClient.NewListByResourceGroupPager.
type ActionGroupsClientListByResourceGroupResponse struct {
	ActionGroupList
}

// ActionGroupsClientListBySubscriptionIDResponse contains the response from method ActionGroupsClient.NewListBySubscriptionIDPager.
type ActionGroupsClientListBySubscriptionIDResponse struct {
	ActionGroupList
}

// ActionGroupsClientUpdateResponse contains the response from method ActionGroupsClient.Update.
type ActionGroupsClientUpdateResponse struct {
	ActionGroupResource
}

// ActivityLogAlertsClientCreateOrUpdateResponse contains the response from method ActivityLogAlertsClient.CreateOrUpdate.
type ActivityLogAlertsClientCreateOrUpdateResponse struct {
	ActivityLogAlertResource
}

// ActivityLogAlertsClientDeleteResponse contains the response from method ActivityLogAlertsClient.Delete.
type ActivityLogAlertsClientDeleteResponse struct {
	// placeholder for future response values
}

// ActivityLogAlertsClientGetResponse contains the response from method ActivityLogAlertsClient.Get.
type ActivityLogAlertsClientGetResponse struct {
	ActivityLogAlertResource
}

// ActivityLogAlertsClientListByResourceGroupResponse contains the response from method ActivityLogAlertsClient.NewListByResourceGroupPager.
type ActivityLogAlertsClientListByResourceGroupResponse struct {
	AlertRuleList
}

// ActivityLogAlertsClientListBySubscriptionIDResponse contains the response from method ActivityLogAlertsClient.NewListBySubscriptionIDPager.
type ActivityLogAlertsClientListBySubscriptionIDResponse struct {
	AlertRuleList
}

// ActivityLogAlertsClientUpdateResponse contains the response from method ActivityLogAlertsClient.Update.
type ActivityLogAlertsClientUpdateResponse struct {
	ActivityLogAlertResource
}

// ActivityLogsClientListResponse contains the response from method ActivityLogsClient.NewListPager.
type ActivityLogsClientListResponse struct {
	EventDataCollection
}

// AlertRuleIncidentsClientGetResponse contains the response from method AlertRuleIncidentsClient.Get.
type AlertRuleIncidentsClientGetResponse struct {
	Incident
}

// AlertRuleIncidentsClientListByAlertRuleResponse contains the response from method AlertRuleIncidentsClient.NewListByAlertRulePager.
type AlertRuleIncidentsClientListByAlertRuleResponse struct {
	IncidentListResult
}

// AlertRulesClientCreateOrUpdateResponse contains the response from method AlertRulesClient.CreateOrUpdate.
type AlertRulesClientCreateOrUpdateResponse struct {
	AlertRuleResource
}

// AlertRulesClientDeleteResponse contains the response from method AlertRulesClient.Delete.
type AlertRulesClientDeleteResponse struct {
	// placeholder for future response values
}

// AlertRulesClientGetResponse contains the response from method AlertRulesClient.Get.
type AlertRulesClientGetResponse struct {
	AlertRuleResource
}

// AlertRulesClientListByResourceGroupResponse contains the response from method AlertRulesClient.NewListByResourceGroupPager.
type AlertRulesClientListByResourceGroupResponse struct {
	AlertRuleResourceCollection
}

// AlertRulesClientListBySubscriptionResponse contains the response from method AlertRulesClient.NewListBySubscriptionPager.
type AlertRulesClientListBySubscriptionResponse struct {
	AlertRuleResourceCollection
}

// AlertRulesClientUpdateResponse contains the response from method AlertRulesClient.Update.
type AlertRulesClientUpdateResponse struct {
	AlertRuleResource
}

// AutoscaleSettingsClientCreateOrUpdateResponse contains the response from method AutoscaleSettingsClient.CreateOrUpdate.
type AutoscaleSettingsClientCreateOrUpdateResponse struct {
	AutoscaleSettingResource
}

// AutoscaleSettingsClientDeleteResponse contains the response from method AutoscaleSettingsClient.Delete.
type AutoscaleSettingsClientDeleteResponse struct {
	// placeholder for future response values
}

// AutoscaleSettingsClientGetResponse contains the response from method AutoscaleSettingsClient.Get.
type AutoscaleSettingsClientGetResponse struct {
	AutoscaleSettingResource
}

// AutoscaleSettingsClientListByResourceGroupResponse contains the response from method AutoscaleSettingsClient.NewListByResourceGroupPager.
type AutoscaleSettingsClientListByResourceGroupResponse struct {
	AutoscaleSettingResourceCollection
}

// AutoscaleSettingsClientListBySubscriptionResponse contains the response from method AutoscaleSettingsClient.NewListBySubscriptionPager.
type AutoscaleSettingsClientListBySubscriptionResponse struct {
	AutoscaleSettingResourceCollection
}

// AutoscaleSettingsClientUpdateResponse contains the response from method AutoscaleSettingsClient.Update.
type AutoscaleSettingsClientUpdateResponse struct {
	AutoscaleSettingResource
}

// AzureMonitorWorkspacesClientCreateResponse contains the response from method AzureMonitorWorkspacesClient.Create.
type AzureMonitorWorkspacesClientCreateResponse struct {
	AzureMonitorWorkspaceResource
}

// AzureMonitorWorkspacesClientDeleteResponse contains the response from method AzureMonitorWorkspacesClient.Delete.
type AzureMonitorWorkspacesClientDeleteResponse struct {
	// placeholder for future response values
}

// AzureMonitorWorkspacesClientGetResponse contains the response from method AzureMonitorWorkspacesClient.Get.
type AzureMonitorWorkspacesClientGetResponse struct {
	AzureMonitorWorkspaceResource
}

// AzureMonitorWorkspacesClientListByResourceGroupResponse contains the response from method AzureMonitorWorkspacesClient.NewListByResourceGroupPager.
type AzureMonitorWorkspacesClientListByResourceGroupResponse struct {
	AzureMonitorWorkspaceResourceListResult
}

// AzureMonitorWorkspacesClientListBySubscriptionResponse contains the response from method AzureMonitorWorkspacesClient.NewListBySubscriptionPager.
type AzureMonitorWorkspacesClientListBySubscriptionResponse struct {
	AzureMonitorWorkspaceResourceListResult
}

// AzureMonitorWorkspacesClientUpdateResponse contains the response from method AzureMonitorWorkspacesClient.Update.
type AzureMonitorWorkspacesClientUpdateResponse struct {
	AzureMonitorWorkspaceResource
}

// BaselinesClientListResponse contains the response from method BaselinesClient.NewListPager.
type BaselinesClientListResponse struct {
	MetricBaselinesResponse
}

// DataCollectionEndpointsClientCreateResponse contains the response from method DataCollectionEndpointsClient.Create.
type DataCollectionEndpointsClientCreateResponse struct {
	DataCollectionEndpointResource
}

// DataCollectionEndpointsClientDeleteResponse contains the response from method DataCollectionEndpointsClient.Delete.
type DataCollectionEndpointsClientDeleteResponse struct {
	// placeholder for future response values
}

// DataCollectionEndpointsClientGetResponse contains the response from method DataCollectionEndpointsClient.Get.
type DataCollectionEndpointsClientGetResponse struct {
	DataCollectionEndpointResource
}

// DataCollectionEndpointsClientListByResourceGroupResponse contains the response from method DataCollectionEndpointsClient.NewListByResourceGroupPager.
type DataCollectionEndpointsClientListByResourceGroupResponse struct {
	DataCollectionEndpointResourceListResult
}

// DataCollectionEndpointsClientListBySubscriptionResponse contains the response from method DataCollectionEndpointsClient.NewListBySubscriptionPager.
type DataCollectionEndpointsClientListBySubscriptionResponse struct {
	DataCollectionEndpointResourceListResult
}

// DataCollectionEndpointsClientUpdateResponse contains the response from method DataCollectionEndpointsClient.Update.
type DataCollectionEndpointsClientUpdateResponse struct {
	DataCollectionEndpointResource
}

// DataCollectionRuleAssociationsClientCreateResponse contains the response from method DataCollectionRuleAssociationsClient.Create.
type DataCollectionRuleAssociationsClientCreateResponse struct {
	DataCollectionRuleAssociationProxyOnlyResource
}

// DataCollectionRuleAssociationsClientDeleteResponse contains the response from method DataCollectionRuleAssociationsClient.Delete.
type DataCollectionRuleAssociationsClientDeleteResponse struct {
	// placeholder for future response values
}

// DataCollectionRuleAssociationsClientGetResponse contains the response from method DataCollectionRuleAssociationsClient.Get.
type DataCollectionRuleAssociationsClientGetResponse struct {
	DataCollectionRuleAssociationProxyOnlyResource
}

// DataCollectionRuleAssociationsClientListByDataCollectionEndpointResponse contains the response from method DataCollectionRuleAssociationsClient.NewListByDataCollectionEndpointPager.
type DataCollectionRuleAssociationsClientListByDataCollectionEndpointResponse struct {
	DataCollectionRuleAssociationProxyOnlyResourceListResult
}

// DataCollectionRuleAssociationsClientListByResourceResponse contains the response from method DataCollectionRuleAssociationsClient.NewListByResourcePager.
type DataCollectionRuleAssociationsClientListByResourceResponse struct {
	DataCollectionRuleAssociationProxyOnlyResourceListResult
}

// DataCollectionRuleAssociationsClientListByRuleResponse contains the response from method DataCollectionRuleAssociationsClient.NewListByRulePager.
type DataCollectionRuleAssociationsClientListByRuleResponse struct {
	DataCollectionRuleAssociationProxyOnlyResourceListResult
}

// DataCollectionRulesClientCreateResponse contains the response from method DataCollectionRulesClient.Create.
type DataCollectionRulesClientCreateResponse struct {
	DataCollectionRuleResource
}

// DataCollectionRulesClientDeleteResponse contains the response from method DataCollectionRulesClient.Delete.
type DataCollectionRulesClientDeleteResponse struct {
	// placeholder for future response values
}

// DataCollectionRulesClientGetResponse contains the response from method DataCollectionRulesClient.Get.
type DataCollectionRulesClientGetResponse struct {
	DataCollectionRuleResource
}

// DataCollectionRulesClientListByResourceGroupResponse contains the response from method DataCollectionRulesClient.NewListByResourceGroupPager.
type DataCollectionRulesClientListByResourceGroupResponse struct {
	DataCollectionRuleResourceListResult
}

// DataCollectionRulesClientListBySubscriptionResponse contains the response from method DataCollectionRulesClient.NewListBySubscriptionPager.
type DataCollectionRulesClientListBySubscriptionResponse struct {
	DataCollectionRuleResourceListResult
}

// DataCollectionRulesClientUpdateResponse contains the response from method DataCollectionRulesClient.Update.
type DataCollectionRulesClientUpdateResponse struct {
	DataCollectionRuleResource
}

// DiagnosticSettingsCategoryClientGetResponse contains the response from method DiagnosticSettingsCategoryClient.Get.
type DiagnosticSettingsCategoryClientGetResponse struct {
	DiagnosticSettingsCategoryResource
}

// DiagnosticSettingsCategoryClientListResponse contains the response from method DiagnosticSettingsCategoryClient.NewListPager.
type DiagnosticSettingsCategoryClientListResponse struct {
	DiagnosticSettingsCategoryResourceCollection
}

// DiagnosticSettingsClientCreateOrUpdateResponse contains the response from method DiagnosticSettingsClient.CreateOrUpdate.
type DiagnosticSettingsClientCreateOrUpdateResponse struct {
	DiagnosticSettingsResource
}

// DiagnosticSettingsClientDeleteResponse contains the response from method DiagnosticSettingsClient.Delete.
type DiagnosticSettingsClientDeleteResponse struct {
	// placeholder for future response values
}

// DiagnosticSettingsClientGetResponse contains the response from method DiagnosticSettingsClient.Get.
type DiagnosticSettingsClientGetResponse struct {
	DiagnosticSettingsResource
}

// DiagnosticSettingsClientListResponse contains the response from method DiagnosticSettingsClient.NewListPager.
type DiagnosticSettingsClientListResponse struct {
	DiagnosticSettingsResourceCollection
}

// EventCategoriesClientListResponse contains the response from method EventCategoriesClient.NewListPager.
type EventCategoriesClientListResponse struct {
	EventCategoryCollection
}

// LogProfilesClientCreateOrUpdateResponse contains the response from method LogProfilesClient.CreateOrUpdate.
type LogProfilesClientCreateOrUpdateResponse struct {
	LogProfileResource
}

// LogProfilesClientDeleteResponse contains the response from method LogProfilesClient.Delete.
type LogProfilesClientDeleteResponse struct {
	// placeholder for future response values
}

// LogProfilesClientGetResponse contains the response from method LogProfilesClient.Get.
type LogProfilesClientGetResponse struct {
	LogProfileResource
}

// LogProfilesClientListResponse contains the response from method LogProfilesClient.NewListPager.
type LogProfilesClientListResponse struct {
	LogProfileCollection
}

// LogProfilesClientUpdateResponse contains the response from method LogProfilesClient.Update.
type LogProfilesClientUpdateResponse struct {
	LogProfileResource
}

// MetricAlertsClientCreateOrUpdateResponse contains the response from method MetricAlertsClient.CreateOrUpdate.
type MetricAlertsClientCreateOrUpdateResponse struct {
	MetricAlertResource
}

// MetricAlertsClientDeleteResponse contains the response from method MetricAlertsClient.Delete.
type MetricAlertsClientDeleteResponse struct {
	// placeholder for future response values
}

// MetricAlertsClientGetResponse contains the response from method MetricAlertsClient.Get.
type MetricAlertsClientGetResponse struct {
	MetricAlertResource
}

// MetricAlertsClientListByResourceGroupResponse contains the response from method MetricAlertsClient.NewListByResourceGroupPager.
type MetricAlertsClientListByResourceGroupResponse struct {
	MetricAlertResourceCollection
}

// MetricAlertsClientListBySubscriptionResponse contains the response from method MetricAlertsClient.NewListBySubscriptionPager.
type MetricAlertsClientListBySubscriptionResponse struct {
	MetricAlertResourceCollection
}

// MetricAlertsClientUpdateResponse contains the response from method MetricAlertsClient.Update.
type MetricAlertsClientUpdateResponse struct {
	MetricAlertResource
}

// MetricAlertsStatusClientListByNameResponse contains the response from method MetricAlertsStatusClient.ListByName.
type MetricAlertsStatusClientListByNameResponse struct {
	MetricAlertStatusCollection
}

// MetricAlertsStatusClientListResponse contains the response from method MetricAlertsStatusClient.List.
type MetricAlertsStatusClientListResponse struct {
	MetricAlertStatusCollection
}

// MetricDefinitionsClientListAtSubscriptionScopeResponse contains the response from method MetricDefinitionsClient.NewListAtSubscriptionScopePager.
type MetricDefinitionsClientListAtSubscriptionScopeResponse struct {
	SubscriptionScopeMetricDefinitionCollection
}

// MetricDefinitionsClientListResponse contains the response from method MetricDefinitionsClient.NewListPager.
type MetricDefinitionsClientListResponse struct {
	MetricDefinitionCollection
}

// MetricNamespacesClientListResponse contains the response from method MetricNamespacesClient.NewListPager.
type MetricNamespacesClientListResponse struct {
	MetricNamespaceCollection
}

// MetricsClientListAtSubscriptionScopePostResponse contains the response from method MetricsClient.ListAtSubscriptionScopePost.
type MetricsClientListAtSubscriptionScopePostResponse struct {
	SubscriptionScopeMetricResponse
}

// MetricsClientListAtSubscriptionScopeResponse contains the response from method MetricsClient.ListAtSubscriptionScope.
type MetricsClientListAtSubscriptionScopeResponse struct {
	SubscriptionScopeMetricResponse
}

// MetricsClientListResponse contains the response from method MetricsClient.List.
type MetricsClientListResponse struct {
	Response
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationListResult
}

// OperationsForMonitorClientListResponse contains the response from method OperationsForMonitorClient.NewListPager.
type OperationsForMonitorClientListResponse struct {
	OperationListResultAutoGenerated
}

// PredictiveMetricClientGetResponse contains the response from method PredictiveMetricClient.Get.
type PredictiveMetricClientGetResponse struct {
	PredictiveResponse
}

// PrivateEndpointConnectionsClientCreateOrUpdateResponse contains the response from method PrivateEndpointConnectionsClient.BeginCreateOrUpdate.
type PrivateEndpointConnectionsClientCreateOrUpdateResponse struct {
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientDeleteResponse contains the response from method PrivateEndpointConnectionsClient.BeginDelete.
type PrivateEndpointConnectionsClientDeleteResponse struct {
	// placeholder for future response values
}

// PrivateEndpointConnectionsClientGetResponse contains the response from method PrivateEndpointConnectionsClient.Get.
type PrivateEndpointConnectionsClientGetResponse struct {
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientListByPrivateLinkScopeResponse contains the response from method PrivateEndpointConnectionsClient.ListByPrivateLinkScope.
type PrivateEndpointConnectionsClientListByPrivateLinkScopeResponse struct {
	PrivateEndpointConnectionListResult
}

// PrivateLinkResourcesClientGetResponse contains the response from method PrivateLinkResourcesClient.Get.
type PrivateLinkResourcesClientGetResponse struct {
	PrivateLinkResource
}

// PrivateLinkResourcesClientListByPrivateLinkScopeResponse contains the response from method PrivateLinkResourcesClient.ListByPrivateLinkScope.
type PrivateLinkResourcesClientListByPrivateLinkScopeResponse struct {
	PrivateLinkResourceListResult
}

// PrivateLinkScopeOperationStatusClientGetResponse contains the response from method PrivateLinkScopeOperationStatusClient.Get.
type PrivateLinkScopeOperationStatusClientGetResponse struct {
	OperationStatus
}

// PrivateLinkScopedResourcesClientCreateOrUpdateResponse contains the response from method PrivateLinkScopedResourcesClient.BeginCreateOrUpdate.
type PrivateLinkScopedResourcesClientCreateOrUpdateResponse struct {
	ScopedResource
}

// PrivateLinkScopedResourcesClientDeleteResponse contains the response from method PrivateLinkScopedResourcesClient.BeginDelete.
type PrivateLinkScopedResourcesClientDeleteResponse struct {
	// placeholder for future response values
}

// PrivateLinkScopedResourcesClientGetResponse contains the response from method PrivateLinkScopedResourcesClient.Get.
type PrivateLinkScopedResourcesClientGetResponse struct {
	ScopedResource
}

// PrivateLinkScopedResourcesClientListByPrivateLinkScopeResponse contains the response from method PrivateLinkScopedResourcesClient.NewListByPrivateLinkScopePager.
type PrivateLinkScopedResourcesClientListByPrivateLinkScopeResponse struct {
	ScopedResourceListResult
}

// PrivateLinkScopesClientCreateOrUpdateResponse contains the response from method PrivateLinkScopesClient.CreateOrUpdate.
type PrivateLinkScopesClientCreateOrUpdateResponse struct {
	AzureMonitorPrivateLinkScope
}

// PrivateLinkScopesClientDeleteResponse contains the response from method PrivateLinkScopesClient.BeginDelete.
type PrivateLinkScopesClientDeleteResponse struct {
	// placeholder for future response values
}

// PrivateLinkScopesClientGetResponse contains the response from method PrivateLinkScopesClient.Get.
type PrivateLinkScopesClientGetResponse struct {
	AzureMonitorPrivateLinkScope
}

// PrivateLinkScopesClientListByResourceGroupResponse contains the response from method PrivateLinkScopesClient.NewListByResourceGroupPager.
type PrivateLinkScopesClientListByResourceGroupResponse struct {
	AzureMonitorPrivateLinkScopeListResult
}

// PrivateLinkScopesClientListResponse contains the response from method PrivateLinkScopesClient.NewListPager.
type PrivateLinkScopesClientListResponse struct {
	AzureMonitorPrivateLinkScopeListResult
}

// PrivateLinkScopesClientUpdateTagsResponse contains the response from method PrivateLinkScopesClient.UpdateTags.
type PrivateLinkScopesClientUpdateTagsResponse struct {
	AzureMonitorPrivateLinkScope
}

// ScheduledQueryRulesClientCreateOrUpdateResponse contains the response from method ScheduledQueryRulesClient.CreateOrUpdate.
type ScheduledQueryRulesClientCreateOrUpdateResponse struct {
	ScheduledQueryRuleResource
}

// ScheduledQueryRulesClientDeleteResponse contains the response from method ScheduledQueryRulesClient.Delete.
type ScheduledQueryRulesClientDeleteResponse struct {
	// placeholder for future response values
}

// ScheduledQueryRulesClientGetResponse contains the response from method ScheduledQueryRulesClient.Get.
type ScheduledQueryRulesClientGetResponse struct {
	ScheduledQueryRuleResource
}

// ScheduledQueryRulesClientListByResourceGroupResponse contains the response from method ScheduledQueryRulesClient.NewListByResourceGroupPager.
type ScheduledQueryRulesClientListByResourceGroupResponse struct {
	ScheduledQueryRuleResourceCollection
}

// ScheduledQueryRulesClientListBySubscriptionResponse contains the response from method ScheduledQueryRulesClient.NewListBySubscriptionPager.
type ScheduledQueryRulesClientListBySubscriptionResponse struct {
	ScheduledQueryRuleResourceCollection
}

// ScheduledQueryRulesClientUpdateResponse contains the response from method ScheduledQueryRulesClient.Update.
type ScheduledQueryRulesClientUpdateResponse struct {
	ScheduledQueryRuleResource
}

// TenantActionGroupsClientCreateOrUpdateResponse contains the response from method TenantActionGroupsClient.CreateOrUpdate.
type TenantActionGroupsClientCreateOrUpdateResponse struct {
	TenantActionGroupResource
}

// TenantActionGroupsClientDeleteResponse contains the response from method TenantActionGroupsClient.Delete.
type TenantActionGroupsClientDeleteResponse struct {
	// placeholder for future response values
}

// TenantActionGroupsClientGetResponse contains the response from method TenantActionGroupsClient.Get.
type TenantActionGroupsClientGetResponse struct {
	TenantActionGroupResource
}

// TenantActionGroupsClientListByManagementGroupIDResponse contains the response from method TenantActionGroupsClient.NewListByManagementGroupIDPager.
type TenantActionGroupsClientListByManagementGroupIDResponse struct {
	TenantActionGroupList
}

// TenantActionGroupsClientUpdateResponse contains the response from method TenantActionGroupsClient.Update.
type TenantActionGroupsClientUpdateResponse struct {
	TenantActionGroupResource
}

// TenantActivityLogsClientListResponse contains the response from method TenantActivityLogsClient.NewListPager.
type TenantActivityLogsClientListResponse struct {
	EventDataCollection
}

// VMInsightsClientGetOnboardingStatusResponse contains the response from method VMInsightsClient.GetOnboardingStatus.
type VMInsightsClientGetOnboardingStatusResponse struct {
	VMInsightsOnboardingStatus
}