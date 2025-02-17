// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./certificate";
export * from "./certificateAuthority";
export * from "./certificateAuthorityActivation";
export * from "./permission";

// Import resources to register:
import { Certificate } from "./certificate";
import { CertificateAuthority } from "./certificateAuthority";
import { CertificateAuthorityActivation } from "./certificateAuthorityActivation";
import { Permission } from "./permission";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:acmpca:Certificate":
                return new Certificate(name, <any>undefined, { urn })
            case "aws-native:acmpca:CertificateAuthority":
                return new CertificateAuthority(name, <any>undefined, { urn })
            case "aws-native:acmpca:CertificateAuthorityActivation":
                return new CertificateAuthorityActivation(name, <any>undefined, { urn })
            case "aws-native:acmpca:Permission":
                return new Permission(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "acmpca", _module)
