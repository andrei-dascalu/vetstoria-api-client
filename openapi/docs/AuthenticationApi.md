# \AuthenticationApi

All URIs are relative to *https://virtserver.swaggerhub.com/vetstoria/Booking-API/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PartnersAuthenticationPost**](AuthenticationApi.md#PartnersAuthenticationPost) | **Post** /partners/authentication | 


# **PartnersAuthenticationPost**
> InlineResponse200 PartnersAuthenticationPost(ctx, optional)


Authenticate the partner

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PartnersAuthenticationPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PartnersAuthenticationPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject** | [**optional.Interface of InlineObject**](InlineObject.md)|  | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

