//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcdn

// AFDCustomDomainsClientCreateResponse contains the response from method AFDCustomDomainsClient.Create.
type AFDCustomDomainsClientCreateResponse struct {
	AFDDomain
}

// AFDCustomDomainsClientDeleteResponse contains the response from method AFDCustomDomainsClient.Delete.
type AFDCustomDomainsClientDeleteResponse struct {
	// placeholder for future response values
}

// AFDCustomDomainsClientGetResponse contains the response from method AFDCustomDomainsClient.Get.
type AFDCustomDomainsClientGetResponse struct {
	AFDDomain
}

// AFDCustomDomainsClientListByProfileResponse contains the response from method AFDCustomDomainsClient.ListByProfile.
type AFDCustomDomainsClientListByProfileResponse struct {
	AFDDomainListResult
}

// AFDCustomDomainsClientRefreshValidationTokenResponse contains the response from method AFDCustomDomainsClient.RefreshValidationToken.
type AFDCustomDomainsClientRefreshValidationTokenResponse struct {
	// placeholder for future response values
}

// AFDCustomDomainsClientUpdateResponse contains the response from method AFDCustomDomainsClient.Update.
type AFDCustomDomainsClientUpdateResponse struct {
	AFDDomain
}

// AFDEndpointsClientCreateResponse contains the response from method AFDEndpointsClient.Create.
type AFDEndpointsClientCreateResponse struct {
	AFDEndpoint
}

// AFDEndpointsClientDeleteResponse contains the response from method AFDEndpointsClient.Delete.
type AFDEndpointsClientDeleteResponse struct {
	// placeholder for future response values
}

// AFDEndpointsClientGetResponse contains the response from method AFDEndpointsClient.Get.
type AFDEndpointsClientGetResponse struct {
	AFDEndpoint
}

// AFDEndpointsClientListByProfileResponse contains the response from method AFDEndpointsClient.ListByProfile.
type AFDEndpointsClientListByProfileResponse struct {
	AFDEndpointListResult
}

// AFDEndpointsClientListResourceUsageResponse contains the response from method AFDEndpointsClient.ListResourceUsage.
type AFDEndpointsClientListResourceUsageResponse struct {
	UsagesListResult
}

// AFDEndpointsClientPurgeContentResponse contains the response from method AFDEndpointsClient.PurgeContent.
type AFDEndpointsClientPurgeContentResponse struct {
	// placeholder for future response values
}

// AFDEndpointsClientUpdateResponse contains the response from method AFDEndpointsClient.Update.
type AFDEndpointsClientUpdateResponse struct {
	AFDEndpoint
}

// AFDEndpointsClientValidateCustomDomainResponse contains the response from method AFDEndpointsClient.ValidateCustomDomain.
type AFDEndpointsClientValidateCustomDomainResponse struct {
	ValidateCustomDomainOutput
}

// AFDOriginGroupsClientCreateResponse contains the response from method AFDOriginGroupsClient.Create.
type AFDOriginGroupsClientCreateResponse struct {
	AFDOriginGroup
}

// AFDOriginGroupsClientDeleteResponse contains the response from method AFDOriginGroupsClient.Delete.
type AFDOriginGroupsClientDeleteResponse struct {
	// placeholder for future response values
}

// AFDOriginGroupsClientGetResponse contains the response from method AFDOriginGroupsClient.Get.
type AFDOriginGroupsClientGetResponse struct {
	AFDOriginGroup
}

// AFDOriginGroupsClientListByProfileResponse contains the response from method AFDOriginGroupsClient.ListByProfile.
type AFDOriginGroupsClientListByProfileResponse struct {
	AFDOriginGroupListResult
}

// AFDOriginGroupsClientListResourceUsageResponse contains the response from method AFDOriginGroupsClient.ListResourceUsage.
type AFDOriginGroupsClientListResourceUsageResponse struct {
	UsagesListResult
}

