package consumption

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
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// ReservationsDetailsClient is the consumption management client provides access to consumption resources for Azure
// Enterprise Subscriptions.
type ReservationsDetailsClient struct {
	BaseClient
}

// NewReservationsDetailsClient creates an instance of the ReservationsDetailsClient client.
func NewReservationsDetailsClient(subscriptionID string) ReservationsDetailsClient {
	return NewReservationsDetailsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewReservationsDetailsClientWithBaseURI creates an instance of the ReservationsDetailsClient client.
func NewReservationsDetailsClientWithBaseURI(baseURI string, subscriptionID string) ReservationsDetailsClient {
	return ReservationsDetailsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// ListByReservationOrder lists the reservations details for provided date range.
// Parameters:
// reservationOrderID - order Id of the reservation
// filter - filter reservation details by date range. The properties/UsageDate for start date and end date. The
// filter supports 'le' and  'ge'
func (client ReservationsDetailsClient) ListByReservationOrder(ctx context.Context, reservationOrderID string, filter string) (result ReservationDetailsListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReservationsDetailsClient.ListByReservationOrder")
		defer func() {
			sc := -1
			if result.rdlr.Response.Response != nil {
				sc = result.rdlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByReservationOrderNextResults
	req, err := client.ListByReservationOrderPreparer(ctx, reservationOrderID, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.ReservationsDetailsClient", "ListByReservationOrder", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByReservationOrderSender(req)
	if err != nil {
		result.rdlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "consumption.ReservationsDetailsClient", "ListByReservationOrder", resp, "Failure sending request")
		return
	}

	result.rdlr, err = client.ListByReservationOrderResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.ReservationsDetailsClient", "ListByReservationOrder", resp, "Failure responding to request")
	}

	return
}

// ListByReservationOrderPreparer prepares the ListByReservationOrder request.
func (client ReservationsDetailsClient) ListByReservationOrderPreparer(ctx context.Context, reservationOrderID string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"reservationOrderId": autorest.Encode("path", reservationOrderID),
	}

	const APIVersion = "2018-10-01"
	queryParameters := map[string]interface{}{
		"$filter":     autorest.Encode("query", filter),
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Capacity/reservationorders/{reservationOrderId}/providers/Microsoft.Consumption/reservationDetails", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByReservationOrderSender sends the ListByReservationOrder request. The method will close the
// http.Response Body if it receives an error.
func (client ReservationsDetailsClient) ListByReservationOrderSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListByReservationOrderResponder handles the response to the ListByReservationOrder request. The method always
// closes the http.Response Body.
func (client ReservationsDetailsClient) ListByReservationOrderResponder(resp *http.Response) (result ReservationDetailsListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByReservationOrderNextResults retrieves the next set of results, if any.
func (client ReservationsDetailsClient) listByReservationOrderNextResults(ctx context.Context, lastResults ReservationDetailsListResult) (result ReservationDetailsListResult, err error) {
	req, err := lastResults.reservationDetailsListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "consumption.ReservationsDetailsClient", "listByReservationOrderNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByReservationOrderSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "consumption.ReservationsDetailsClient", "listByReservationOrderNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByReservationOrderResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.ReservationsDetailsClient", "listByReservationOrderNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByReservationOrderComplete enumerates all values, automatically crossing page boundaries as required.
func (client ReservationsDetailsClient) ListByReservationOrderComplete(ctx context.Context, reservationOrderID string, filter string) (result ReservationDetailsListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReservationsDetailsClient.ListByReservationOrder")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByReservationOrder(ctx, reservationOrderID, filter)
	return
}

// ListByReservationOrderAndReservation lists the reservations details for provided date range.
// Parameters:
// reservationOrderID - order Id of the reservation
// reservationID - id of the reservation
// filter - filter reservation details by date range. The properties/UsageDate for start date and end date. The
// filter supports 'le' and  'ge'
func (client ReservationsDetailsClient) ListByReservationOrderAndReservation(ctx context.Context, reservationOrderID string, reservationID string, filter string) (result ReservationDetailsListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReservationsDetailsClient.ListByReservationOrderAndReservation")
		defer func() {
			sc := -1
			if result.rdlr.Response.Response != nil {
				sc = result.rdlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByReservationOrderAndReservationNextResults
	req, err := client.ListByReservationOrderAndReservationPreparer(ctx, reservationOrderID, reservationID, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.ReservationsDetailsClient", "ListByReservationOrderAndReservation", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByReservationOrderAndReservationSender(req)
	if err != nil {
		result.rdlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "consumption.ReservationsDetailsClient", "ListByReservationOrderAndReservation", resp, "Failure sending request")
		return
	}

	result.rdlr, err = client.ListByReservationOrderAndReservationResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.ReservationsDetailsClient", "ListByReservationOrderAndReservation", resp, "Failure responding to request")
	}

	return
}

// ListByReservationOrderAndReservationPreparer prepares the ListByReservationOrderAndReservation request.
func (client ReservationsDetailsClient) ListByReservationOrderAndReservationPreparer(ctx context.Context, reservationOrderID string, reservationID string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"reservationId":      autorest.Encode("path", reservationID),
		"reservationOrderId": autorest.Encode("path", reservationOrderID),
	}

	const APIVersion = "2018-10-01"
	queryParameters := map[string]interface{}{
		"$filter":     autorest.Encode("query", filter),
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Capacity/reservationorders/{reservationOrderId}/reservations/{reservationId}/providers/Microsoft.Consumption/reservationDetails", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByReservationOrderAndReservationSender sends the ListByReservationOrderAndReservation request. The method will close the
// http.Response Body if it receives an error.
func (client ReservationsDetailsClient) ListByReservationOrderAndReservationSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListByReservationOrderAndReservationResponder handles the response to the ListByReservationOrderAndReservation request. The method always
// closes the http.Response Body.
func (client ReservationsDetailsClient) ListByReservationOrderAndReservationResponder(resp *http.Response) (result ReservationDetailsListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByReservationOrderAndReservationNextResults retrieves the next set of results, if any.
func (client ReservationsDetailsClient) listByReservationOrderAndReservationNextResults(ctx context.Context, lastResults ReservationDetailsListResult) (result ReservationDetailsListResult, err error) {
	req, err := lastResults.reservationDetailsListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "consumption.ReservationsDetailsClient", "listByReservationOrderAndReservationNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByReservationOrderAndReservationSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "consumption.ReservationsDetailsClient", "listByReservationOrderAndReservationNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByReservationOrderAndReservationResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.ReservationsDetailsClient", "listByReservationOrderAndReservationNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByReservationOrderAndReservationComplete enumerates all values, automatically crossing page boundaries as required.
func (client ReservationsDetailsClient) ListByReservationOrderAndReservationComplete(ctx context.Context, reservationOrderID string, reservationID string, filter string) (result ReservationDetailsListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReservationsDetailsClient.ListByReservationOrderAndReservation")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByReservationOrderAndReservation(ctx, reservationOrderID, reservationID, filter)
	return
}
