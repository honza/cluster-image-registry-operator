package serialconsole

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

// ConsoleClient is the azure Virtual Machine Serial Console allows you to access serial console of a Virtual Machine
type ConsoleClient struct {
	BaseClient
}

// NewConsoleClient creates an instance of the ConsoleClient client.
func NewConsoleClient(subscriptionID string) ConsoleClient {
	return NewConsoleClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewConsoleClientWithBaseURI creates an instance of the ConsoleClient client.
func NewConsoleClientWithBaseURI(baseURI string, subscriptionID string) ConsoleClient {
	return ConsoleClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// DisableConsole disables Serial Console for a subscription
func (client ConsoleClient) DisableConsole(ctx context.Context) (result SetDisabledResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConsoleClient.DisableConsole")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DisableConsolePreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "serialconsole.ConsoleClient", "DisableConsole", nil, "Failure preparing request")
		return
	}

	resp, err := client.DisableConsoleSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "serialconsole.ConsoleClient", "DisableConsole", resp, "Failure sending request")
		return
	}

	result, err = client.DisableConsoleResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "serialconsole.ConsoleClient", "DisableConsole", resp, "Failure responding to request")
	}

	return
}

// DisableConsolePreparer prepares the DisableConsole request.
func (client ConsoleClient) DisableConsolePreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"default":        autorest.Encode("path", "default"),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.SerialConsole/consoleServices/{default}/disableConsole", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DisableConsoleSender sends the DisableConsole request. The method will close the
// http.Response Body if it receives an error.
func (client ConsoleClient) DisableConsoleSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// DisableConsoleResponder handles the response to the DisableConsole request. The method always
// closes the http.Response Body.
func (client ConsoleClient) DisableConsoleResponder(resp *http.Response) (result SetDisabledResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// EnableConsole enables Serial Console for a subscription
func (client ConsoleClient) EnableConsole(ctx context.Context) (result SetDisabledResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConsoleClient.EnableConsole")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.EnableConsolePreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "serialconsole.ConsoleClient", "EnableConsole", nil, "Failure preparing request")
		return
	}

	resp, err := client.EnableConsoleSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "serialconsole.ConsoleClient", "EnableConsole", resp, "Failure sending request")
		return
	}

	result, err = client.EnableConsoleResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "serialconsole.ConsoleClient", "EnableConsole", resp, "Failure responding to request")
	}

	return
}

// EnableConsolePreparer prepares the EnableConsole request.
func (client ConsoleClient) EnableConsolePreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"default":        autorest.Encode("path", "default"),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.SerialConsole/consoleServices/{default}/enableConsole", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// EnableConsoleSender sends the EnableConsole request. The method will close the
// http.Response Body if it receives an error.
func (client ConsoleClient) EnableConsoleSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// EnableConsoleResponder handles the response to the EnableConsole request. The method always
// closes the http.Response Body.
func (client ConsoleClient) EnableConsoleResponder(resp *http.Response) (result SetDisabledResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