// AFDOriginGroupsClientUpdateResponse contains the response from method AFDOriginGroupsClient.Update.
type AFDOriginGroupsClientUpdateResponse struct {
	AFDOriginGroup
}

// AFDOriginsClientCreateResponse contains the response from method AFDOriginsClient.Create.
type AFDOriginsClientCreateResponse struct {
	AFDOrigin
}

// AFDOriginsClientDeleteResponse contains the response from method AFDOriginsClient.Delete.
type AFDOriginsClientDeleteResponse struct {
	// placeholder for future response values
}

// AFDOriginsClientGetResponse contains the response from method AFDOriginsClient.Get.
type AFDOriginsClientGetResponse struct {
	AFDOrigin
}

// AFDOriginsClientListByOriginGroupResponse contains the response from method AFDOriginsClient.ListByOriginGroup.
type AFDOriginsClientListByOriginGroupResponse struct {
	AFDOriginListResult
}

// AFDOriginsClientUpdateResponse contains the response from method AFDOriginsClient.Update.
type AFDOriginsClientUpdateResponse struct {
	AFDOrigin
}

// AFDProfilesClientCheckHostNameAvailabilityResponse contains the response from method AFDProfilesClient.CheckHostNameAvailability.
type AFDProfilesClientCheckHostNameAvailabilityResponse struct {
	CheckNameAvailabilityOutput
}

// AFDProfilesClientListResourceUsageResponse contains the response from method AFDProfilesClient.ListResourceUsage.
type AFDProfilesClientListResourceUsageResponse struct {
	UsagesListResult
}

// CustomDomainsClientCreateResponse contains the response from method CustomDomainsClient.Create.
type CustomDomainsClientCreateResponse struct {
	CustomDomain
}

// CustomDomainsClientDeleteResponse contains the response from method CustomDomainsClient.Delete.
type CustomDomainsClientDeleteResponse struct {
	CustomDomain
}

// CustomDomainsClientDisableCustomHTTPSResponse contains the response from method CustomDomainsClient.DisableCustomHTTPS.
type CustomDomainsClientDisableCustomHTTPSResponse struct {
	CustomDomain
}

// CustomDomainsClientEnableCustomHTTPSResponse contains the response from method CustomDomainsClient.EnableCustomHTTPS.
type CustomDomainsClientEnableCustomHTTPSResponse struct {
	CustomDomain
}

// CustomDomainsClientGetResponse contains the response from method CustomDomainsClient.Get.
type CustomDomainsClientGetResponse struct {
	CustomDomain
}

// CustomDomainsClientListByEndpointResponse contains the response from method CustomDomainsClient.ListByEndpoint.
type CustomDomainsClientListByEndpointResponse struct {
	CustomDomainListResult
}

// EdgeNodesClientListResponse contains the response from method EdgeNodesClient.List.
type EdgeNodesClientListResponse struct {
	EdgenodeResult
}

// EndpointsClientCreateResponse contains the response from method EndpointsClient.Create.
type EndpointsClientCreateResponse struct {
	Endpoint
}

// EndpointsClientDeleteResponse contains the response from method EndpointsClient.Delete.
type EndpointsClientDeleteResponse struct {
	// placeholder for future response values
}

// EndpointsClientGetResponse contains the response from method EndpointsClient.Get.
type EndpointsClientGetResponse struct {
	Endpoint
}

// EndpointsClientListByProfileResponse contains the response from method EndpointsClient.ListByProfile.
type EndpointsClientListByProfileResponse struct {
	EndpointListResult
}

// EndpointsClientListResourceUsageResponse contains the response from method EndpointsClient.ListResourceUsage.
type EndpointsClientListResourceUsageResponse struct {
	ResourceUsageListResult
}

// EndpointsClientLoadContentResponse contains the response from method EndpointsClient.LoadContent.
type EndpointsClientLoadContentResponse struct {
	// placeholder for future response values
}

