// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./api";
export * from "./apiGatewayManagedOverrides";
export * from "./apiMapping";
export * from "./authorizer";
export * from "./deployment";
export * from "./domainName";
export * from "./integration";
export * from "./integrationResponse";
export * from "./model";
export * from "./route";
export * from "./routeResponse";
export * from "./stage";
export * from "./vpcLink";

// Import resources to register:
import { Api } from "./api";
import { ApiGatewayManagedOverrides } from "./apiGatewayManagedOverrides";
import { ApiMapping } from "./apiMapping";
import { Authorizer } from "./authorizer";
import { Deployment } from "./deployment";
import { DomainName } from "./domainName";
import { Integration } from "./integration";
import { IntegrationResponse } from "./integrationResponse";
import { Model } from "./model";
import { Route } from "./route";
import { RouteResponse } from "./routeResponse";
import { Stage } from "./stage";
import { VpcLink } from "./vpcLink";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:apigatewayv2:Api":
                return new Api(name, <any>undefined, { urn })
            case "aws-native:apigatewayv2:ApiGatewayManagedOverrides":
                return new ApiGatewayManagedOverrides(name, <any>undefined, { urn })
            case "aws-native:apigatewayv2:ApiMapping":
                return new ApiMapping(name, <any>undefined, { urn })
            case "aws-native:apigatewayv2:Authorizer":
                return new Authorizer(name, <any>undefined, { urn })
            case "aws-native:apigatewayv2:Deployment":
                return new Deployment(name, <any>undefined, { urn })
            case "aws-native:apigatewayv2:DomainName":
                return new DomainName(name, <any>undefined, { urn })
            case "aws-native:apigatewayv2:Integration":
                return new Integration(name, <any>undefined, { urn })
            case "aws-native:apigatewayv2:IntegrationResponse":
                return new IntegrationResponse(name, <any>undefined, { urn })
            case "aws-native:apigatewayv2:Model":
                return new Model(name, <any>undefined, { urn })
            case "aws-native:apigatewayv2:Route":
                return new Route(name, <any>undefined, { urn })
            case "aws-native:apigatewayv2:RouteResponse":
                return new RouteResponse(name, <any>undefined, { urn })
            case "aws-native:apigatewayv2:Stage":
                return new Stage(name, <any>undefined, { urn })
            case "aws-native:apigatewayv2:VpcLink":
                return new VpcLink(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "apigatewayv2", _module)
