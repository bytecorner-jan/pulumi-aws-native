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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-dashboardsourcetemplate.html
    /// </summary>
    public sealed class DashboardDashboardSourceTemplateArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-dashboardsourcetemplate.html#cfn-quicksight-dashboard-dashboardsourcetemplate-arn
        /// </summary>
        [Input("arn", required: true)]
        public Input<string> Arn { get; set; } = null!;

        [Input("dataSetReferences", required: true)]
        private InputList<Inputs.DashboardDataSetReferenceArgs>? _dataSetReferences;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-dashboardsourcetemplate.html#cfn-quicksight-dashboard-dashboardsourcetemplate-datasetreferences
        /// </summary>
        public InputList<Inputs.DashboardDataSetReferenceArgs> DataSetReferences
        {
            get => _dataSetReferences ?? (_dataSetReferences = new InputList<Inputs.DashboardDataSetReferenceArgs>());
            set => _dataSetReferences = value;
        }

        public DashboardDashboardSourceTemplateArgs()
        {
        }
    }
}
