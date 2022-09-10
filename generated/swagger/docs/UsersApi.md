# {{classname}}

All URIs are relative to *http://localhost:{port}/{version}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersUserIdAttributesGet**](UsersApi.md#UsersUserIdAttributesGet) | **Get** /users/{userId}/attributes | 属性情報群取得

# **UsersUserIdAttributesGet**
> UserAttribute UsersUserIdAttributesGet(ctx, userId)
属性情報群取得

氏名や生年月日等の属性情報群を取得する。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | [**string**](.md)|  | 

### Return type

[**UserAttribute**](UserAttribute.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

