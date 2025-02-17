// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DocDB
{
    /// <summary>
    /// Resource Type definition for AWS::DocDB::DBInstance
    /// </summary>
    [Obsolete(@"DBInstance is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")]
    [AwsNativeResourceType("aws-native:docdb:DBInstance")]
    public partial class DBInstance : Pulumi.CustomResource
    {
        [Output("autoMinorVersionUpgrade")]
        public Output<bool?> AutoMinorVersionUpgrade { get; private set; } = null!;

        [Output("availabilityZone")]
        public Output<string?> AvailabilityZone { get; private set; } = null!;

        [Output("dBClusterIdentifier")]
        public Output<string> DBClusterIdentifier { get; private set; } = null!;

        [Output("dBInstanceClass")]
        public Output<string> DBInstanceClass { get; private set; } = null!;

        [Output("dBInstanceIdentifier")]
        public Output<string?> DBInstanceIdentifier { get; private set; } = null!;

        [Output("endpoint")]
        public Output<string> Endpoint { get; private set; } = null!;

        [Output("port")]
        public Output<string> Port { get; private set; } = null!;

        [Output("preferredMaintenanceWindow")]
        public Output<string?> PreferredMaintenanceWindow { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableArray<Outputs.DBInstanceTag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a DBInstance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DBInstance(string name, DBInstanceArgs args, CustomResourceOptions? options = null)
            : base("aws-native:docdb:DBInstance", name, args ?? new DBInstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DBInstance(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:docdb:DBInstance", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing DBInstance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DBInstance Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new DBInstance(name, id, options);
        }
    }

    public sealed class DBInstanceArgs : Pulumi.ResourceArgs
    {
        [Input("autoMinorVersionUpgrade")]
        public Input<bool>? AutoMinorVersionUpgrade { get; set; }

        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        [Input("dBClusterIdentifier", required: true)]
        public Input<string> DBClusterIdentifier { get; set; } = null!;

        [Input("dBInstanceClass", required: true)]
        public Input<string> DBInstanceClass { get; set; } = null!;

        [Input("dBInstanceIdentifier")]
        public Input<string>? DBInstanceIdentifier { get; set; }

        [Input("preferredMaintenanceWindow")]
        public Input<string>? PreferredMaintenanceWindow { get; set; }

        [Input("tags")]
        private InputList<Inputs.DBInstanceTagArgs>? _tags;
        public InputList<Inputs.DBInstanceTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.DBInstanceTagArgs>());
            set => _tags = value;
        }

        public DBInstanceArgs()
        {
        }
    }
}
