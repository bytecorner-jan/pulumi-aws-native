// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-datetimeparameter.html
    /// </summary>
    public sealed class DashboardDateTimeParameterArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-datetimeparameter.html#cfn-quicksight-dashboard-datetimeparameter-name
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("values", required: true)]
        private InputList<string>? _values;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-datetimeparameter.html#cfn-quicksight-dashboard-datetimeparameter-values
        /// </summary>
        public InputList<string> Values
        {
            get => _values ?? (_values = new InputList<string>());
            set => _values = value;
        }

        public DashboardDateTimeParameterArgs()
        {
        }
    }
}
