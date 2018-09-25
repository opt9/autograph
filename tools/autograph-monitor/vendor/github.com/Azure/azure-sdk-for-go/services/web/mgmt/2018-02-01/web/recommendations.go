package web

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
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// RecommendationsClient is the webSite Management Client
type RecommendationsClient struct {
	BaseClient
}

// NewRecommendationsClient creates an instance of the RecommendationsClient client.
func NewRecommendationsClient(subscriptionID string) RecommendationsClient {
	return NewRecommendationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewRecommendationsClientWithBaseURI creates an instance of the RecommendationsClient client.
func NewRecommendationsClientWithBaseURI(baseURI string, subscriptionID string) RecommendationsClient {
	return RecommendationsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// DisableAllForWebApp disable all recommendations for an app.
// Parameters:
// resourceGroupName - name of the resource group to which the resource belongs.
// siteName - name of the app.
func (client RecommendationsClient) DisableAllForWebApp(ctx context.Context, resourceGroupName string, siteName string) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+[^\.]$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("web.RecommendationsClient", "DisableAllForWebApp", err.Error())
	}

	req, err := client.DisableAllForWebAppPreparer(ctx, resourceGroupName, siteName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.RecommendationsClient", "DisableAllForWebApp", nil, "Failure preparing request")
		return
	}

	resp, err := client.DisableAllForWebAppSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "web.RecommendationsClient", "DisableAllForWebApp", resp, "Failure sending request")
		return
	}

	result, err = client.DisableAllForWebAppResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.RecommendationsClient", "DisableAllForWebApp", resp, "Failure responding to request")
	}

	return
}

// DisableAllForWebAppPreparer prepares the DisableAllForWebApp request.
func (client RecommendationsClient) DisableAllForWebAppPreparer(ctx context.Context, resourceGroupName string, siteName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"siteName":          autorest.Encode("path", siteName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-02-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/recommendations/disable", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DisableAllForWebAppSender sends the DisableAllForWebApp request. The method will close the
// http.Response Body if it receives an error.
func (client RecommendationsClient) DisableAllForWebAppSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// DisableAllForWebAppResponder handles the response to the DisableAllForWebApp request. The method always
// closes the http.Response Body.
func (client RecommendationsClient) DisableAllForWebAppResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// DisableRecommendationForSite disables the specific rule for a web site permanently.
// Parameters:
// resourceGroupName - name of the resource group to which the resource belongs.
// siteName - site name
// name - rule name
func (client RecommendationsClient) DisableRecommendationForSite(ctx context.Context, resourceGroupName string, siteName string, name string) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+[^\.]$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("web.RecommendationsClient", "DisableRecommendationForSite", err.Error())
	}

	req, err := client.DisableRecommendationForSitePreparer(ctx, resourceGroupName, siteName, name)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.RecommendationsClient", "DisableRecommendationForSite", nil, "Failure preparing request")
		return
	}

	resp, err := client.DisableRecommendationForSiteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "web.RecommendationsClient", "DisableRecommendationForSite", resp, "Failure sending request")
		return
	}

	result, err = client.DisableRecommendationForSiteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.RecommendationsClient", "DisableRecommendationForSite", resp, "Failure responding to request")
	}

	return
}

// DisableRecommendationForSitePreparer prepares the DisableRecommendationForSite request.
func (client RecommendationsClient) DisableRecommendationForSitePreparer(ctx context.Context, resourceGroupName string, siteName string, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"siteName":          autorest.Encode("path", siteName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-02-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/recommendations/{name}/disable", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DisableRecommendationForSiteSender sends the DisableRecommendationForSite request. The method will close the
// http.Response Body if it receives an error.
func (client RecommendationsClient) DisableRecommendationForSiteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// DisableRecommendationForSiteResponder handles the response to the DisableRecommendationForSite request. The method always
// closes the http.Response Body.
func (client RecommendationsClient) DisableRecommendationForSiteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// DisableRecommendationForSubscription disables the specified rule so it will not apply to a subscription in the
// future.
// Parameters:
// name - rule name
func (client RecommendationsClient) DisableRecommendationForSubscription(ctx context.Context, name string) (result autorest.Response, err error) {
	req, err := client.DisableRecommendationForSubscriptionPreparer(ctx, name)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.RecommendationsClient", "DisableRecommendationForSubscription", nil, "Failure preparing request")
		return
	}

	resp, err := client.DisableRecommendationForSubscriptionSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "web.RecommendationsClient", "DisableRecommendationForSubscription", resp, "Failure sending request")
		return
	}

	result, err = client.DisableRecommendationForSubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.RecommendationsClient", "DisableRecommendationForSubscription", resp, "Failure responding to request")
	}

	return
}

