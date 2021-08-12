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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-columngroup.html
    /// </summary>
    public sealed class DataSetColumnGroupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-columngroup.html#cfn-quicksight-dataset-columngroup-geospatialcolumngroup
        /// </summary>
        [Input("geoSpatialColumnGroup")]
        public Input<Inputs.DataSetGeoSpatialColumnGroupArgs>? GeoSpatialColumnGroup { get; set; }

        public DataSetColumnGroupArgs()
        {
        }
    }
}