// EndpointsClientPurgeContentResponse contains the response from method EndpointsClient.PurgeContent.
type EndpointsClientPurgeContentResponse struct {
	// placeholder for future response values
}

// EndpointsClientStartResponse contains the response from method EndpointsClient.Start.
type EndpointsClientStartResponse struct {
	Endpoint
}

// EndpointsClientStopResponse contains the response from method EndpointsClient.Stop.
type EndpointsClientStopResponse struct {
	Endpoint
}

// EndpointsClientUpdateResponse contains the response from method EndpointsClient.Update.
type EndpointsClientUpdateResponse struct {
	Endpoint
}

// EndpointsClientValidateCustomDomainResponse contains the response from method EndpointsClient.ValidateCustomDomain.
type EndpointsClientValidateCustomDomainResponse struct {
	ValidateCustomDomainOutput
}

// LogAnalyticsClientGetLogAnalyticsLocationsResponse contains the response from method LogAnalyticsClient.GetLogAnalyticsLocations.
type LogAnalyticsClientGetLogAnalyticsLocationsResponse struct {
	ContinentsResponse
}

// LogAnalyticsClientGetLogAnalyticsMetricsResponse contains the response from method LogAnalyticsClient.GetLogAnalyticsMetrics.
type LogAnalyticsClientGetLogAnalyticsMetricsResponse struct {
	MetricsResponse
}

// LogAnalyticsClientGetLogAnalyticsRankingsResponse contains the response from method LogAnalyticsClient.GetLogAnalyticsRankings.
type LogAnalyticsClientGetLogAnalyticsRankingsResponse struct {
	RankingsResponse
}

// LogAnalyticsClientGetLogAnalyticsResourcesResponse contains the response from method LogAnalyticsClient.GetLogAnalyticsResources.
type LogAnalyticsClientGetLogAnalyticsResourcesResponse struct {
	ResourcesResponse
}

// LogAnalyticsClientGetWafLogAnalyticsMetricsResponse contains the response from method LogAnalyticsClient.GetWafLogAnalyticsMetrics.
type LogAnalyticsClientGetWafLogAnalyticsMetricsResponse struct {
	WafMetricsResponse
}

// LogAnalyticsClientGetWafLogAnalyticsRankingsResponse contains the response from method LogAnalyticsClient.GetWafLogAnalyticsRankings.
type LogAnalyticsClientGetWafLogAnalyticsRankingsResponse struct {
	WafRankingsResponse
}

// ManagedRuleSetsClientListResponse contains the response from method ManagedRuleSetsClient.List.
type ManagedRuleSetsClientListResponse struct {
	ManagedRuleSetDefinitionList
}

// ManagementClientCheckEndpointNameAvailabilityResponse contains the response from method ManagementClient.CheckEndpointNameAvailability.
type ManagementClientCheckEndpointNameAvailabilityResponse struct {
	CheckEndpointNameAvailabilityOutput
}

// ManagementClientCheckNameAvailabilityResponse contains the response from method ManagementClient.CheckNameAvailability.
type ManagementClientCheckNameAvailabilityResponse struct {
	CheckNameAvailabilityOutput
}

// ManagementClientCheckNameAvailabilityWithSubscriptionResponse contains the response from method ManagementClient.CheckNameAvailabilityWithSubscription.
type ManagementClientCheckNameAvailabilityWithSubscriptionResponse struct {
	CheckNameAvailabilityOutput
}

// ManagementClientValidateProbeResponse contains the response from method ManagementClient.ValidateProbe.
type ManagementClientValidateProbeResponse struct {
	ValidateProbeOutput
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationsListResult
}

// OriginGroupsClientCreateResponse contains the response from method OriginGroupsClient.Create.
type OriginGroupsClientCreateResponse struct {
	OriginGroup
}

// OriginGroupsClientDeleteResponse contains the response from method OriginGroupsClient.Delete.
type OriginGroupsClientDeleteResponse struct {
	// placeholder for future response values
}

