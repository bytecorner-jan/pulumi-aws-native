// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./httpNamespace";
export * from "./instance";
export * from "./privateDnsNamespace";
export * from "./publicDnsNamespace";
export * from "./service";

// Import resources to register:
import { HttpNamespace } from "./httpNamespace";
import { Instance } from "./instance";
import { PrivateDnsNamespace } from "./privateDnsNamespace";
import { PublicDnsNamespace } from "./publicDnsNamespace";
import { Service } from "./service";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:servicediscovery:HttpNamespace":
                return new HttpNamespace(name, <any>undefined, { urn })
            case "aws-native:servicediscovery:Instance":
                return new Instance(name, <any>undefined, { urn })
            case "aws-native:servicediscovery:PrivateDnsNamespace":
                return new PrivateDnsNamespace(name, <any>undefined, { urn })
            case "aws-native:servicediscovery:PublicDnsNamespace":
                return new PublicDnsNamespace(name, <any>undefined, { urn })
            case "aws-native:servicediscovery:Service":
                return new Service(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "servicediscovery", _module)
