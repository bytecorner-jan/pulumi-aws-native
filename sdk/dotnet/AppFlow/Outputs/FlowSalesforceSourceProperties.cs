// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppFlow.Outputs
{

    [OutputType]
    public sealed class FlowSalesforceSourceProperties
    {
        public readonly bool? EnableDynamicFieldUpdate;
        public readonly bool? IncludeDeletedRecords;
        public readonly string Object;

        [OutputConstructor]
        private FlowSalesforceSourceProperties(
            bool? enableDynamicFieldUpdate,

            bool? includeDeletedRecords,

            string @object)
        {
            EnableDynamicFieldUpdate = enableDynamicFieldUpdate;
            IncludeDeletedRecords = includeDeletedRecords;
            Object = @object;
        }
    }
}
