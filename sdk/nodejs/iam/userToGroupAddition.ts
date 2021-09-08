// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-addusertogroup.html
 */
export class UserToGroupAddition extends pulumi.CustomResource {
    /**
     * Get an existing UserToGroupAddition resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): UserToGroupAddition {
        return new UserToGroupAddition(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:iam:UserToGroupAddition';

    /**
     * Returns true if the given object is an instance of UserToGroupAddition.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UserToGroupAddition {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UserToGroupAddition.__pulumiType;
    }

    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-addusertogroup.html#cfn-iam-addusertogroup-groupname
     */
    public readonly groupName!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-addusertogroup.html#cfn-iam-addusertogroup-users
     */
    public readonly users!: pulumi.Output<string[]>;

    /**
     * Create a UserToGroupAddition resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UserToGroupAdditionArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.groupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'groupName'");
            }
            if ((!args || args.users === undefined) && !opts.urn) {
                throw new Error("Missing required property 'users'");
            }
            inputs["groupName"] = args ? args.groupName : undefined;
            inputs["users"] = args ? args.users : undefined;
        } else {
            inputs["groupName"] = undefined /*out*/;
            inputs["users"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(UserToGroupAddition.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a UserToGroupAddition resource.
 */
export interface UserToGroupAdditionArgs {
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-addusertogroup.html#cfn-iam-addusertogroup-groupname
     */
    groupName: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-addusertogroup.html#cfn-iam-addusertogroup-users
     */
    users: pulumi.Input<pulumi.Input<string>[]>;
}
