// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./dbcluster";
export * from "./dbclusterParameterGroup";
export * from "./dbinstance";
export * from "./dbsubnetGroup";

// Import resources to register:
import { DBCluster } from "./dbcluster";
import { DBClusterParameterGroup } from "./dbclusterParameterGroup";
import { DBInstance } from "./dbinstance";
import { DBSubnetGroup } from "./dbsubnetGroup";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:docdb:DBCluster":
                return new DBCluster(name, <any>undefined, { urn })
            case "aws-native:docdb:DBClusterParameterGroup":
                return new DBClusterParameterGroup(name, <any>undefined, { urn })
            case "aws-native:docdb:DBInstance":
                return new DBInstance(name, <any>undefined, { urn })
            case "aws-native:docdb:DBSubnetGroup":
                return new DBSubnetGroup(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "docdb", _module)
