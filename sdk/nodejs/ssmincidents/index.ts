// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./replicationSet";
export * from "./responsePlan";

// Export enums:
export * from "../types/enums/ssmincidents";

// Import resources to register:
import { ReplicationSet } from "./replicationSet";
import { ResponsePlan } from "./responsePlan";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:ssmincidents:ReplicationSet":
                return new ReplicationSet(name, <any>undefined, { urn })
            case "aws-native:ssmincidents:ResponsePlan":
                return new ResponsePlan(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "ssmincidents", _module)
