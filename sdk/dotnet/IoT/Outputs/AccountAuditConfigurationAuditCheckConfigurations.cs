// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoT.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-accountauditconfiguration-auditcheckconfigurations.html
    /// </summary>
    [OutputType]
    public sealed class AccountAuditConfigurationAuditCheckConfigurations
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-accountauditconfiguration-auditcheckconfigurations.html#cfn-iot-accountauditconfiguration-auditcheckconfigurations-authenticatedcognitoroleoverlypermissivecheck
        /// </summary>
        public readonly Outputs.AccountAuditConfigurationAuditCheckConfiguration? AuthenticatedCognitoRoleOverlyPermissiveCheck;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-accountauditconfiguration-auditcheckconfigurations.html#cfn-iot-accountauditconfiguration-auditcheckconfigurations-cacertificateexpiringcheck
        /// </summary>
        public readonly Outputs.AccountAuditConfigurationAuditCheckConfiguration? CaCertificateExpiringCheck;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-accountauditconfiguration-auditcheckconfigurations.html#cfn-iot-accountauditconfiguration-auditcheckconfigurations-cacertificatekeyqualitycheck
        /// </summary>
        public readonly Outputs.AccountAuditConfigurationAuditCheckConfiguration? CaCertificateKeyQualityCheck;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-accountauditconfiguration-auditcheckconfigurations.html#cfn-iot-accountauditconfiguration-auditcheckconfigurations-conflictingclientidscheck
        /// </summary>
        public readonly Outputs.AccountAuditConfigurationAuditCheckConfiguration? ConflictingClientIdsCheck;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-accountauditconfiguration-auditcheckconfigurations.html#cfn-iot-accountauditconfiguration-auditcheckconfigurations-devicecertificateexpiringcheck
        /// </summary>
        public readonly Outputs.AccountAuditConfigurationAuditCheckConfiguration? DeviceCertificateExpiringCheck;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-accountauditconfiguration-auditcheckconfigurations.html#cfn-iot-accountauditconfiguration-auditcheckconfigurations-devicecertificatekeyqualitycheck
        /// </summary>
        public readonly Outputs.AccountAuditConfigurationAuditCheckConfiguration? DeviceCertificateKeyQualityCheck;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-accountauditconfiguration-auditcheckconfigurations.html#cfn-iot-accountauditconfiguration-auditcheckconfigurations-devicecertificatesharedcheck
        /// </summary>
        public readonly Outputs.AccountAuditConfigurationAuditCheckConfiguration? DeviceCertificateSharedCheck;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-accountauditconfiguration-auditcheckconfigurations.html#cfn-iot-accountauditconfiguration-auditcheckconfigurations-iotpolicyoverlypermissivecheck
        /// </summary>
        public readonly Outputs.AccountAuditConfigurationAuditCheckConfiguration? IotPolicyOverlyPermissiveCheck;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-accountauditconfiguration-auditcheckconfigurations.html#cfn-iot-accountauditconfiguration-auditcheckconfigurations-iotrolealiasallowsaccesstounusedservicescheck
        /// </summary>
        public readonly Outputs.AccountAuditConfigurationAuditCheckConfiguration? IotRoleAliasAllowsAccessToUnusedServicesCheck;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-accountauditconfiguration-auditcheckconfigurations.html#cfn-iot-accountauditconfiguration-auditcheckconfigurations-iotrolealiasoverlypermissivecheck
        /// </summary>
        public readonly Outputs.AccountAuditConfigurationAuditCheckConfiguration? IotRoleAliasOverlyPermissiveCheck;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-accountauditconfiguration-auditcheckconfigurations.html#cfn-iot-accountauditconfiguration-auditcheckconfigurations-loggingdisabledcheck
        /// </summary>
        public readonly Outputs.AccountAuditConfigurationAuditCheckConfiguration? LoggingDisabledCheck;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-accountauditconfiguration-auditcheckconfigurations.html#cfn-iot-accountauditconfiguration-auditcheckconfigurations-revokedcacertificatestillactivecheck
        /// </summary>
        public readonly Outputs.AccountAuditConfigurationAuditCheckConfiguration? RevokedCaCertificateStillActiveCheck;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-accountauditconfiguration-auditcheckconfigurations.html#cfn-iot-accountauditconfiguration-auditcheckconfigurations-revokeddevicecertificatestillactivecheck
        /// </summary>
        public readonly Outputs.AccountAuditConfigurationAuditCheckConfiguration? RevokedDeviceCertificateStillActiveCheck;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-accountauditconfiguration-auditcheckconfigurations.html#cfn-iot-accountauditconfiguration-auditcheckconfigurations-unauthenticatedcognitoroleoverlypermissivecheck
        /// </summary>
        public readonly Outputs.AccountAuditConfigurationAuditCheckConfiguration? UnauthenticatedCognitoRoleOverlyPermissiveCheck;

        [OutputConstructor]
        private AccountAuditConfigurationAuditCheckConfigurations(
            Outputs.AccountAuditConfigurationAuditCheckConfiguration? authenticatedCognitoRoleOverlyPermissiveCheck,

            Outputs.AccountAuditConfigurationAuditCheckConfiguration? caCertificateExpiringCheck,

            Outputs.AccountAuditConfigurationAuditCheckConfiguration? caCertificateKeyQualityCheck,

            Outputs.AccountAuditConfigurationAuditCheckConfiguration? conflictingClientIdsCheck,

            Outputs.AccountAuditConfigurationAuditCheckConfiguration? deviceCertificateExpiringCheck,

            Outputs.AccountAuditConfigurationAuditCheckConfiguration? deviceCertificateKeyQualityCheck,

            Outputs.AccountAuditConfigurationAuditCheckConfiguration? deviceCertificateSharedCheck,

            Outputs.AccountAuditConfigurationAuditCheckConfiguration? iotPolicyOverlyPermissiveCheck,

            Outputs.AccountAuditConfigurationAuditCheckConfiguration? iotRoleAliasAllowsAccessToUnusedServicesCheck,

            Outputs.AccountAuditConfigurationAuditCheckConfiguration? iotRoleAliasOverlyPermissiveCheck,

            Outputs.AccountAuditConfigurationAuditCheckConfiguration? loggingDisabledCheck,

            Outputs.AccountAuditConfigurationAuditCheckConfiguration? revokedCaCertificateStillActiveCheck,

            Outputs.AccountAuditConfigurationAuditCheckConfiguration? revokedDeviceCertificateStillActiveCheck,

            Outputs.AccountAuditConfigurationAuditCheckConfiguration? unauthenticatedCognitoRoleOverlyPermissiveCheck)
        {
            AuthenticatedCognitoRoleOverlyPermissiveCheck = authenticatedCognitoRoleOverlyPermissiveCheck;
            CaCertificateExpiringCheck = caCertificateExpiringCheck;
            CaCertificateKeyQualityCheck = caCertificateKeyQualityCheck;
            ConflictingClientIdsCheck = conflictingClientIdsCheck;
            DeviceCertificateExpiringCheck = deviceCertificateExpiringCheck;
            DeviceCertificateKeyQualityCheck = deviceCertificateKeyQualityCheck;
            DeviceCertificateSharedCheck = deviceCertificateSharedCheck;
            IotPolicyOverlyPermissiveCheck = iotPolicyOverlyPermissiveCheck;
            IotRoleAliasAllowsAccessToUnusedServicesCheck = iotRoleAliasAllowsAccessToUnusedServicesCheck;
            IotRoleAliasOverlyPermissiveCheck = iotRoleAliasOverlyPermissiveCheck;
            LoggingDisabledCheck = loggingDisabledCheck;
            RevokedCaCertificateStillActiveCheck = revokedCaCertificateStillActiveCheck;
            RevokedDeviceCertificateStillActiveCheck = revokedDeviceCertificateStillActiveCheck;
            UnauthenticatedCognitoRoleOverlyPermissiveCheck = unauthenticatedCognitoRoleOverlyPermissiveCheck;
        }
    }
}
