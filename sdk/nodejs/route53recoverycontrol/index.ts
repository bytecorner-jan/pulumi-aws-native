// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./cluster";
export * from "./controlPanel";
export * from "./routingControl";
export * from "./safetyRule";

// Export enums:
export * from "../types/enums/route53recoverycontrol";

// Import resources to register:
import { Cluster } from "./cluster";
import { ControlPanel } from "./controlPanel";
import { RoutingControl } from "./routingControl";
import { SafetyRule } from "./safetyRule";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:route53recoverycontrol:Cluster":
                return new Cluster(name, <any>undefined, { urn })
            case "aws-native:route53recoverycontrol:ControlPanel":
                return new ControlPanel(name, <any>undefined, { urn })
            case "aws-native:route53recoverycontrol:RoutingControl":
                return new RoutingControl(name, <any>undefined, { urn })
            case "aws-native:route53recoverycontrol:SafetyRule":
                return new SafetyRule(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "route53recoverycontrol", _module)
