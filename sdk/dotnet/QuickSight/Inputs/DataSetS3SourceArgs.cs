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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-s3source.html
    /// </summary>
    public sealed class DataSetS3SourceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-s3source.html#cfn-quicksight-dataset-s3source-datasourcearn
        /// </summary>
        [Input("dataSourceArn", required: true)]
        public Input<string> DataSourceArn { get; set; } = null!;

        [Input("inputColumns", required: true)]
        private InputList<Inputs.DataSetInputColumnArgs>? _inputColumns;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-s3source.html#cfn-quicksight-dataset-s3source-inputcolumns
        /// </summary>
        public InputList<Inputs.DataSetInputColumnArgs> InputColumns
        {
            get => _inputColumns ?? (_inputColumns = new InputList<Inputs.DataSetInputColumnArgs>());
            set => _inputColumns = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-s3source.html#cfn-quicksight-dataset-s3source-uploadsettings
        /// </summary>
        [Input("uploadSettings")]
        public Input<Inputs.DataSetUploadSettingsArgs>? UploadSettings { get; set; }

        public DataSetS3SourceArgs()
        {
        }
    }
}
