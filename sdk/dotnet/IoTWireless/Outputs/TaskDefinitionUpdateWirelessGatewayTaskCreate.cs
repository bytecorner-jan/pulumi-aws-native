// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTWireless.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-taskdefinition-updatewirelessgatewaytaskcreate.html
    /// </summary>
    [OutputType]
    public sealed class TaskDefinitionUpdateWirelessGatewayTaskCreate
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-taskdefinition-updatewirelessgatewaytaskcreate.html#cfn-iotwireless-taskdefinition-updatewirelessgatewaytaskcreate-lorawan
        /// </summary>
        public readonly Outputs.TaskDefinitionLoRaWANUpdateGatewayTaskCreate? LoRaWAN;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-taskdefinition-updatewirelessgatewaytaskcreate.html#cfn-iotwireless-taskdefinition-updatewirelessgatewaytaskcreate-updatedatarole
        /// </summary>
        public readonly string? UpdateDataRole;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-taskdefinition-updatewirelessgatewaytaskcreate.html#cfn-iotwireless-taskdefinition-updatewirelessgatewaytaskcreate-updatedatasource
        /// </summary>
        public readonly string? UpdateDataSource;

        [OutputConstructor]
        private TaskDefinitionUpdateWirelessGatewayTaskCreate(
            Outputs.TaskDefinitionLoRaWANUpdateGatewayTaskCreate? loRaWAN,

            string? updateDataRole,

            string? updateDataSource)
        {
            LoRaWAN = loRaWAN;
            UpdateDataRole = updateDataRole;
            UpdateDataSource = updateDataSource;
        }
    }
}
