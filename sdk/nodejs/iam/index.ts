// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./accessKey";
export * from "./group";
export * from "./instanceProfile";
export * from "./managedPolicy";
export * from "./oidcprovider";
export * from "./policy";
export * from "./role";
export * from "./samlprovider";
export * from "./serverCertificate";
export * from "./serviceLinkedRole";
export * from "./user";
export * from "./userToGroupAddition";
export * from "./virtualMFADevice";

// Import resources to register:
import { AccessKey } from "./accessKey";
import { Group } from "./group";
import { InstanceProfile } from "./instanceProfile";
import { ManagedPolicy } from "./managedPolicy";
import { OIDCProvider } from "./oidcprovider";
import { Policy } from "./policy";
import { Role } from "./role";
import { SAMLProvider } from "./samlprovider";
import { ServerCertificate } from "./serverCertificate";
import { ServiceLinkedRole } from "./serviceLinkedRole";
import { User } from "./user";
import { UserToGroupAddition } from "./userToGroupAddition";
import { VirtualMFADevice } from "./virtualMFADevice";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:iam:AccessKey":
                return new AccessKey(name, <any>undefined, { urn })
            case "aws-native:iam:Group":
                return new Group(name, <any>undefined, { urn })
            case "aws-native:iam:InstanceProfile":
                return new InstanceProfile(name, <any>undefined, { urn })
            case "aws-native:iam:ManagedPolicy":
                return new ManagedPolicy(name, <any>undefined, { urn })
            case "aws-native:iam:OIDCProvider":
                return new OIDCProvider(name, <any>undefined, { urn })
            case "aws-native:iam:Policy":
                return new Policy(name, <any>undefined, { urn })
            case "aws-native:iam:Role":
                return new Role(name, <any>undefined, { urn })
            case "aws-native:iam:SAMLProvider":
                return new SAMLProvider(name, <any>undefined, { urn })
            case "aws-native:iam:ServerCertificate":
                return new ServerCertificate(name, <any>undefined, { urn })
            case "aws-native:iam:ServiceLinkedRole":
                return new ServiceLinkedRole(name, <any>undefined, { urn })
            case "aws-native:iam:User":
                return new User(name, <any>undefined, { urn })
            case "aws-native:iam:UserToGroupAddition":
                return new UserToGroupAddition(name, <any>undefined, { urn })
            case "aws-native:iam:VirtualMFADevice":
                return new VirtualMFADevice(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "iam", _module)
