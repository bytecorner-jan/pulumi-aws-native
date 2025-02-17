// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::DynamoDB::Table
 *
 * @deprecated Table is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class Table extends pulumi.CustomResource {
    /**
     * Get an existing Table resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Table {
        pulumi.log.warn("Table is deprecated: Table is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new Table(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:dynamodb:Table';

    /**
     * Returns true if the given object is an instance of Table.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Table {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Table.__pulumiType;
    }

    public /*out*/ readonly arn!: pulumi.Output<string>;
    public readonly attributeDefinitions!: pulumi.Output<outputs.dynamodb.TableAttributeDefinition[] | undefined>;
    public readonly billingMode!: pulumi.Output<string | undefined>;
    public readonly contributorInsightsSpecification!: pulumi.Output<outputs.dynamodb.TableContributorInsightsSpecification | undefined>;
    public readonly globalSecondaryIndexes!: pulumi.Output<outputs.dynamodb.TableGlobalSecondaryIndex[] | undefined>;
    public readonly keySchema!: pulumi.Output<outputs.dynamodb.TableKeySchema[]>;
    public readonly kinesisStreamSpecification!: pulumi.Output<outputs.dynamodb.TableKinesisStreamSpecification | undefined>;
    public readonly localSecondaryIndexes!: pulumi.Output<outputs.dynamodb.TableLocalSecondaryIndex[] | undefined>;
    public readonly pointInTimeRecoverySpecification!: pulumi.Output<outputs.dynamodb.TablePointInTimeRecoverySpecification | undefined>;
    public readonly provisionedThroughput!: pulumi.Output<outputs.dynamodb.TableProvisionedThroughput | undefined>;
    public readonly sSESpecification!: pulumi.Output<outputs.dynamodb.TableSSESpecification | undefined>;
    public /*out*/ readonly streamArn!: pulumi.Output<string>;
    public readonly streamSpecification!: pulumi.Output<outputs.dynamodb.TableStreamSpecification | undefined>;
    public readonly tableName!: pulumi.Output<string | undefined>;
    public readonly tags!: pulumi.Output<outputs.dynamodb.TableTag[] | undefined>;
    public readonly timeToLiveSpecification!: pulumi.Output<outputs.dynamodb.TableTimeToLiveSpecification | undefined>;

    /**
     * Create a Table resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated Table is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: TableArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("Table is deprecated: Table is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.keySchema === undefined) && !opts.urn) {
                throw new Error("Missing required property 'keySchema'");
            }
            inputs["attributeDefinitions"] = args ? args.attributeDefinitions : undefined;
            inputs["billingMode"] = args ? args.billingMode : undefined;
            inputs["contributorInsightsSpecification"] = args ? args.contributorInsightsSpecification : undefined;
            inputs["globalSecondaryIndexes"] = args ? args.globalSecondaryIndexes : undefined;
            inputs["keySchema"] = args ? args.keySchema : undefined;
            inputs["kinesisStreamSpecification"] = args ? args.kinesisStreamSpecification : undefined;
            inputs["localSecondaryIndexes"] = args ? args.localSecondaryIndexes : undefined;
            inputs["pointInTimeRecoverySpecification"] = args ? args.pointInTimeRecoverySpecification : undefined;
            inputs["provisionedThroughput"] = args ? args.provisionedThroughput : undefined;
            inputs["sSESpecification"] = args ? args.sSESpecification : undefined;
            inputs["streamSpecification"] = args ? args.streamSpecification : undefined;
            inputs["tableName"] = args ? args.tableName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["timeToLiveSpecification"] = args ? args.timeToLiveSpecification : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["streamArn"] = undefined /*out*/;
        } else {
            inputs["arn"] = undefined /*out*/;
            inputs["attributeDefinitions"] = undefined /*out*/;
            inputs["billingMode"] = undefined /*out*/;
            inputs["contributorInsightsSpecification"] = undefined /*out*/;
            inputs["globalSecondaryIndexes"] = undefined /*out*/;
            inputs["keySchema"] = undefined /*out*/;
            inputs["kinesisStreamSpecification"] = undefined /*out*/;
            inputs["localSecondaryIndexes"] = undefined /*out*/;
            inputs["pointInTimeRecoverySpecification"] = undefined /*out*/;
            inputs["provisionedThroughput"] = undefined /*out*/;
            inputs["sSESpecification"] = undefined /*out*/;
            inputs["streamArn"] = undefined /*out*/;
            inputs["streamSpecification"] = undefined /*out*/;
            inputs["tableName"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["timeToLiveSpecification"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Table.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Table resource.
 */
export interface TableArgs {
    attributeDefinitions?: pulumi.Input<pulumi.Input<inputs.dynamodb.TableAttributeDefinitionArgs>[]>;
    billingMode?: pulumi.Input<string>;
    contributorInsightsSpecification?: pulumi.Input<inputs.dynamodb.TableContributorInsightsSpecificationArgs>;
    globalSecondaryIndexes?: pulumi.Input<pulumi.Input<inputs.dynamodb.TableGlobalSecondaryIndexArgs>[]>;
    keySchema: pulumi.Input<pulumi.Input<inputs.dynamodb.TableKeySchemaArgs>[]>;
    kinesisStreamSpecification?: pulumi.Input<inputs.dynamodb.TableKinesisStreamSpecificationArgs>;
    localSecondaryIndexes?: pulumi.Input<pulumi.Input<inputs.dynamodb.TableLocalSecondaryIndexArgs>[]>;
    pointInTimeRecoverySpecification?: pulumi.Input<inputs.dynamodb.TablePointInTimeRecoverySpecificationArgs>;
    provisionedThroughput?: pulumi.Input<inputs.dynamodb.TableProvisionedThroughputArgs>;
    sSESpecification?: pulumi.Input<inputs.dynamodb.TableSSESpecificationArgs>;
    streamSpecification?: pulumi.Input<inputs.dynamodb.TableStreamSpecificationArgs>;
    tableName?: pulumi.Input<string>;
    tags?: pulumi.Input<pulumi.Input<inputs.dynamodb.TableTagArgs>[]>;
    timeToLiveSpecification?: pulumi.Input<inputs.dynamodb.TableTimeToLiveSpecificationArgs>;
}
