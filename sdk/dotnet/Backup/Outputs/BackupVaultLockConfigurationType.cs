// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Backup.Outputs
{

    [OutputType]
    public sealed class BackupVaultLockConfigurationType
    {
        public readonly double? ChangeableForDays;
        public readonly double? MaxRetentionDays;
        public readonly double MinRetentionDays;

        [OutputConstructor]
        private BackupVaultLockConfigurationType(
            double? changeableForDays,

            double? maxRetentionDays,

            double minRetentionDays)
        {
            ChangeableForDays = changeableForDays;
            MaxRetentionDays = maxRetentionDays;
            MinRetentionDays = minRetentionDays;
        }
    }
}
