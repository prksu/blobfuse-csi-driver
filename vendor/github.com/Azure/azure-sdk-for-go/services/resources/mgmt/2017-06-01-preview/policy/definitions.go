package policy

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// DefinitionsClient is the to manage and control access to your resources, you can define customized policies and
// assign them at a scope.
type DefinitionsClient struct {
	BaseClient
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// NewDefinitionsClient creates an instance of the DefinitionsClient client.
func NewDefinitionsClient(subscriptionID string) DefinitionsClient {
	return NewDefinitionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// NewDefinitionsClientWithBaseURI creates an instance of the DefinitionsClient client.
func NewDefinitionsClientWithBaseURI(baseURI string, subscriptionID string) DefinitionsClient {
	return DefinitionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// CreateOrUpdate creates or updates a policy definition.
//
// policyDefinitionName is the name of the policy definition to create. parameters is the policy definition
// properties.
func (client DefinitionsClient) CreateOrUpdate(ctx context.Context, policyDefinitionName string, parameters Definition) (result Definition, err error) {
	req, err := client.CreateOrUpdatePreparer(ctx, policyDefinitionName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client DefinitionsClient) CreateOrUpdatePreparer(ctx context.Context, policyDefinitionName string, parameters Definition) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"policyDefinitionName": autorest.Encode("path", policyDefinitionName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policyDefinitions/{policyDefinitionName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client DefinitionsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client DefinitionsClient) CreateOrUpdateResponder(resp *http.Response) (result Definition, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// CreateOrUpdateAtManagementGroup creates or updates a policy definition at management group level.
//
// policyDefinitionName is the name of the policy definition to create. parameters is the policy definition
// properties. managementGroupID is the ID of the management group.
func (client DefinitionsClient) CreateOrUpdateAtManagementGroup(ctx context.Context, policyDefinitionName string, parameters Definition, managementGroupID string) (result Definition, err error) {
	req, err := client.CreateOrUpdateAtManagementGroupPreparer(ctx, policyDefinitionName, parameters, managementGroupID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "CreateOrUpdateAtManagementGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateAtManagementGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "CreateOrUpdateAtManagementGroup", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateAtManagementGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "CreateOrUpdateAtManagementGroup", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// CreateOrUpdateAtManagementGroupPreparer prepares the CreateOrUpdateAtManagementGroup request.
func (client DefinitionsClient) CreateOrUpdateAtManagementGroupPreparer(ctx context.Context, policyDefinitionName string, parameters Definition, managementGroupID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"managementGroupId":    autorest.Encode("path", managementGroupID),
		"policyDefinitionName": autorest.Encode("path", policyDefinitionName),
	}

	const APIVersion = "2016-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Management/managementgroups/{managementGroupId}/providers/Microsoft.Authorization/policyDefinitions/{policyDefinitionName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// CreateOrUpdateAtManagementGroupSender sends the CreateOrUpdateAtManagementGroup request. The method will close the
// http.Response Body if it receives an error.
func (client DefinitionsClient) CreateOrUpdateAtManagementGroupSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// CreateOrUpdateAtManagementGroupResponder handles the response to the CreateOrUpdateAtManagementGroup request. The method always
// closes the http.Response Body.
func (client DefinitionsClient) CreateOrUpdateAtManagementGroupResponder(resp *http.Response) (result Definition, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// Delete deletes a policy definition.
//
// policyDefinitionName is the name of the policy definition to delete.
func (client DefinitionsClient) Delete(ctx context.Context, policyDefinitionName string) (result autorest.Response, err error) {
	req, err := client.DeletePreparer(ctx, policyDefinitionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// DeletePreparer prepares the Delete request.
func (client DefinitionsClient) DeletePreparer(ctx context.Context, policyDefinitionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"policyDefinitionName": autorest.Encode("path", policyDefinitionName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policyDefinitions/{policyDefinitionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client DefinitionsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client DefinitionsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// DeleteAtManagementGroup deletes a policy definition at management group level.
//
// policyDefinitionName is the name of the policy definition to delete. managementGroupID is the ID of the
// management group.
func (client DefinitionsClient) DeleteAtManagementGroup(ctx context.Context, policyDefinitionName string, managementGroupID string) (result autorest.Response, err error) {
	req, err := client.DeleteAtManagementGroupPreparer(ctx, policyDefinitionName, managementGroupID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "DeleteAtManagementGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteAtManagementGroupSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "DeleteAtManagementGroup", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteAtManagementGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "DeleteAtManagementGroup", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// DeleteAtManagementGroupPreparer prepares the DeleteAtManagementGroup request.
func (client DefinitionsClient) DeleteAtManagementGroupPreparer(ctx context.Context, policyDefinitionName string, managementGroupID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"managementGroupId":    autorest.Encode("path", managementGroupID),
		"policyDefinitionName": autorest.Encode("path", policyDefinitionName),
	}

	const APIVersion = "2016-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Management/managementgroups/{managementGroupId}/providers/Microsoft.Authorization/policyDefinitions/{policyDefinitionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// DeleteAtManagementGroupSender sends the DeleteAtManagementGroup request. The method will close the
// http.Response Body if it receives an error.
func (client DefinitionsClient) DeleteAtManagementGroupSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// DeleteAtManagementGroupResponder handles the response to the DeleteAtManagementGroup request. The method always
// closes the http.Response Body.
func (client DefinitionsClient) DeleteAtManagementGroupResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// Get gets the policy definition.
//
// policyDefinitionName is the name of the policy definition to get.
func (client DefinitionsClient) Get(ctx context.Context, policyDefinitionName string) (result Definition, err error) {
	req, err := client.GetPreparer(ctx, policyDefinitionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// GetPreparer prepares the Get request.
func (client DefinitionsClient) GetPreparer(ctx context.Context, policyDefinitionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"policyDefinitionName": autorest.Encode("path", policyDefinitionName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policyDefinitions/{policyDefinitionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client DefinitionsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client DefinitionsClient) GetResponder(resp *http.Response) (result Definition, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// GetAtManagementGroup gets the policy definition at management group level.
//
// policyDefinitionName is the name of the policy definition to get. managementGroupID is the ID of the management
// group.
func (client DefinitionsClient) GetAtManagementGroup(ctx context.Context, policyDefinitionName string, managementGroupID string) (result Definition, err error) {
	req, err := client.GetAtManagementGroupPreparer(ctx, policyDefinitionName, managementGroupID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "GetAtManagementGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetAtManagementGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "GetAtManagementGroup", resp, "Failure sending request")
		return
	}

	result, err = client.GetAtManagementGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "GetAtManagementGroup", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// GetAtManagementGroupPreparer prepares the GetAtManagementGroup request.
func (client DefinitionsClient) GetAtManagementGroupPreparer(ctx context.Context, policyDefinitionName string, managementGroupID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"managementGroupId":    autorest.Encode("path", managementGroupID),
		"policyDefinitionName": autorest.Encode("path", policyDefinitionName),
	}

	const APIVersion = "2016-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Management/managementgroups/{managementGroupId}/providers/Microsoft.Authorization/policyDefinitions/{policyDefinitionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// GetAtManagementGroupSender sends the GetAtManagementGroup request. The method will close the
// http.Response Body if it receives an error.
func (client DefinitionsClient) GetAtManagementGroupSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// GetAtManagementGroupResponder handles the response to the GetAtManagementGroup request. The method always
// closes the http.Response Body.
func (client DefinitionsClient) GetAtManagementGroupResponder(resp *http.Response) (result Definition, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// GetBuiltIn gets the built in policy definition.
//
// policyDefinitionName is the name of the built in policy definition to get.
func (client DefinitionsClient) GetBuiltIn(ctx context.Context, policyDefinitionName string) (result Definition, err error) {
	req, err := client.GetBuiltInPreparer(ctx, policyDefinitionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "GetBuiltIn", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetBuiltInSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "GetBuiltIn", resp, "Failure sending request")
		return
	}

	result, err = client.GetBuiltInResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "GetBuiltIn", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// GetBuiltInPreparer prepares the GetBuiltIn request.
func (client DefinitionsClient) GetBuiltInPreparer(ctx context.Context, policyDefinitionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"policyDefinitionName": autorest.Encode("path", policyDefinitionName),
	}

	const APIVersion = "2016-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Authorization/policyDefinitions/{policyDefinitionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// GetBuiltInSender sends the GetBuiltIn request. The method will close the
// http.Response Body if it receives an error.
func (client DefinitionsClient) GetBuiltInSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// GetBuiltInResponder handles the response to the GetBuiltIn request. The method always
// closes the http.Response Body.
func (client DefinitionsClient) GetBuiltInResponder(resp *http.Response) (result Definition, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// List gets all the policy definitions for a subscription.
func (client DefinitionsClient) List(ctx context.Context) (result DefinitionListResultPage, err error) {
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.dlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "List", resp, "Failure sending request")
		return
	}

	result.dlr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "List", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// ListPreparer prepares the List request.
func (client DefinitionsClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policyDefinitions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client DefinitionsClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client DefinitionsClient) ListResponder(resp *http.Response) (result DefinitionListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client DefinitionsClient) listNextResults(lastResults DefinitionListResult) (result DefinitionListResult, err error) {
	req, err := lastResults.definitionListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "policy.DefinitionsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "policy.DefinitionsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client DefinitionsClient) ListComplete(ctx context.Context) (result DefinitionListResultIterator, err error) {
	result.page, err = client.List(ctx)
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// ListBuiltIn gets all the built in policy definitions.
func (client DefinitionsClient) ListBuiltIn(ctx context.Context) (result DefinitionListResultPage, err error) {
	result.fn = client.listBuiltInNextResults
	req, err := client.ListBuiltInPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "ListBuiltIn", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBuiltInSender(req)
	if err != nil {
		result.dlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "ListBuiltIn", resp, "Failure sending request")
		return
	}

	result.dlr, err = client.ListBuiltInResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "ListBuiltIn", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// ListBuiltInPreparer prepares the ListBuiltIn request.
func (client DefinitionsClient) ListBuiltInPreparer(ctx context.Context) (*http.Request, error) {
	const APIVersion = "2016-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Authorization/policyDefinitions"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// ListBuiltInSender sends the ListBuiltIn request. The method will close the
// http.Response Body if it receives an error.
func (client DefinitionsClient) ListBuiltInSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// ListBuiltInResponder handles the response to the ListBuiltIn request. The method always
// closes the http.Response Body.
func (client DefinitionsClient) ListBuiltInResponder(resp *http.Response) (result DefinitionListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listBuiltInNextResults retrieves the next set of results, if any.
func (client DefinitionsClient) listBuiltInNextResults(lastResults DefinitionListResult) (result DefinitionListResult, err error) {
	req, err := lastResults.definitionListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "policy.DefinitionsClient", "listBuiltInNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListBuiltInSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "policy.DefinitionsClient", "listBuiltInNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListBuiltInResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "listBuiltInNextResults", resp, "Failure responding to next results request")
	}
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// ListBuiltInComplete enumerates all values, automatically crossing page boundaries as required.
func (client DefinitionsClient) ListBuiltInComplete(ctx context.Context) (result DefinitionListResultIterator, err error) {
	result.page, err = client.ListBuiltIn(ctx)
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// ListByManagementGroup gets all the policy definitions for a subscription at management group level.
//
// managementGroupID is the ID of the management group.
func (client DefinitionsClient) ListByManagementGroup(ctx context.Context, managementGroupID string) (result DefinitionListResultPage, err error) {
	result.fn = client.listByManagementGroupNextResults
	req, err := client.ListByManagementGroupPreparer(ctx, managementGroupID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "ListByManagementGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByManagementGroupSender(req)
	if err != nil {
		result.dlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "ListByManagementGroup", resp, "Failure sending request")
		return
	}

	result.dlr, err = client.ListByManagementGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "ListByManagementGroup", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// ListByManagementGroupPreparer prepares the ListByManagementGroup request.
func (client DefinitionsClient) ListByManagementGroupPreparer(ctx context.Context, managementGroupID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"managementGroupId": autorest.Encode("path", managementGroupID),
	}

	const APIVersion = "2016-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Management/managementgroups/{managementGroupId}/providers/Microsoft.Authorization/policyDefinitions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// ListByManagementGroupSender sends the ListByManagementGroup request. The method will close the
// http.Response Body if it receives an error.
func (client DefinitionsClient) ListByManagementGroupSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// ListByManagementGroupResponder handles the response to the ListByManagementGroup request. The method always
// closes the http.Response Body.
func (client DefinitionsClient) ListByManagementGroupResponder(resp *http.Response) (result DefinitionListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByManagementGroupNextResults retrieves the next set of results, if any.
func (client DefinitionsClient) listByManagementGroupNextResults(lastResults DefinitionListResult) (result DefinitionListResult, err error) {
	req, err := lastResults.definitionListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "policy.DefinitionsClient", "listByManagementGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByManagementGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "policy.DefinitionsClient", "listByManagementGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByManagementGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "listByManagementGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// ListByManagementGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client DefinitionsClient) ListByManagementGroupComplete(ctx context.Context, managementGroupID string) (result DefinitionListResultIterator, err error) {
	result.page, err = client.ListByManagementGroup(ctx, managementGroupID)
	return
}
