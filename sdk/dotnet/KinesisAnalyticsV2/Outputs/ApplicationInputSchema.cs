// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.KinesisAnalyticsV2.Outputs
{

    [OutputType]
    public sealed class ApplicationInputSchema
    {
        public readonly ImmutableArray<Outputs.ApplicationRecordColumn> RecordColumns;
        public readonly string? RecordEncoding;
        public readonly Outputs.ApplicationRecordFormat RecordFormat;

        [OutputConstructor]
        private ApplicationInputSchema(
            ImmutableArray<Outputs.ApplicationRecordColumn> recordColumns,

            string? recordEncoding,

            Outputs.ApplicationRecordFormat recordFormat)
        {
            RecordColumns = recordColumns;
            RecordEncoding = recordEncoding;
            RecordFormat = recordFormat;
        }
    }
}
