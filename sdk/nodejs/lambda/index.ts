// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./alias";
export * from "./codeSigningConfig";
export * from "./eventInvokeConfig";
export * from "./eventSourceMapping";
export * from "./function";
export * from "./layerVersion";
export * from "./layerVersionPermission";
export * from "./permission";
export * from "./version";

// Import resources to register:
import { Alias } from "./alias";
import { CodeSigningConfig } from "./codeSigningConfig";
import { EventInvokeConfig } from "./eventInvokeConfig";
import { EventSourceMapping } from "./eventSourceMapping";
import { Function } from "./function";
import { LayerVersion } from "./layerVersion";
import { LayerVersionPermission } from "./layerVersionPermission";
import { Permission } from "./permission";
import { Version } from "./version";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "cloudformation:Lambda:Alias":
                return new Alias(name, <any>undefined, { urn })
            case "cloudformation:Lambda:CodeSigningConfig":
                return new CodeSigningConfig(name, <any>undefined, { urn })
            case "cloudformation:Lambda:EventInvokeConfig":
                return new EventInvokeConfig(name, <any>undefined, { urn })
            case "cloudformation:Lambda:EventSourceMapping":
                return new EventSourceMapping(name, <any>undefined, { urn })
            case "cloudformation:Lambda:Function":
                return new Function(name, <any>undefined, { urn })
            case "cloudformation:Lambda:LayerVersion":
                return new LayerVersion(name, <any>undefined, { urn })
            case "cloudformation:Lambda:LayerVersionPermission":
                return new LayerVersionPermission(name, <any>undefined, { urn })
            case "cloudformation:Lambda:Permission":
                return new Permission(name, <any>undefined, { urn })
            case "cloudformation:Lambda:Version":
                return new Version(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("cloudformation", "Lambda", _module)
