// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DynamoDB.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-globalsecondaryindex.html
    /// </summary>
    public sealed class GlobalTableGlobalSecondaryIndexArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-globalsecondaryindex.html#cfn-dynamodb-globaltable-globalsecondaryindex-indexname
        /// </summary>
        [Input("indexName", required: true)]
        public Input<string> IndexName { get; set; } = null!;

        [Input("keySchema", required: true)]
        private InputList<Inputs.GlobalTableKeySchemaArgs>? _keySchema;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-globalsecondaryindex.html#cfn-dynamodb-globaltable-globalsecondaryindex-keyschema
        /// </summary>
        public InputList<Inputs.GlobalTableKeySchemaArgs> KeySchema
        {
            get => _keySchema ?? (_keySchema = new InputList<Inputs.GlobalTableKeySchemaArgs>());
            set => _keySchema = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-globalsecondaryindex.html#cfn-dynamodb-globaltable-globalsecondaryindex-projection
        /// </summary>
        [Input("projection", required: true)]
        public Input<Inputs.GlobalTableProjectionArgs> Projection { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-globalsecondaryindex.html#cfn-dynamodb-globaltable-globalsecondaryindex-writeprovisionedthroughputsettings
        /// </summary>
        [Input("writeProvisionedThroughputSettings")]
        public Input<Inputs.GlobalTableWriteProvisionedThroughputSettingsArgs>? WriteProvisionedThroughputSettings { get; set; }

        public GlobalTableGlobalSecondaryIndexArgs()
        {
        }
    }
}
