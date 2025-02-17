// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataBrew.Outputs
{

    /// <summary>
    /// Input
    /// </summary>
    [OutputType]
    public sealed class DatasetInput
    {
        public readonly Outputs.DatasetDataCatalogInputDefinition? DataCatalogInputDefinition;
        public readonly Outputs.DatasetDatabaseInputDefinition? DatabaseInputDefinition;
        public readonly Outputs.DatasetS3Location? S3InputDefinition;

        [OutputConstructor]
        private DatasetInput(
            Outputs.DatasetDataCatalogInputDefinition? dataCatalogInputDefinition,

            Outputs.DatasetDatabaseInputDefinition? databaseInputDefinition,

            Outputs.DatasetS3Location? s3InputDefinition)
        {
            DataCatalogInputDefinition = dataCatalogInputDefinition;
            DatabaseInputDefinition = databaseInputDefinition;
            S3InputDefinition = s3InputDefinition;
        }
    }
}
