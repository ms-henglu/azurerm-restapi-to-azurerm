//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsecurityinsights

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// WatchlistItemsClient contains the methods for the WatchlistItems group.
// Don't use this type directly, use NewWatchlistItemsClient() instead.
type WatchlistItemsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewWatchlistItemsClient creates a new instance of WatchlistItemsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewWatchlistItemsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*WatchlistItemsClient, error) {
	cl, err := arm.NewClient(moduleName+".WatchlistItemsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &WatchlistItemsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates a watchlist item.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - watchlistAlias - Watchlist Alias
//   - watchlistItemID - Watchlist Item Id (GUID)
//   - watchlistItem - The watchlist item
//   - options - WatchlistItemsClientCreateOrUpdateOptions contains the optional parameters for the WatchlistItemsClient.CreateOrUpdate
//     method.
func (client *WatchlistItemsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, watchlistAlias string, watchlistItemID string, watchlistItem WatchlistItem, options *WatchlistItemsClientCreateOrUpdateOptions) (WatchlistItemsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, watchlistAlias, watchlistItemID, watchlistItem, options)
	if err != nil {
		return WatchlistItemsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WatchlistItemsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return WatchlistItemsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *WatchlistItemsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, watchlistAlias string, watchlistItemID string, watchlistItem WatchlistItem, options *WatchlistItemsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/watchlists/{watchlistAlias}/watchlistItems/{watchlistItemId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if watchlistAlias == "" {
		return nil, errors.New("parameter watchlistAlias cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{watchlistAlias}", url.PathEscape(watchlistAlias))
	if watchlistItemID == "" {
		return nil, errors.New("parameter watchlistItemID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{watchlistItemId}", url.PathEscape(watchlistItemID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, watchlistItem)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *WatchlistItemsClient) createOrUpdateHandleResponse(resp *http.Response) (WatchlistItemsClientCreateOrUpdateResponse, error) {
	result := WatchlistItemsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WatchlistItem); err != nil {
		return WatchlistItemsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete a watchlist item.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - watchlistAlias - Watchlist Alias
//   - watchlistItemID - Watchlist Item Id (GUID)
//   - options - WatchlistItemsClientDeleteOptions contains the optional parameters for the WatchlistItemsClient.Delete method.
func (client *WatchlistItemsClient) Delete(ctx context.Context, resourceGroupName string, workspaceName string, watchlistAlias string, watchlistItemID string, options *WatchlistItemsClientDeleteOptions) (WatchlistItemsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, watchlistAlias, watchlistItemID, options)
	if err != nil {
		return WatchlistItemsClientDeleteResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WatchlistItemsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return WatchlistItemsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return WatchlistItemsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *WatchlistItemsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, watchlistAlias string, watchlistItemID string, options *WatchlistItemsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/watchlists/{watchlistAlias}/watchlistItems/{watchlistItemId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if watchlistAlias == "" {
		return nil, errors.New("parameter watchlistAlias cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{watchlistAlias}", url.PathEscape(watchlistAlias))
	if watchlistItemID == "" {
		return nil, errors.New("parameter watchlistItemID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{watchlistItemId}", url.PathEscape(watchlistItemID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a watchlist, without its watchlist items.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - watchlistAlias - Watchlist Alias
//   - watchlistItemID - Watchlist Item Id (GUID)
//   - options - WatchlistItemsClientGetOptions contains the optional parameters for the WatchlistItemsClient.Get method.
func (client *WatchlistItemsClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, watchlistAlias string, watchlistItemID string, options *WatchlistItemsClientGetOptions) (WatchlistItemsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, watchlistAlias, watchlistItemID, options)
	if err != nil {
		return WatchlistItemsClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WatchlistItemsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WatchlistItemsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *WatchlistItemsClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, watchlistAlias string, watchlistItemID string, options *WatchlistItemsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/watchlists/{watchlistAlias}/watchlistItems/{watchlistItemId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if watchlistAlias == "" {
		return nil, errors.New("parameter watchlistAlias cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{watchlistAlias}", url.PathEscape(watchlistAlias))
	if watchlistItemID == "" {
		return nil, errors.New("parameter watchlistItemID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{watchlistItemId}", url.PathEscape(watchlistItemID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *WatchlistItemsClient) getHandleResponse(resp *http.Response) (WatchlistItemsClientGetResponse, error) {
	result := WatchlistItemsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WatchlistItem); err != nil {
		return WatchlistItemsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets all watchlist Items.
//
// Generated from API version 2022-09-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - watchlistAlias - Watchlist Alias
//   - options - WatchlistItemsClientListOptions contains the optional parameters for the WatchlistItemsClient.NewListPager method.
func (client *WatchlistItemsClient) NewListPager(resourceGroupName string, workspaceName string, watchlistAlias string, options *WatchlistItemsClientListOptions) *runtime.Pager[WatchlistItemsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[WatchlistItemsClientListResponse]{
		More: func(page WatchlistItemsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *WatchlistItemsClientListResponse) (WatchlistItemsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, workspaceName, watchlistAlias, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return WatchlistItemsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return WatchlistItemsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return WatchlistItemsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *WatchlistItemsClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, watchlistAlias string, options *WatchlistItemsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/watchlists/{watchlistAlias}/watchlistItems"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if watchlistAlias == "" {
		return nil, errors.New("parameter watchlistAlias cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{watchlistAlias}", url.PathEscape(watchlistAlias))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *WatchlistItemsClient) listHandleResponse(resp *http.Response) (WatchlistItemsClientListResponse, error) {
	result := WatchlistItemsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WatchlistItemList); err != nil {
		return WatchlistItemsClientListResponse{}, err
	}
	return result, nil
}