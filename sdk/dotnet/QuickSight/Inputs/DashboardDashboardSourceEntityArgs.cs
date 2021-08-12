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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-dashboardsourceentity.html
    /// </summary>
    public sealed class DashboardDashboardSourceEntityArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-dashboardsourceentity.html#cfn-quicksight-dashboard-dashboardsourceentity-sourcetemplate
        /// </summary>
        [Input("sourceTemplate")]
        public Input<Inputs.DashboardDashboardSourceTemplateArgs>? SourceTemplate { get; set; }

        public DashboardDashboardSourceEntityArgs()
        {
        }
    }
}
