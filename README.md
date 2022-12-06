# Go API client for holodex

Holodex Public API. Successor to the HoloAPI v1

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 2.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import holodex "github.com/watsonindustries/go-holodex"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), holodex.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), holodex.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), holodex.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), holodex.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://holodex.net/api/v2*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**GetCachedLive**](docs/DefaultApi.md#getcachedlive) | **Get** /users/live | Quickly Access Live / Upcoming for a set of Channels
*DefaultApi* | [**GetChannels**](docs/DefaultApi.md#getchannels) | **Get** /channels | List Channels
*DefaultApi* | [**GetV2ChannelsChannelId**](docs/DefaultApi.md#getv2channelschannelid) | **Get** /channels/{channelId} | Get Channel Information
*DefaultApi* | [**GetV2ChannelsChannelIdClips**](docs/DefaultApi.md#getv2channelschannelidclips) | **Get** /channels/{channelId}/{type} | Query Videos Related to Channel
*DefaultApi* | [**GetVideosVideoId**](docs/DefaultApi.md#getvideosvideoid) | **Get** /videos/{videoId} | Get a single Video&#39;s metadata
*DefaultApi* | [**LiveGet**](docs/DefaultApi.md#liveget) | **Get** /live | Query Live and Upcoming Videos
*DefaultApi* | [**PostSearchCommentSearch**](docs/DefaultApi.md#postsearchcommentsearch) | **Post** /search/commentSearch | 
*DefaultApi* | [**PostSearchVideoSearch**](docs/DefaultApi.md#postsearchvideosearch) | **Post** /search/videoSearch | 
*DefaultApi* | [**VideosGet**](docs/DefaultApi.md#videosget) | **Get** /videos | Query Videos


## Documentation For Models

 - [Channel](docs/Channel.md)
 - [ChannelMin](docs/ChannelMin.md)
 - [Comment](docs/Comment.md)
 - [GetV2ChannelsChannelIdClips200Response](docs/GetV2ChannelsChannelIdClips200Response.md)
 - [GetV2ChannelsChannelIdClips200ResponseOneOf](docs/GetV2ChannelsChannelIdClips200ResponseOneOf.md)
 - [GetVideosVideoId200Response](docs/GetVideosVideoId200Response.md)
 - [GetVideosVideoId200ResponseAllOf](docs/GetVideosVideoId200ResponseAllOf.md)
 - [LiveGet200Response](docs/LiveGet200Response.md)
 - [LiveGet200ResponseOneOf](docs/LiveGet200ResponseOneOf.md)
 - [PostSearchCommentSearch200Response](docs/PostSearchCommentSearch200Response.md)
 - [PostSearchCommentSearch200ResponseOneOf](docs/PostSearchCommentSearch200ResponseOneOf.md)
 - [PostSearchCommentSearch200ResponseOneOfInner](docs/PostSearchCommentSearch200ResponseOneOfInner.md)
 - [PostSearchCommentSearch200ResponseOneOfInnerAllOf](docs/PostSearchCommentSearch200ResponseOneOfInnerAllOf.md)
 - [PostSearchCommentSearchRequest](docs/PostSearchCommentSearchRequest.md)
 - [PostSearchVideoSearch200Response](docs/PostSearchVideoSearch200Response.md)
 - [PostSearchVideoSearch200ResponseOneOf](docs/PostSearchVideoSearch200ResponseOneOf.md)
 - [PostSearchVideoSearchRequest](docs/PostSearchVideoSearchRequest.md)
 - [PostSearchVideoSearchRequestConditionsInner](docs/PostSearchVideoSearchRequestConditionsInner.md)
 - [Video](docs/Video.md)
 - [VideoFull](docs/VideoFull.md)
 - [VideoFullAllOf](docs/VideoFullAllOf.md)
 - [VideoWithChannel](docs/VideoWithChannel.md)
 - [VideoWithChannelAllOf](docs/VideoWithChannelAllOf.md)


## Documentation For Authorization



### STOPLIGHT

- **Type**: API key
- **API key parameter name**: X-APIKEY
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: X-APIKEY and passed in as the auth context for each request.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



