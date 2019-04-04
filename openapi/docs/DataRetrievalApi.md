# \DataRetrievalApi

All URIs are relative to *https://virtserver.swaggerhub.com/vetstoria/Booking-API/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LocationsGet**](DataRetrievalApi.md#LocationsGet) | **Get** /locations | 
[**LocationsIdAppointmenttypesGet**](DataRetrievalApi.md#LocationsIdAppointmenttypesGet) | **Get** /locations/{id}/appointmenttypes | 
[**LocationsIdSchedulesGet**](DataRetrievalApi.md#LocationsIdSchedulesGet) | **Get** /locations/{id}/schedules | 
[**LocationsIdSpeciesGet**](DataRetrievalApi.md#LocationsIdSpeciesGet) | **Get** /locations/{id}/species | Species
[**PreferencesGet**](DataRetrievalApi.md#PreferencesGet) | **Get** /preferences | 
[**SlotsPost**](DataRetrievalApi.md#SlotsPost) | **Post** /slots | 


# **LocationsGet**
> []LocationDto LocationsGet(ctx, siteHash)


Get clinic locations

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteHash** | **string**| Site-Hash of the Vet Practice&#39;s account on Vetstoria platform  | 

### Return type

[**[]LocationDto**](LocationDTO.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocationsIdAppointmenttypesGet**
> []AppointmentTypeDto LocationsIdAppointmenttypesGet(ctx, siteHash, id)


Get appointment types

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteHash** | **string**| Site-Hash of the Vet Practice&#39;s account on Vetstoria platform  | 
  **id** | **int32**| Location ID | 

### Return type

[**[]AppointmentTypeDto**](AppointmentTypeDTO.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocationsIdSchedulesGet**
> []ScheduleDto LocationsIdSchedulesGet(ctx, siteHash, id)


Get schedules

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteHash** | **string**| Site-Hash of the Vet Practice&#39;s account on Vetstoria platform  | 
  **id** | **int32**| Location ID | 

### Return type

[**[]ScheduleDto**](ScheduleDTO.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocationsIdSpeciesGet**
> []SpeciesDto LocationsIdSpeciesGet(ctx, siteHash, id)
Species

Get species

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteHash** | **string**| Site-Hash of the Vet Practice&#39;s account on Vetstoria platform  | 
  **id** | **int32**| Location ID | 

### Return type

[**[]SpeciesDto**](SpeciesDTO.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PreferencesGet**
> InlineResponse2001 PreferencesGet(ctx, siteHash)


Get metadata relevant to the booking process

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteHash** | **string**| Site-Hash of the Vet Practice&#39;s account on Vetstoria platform  | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SlotsPost**
> []AvailabilityResponseSlotsDto SlotsPost(ctx, siteHash, availabilityRequestDto)


Find slots for booking based on the specified parameters

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteHash** | **string**| Site-Hash of the Vet Practice&#39;s account on Vetstoria platform  | 
  **availabilityRequestDto** | [**AvailabilityRequestDto**](AvailabilityRequestDto.md)|  | 

### Return type

[**[]AvailabilityResponseSlotsDto**](AvailabilityResponseSlotsDTO.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

