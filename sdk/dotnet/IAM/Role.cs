// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IAM
{
    /// <summary>
    /// Resource Type definition for AWS::IAM::Role
    /// </summary>
    [AwsNativeResourceType("aws-native:iam:Role")]
    public partial class Role : Pulumi.CustomResource
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) for the role.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The trust policy that is associated with this role.
        /// </summary>
        [Output("assumeRolePolicyDocument")]
        public Output<object> AssumeRolePolicyDocument { get; private set; } = null!;

        /// <summary>
        /// A description of the role that you provide.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// A list of Amazon Resource Names (ARNs) of the IAM managed policies that you want to attach to the role. 
        /// </summary>
        [Output("managedPolicyArns")]
        public Output<ImmutableArray<string>> ManagedPolicyArns { get; private set; } = null!;

        /// <summary>
        /// The maximum session duration (in seconds) that you want to set for the specified role. If you do not specify a value for this setting, the default maximum of one hour is applied. This setting can have a value from 1 hour to 12 hours. 
        /// </summary>
        [Output("maxSessionDuration")]
        public Output<int?> MaxSessionDuration { get; private set; } = null!;

        /// <summary>
        /// The path to the role.
        /// </summary>
        [Output("path")]
        public Output<string?> Path { get; private set; } = null!;

        /// <summary>
        /// The ARN of the policy used to set the permissions boundary for the role.
        /// </summary>
        [Output("permissionsBoundary")]
        public Output<string?> PermissionsBoundary { get; private set; } = null!;

        /// <summary>
        /// Adds or updates an inline policy document that is embedded in the specified IAM role. 
        /// </summary>
        [Output("policies")]
        public Output<ImmutableArray<Outputs.RolePolicy>> Policies { get; private set; } = null!;

        /// <summary>
        /// The stable and unique string identifying the role.
        /// </summary>
        [Output("roleId")]
        public Output<string> RoleId { get; private set; } = null!;

        /// <summary>
        /// A name for the IAM role, up to 64 characters in length.
        /// </summary>
        [Output("roleName")]
        public Output<string?> RoleName { get; private set; } = null!;

        /// <summary>
        /// A list of tags that are attached to the role.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.RoleTag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Role resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Role(string name, RoleArgs args, CustomResourceOptions? options = null)
            : base("aws-native:iam:Role", name, args ?? new RoleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Role(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:iam:Role", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Role resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Role Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Role(name, id, options);
        }
    }

    public sealed class RoleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The trust policy that is associated with this role.
        /// </summary>
        [Input("assumeRolePolicyDocument", required: true)]
        public Input<object> AssumeRolePolicyDocument { get; set; } = null!;

        /// <summary>
        /// A description of the role that you provide.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("managedPolicyArns")]
        private InputList<string>? _managedPolicyArns;

        /// <summary>
        /// A list of Amazon Resource Names (ARNs) of the IAM managed policies that you want to attach to the role. 
        /// </summary>
        public InputList<string> ManagedPolicyArns
        {
            get => _managedPolicyArns ?? (_managedPolicyArns = new InputList<string>());
            set => _managedPolicyArns = value;
        }

        /// <summary>
        /// The maximum session duration (in seconds) that you want to set for the specified role. If you do not specify a value for this setting, the default maximum of one hour is applied. This setting can have a value from 1 hour to 12 hours. 
        /// </summary>
        [Input("maxSessionDuration")]
        public Input<int>? MaxSessionDuration { get; set; }

        /// <summary>
        /// The path to the role.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// The ARN of the policy used to set the permissions boundary for the role.
        /// </summary>
        [Input("permissionsBoundary")]
        public Input<string>? PermissionsBoundary { get; set; }

        [Input("policies")]
        private InputList<Inputs.RolePolicyArgs>? _policies;

        /// <summary>
        /// Adds or updates an inline policy document that is embedded in the specified IAM role. 
        /// </summary>
        public InputList<Inputs.RolePolicyArgs> Policies
        {
            get => _policies ?? (_policies = new InputList<Inputs.RolePolicyArgs>());
            set => _policies = value;
        }

        /// <summary>
        /// A name for the IAM role, up to 64 characters in length.
        /// </summary>
        [Input("roleName")]
        public Input<string>? RoleName { get; set; }

        [Input("tags")]
        private InputList<Inputs.RoleTagArgs>? _tags;

        /// <summary>
        /// A list of tags that are attached to the role.
        /// </summary>
        public InputList<Inputs.RoleTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.RoleTagArgs>());
            set => _tags = value;
        }

        public RoleArgs()
        {
        }
    }
}
