// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTSiteWise.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-accesspolicy-accesspolicyidentity.html
    /// </summary>
    [OutputType]
    public sealed class AccessPolicyAccessPolicyIdentity
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-accesspolicy-accesspolicyidentity.html#cfn-iotsitewise-accesspolicy-accesspolicyidentity-iamrole
        /// </summary>
        public readonly Outputs.AccessPolicyIamRole? IamRole;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-accesspolicy-accesspolicyidentity.html#cfn-iotsitewise-accesspolicy-accesspolicyidentity-iamuser
        /// </summary>
        public readonly Outputs.AccessPolicyIamUser? IamUser;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-accesspolicy-accesspolicyidentity.html#cfn-iotsitewise-accesspolicy-accesspolicyidentity-user
        /// </summary>
        public readonly Outputs.AccessPolicyUser? User;

        [OutputConstructor]
        private AccessPolicyAccessPolicyIdentity(
            Outputs.AccessPolicyIamRole? iamRole,

            Outputs.AccessPolicyIamUser? iamUser,

            Outputs.AccessPolicyUser? user)
        {
            IamRole = iamRole;
            IamUser = iamUser;
            User = user;
        }
    }
}
