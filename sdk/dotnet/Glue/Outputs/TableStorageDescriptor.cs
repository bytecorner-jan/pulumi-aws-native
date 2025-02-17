// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Glue.Outputs
{

    [OutputType]
    public sealed class TableStorageDescriptor
    {
        public readonly ImmutableArray<string> BucketColumns;
        public readonly ImmutableArray<Outputs.TableColumn> Columns;
        public readonly bool? Compressed;
        public readonly string? InputFormat;
        public readonly string? Location;
        public readonly int? NumberOfBuckets;
        public readonly string? OutputFormat;
        public readonly object? Parameters;
        public readonly Outputs.TableSchemaReference? SchemaReference;
        public readonly Outputs.TableSerdeInfo? SerdeInfo;
        public readonly Outputs.TableSkewedInfo? SkewedInfo;
        public readonly ImmutableArray<Outputs.TableOrder> SortColumns;
        public readonly bool? StoredAsSubDirectories;

        [OutputConstructor]
        private TableStorageDescriptor(
            ImmutableArray<string> bucketColumns,

            ImmutableArray<Outputs.TableColumn> columns,

            bool? compressed,

            string? inputFormat,

            string? location,

            int? numberOfBuckets,

            string? outputFormat,

            object? parameters,

            Outputs.TableSchemaReference? schemaReference,

            Outputs.TableSerdeInfo? serdeInfo,

            Outputs.TableSkewedInfo? skewedInfo,

            ImmutableArray<Outputs.TableOrder> sortColumns,

            bool? storedAsSubDirectories)
        {
            BucketColumns = bucketColumns;
            Columns = columns;
            Compressed = compressed;
            InputFormat = inputFormat;
            Location = location;
            NumberOfBuckets = numberOfBuckets;
            OutputFormat = outputFormat;
            Parameters = parameters;
            SchemaReference = schemaReference;
            SerdeInfo = serdeInfo;
            SkewedInfo = skewedInfo;
            SortColumns = sortColumns;
            StoredAsSubDirectories = storedAsSubDirectories;
        }
    }
}
