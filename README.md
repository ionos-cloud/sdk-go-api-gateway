# Go API client for ionoscloud

API Gateway is an application that acts as a \"front door\" for backend services and APIs, handling client requests and routing them to the appropriate backend.


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 0.0.1
- Package version: 1.1.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import ionoscloud "github.com/ionos-cloud/sdk-go-api-gateway"
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
ctx := context.WithValue(context.Background(), ionoscloud.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), ionoscloud.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

## Documentation for API Endpoints

All URIs are relative to *https://apigateway.de-txl.ionos.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*APIGatewaysApi* | [**ApigatewaysDelete**](docs/api/APIGatewaysApi.md#apigatewaysdelete) | **Delete** /gateways/{apigatewayId} | Delete Gateway
*APIGatewaysApi* | [**ApigatewaysFindById**](docs/api/APIGatewaysApi.md#apigatewaysfindbyid) | **Get** /gateways/{apigatewayId} | Retrieve Gateway
*APIGatewaysApi* | [**ApigatewaysGet**](docs/api/APIGatewaysApi.md#apigatewaysget) | **Get** /gateways | Retrieve all Apigateways
*APIGatewaysApi* | [**ApigatewaysPost**](docs/api/APIGatewaysApi.md#apigatewayspost) | **Post** /gateways | Create Gateway
*APIGatewaysApi* | [**ApigatewaysPut**](docs/api/APIGatewaysApi.md#apigatewaysput) | **Put** /gateways/{apigatewayId} | Ensure Gateway
*RoutesApi* | [**ApigatewaysRoutesDelete**](docs/api/RoutesApi.md#apigatewaysroutesdelete) | **Delete** /gateways/{apigatewayId}/routes/{routeId} | Delete Route
*RoutesApi* | [**ApigatewaysRoutesFindById**](docs/api/RoutesApi.md#apigatewaysroutesfindbyid) | **Get** /gateways/{apigatewayId}/routes/{routeId} | Retrieve Route
*RoutesApi* | [**ApigatewaysRoutesGet**](docs/api/RoutesApi.md#apigatewaysroutesget) | **Get** /gateways/{apigatewayId}/routes | Retrieve all Routes
*RoutesApi* | [**ApigatewaysRoutesPost**](docs/api/RoutesApi.md#apigatewaysroutespost) | **Post** /gateways/{apigatewayId}/routes | Create Route
*RoutesApi* | [**ApigatewaysRoutesPut**](docs/api/RoutesApi.md#apigatewaysroutesput) | **Put** /gateways/{apigatewayId}/routes/{routeId} | Ensure Route


## Documentation For Models

 - [Error](docs/models/Error.md)
 - [ErrorMessages](docs/models/ErrorMessages.md)
 - [Gateway](docs/models/Gateway.md)
 - [GatewayCreate](docs/models/GatewayCreate.md)
 - [GatewayCustomDomains](docs/models/GatewayCustomDomains.md)
 - [GatewayEnsure](docs/models/GatewayEnsure.md)
 - [GatewayRead](docs/models/GatewayRead.md)
 - [GatewayReadList](docs/models/GatewayReadList.md)
 - [GatewayReadListAllOf](docs/models/GatewayReadListAllOf.md)
 - [Links](docs/models/Links.md)
 - [Metadata](docs/models/Metadata.md)
 - [MetadataWithEndpoint](docs/models/MetadataWithEndpoint.md)
 - [MetadataWithEndpointAllOf](docs/models/MetadataWithEndpointAllOf.md)
 - [MetadataWithStatus](docs/models/MetadataWithStatus.md)
 - [MetadataWithStatusAllOf](docs/models/MetadataWithStatusAllOf.md)
 - [Pagination](docs/models/Pagination.md)
 - [Route](docs/models/Route.md)
 - [RouteCreate](docs/models/RouteCreate.md)
 - [RouteEnsure](docs/models/RouteEnsure.md)
 - [RouteRead](docs/models/RouteRead.md)
 - [RouteReadList](docs/models/RouteReadList.md)
 - [RouteReadListAllOf](docs/models/RouteReadListAllOf.md)
 - [RouteUpstreams](docs/models/RouteUpstreams.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### tokenAuth

- **Type**: HTTP Bearer token authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "BEARER_TOKEN_STRING")
r, err := client.Service.Operation(auth, args)
```


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



