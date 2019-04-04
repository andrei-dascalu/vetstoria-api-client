# \BookingApi

All URIs are relative to *https://virtserver.swaggerhub.com/vetstoria/Booking-API/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BookingsIdDelete**](BookingApi.md#BookingsIdDelete) | **Delete** /bookings/{id} | 
[**BookingsIdGet**](BookingApi.md#BookingsIdGet) | **Get** /bookings/{id} | 
[**BookingsIdPut**](BookingApi.md#BookingsIdPut) | **Put** /bookings/{id} | 
[**BookingsPost**](BookingApi.md#BookingsPost) | **Post** /bookings | 


# **BookingsIdDelete**
> BookingsIdDelete(ctx, id, siteHash)


Delete booking

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Booking ID | 
  **siteHash** | **string**| Site-Hash of the Vet Practice&#39;s account on Vetstoria platform  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BookingsIdGet**
> BookingDto BookingsIdGet(ctx, id, siteHash)


Get booking

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Booking ID | 
  **siteHash** | **string**| Site-Hash of the Vet Practice&#39;s account on Vetstoria platform  | 

### Return type

[**BookingDto**](BookingDTO.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BookingsIdPut**
> BookingDto BookingsIdPut(ctx, id, siteHash, bookingRequestDto)


Update booking

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Booking ID | 
  **siteHash** | **string**| Site-Hash of the Vet Practice&#39;s account on Vetstoria platform  | 
  **bookingRequestDto** | [**BookingRequestDto**](BookingRequestDto.md)|  | 

### Return type

[**BookingDto**](BookingDTO.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BookingsPost**
> BookingDto BookingsPost(ctx, siteHash, bookingRequestDto)


Create booking

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteHash** | **string**| Site-Hash of the Vet Practice&#39;s account on Vetstoria platform  | 
  **bookingRequestDto** | [**BookingRequestDto**](BookingRequestDto.md)|  | 

### Return type

[**BookingDto**](BookingDTO.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

