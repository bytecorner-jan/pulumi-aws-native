// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./apiDestination";
export * from "./archive";
export * from "./connection";
export * from "./eventBus";
export * from "./eventBusPolicy";
export * from "./rule";

// Import resources to register:
import { ApiDestination } from "./apiDestination";
import { Archive } from "./archive";
import { Connection } from "./connection";
import { EventBus } from "./eventBus";
import { EventBusPolicy } from "./eventBusPolicy";
import { Rule } from "./rule";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:events:ApiDestination":
                return new ApiDestination(name, <any>undefined, { urn })
            case "aws-native:events:Archive":
                return new Archive(name, <any>undefined, { urn })
            case "aws-native:events:Connection":
                return new Connection(name, <any>undefined, { urn })
            case "aws-native:events:EventBus":
                return new EventBus(name, <any>undefined, { urn })
            case "aws-native:events:EventBusPolicy":
                return new EventBusPolicy(name, <any>undefined, { urn })
            case "aws-native:events:Rule":
                return new Rule(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "events", _module)
