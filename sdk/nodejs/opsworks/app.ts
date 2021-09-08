// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html
 */
export class App extends pulumi.CustomResource {
    /**
     * Get an existing App resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): App {
        return new App(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:opsworks:App';

    /**
     * Returns true if the given object is an instance of App.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is App {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === App.__pulumiType;
    }

    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-appsource
     */
    public readonly appSource!: pulumi.Output<outputs.opsworks.AppSource | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-attributes
     */
    public readonly attributes!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-datasources
     */
    public readonly dataSources!: pulumi.Output<outputs.opsworks.AppDataSource[] | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-description
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-domains
     */
    public readonly domains!: pulumi.Output<string[] | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-enablessl
     */
    public readonly enableSsl!: pulumi.Output<boolean | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-environment
     */
    public readonly environment!: pulumi.Output<outputs.opsworks.AppEnvironmentVariable[] | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-name
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-shortname
     */
    public readonly shortname!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-sslconfiguration
     */
    public readonly sslConfiguration!: pulumi.Output<outputs.opsworks.AppSslConfiguration | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-stackid
     */
    public readonly stackId!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-type
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a App resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AppArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            if ((!args || args.stackId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'stackId'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            inputs["appSource"] = args ? args.appSource : undefined;
            inputs["attributes"] = args ? args.attributes : undefined;
            inputs["dataSources"] = args ? args.dataSources : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["domains"] = args ? args.domains : undefined;
            inputs["enableSsl"] = args ? args.enableSsl : undefined;
            inputs["environment"] = args ? args.environment : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["shortname"] = args ? args.shortname : undefined;
            inputs["sslConfiguration"] = args ? args.sslConfiguration : undefined;
            inputs["stackId"] = args ? args.stackId : undefined;
            inputs["type"] = args ? args.type : undefined;
        } else {
            inputs["appSource"] = undefined /*out*/;
            inputs["attributes"] = undefined /*out*/;
            inputs["dataSources"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["domains"] = undefined /*out*/;
            inputs["enableSsl"] = undefined /*out*/;
            inputs["environment"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["shortname"] = undefined /*out*/;
            inputs["sslConfiguration"] = undefined /*out*/;
            inputs["stackId"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(App.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a App resource.
 */
export interface AppArgs {
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-appsource
     */
    appSource?: pulumi.Input<inputs.opsworks.AppSourceArgs>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-attributes
     */
    attributes?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-datasources
     */
    dataSources?: pulumi.Input<pulumi.Input<inputs.opsworks.AppDataSourceArgs>[]>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-description
     */
    description?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-domains
     */
    domains?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-enablessl
     */
    enableSsl?: pulumi.Input<boolean>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-environment
     */
    environment?: pulumi.Input<pulumi.Input<inputs.opsworks.AppEnvironmentVariableArgs>[]>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-name
     */
    name: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-shortname
     */
    shortname?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-sslconfiguration
     */
    sslConfiguration?: pulumi.Input<inputs.opsworks.AppSslConfigurationArgs>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-stackid
     */
    stackId: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-type
     */
    type: pulumi.Input<string>;
}
