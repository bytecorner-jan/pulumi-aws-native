// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Logs
{
    /// <summary>
    /// Resource schema for AWS::Logs::LogGroup
    /// </summary>
    [AwsNativeResourceType("aws-native:logs:LogGroup")]
    public partial class LogGroup : Pulumi.CustomResource
    {
        /// <summary>
        /// The CloudWatch log group ARN.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the CMK to use when encrypting log data.
        /// </summary>
        [Output("kmsKeyId")]
        public Output<string?> KmsKeyId { get; private set; } = null!;

        /// <summary>
        /// The name of the log group. If you don't specify a name, AWS CloudFormation generates a unique ID for the log group.
        /// </summary>
        [Output("logGroupName")]
        public Output<string?> LogGroupName { get; private set; } = null!;

        /// <summary>
        /// The number of days to retain the log events in the specified log group. Possible values are: 1, 3, 5, 7, 14, 30, 60, 90, 120, 150, 180, 365, 400, 545, 731, 1827, and 3653.
        /// </summary>
        [Output("retentionInDays")]
        public Output<int?> RetentionInDays { get; private set; } = null!;


        /// <summary>
        /// Create a LogGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LogGroup(string name, LogGroupArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:logs:LogGroup", name, args ?? new LogGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LogGroup(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:logs:LogGroup", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing LogGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LogGroup Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new LogGroup(name, id, options);
        }
    }

    public sealed class LogGroupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the CMK to use when encrypting log data.
        /// </summary>
        [Input("kmsKeyId")]
        public Input<string>? KmsKeyId { get; set; }

        /// <summary>
        /// The name of the log group. If you don't specify a name, AWS CloudFormation generates a unique ID for the log group.
        /// </summary>
        [Input("logGroupName")]
        public Input<string>? LogGroupName { get; set; }

        /// <summary>
        /// The number of days to retain the log events in the specified log group. Possible values are: 1, 3, 5, 7, 14, 30, 60, 90, 120, 150, 180, 365, 400, 545, 731, 1827, and 3653.
        /// </summary>
        [Input("retentionInDays")]
        public Input<int>? RetentionInDays { get; set; }

        public LogGroupArgs()
        {
        }
    }
}