// DisableRecommendationForSubscriptionPreparer prepares the DisableRecommendationForSubscription request.
func (client RecommendationsClient) DisableRecommendationForSubscriptionPreparer(ctx context.Context, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":           autorest.Encode("path", name),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-02-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Web/recommendations/{name}/disable", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DisableRecommendationForSubscriptionSender sends the DisableRecommendationForSubscription request. The method will close the
// http.Response Body if it receives an error.
func (client RecommendationsClient) DisableRecommendationForSubscriptionSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// DisableRecommendationForSubscriptionResponder handles the response to the DisableRecommendationForSubscription request. The method always
// closes the http.Response Body.
func (client RecommendationsClient) DisableRecommendationForSubscriptionResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// GetRuleDetailsByWebApp get a recommendation rule for an app.
// Parameters:
// resourceGroupName - name of the resource group to which the resource belongs.
// siteName - name of the app.
// name - name of the recommendation.
// updateSeen - specify <code>true</code> to update the last-seen timestamp of the recommendation object.
// recommendationID - the GUID of the recommedation object if you query an expired one. You don't need to
// specify it to query an active entry.
func (client RecommendationsClient) GetRuleDetailsByWebApp(ctx context.Context, resourceGroupName string, siteName string, name string, updateSeen *bool, recommendationID string) (result RecommendationRule, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+[^\.]$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("web.RecommendationsClient", "GetRuleDetailsByWebApp", err.Error())
	}

	req, err := client.GetRuleDetailsByWebAppPreparer(ctx, resourceGroupName, siteName, name, updateSeen, recommendationID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.RecommendationsClient", "GetRuleDetailsByWebApp", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetRuleDetailsByWebAppSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "web.RecommendationsClient", "GetRuleDetailsByWebApp", resp, "Failure sending request")
		return
	}

	result, err = client.GetRuleDetailsByWebAppResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.RecommendationsClient", "GetRuleDetailsByWebApp", resp, "Failure responding to request")
	}

	return
}

