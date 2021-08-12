// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.S3Outposts.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3outposts-bucket-lifecycleconfiguration.html
    /// </summary>
    public sealed class BucketLifecycleConfigurationArgs : Pulumi.ResourceArgs
    {
        [Input("rules", required: true)]
        private InputList<Inputs.BucketRuleArgs>? _rules;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3outposts-bucket-lifecycleconfiguration.html#cfn-s3outposts-bucket-lifecycleconfiguration-rules
        /// </summary>
        public InputList<Inputs.BucketRuleArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.BucketRuleArgs>());
            set => _rules = value;
        }

        public BucketLifecycleConfigurationArgs()
        {
        }
    }
}
