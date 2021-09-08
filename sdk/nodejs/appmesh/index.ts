// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./gatewayRoute";
export * from "./mesh";
export * from "./route";
export * from "./virtualGateway";
export * from "./virtualNode";
export * from "./virtualRouter";
export * from "./virtualService";

// Import resources to register:
import { GatewayRoute } from "./gatewayRoute";
import { Mesh } from "./mesh";
import { Route } from "./route";
import { VirtualGateway } from "./virtualGateway";
import { VirtualNode } from "./virtualNode";
import { VirtualRouter } from "./virtualRouter";
import { VirtualService } from "./virtualService";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:appmesh:GatewayRoute":
                return new GatewayRoute(name, <any>undefined, { urn })
            case "aws-native:appmesh:Mesh":
                return new Mesh(name, <any>undefined, { urn })
            case "aws-native:appmesh:Route":
                return new Route(name, <any>undefined, { urn })
            case "aws-native:appmesh:VirtualGateway":
                return new VirtualGateway(name, <any>undefined, { urn })
            case "aws-native:appmesh:VirtualNode":
                return new VirtualNode(name, <any>undefined, { urn })
            case "aws-native:appmesh:VirtualRouter":
                return new VirtualRouter(name, <any>undefined, { urn })
            case "aws-native:appmesh:VirtualService":
                return new VirtualService(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "appmesh", _module)
