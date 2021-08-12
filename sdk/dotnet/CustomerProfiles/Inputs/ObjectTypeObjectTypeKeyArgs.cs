// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CustomerProfiles.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-objecttype-objecttypekey.html
    /// </summary>
    public sealed class ObjectTypeObjectTypeKeyArgs : Pulumi.ResourceArgs
    {
        [Input("fieldNames")]
        private InputList<string>? _fieldNames;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-objecttype-objecttypekey.html#cfn-customerprofiles-objecttype-objecttypekey-fieldnames
        /// </summary>
        public InputList<string> FieldNames
        {
            get => _fieldNames ?? (_fieldNames = new InputList<string>());
            set => _fieldNames = value;
        }

        [Input("standardIdentifiers")]
        private InputList<string>? _standardIdentifiers;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-objecttype-objecttypekey.html#cfn-customerprofiles-objecttype-objecttypekey-standardidentifiers
        /// </summary>
        public InputList<string> StandardIdentifiers
        {
            get => _standardIdentifiers ?? (_standardIdentifiers = new InputList<string>());
            set => _standardIdentifiers = value;
        }

        public ObjectTypeObjectTypeKeyArgs()
        {
        }
    }
}
