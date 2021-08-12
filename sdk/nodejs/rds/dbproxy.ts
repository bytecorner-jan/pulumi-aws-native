// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxy.html
 */
export class DBProxy extends pulumi.CustomResource {
    /**
     * Get an existing DBProxy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DBProxy {
        return new DBProxy(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:RDS:DBProxy';

    /**
     * Returns true if the given object is an instance of DBProxy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DBProxy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DBProxy.__pulumiType;
    }

    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxy.html#cfn-rds-dbproxy-auth
     */
    public readonly auth!: pulumi.Output<outputs.RDS.DBProxyAuthFormat[]>;
    public /*out*/ readonly dBProxyArn!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxy.html#cfn-rds-dbproxy-dbproxyname
     */
    public readonly dBProxyName!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxy.html#cfn-rds-dbproxy-debuglogging
     */
    public readonly debugLogging!: pulumi.Output<boolean | undefined>;
    public /*out*/ readonly endpoint!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxy.html#cfn-rds-dbproxy-enginefamily
     */
    public readonly engineFamily!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxy.html#cfn-rds-dbproxy-idleclienttimeout
     */
    public readonly idleClientTimeout!: pulumi.Output<number | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxy.html#cfn-rds-dbproxy-requiretls
     */
    public readonly requireTLS!: pulumi.Output<boolean | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxy.html#cfn-rds-dbproxy-rolearn
     */
    public readonly roleArn!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxy.html#cfn-rds-dbproxy-tags
     */
    public readonly tags!: pulumi.Output<outputs.RDS.DBProxyTagFormat[] | undefined>;
    public /*out*/ readonly vpcId!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxy.html#cfn-rds-dbproxy-vpcsecuritygroupids
     */
    public readonly vpcSecurityGroupIds!: pulumi.Output<string[] | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxy.html#cfn-rds-dbproxy-vpcsubnetids
     */
    public readonly vpcSubnetIds!: pulumi.Output<string[]>;

    /**
     * Create a DBProxy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DBProxyArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.auth === undefined) && !opts.urn) {
                throw new Error("Missing required property 'auth'");
            }
            if ((!args || args.dBProxyName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dBProxyName'");
            }
            if ((!args || args.engineFamily === undefined) && !opts.urn) {
                throw new Error("Missing required property 'engineFamily'");
            }
            if ((!args || args.roleArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'roleArn'");
            }
            if ((!args || args.vpcSubnetIds === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcSubnetIds'");
            }
            inputs["auth"] = args ? args.auth : undefined;
            inputs["dBProxyName"] = args ? args.dBProxyName : undefined;
            inputs["debugLogging"] = args ? args.debugLogging : undefined;
            inputs["engineFamily"] = args ? args.engineFamily : undefined;
            inputs["idleClientTimeout"] = args ? args.idleClientTimeout : undefined;
            inputs["requireTLS"] = args ? args.requireTLS : undefined;
            inputs["roleArn"] = args ? args.roleArn : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["vpcSecurityGroupIds"] = args ? args.vpcSecurityGroupIds : undefined;
            inputs["vpcSubnetIds"] = args ? args.vpcSubnetIds : undefined;
            inputs["dBProxyArn"] = undefined /*out*/;
            inputs["endpoint"] = undefined /*out*/;
            inputs["vpcId"] = undefined /*out*/;
        } else {
            inputs["auth"] = undefined /*out*/;
            inputs["dBProxyArn"] = undefined /*out*/;
            inputs["dBProxyName"] = undefined /*out*/;
            inputs["debugLogging"] = undefined /*out*/;
            inputs["endpoint"] = undefined /*out*/;
            inputs["engineFamily"] = undefined /*out*/;
            inputs["idleClientTimeout"] = undefined /*out*/;
            inputs["requireTLS"] = undefined /*out*/;
            inputs["roleArn"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["vpcId"] = undefined /*out*/;
            inputs["vpcSecurityGroupIds"] = undefined /*out*/;
            inputs["vpcSubnetIds"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(DBProxy.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a DBProxy resource.
 */
export interface DBProxyArgs {
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxy.html#cfn-rds-dbproxy-auth
     */
    auth: pulumi.Input<pulumi.Input<inputs.RDS.DBProxyAuthFormatArgs>[]>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxy.html#cfn-rds-dbproxy-dbproxyname
     */
    dBProxyName: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxy.html#cfn-rds-dbproxy-debuglogging
     */
    debugLogging?: pulumi.Input<boolean>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxy.html#cfn-rds-dbproxy-enginefamily
     */
    engineFamily: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxy.html#cfn-rds-dbproxy-idleclienttimeout
     */
    idleClientTimeout?: pulumi.Input<number>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxy.html#cfn-rds-dbproxy-requiretls
     */
    requireTLS?: pulumi.Input<boolean>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxy.html#cfn-rds-dbproxy-rolearn
     */
    roleArn: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxy.html#cfn-rds-dbproxy-tags
     */
    tags?: pulumi.Input<pulumi.Input<inputs.RDS.DBProxyTagFormatArgs>[]>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxy.html#cfn-rds-dbproxy-vpcsecuritygroupids
     */
    vpcSecurityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxy.html#cfn-rds-dbproxy-vpcsubnetids
     */
    vpcSubnetIds: pulumi.Input<pulumi.Input<string>[]>;
}
