// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./account";
export * from "./apiKey";
export * from "./authorizer";
export * from "./basePathMapping";
export * from "./clientCertificate";
export * from "./deployment";
export * from "./documentationPart";
export * from "./documentationVersion";
export * from "./domainName";
export * from "./gatewayResponse";
export * from "./method";
export * from "./model";
export * from "./requestValidator";
export * from "./resource";
export * from "./restApi";
export * from "./stage";
export * from "./usagePlan";
export * from "./usagePlanKey";
export * from "./vpcLink";

// Import resources to register:
import { Account } from "./account";
import { ApiKey } from "./apiKey";
import { Authorizer } from "./authorizer";
import { BasePathMapping } from "./basePathMapping";
import { ClientCertificate } from "./clientCertificate";
import { Deployment } from "./deployment";
import { DocumentationPart } from "./documentationPart";
import { DocumentationVersion } from "./documentationVersion";
import { DomainName } from "./domainName";
import { GatewayResponse } from "./gatewayResponse";
import { Method } from "./method";
import { Model } from "./model";
import { RequestValidator } from "./requestValidator";
import { Resource } from "./resource";
import { RestApi } from "./restApi";
import { Stage } from "./stage";
import { UsagePlan } from "./usagePlan";
import { UsagePlanKey } from "./usagePlanKey";
import { VpcLink } from "./vpcLink";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "cloudformation:ApiGateway:Account":
                return new Account(name, <any>undefined, { urn })
            case "cloudformation:ApiGateway:ApiKey":
                return new ApiKey(name, <any>undefined, { urn })
            case "cloudformation:ApiGateway:Authorizer":
                return new Authorizer(name, <any>undefined, { urn })
            case "cloudformation:ApiGateway:BasePathMapping":
                return new BasePathMapping(name, <any>undefined, { urn })
            case "cloudformation:ApiGateway:ClientCertificate":
                return new ClientCertificate(name, <any>undefined, { urn })
            case "cloudformation:ApiGateway:Deployment":
                return new Deployment(name, <any>undefined, { urn })
            case "cloudformation:ApiGateway:DocumentationPart":
                return new DocumentationPart(name, <any>undefined, { urn })
            case "cloudformation:ApiGateway:DocumentationVersion":
                return new DocumentationVersion(name, <any>undefined, { urn })
            case "cloudformation:ApiGateway:DomainName":
                return new DomainName(name, <any>undefined, { urn })
            case "cloudformation:ApiGateway:GatewayResponse":
                return new GatewayResponse(name, <any>undefined, { urn })
            case "cloudformation:ApiGateway:Method":
                return new Method(name, <any>undefined, { urn })
            case "cloudformation:ApiGateway:Model":
                return new Model(name, <any>undefined, { urn })
            case "cloudformation:ApiGateway:RequestValidator":
                return new RequestValidator(name, <any>undefined, { urn })
            case "cloudformation:ApiGateway:Resource":
                return new Resource(name, <any>undefined, { urn })
            case "cloudformation:ApiGateway:RestApi":
                return new RestApi(name, <any>undefined, { urn })
            case "cloudformation:ApiGateway:Stage":
                return new Stage(name, <any>undefined, { urn })
            case "cloudformation:ApiGateway:UsagePlan":
                return new UsagePlan(name, <any>undefined, { urn })
            case "cloudformation:ApiGateway:UsagePlanKey":
                return new UsagePlanKey(name, <any>undefined, { urn })
            case "cloudformation:ApiGateway:VpcLink":
                return new VpcLink(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("cloudformation", "ApiGateway", _module)
