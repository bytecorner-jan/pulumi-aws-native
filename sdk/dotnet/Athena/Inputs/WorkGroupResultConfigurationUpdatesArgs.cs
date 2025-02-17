// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Athena.Inputs
{

    /// <summary>
    /// The result configuration information about the queries in this workgroup that will be updated. Includes the updated results location and an updated option for encrypting query results. 
    /// </summary>
    public sealed class WorkGroupResultConfigurationUpdatesArgs : Pulumi.ResourceArgs
    {
        [Input("encryptionConfiguration")]
        public Input<Inputs.WorkGroupEncryptionConfigurationArgs>? EncryptionConfiguration { get; set; }

        [Input("outputLocation")]
        public Input<string>? OutputLocation { get; set; }

        [Input("removeEncryptionConfiguration")]
        public Input<bool>? RemoveEncryptionConfiguration { get; set; }

        [Input("removeOutputLocation")]
        public Input<bool>? RemoveOutputLocation { get; set; }

        public WorkGroupResultConfigurationUpdatesArgs()
        {
        }
    }
}
