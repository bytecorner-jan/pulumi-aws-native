// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./customDataIdentifier";
export * from "./findingsFilter";
export * from "./session";

// Export enums:
export * from "../types/enums/macie";

// Import resources to register:
import { CustomDataIdentifier } from "./customDataIdentifier";
import { FindingsFilter } from "./findingsFilter";
import { Session } from "./session";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:macie:CustomDataIdentifier":
                return new CustomDataIdentifier(name, <any>undefined, { urn })
            case "aws-native:macie:FindingsFilter":
                return new FindingsFilter(name, <any>undefined, { urn })
            case "aws-native:macie:Session":
                return new Session(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "macie", _module)
