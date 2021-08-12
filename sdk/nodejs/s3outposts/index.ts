// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./accessPoint";
export * from "./bucket";
export * from "./bucketPolicy";
export * from "./endpoint";

// Import resources to register:
import { AccessPoint } from "./accessPoint";
import { Bucket } from "./bucket";
import { BucketPolicy } from "./bucketPolicy";
import { Endpoint } from "./endpoint";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:S3Outposts:AccessPoint":
                return new AccessPoint(name, <any>undefined, { urn })
            case "aws-native:S3Outposts:Bucket":
                return new Bucket(name, <any>undefined, { urn })
            case "aws-native:S3Outposts:BucketPolicy":
                return new BucketPolicy(name, <any>undefined, { urn })
            case "aws-native:S3Outposts:Endpoint":
                return new Endpoint(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "S3Outposts", _module)
