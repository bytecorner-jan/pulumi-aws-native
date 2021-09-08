// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./cluster";
export * from "./parameterGroup";
export * from "./subnetGroup";

// Import resources to register:
import { Cluster } from "./cluster";
import { ParameterGroup } from "./parameterGroup";
import { SubnetGroup } from "./subnetGroup";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:dax:Cluster":
                return new Cluster(name, <any>undefined, { urn })
            case "aws-native:dax:ParameterGroup":
                return new ParameterGroup(name, <any>undefined, { urn })
            case "aws-native:dax:SubnetGroup":
                return new SubnetGroup(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "dax", _module)
