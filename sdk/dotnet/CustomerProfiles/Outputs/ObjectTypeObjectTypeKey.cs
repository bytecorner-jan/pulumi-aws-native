// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CustomerProfiles.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-objecttype-objecttypekey.html
    /// </summary>
    [OutputType]
    public sealed class ObjectTypeObjectTypeKey
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-objecttype-objecttypekey.html#cfn-customerprofiles-objecttype-objecttypekey-fieldnames
        /// </summary>
        public readonly ImmutableArray<string> FieldNames;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-objecttype-objecttypekey.html#cfn-customerprofiles-objecttype-objecttypekey-standardidentifiers
        /// </summary>
        public readonly ImmutableArray<string> StandardIdentifiers;

        [OutputConstructor]
        private ObjectTypeObjectTypeKey(
            ImmutableArray<string> fieldNames,

            ImmutableArray<string> standardIdentifiers)
        {
            FieldNames = fieldNames;
            StandardIdentifiers = standardIdentifiers;
        }
    }
}
