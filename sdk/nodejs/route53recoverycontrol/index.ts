// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./cluster";
export * from "./controlPanel";
export * from "./routingControl";
export * from "./safetyRule";

// Import resources to register:
import { Cluster } from "./cluster";
import { ControlPanel } from "./controlPanel";
import { RoutingControl } from "./routingControl";
import { SafetyRule } from "./safetyRule";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:Route53RecoveryControl:Cluster":
                return new Cluster(name, <any>undefined, { urn })
            case "aws-native:Route53RecoveryControl:ControlPanel":
                return new ControlPanel(name, <any>undefined, { urn })
            case "aws-native:Route53RecoveryControl:RoutingControl":
                return new RoutingControl(name, <any>undefined, { urn })
            case "aws-native:Route53RecoveryControl:SafetyRule":
                return new SafetyRule(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "Route53RecoveryControl", _module)