// OriginGroupsClientGetResponse contains the response from method OriginGroupsClient.Get.
type OriginGroupsClientGetResponse struct {
	OriginGroup
}

// OriginGroupsClientListByEndpointResponse contains the response from method OriginGroupsClient.ListByEndpoint.
type OriginGroupsClientListByEndpointResponse struct {
	OriginGroupListResult
}

// OriginGroupsClientUpdateResponse contains the response from method OriginGroupsClient.Update.
type OriginGroupsClientUpdateResponse struct {
	OriginGroup
}

// OriginsClientCreateResponse contains the response from method OriginsClient.Create.
type OriginsClientCreateResponse struct {
	Origin
}

// OriginsClientDeleteResponse contains the response from method OriginsClient.Delete.
type OriginsClientDeleteResponse struct {
	// placeholder for future response values
}

// OriginsClientGetResponse contains the response from method OriginsClient.Get.
type OriginsClientGetResponse struct {
	Origin
}

// OriginsClientListByEndpointResponse contains the response from method OriginsClient.ListByEndpoint.
type OriginsClientListByEndpointResponse struct {
	OriginListResult
}

// OriginsClientUpdateResponse contains the response from method OriginsClient.Update.
type OriginsClientUpdateResponse struct {
	Origin
}

// PoliciesClientCreateOrUpdateResponse contains the response from method PoliciesClient.CreateOrUpdate.
type PoliciesClientCreateOrUpdateResponse struct {
	WebApplicationFirewallPolicy
}

// PoliciesClientDeleteResponse contains the response from method PoliciesClient.Delete.
type PoliciesClientDeleteResponse struct {
	// placeholder for future response values
}

// PoliciesClientGetResponse contains the response from method PoliciesClient.Get.
type PoliciesClientGetResponse struct {
	WebApplicationFirewallPolicy
}

// PoliciesClientListResponse contains the response from method PoliciesClient.List.
type PoliciesClientListResponse struct {
	WebApplicationFirewallPolicyList
}

// PoliciesClientUpdateResponse contains the response from method PoliciesClient.Update.
type PoliciesClientUpdateResponse struct {
	WebApplicationFirewallPolicy
}

// ProfilesClientCreateResponse contains the response from method ProfilesClient.Create.
type ProfilesClientCreateResponse struct {
	Profile
}

// ProfilesClientDeleteResponse contains the response from method ProfilesClient.Delete.
type ProfilesClientDeleteResponse struct {
	// placeholder for future response values
}

// ProfilesClientGenerateSsoURIResponse contains the response from method ProfilesClient.GenerateSsoURI.
type ProfilesClientGenerateSsoURIResponse struct {
	SsoURI
}

// ProfilesClientGetResponse contains the response from method ProfilesClient.Get.
type ProfilesClientGetResponse struct {
	Profile
}

// ProfilesClientListByResourceGroupResponse contains the response from method ProfilesClient.ListByResourceGroup.
type ProfilesClientListByResourceGroupResponse struct {
	ProfileListResult
}

// ProfilesClientListResourceUsageResponse contains the response from method ProfilesClient.ListResourceUsage.
type ProfilesClientListResourceUsageResponse struct {
	ResourceUsageListResult
}

// ProfilesClientListResponse contains the response from method ProfilesClient.List.
type ProfilesClientListResponse struct {
	ProfileListResult
}

// ProfilesClientListSupportedOptimizationTypesResponse contains the response from method ProfilesClient.ListSupportedOptimizationTypes.
type ProfilesClientListSupportedOptimizationTypesResponse struct {
	SupportedOptimizationTypesListResult
}

// ProfilesClientUpdateResponse contains the response from method ProfilesClient.Update.
type ProfilesClientUpdateResponse struct {
	Profile
}

// ResourceUsageClientListResponse contains the response from method ResourceUsageClient.List.
type ResourceUsageClientListResponse struct {
	ResourceUsageListResult
}

// RoutesClientCreateResponse contains the response from method RoutesClient.Create.
type RoutesClientCreateResponse struct {
	Route
}

