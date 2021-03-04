// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./application";
export * from "./applicationOutput";
export * from "./applicationReferenceDataSource";

// Import resources to register:
import { Application } from "./application";
import { ApplicationOutput } from "./applicationOutput";
import { ApplicationReferenceDataSource } from "./applicationReferenceDataSource";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:KinesisAnalytics:Application":
                return new Application(name, <any>undefined, { urn })
            case "aws-native:KinesisAnalytics:ApplicationOutput":
                return new ApplicationOutput(name, <any>undefined, { urn })
            case "aws-native:KinesisAnalytics:ApplicationReferenceDataSource":
                return new ApplicationReferenceDataSource(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "KinesisAnalytics", _module)
