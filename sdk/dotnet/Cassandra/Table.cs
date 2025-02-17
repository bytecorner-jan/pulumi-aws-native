// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Cassandra
{
    /// <summary>
    /// Resource schema for AWS::Cassandra::Table
    /// </summary>
    [AwsNativeResourceType("aws-native:cassandra:Table")]
    public partial class Table : Pulumi.CustomResource
    {
        [Output("billingMode")]
        public Output<Outputs.TableBillingMode?> BillingMode { get; private set; } = null!;

        /// <summary>
        /// Clustering key columns of the table
        /// </summary>
        [Output("clusteringKeyColumns")]
        public Output<ImmutableArray<Outputs.TableClusteringKeyColumn>> ClusteringKeyColumns { get; private set; } = null!;

        /// <summary>
        /// Default TTL (Time To Live) in seconds, where zero is disabled. If the value is greater than zero, TTL is enabled for the entire table and an expiration timestamp is added to each column.
        /// </summary>
        [Output("defaultTimeToLive")]
        public Output<int?> DefaultTimeToLive { get; private set; } = null!;

        [Output("encryptionSpecification")]
        public Output<Outputs.TableEncryptionSpecification?> EncryptionSpecification { get; private set; } = null!;

        /// <summary>
        /// Name for Cassandra keyspace
        /// </summary>
        [Output("keyspaceName")]
        public Output<string> KeyspaceName { get; private set; } = null!;

        /// <summary>
        /// Partition key columns of the table
        /// </summary>
        [Output("partitionKeyColumns")]
        public Output<ImmutableArray<Outputs.TableColumn>> PartitionKeyColumns { get; private set; } = null!;

        /// <summary>
        /// Indicates whether point in time recovery is enabled (true) or disabled (false) on the table
        /// </summary>
        [Output("pointInTimeRecoveryEnabled")]
        public Output<bool?> PointInTimeRecoveryEnabled { get; private set; } = null!;

        /// <summary>
        /// Non-key columns of the table
        /// </summary>
        [Output("regularColumns")]
        public Output<ImmutableArray<Outputs.TableColumn>> RegularColumns { get; private set; } = null!;

        /// <summary>
        /// Name for Cassandra table
        /// </summary>
        [Output("tableName")]
        public Output<string?> TableName { get; private set; } = null!;

        /// <summary>
        /// An array of key-value pairs to apply to this resource
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.TableTag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Table resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Table(string name, TableArgs args, CustomResourceOptions? options = null)
            : base("aws-native:cassandra:Table", name, args ?? new TableArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Table(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:cassandra:Table", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Table resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Table Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Table(name, id, options);
        }
    }

    public sealed class TableArgs : Pulumi.ResourceArgs
    {
        [Input("billingMode")]
        public Input<Inputs.TableBillingModeArgs>? BillingMode { get; set; }

        [Input("clusteringKeyColumns")]
        private InputList<Inputs.TableClusteringKeyColumnArgs>? _clusteringKeyColumns;

        /// <summary>
        /// Clustering key columns of the table
        /// </summary>
        public InputList<Inputs.TableClusteringKeyColumnArgs> ClusteringKeyColumns
        {
            get => _clusteringKeyColumns ?? (_clusteringKeyColumns = new InputList<Inputs.TableClusteringKeyColumnArgs>());
            set => _clusteringKeyColumns = value;
        }

        /// <summary>
        /// Default TTL (Time To Live) in seconds, where zero is disabled. If the value is greater than zero, TTL is enabled for the entire table and an expiration timestamp is added to each column.
        /// </summary>
        [Input("defaultTimeToLive")]
        public Input<int>? DefaultTimeToLive { get; set; }

        [Input("encryptionSpecification")]
        public Input<Inputs.TableEncryptionSpecificationArgs>? EncryptionSpecification { get; set; }

        /// <summary>
        /// Name for Cassandra keyspace
        /// </summary>
        [Input("keyspaceName", required: true)]
        public Input<string> KeyspaceName { get; set; } = null!;

        [Input("partitionKeyColumns", required: true)]
        private InputList<Inputs.TableColumnArgs>? _partitionKeyColumns;

        /// <summary>
        /// Partition key columns of the table
        /// </summary>
        public InputList<Inputs.TableColumnArgs> PartitionKeyColumns
        {
            get => _partitionKeyColumns ?? (_partitionKeyColumns = new InputList<Inputs.TableColumnArgs>());
            set => _partitionKeyColumns = value;
        }

        /// <summary>
        /// Indicates whether point in time recovery is enabled (true) or disabled (false) on the table
        /// </summary>
        [Input("pointInTimeRecoveryEnabled")]
        public Input<bool>? PointInTimeRecoveryEnabled { get; set; }

        [Input("regularColumns")]
        private InputList<Inputs.TableColumnArgs>? _regularColumns;

        /// <summary>
        /// Non-key columns of the table
        /// </summary>
        public InputList<Inputs.TableColumnArgs> RegularColumns
        {
            get => _regularColumns ?? (_regularColumns = new InputList<Inputs.TableColumnArgs>());
            set => _regularColumns = value;
        }

        /// <summary>
        /// Name for Cassandra table
        /// </summary>
        [Input("tableName")]
        public Input<string>? TableName { get; set; }

        [Input("tags")]
        private InputList<Inputs.TableTagArgs>? _tags;

        /// <summary>
        /// An array of key-value pairs to apply to this resource
        /// </summary>
        public InputList<Inputs.TableTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.TableTagArgs>());
            set => _tags = value;
        }

        public TableArgs()
        {
        }
    }
}