// RoutesClientDeleteResponse contains the response from method RoutesClient.Delete.
type RoutesClientDeleteResponse struct {
	// placeholder for future response values
}

// RoutesClientGetResponse contains the response from method RoutesClient.Get.
type RoutesClientGetResponse struct {
	Route
}

// RoutesClientListByEndpointResponse contains the response from method RoutesClient.ListByEndpoint.
type RoutesClientListByEndpointResponse struct {
	RouteListResult
}

// RoutesClientUpdateResponse contains the response from method RoutesClient.Update.
type RoutesClientUpdateResponse struct {
	Route
}

// RuleSetsClientCreateResponse contains the response from method RuleSetsClient.Create.
type RuleSetsClientCreateResponse struct {
	RuleSet
}

// RuleSetsClientDeleteResponse contains the response from method RuleSetsClient.Delete.
type RuleSetsClientDeleteResponse struct {
	// placeholder for future response values
}

// RuleSetsClientGetResponse contains the response from method RuleSetsClient.Get.
type RuleSetsClientGetResponse struct {
	RuleSet
}

// RuleSetsClientListByProfileResponse contains the response from method RuleSetsClient.ListByProfile.
type RuleSetsClientListByProfileResponse struct {
	RuleSetListResult
}

// RuleSetsClientListResourceUsageResponse contains the response from method RuleSetsClient.ListResourceUsage.
type RuleSetsClientListResourceUsageResponse struct {
	UsagesListResult
}

// RulesClientCreateResponse contains the response from method RulesClient.Create.
type RulesClientCreateResponse struct {
	Rule
}

// RulesClientDeleteResponse contains the response from method RulesClient.Delete.
type RulesClientDeleteResponse struct {
	// placeholder for future response values
}

// RulesClientGetResponse contains the response from method RulesClient.Get.
type RulesClientGetResponse struct {
	Rule
}

// RulesClientListByRuleSetResponse contains the response from method RulesClient.ListByRuleSet.
type RulesClientListByRuleSetResponse struct {
	RuleListResult
}

// RulesClientUpdateResponse contains the response from method RulesClient.Update.
type RulesClientUpdateResponse struct {
	Rule
}

// SecretsClientCreateResponse contains the response from method SecretsClient.Create.
type SecretsClientCreateResponse struct {
	Secret
}

// SecretsClientDeleteResponse contains the response from method SecretsClient.Delete.
type SecretsClientDeleteResponse struct {
	// placeholder for future response values
}

// SecretsClientGetResponse contains the response from method SecretsClient.Get.
type SecretsClientGetResponse struct {
	Secret
}

// SecretsClientListByProfileResponse contains the response from method SecretsClient.ListByProfile.
type SecretsClientListByProfileResponse struct {
	SecretListResult
}

// SecurityPoliciesClientCreateResponse contains the response from method SecurityPoliciesClient.Create.
type SecurityPoliciesClientCreateResponse struct {
	SecurityPolicy
}

// SecurityPoliciesClientDeleteResponse contains the response from method SecurityPoliciesClient.Delete.
type SecurityPoliciesClientDeleteResponse struct {
	// placeholder for future response values
}

// SecurityPoliciesClientGetResponse contains the response from method SecurityPoliciesClient.Get.
type SecurityPoliciesClientGetResponse struct {
	SecurityPolicy
}

// SecurityPoliciesClientListByProfileResponse contains the response from method SecurityPoliciesClient.ListByProfile.
type SecurityPoliciesClientListByProfileResponse struct {
	SecurityPolicyListResult
}

// SecurityPoliciesClientPatchResponse contains the response from method SecurityPoliciesClient.Patch.
type SecurityPoliciesClientPatchResponse struct {
	SecurityPolicy
}

// ValidateClientSecretResponse contains the response from method ValidateClient.Secret.
type ValidateClientSecretResponse struct {
	ValidateSecretOutput
}