// GetRuleDetailsByWebAppPreparer prepares the GetRuleDetailsByWebApp request.
func (client RecommendationsClient) GetRuleDetailsByWebAppPreparer(ctx context.Context, resourceGroupName string, siteName string, name string, updateSeen *bool, recommendationID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"siteName":          autorest.Encode("path", siteName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-02-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if updateSeen != nil {
		queryParameters["updateSeen"] = autorest.Encode("query", *updateSeen)
	}
	if len(recommendationID) > 0 {
		queryParameters["recommendationId"] = autorest.Encode("query", recommendationID)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/recommendations/{name}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetRuleDetailsByWebAppSender sends the GetRuleDetailsByWebApp request. The method will close the
// http.Response Body if it receives an error.
func (client RecommendationsClient) GetRuleDetailsByWebAppSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetRuleDetailsByWebAppResponder handles the response to the GetRuleDetailsByWebApp request. The method always
// closes the http.Response Body.
func (client RecommendationsClient) GetRuleDetailsByWebAppResponder(resp *http.Response) (result RecommendationRule, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list all recommendations for a subscription.
// Parameters:
// featured - specify <code>true</code> to return only the most critical recommendations. The default is
// <code>false</code>, which returns all recommendations.
// filter - filter is specified by using OData syntax. Example: $filter=channel eq 'Api' or channel eq
// 'Notification' and startTime eq 2014-01-01T00:00:00Z and endTime eq 2014-12-31T23:59:59Z and timeGrain eq
// duration'[PT1H|PT1M|P1D]
func (client RecommendationsClient) List(ctx context.Context, featured *bool, filter string) (result RecommendationCollectionPage, err error) {
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, featured, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.RecommendationsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.rc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "web.RecommendationsClient", "List", resp, "Failure sending request")
		return
	}

	result.rc, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.RecommendationsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client RecommendationsClient) ListPreparer(ctx context.Context, featured *bool, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-02-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if featured != nil {
		queryParameters["featured"] = autorest.Encode("query", *featured)
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = filter
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Web/recommendations", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client RecommendationsClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client RecommendationsClient) ListResponder(resp *http.Response) (result RecommendationCollection, err error) {
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
func (client RecommendationsClient) listNextResults(lastResults RecommendationCollection) (result RecommendationCollection, err error) {
	req, err := lastResults.recommendationCollectionPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web.RecommendationsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web.RecommendationsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.RecommendationsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client RecommendationsClient) ListComplete(ctx context.Context, featured *bool, filter string) (result RecommendationCollectionIterator, err error) {
	result.page, err = client.List(ctx, featured, filter)
	return
}

// ListHistoryForWebApp get past recommendations for an app, optionally specified by the time range.
// Parameters:
// resourceGroupName - name of the resource group to which the resource belongs.
// siteName - name of the app.
// expiredOnly - specify <code>false</code> to return all recommendations. The default is <code>true</code>,
// which returns only expired recommendations.
// filter - filter is specified by using OData syntax. Example: $filter=channel eq 'Api' or channel eq
// 'Notification' and startTime eq 2014-01-01T00:00:00Z and endTime eq 2014-12-31T23:59:59Z and timeGrain eq
// duration'[PT1H|PT1M|P1D]
func (client RecommendationsClient) ListHistoryForWebApp(ctx context.Context, resourceGroupName string, siteName string, expiredOnly *bool, filter string) (result RecommendationCollectionPage, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+[^\.]$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("web.RecommendationsClient", "ListHistoryForWebApp", err.Error())
	}

	result.fn = client.listHistoryForWebAppNextResults
	req, err := client.ListHistoryForWebAppPreparer(ctx, resourceGroupName, siteName, expiredOnly, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.RecommendationsClient", "ListHistoryForWebApp", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListHistoryForWebAppSender(req)
	if err != nil {
		result.rc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "web.RecommendationsClient", "ListHistoryForWebApp", resp, "Failure sending request")
		return
	}

	result.rc, err = client.ListHistoryForWebAppResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.RecommendationsClient", "ListHistoryForWebApp", resp, "Failure responding to request")
	}

	return
}

// ListHistoryForWebAppPreparer prepares the ListHistoryForWebApp request.
func (client RecommendationsClient) ListHistoryForWebAppPreparer(ctx context.Context, resourceGroupName string, siteName string, expiredOnly *bool, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"siteName":          autorest.Encode("path", siteName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-02-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if expiredOnly != nil {
		queryParameters["expiredOnly"] = autorest.Encode("query", *expiredOnly)
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = filter
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/recommendationHistory", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListHistoryForWebAppSender sends the ListHistoryForWebApp request. The method will close the
// http.Response Body if it receives an error.
func (client RecommendationsClient) ListHistoryForWebAppSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListHistoryForWebAppResponder handles the response to the ListHistoryForWebApp request. The method always
// closes the http.Response Body.
func (client RecommendationsClient) ListHistoryForWebAppResponder(resp *http.Response) (result RecommendationCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listHistoryForWebAppNextResults retrieves the next set of results, if any.
func (client RecommendationsClient) listHistoryForWebAppNextResults(lastResults RecommendationCollection) (result RecommendationCollection, err error) {
	req, err := lastResults.recommendationCollectionPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web.RecommendationsClient", "listHistoryForWebAppNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListHistoryForWebAppSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web.RecommendationsClient", "listHistoryForWebAppNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListHistoryForWebAppResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.RecommendationsClient", "listHistoryForWebAppNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListHistoryForWebAppComplete enumerates all values, automatically crossing page boundaries as required.
func (client RecommendationsClient) ListHistoryForWebAppComplete(ctx context.Context, resourceGroupName string, siteName string, expiredOnly *bool, filter string) (result RecommendationCollectionIterator, err error) {
	result.page, err = client.ListHistoryForWebApp(ctx, resourceGroupName, siteName, expiredOnly, filter)
	return
}

// ListRecommendedRulesForWebApp get all recommendations for an app.
// Parameters:
// resourceGroupName - name of the resource group to which the resource belongs.
// siteName - name of the app.
// featured - specify <code>true</code> to return only the most critical recommendations. The default is
// <code>false</code>, which returns all recommendations.
// filter - return only channels specified in the filter. Filter is specified by using OData syntax. Example:
// $filter=channel eq 'Api' or channel eq 'Notification'
func (client RecommendationsClient) ListRecommendedRulesForWebApp(ctx context.Context, resourceGroupName string, siteName string, featured *bool, filter string) (result RecommendationCollectionPage, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+[^\.]$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("web.RecommendationsClient", "ListRecommendedRulesForWebApp", err.Error())
	}

	result.fn = client.listRecommendedRulesForWebAppNextResults
	req, err := client.ListRecommendedRulesForWebAppPreparer(ctx, resourceGroupName, siteName, featured, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.RecommendationsClient", "ListRecommendedRulesForWebApp", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListRecommendedRulesForWebAppSender(req)
	if err != nil {
		result.rc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "web.RecommendationsClient", "ListRecommendedRulesForWebApp", resp, "Failure sending request")
		return
	}

	result.rc, err = client.ListRecommendedRulesForWebAppResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.RecommendationsClient", "ListRecommendedRulesForWebApp", resp, "Failure responding to request")
	}

	return
}

// ListRecommendedRulesForWebAppPreparer prepares the ListRecommendedRulesForWebApp request.
func (client RecommendationsClient) ListRecommendedRulesForWebAppPreparer(ctx context.Context, resourceGroupName string, siteName string, featured *bool, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"siteName":          autorest.Encode("path", siteName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-02-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if featured != nil {
		queryParameters["featured"] = autorest.Encode("query", *featured)
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = filter
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/recommendations", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListRecommendedRulesForWebAppSender sends the ListRecommendedRulesForWebApp request. The method will close the
// http.Response Body if it receives an error.
func (client RecommendationsClient) ListRecommendedRulesForWebAppSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListRecommendedRulesForWebAppResponder handles the response to the ListRecommendedRulesForWebApp request. The method always
// closes the http.Response Body.
func (client RecommendationsClient) ListRecommendedRulesForWebAppResponder(resp *http.Response) (result RecommendationCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listRecommendedRulesForWebAppNextResults retrieves the next set of results, if any.
func (client RecommendationsClient) listRecommendedRulesForWebAppNextResults(lastResults RecommendationCollection) (result RecommendationCollection, err error) {
	req, err := lastResults.recommendationCollectionPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web.RecommendationsClient", "listRecommendedRulesForWebAppNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListRecommendedRulesForWebAppSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web.RecommendationsClient", "listRecommendedRulesForWebAppNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListRecommendedRulesForWebAppResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.RecommendationsClient", "listRecommendedRulesForWebAppNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListRecommendedRulesForWebAppComplete enumerates all values, automatically crossing page boundaries as required.
func (client RecommendationsClient) ListRecommendedRulesForWebAppComplete(ctx context.Context, resourceGroupName string, siteName string, featured *bool, filter string) (result RecommendationCollectionIterator, err error) {
	result.page, err = client.ListRecommendedRulesForWebApp(ctx, resourceGroupName, siteName, featured, filter)
	return
}

// ResetAllFilters reset all recommendation opt-out settings for a subscription.
func (client RecommendationsClient) ResetAllFilters(ctx context.Context) (result autorest.Response, err error) {
	req, err := client.ResetAllFiltersPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.RecommendationsClient", "ResetAllFilters", nil, "Failure preparing request")
		return
	}

	resp, err := client.ResetAllFiltersSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "web.RecommendationsClient", "ResetAllFilters", resp, "Failure sending request")
		return
	}

	result, err = client.ResetAllFiltersResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.RecommendationsClient", "ResetAllFilters", resp, "Failure responding to request")
	}

	return
}

// ResetAllFiltersPreparer prepares the ResetAllFilters request.
func (client RecommendationsClient) ResetAllFiltersPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-02-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Web/recommendations/reset", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ResetAllFiltersSender sends the ResetAllFilters request. The method will close the
// http.Response Body if it receives an error.
func (client RecommendationsClient) ResetAllFiltersSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ResetAllFiltersResponder handles the response to the ResetAllFilters request. The method always
// closes the http.Response Body.
func (client RecommendationsClient) ResetAllFiltersResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// ResetAllFiltersForWebApp reset all recommendation opt-out settings for an app.
// Parameters:
// resourceGroupName - name of the resource group to which the resource belongs.
// siteName - name of the app.
func (client RecommendationsClient) ResetAllFiltersForWebApp(ctx context.Context, resourceGroupName string, siteName string) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+[^\.]$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("web.RecommendationsClient", "ResetAllFiltersForWebApp", err.Error())
	}

	req, err := client.ResetAllFiltersForWebAppPreparer(ctx, resourceGroupName, siteName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.RecommendationsClient", "ResetAllFiltersForWebApp", nil, "Failure preparing request")
		return
	}

	resp, err := client.ResetAllFiltersForWebAppSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "web.RecommendationsClient", "ResetAllFiltersForWebApp", resp, "Failure sending request")
		return
	}

	result, err = client.ResetAllFiltersForWebAppResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.RecommendationsClient", "ResetAllFiltersForWebApp", resp, "Failure responding to request")
	}

	return
}

// ResetAllFiltersForWebAppPreparer prepares the ResetAllFiltersForWebApp request.
func (client RecommendationsClient) ResetAllFiltersForWebAppPreparer(ctx context.Context, resourceGroupName string, siteName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"siteName":          autorest.Encode("path", siteName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-02-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/recommendations/reset", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ResetAllFiltersForWebAppSender sends the ResetAllFiltersForWebApp request. The method will close the
// http.Response Body if it receives an error.
func (client RecommendationsClient) ResetAllFiltersForWebAppSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ResetAllFiltersForWebAppResponder handles the response to the ResetAllFiltersForWebApp request. The method always
// closes the http.Response Body.
func (client RecommendationsClient) ResetAllFiltersForWebAppResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}
