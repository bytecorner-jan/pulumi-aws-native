// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.FSx.Inputs
{

    public sealed class FileSystemAuditLogConfigurationArgs : Pulumi.ResourceArgs
    {
        [Input("auditLogDestination")]
        public Input<string>? AuditLogDestination { get; set; }

        [Input("fileAccessAuditLogLevel", required: true)]
        public Input<string> FileAccessAuditLogLevel { get; set; } = null!;

        [Input("fileShareAccessAuditLogLevel", required: true)]
        public Input<string> FileShareAccessAuditLogLevel { get; set; } = null!;

        public FileSystemAuditLogConfigurationArgs()
        {
        }
    }
}
