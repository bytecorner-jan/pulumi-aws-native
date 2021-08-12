// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.NimbleStudio.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-studiocomponent-activedirectorycomputerattribute.html
    /// </summary>
    [OutputType]
    public sealed class StudioComponentActiveDirectoryComputerAttribute
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-studiocomponent-activedirectorycomputerattribute.html#cfn-nimblestudio-studiocomponent-activedirectorycomputerattribute-name
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-studiocomponent-activedirectorycomputerattribute.html#cfn-nimblestudio-studiocomponent-activedirectorycomputerattribute-value
        /// </summary>
        public readonly string? Value;

        [OutputConstructor]
        private StudioComponentActiveDirectoryComputerAttribute(
            string? name,

            string? value)
        {
            Name = name;
            Value = value;
        }
    }
}
