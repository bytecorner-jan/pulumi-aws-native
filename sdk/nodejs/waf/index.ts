// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./byteMatchSet";
export * from "./ipset";
export * from "./rule";
export * from "./sizeConstraintSet";
export * from "./sqlInjectionMatchSet";
export * from "./webACL";
export * from "./xssMatchSet";

// Import resources to register:
import { ByteMatchSet } from "./byteMatchSet";
import { IPSet } from "./ipset";
import { Rule } from "./rule";
import { SizeConstraintSet } from "./sizeConstraintSet";
import { SqlInjectionMatchSet } from "./sqlInjectionMatchSet";
import { WebACL } from "./webACL";
import { XssMatchSet } from "./xssMatchSet";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:waf:ByteMatchSet":
                return new ByteMatchSet(name, <any>undefined, { urn })
            case "aws-native:waf:IPSet":
                return new IPSet(name, <any>undefined, { urn })
            case "aws-native:waf:Rule":
                return new Rule(name, <any>undefined, { urn })
            case "aws-native:waf:SizeConstraintSet":
                return new SizeConstraintSet(name, <any>undefined, { urn })
            case "aws-native:waf:SqlInjectionMatchSet":
                return new SqlInjectionMatchSet(name, <any>undefined, { urn })
            case "aws-native:waf:WebACL":
                return new WebACL(name, <any>undefined, { urn })
            case "aws-native:waf:XssMatchSet":
                return new XssMatchSet(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "waf", _module)
