// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-parameters.html
    /// </summary>
    [OutputType]
    public sealed class DashboardParameters
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-parameters.html#cfn-quicksight-dashboard-parameters-datetimeparameters
        /// </summary>
        public readonly ImmutableArray<Outputs.DashboardDateTimeParameter> DateTimeParameters;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-parameters.html#cfn-quicksight-dashboard-parameters-decimalparameters
        /// </summary>
        public readonly ImmutableArray<Outputs.DashboardDecimalParameter> DecimalParameters;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-parameters.html#cfn-quicksight-dashboard-parameters-integerparameters
        /// </summary>
        public readonly ImmutableArray<Outputs.DashboardIntegerParameter> IntegerParameters;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-parameters.html#cfn-quicksight-dashboard-parameters-stringparameters
        /// </summary>
        public readonly ImmutableArray<Outputs.DashboardStringParameter> StringParameters;

        [OutputConstructor]
        private DashboardParameters(
            ImmutableArray<Outputs.DashboardDateTimeParameter> dateTimeParameters,

            ImmutableArray<Outputs.DashboardDecimalParameter> decimalParameters,

            ImmutableArray<Outputs.DashboardIntegerParameter> integerParameters,

            ImmutableArray<Outputs.DashboardStringParameter> stringParameters)
        {
            DateTimeParameters = dateTimeParameters;
            DecimalParameters = decimalParameters;
            IntegerParameters = integerParameters;
            StringParameters = stringParameters;
        }
    }
}
