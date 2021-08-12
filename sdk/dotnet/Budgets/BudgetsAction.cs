// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Budgets
{
    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-budgets-budgetsaction.html
    /// </summary>
    [AwsNativeResourceType("aws-native:Budgets:BudgetsAction")]
    public partial class BudgetsAction : Pulumi.CustomResource
    {
        [Output("actionId")]
        public Output<string> ActionId { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-budgets-budgetsaction.html#cfn-budgets-budgetsaction-actionthreshold
        /// </summary>
        [Output("actionThreshold")]
        public Output<Outputs.BudgetsActionActionThreshold> ActionThreshold { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-budgets-budgetsaction.html#cfn-budgets-budgetsaction-actiontype
        /// </summary>
        [Output("actionType")]
        public Output<string> ActionType { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-budgets-budgetsaction.html#cfn-budgets-budgetsaction-approvalmodel
        /// </summary>
        [Output("approvalModel")]
        public Output<string?> ApprovalModel { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-budgets-budgetsaction.html#cfn-budgets-budgetsaction-budgetname
        /// </summary>
        [Output("budgetName")]
        public Output<string> BudgetName { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-budgets-budgetsaction.html#cfn-budgets-budgetsaction-definition
        /// </summary>
        [Output("definition")]
        public Output<Outputs.BudgetsActionDefinition> Definition { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-budgets-budgetsaction.html#cfn-budgets-budgetsaction-executionrolearn
        /// </summary>
        [Output("executionRoleArn")]
        public Output<string> ExecutionRoleArn { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-budgets-budgetsaction.html#cfn-budgets-budgetsaction-notificationtype
        /// </summary>
        [Output("notificationType")]
        public Output<string> NotificationType { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-budgets-budgetsaction.html#cfn-budgets-budgetsaction-subscribers
        /// </summary>
        [Output("subscribers")]
        public Output<ImmutableArray<Outputs.BudgetsActionSubscriber>> Subscribers { get; private set; } = null!;


        /// <summary>
        /// Create a BudgetsAction resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BudgetsAction(string name, BudgetsActionArgs args, CustomResourceOptions? options = null)
            : base("aws-native:Budgets:BudgetsAction", name, args ?? new BudgetsActionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BudgetsAction(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:Budgets:BudgetsAction", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing BudgetsAction resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BudgetsAction Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new BudgetsAction(name, id, options);
        }
    }

    public sealed class BudgetsActionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-budgets-budgetsaction.html#cfn-budgets-budgetsaction-actionthreshold
        /// </summary>
        [Input("actionThreshold", required: true)]
        public Input<Inputs.BudgetsActionActionThresholdArgs> ActionThreshold { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-budgets-budgetsaction.html#cfn-budgets-budgetsaction-actiontype
        /// </summary>
        [Input("actionType", required: true)]
        public Input<string> ActionType { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-budgets-budgetsaction.html#cfn-budgets-budgetsaction-approvalmodel
        /// </summary>
        [Input("approvalModel")]
        public Input<string>? ApprovalModel { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-budgets-budgetsaction.html#cfn-budgets-budgetsaction-budgetname
        /// </summary>
        [Input("budgetName", required: true)]
        public Input<string> BudgetName { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-budgets-budgetsaction.html#cfn-budgets-budgetsaction-definition
        /// </summary>
        [Input("definition", required: true)]
        public Input<Inputs.BudgetsActionDefinitionArgs> Definition { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-budgets-budgetsaction.html#cfn-budgets-budgetsaction-executionrolearn
        /// </summary>
        [Input("executionRoleArn", required: true)]
        public Input<string> ExecutionRoleArn { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-budgets-budgetsaction.html#cfn-budgets-budgetsaction-notificationtype
        /// </summary>
        [Input("notificationType", required: true)]
        public Input<string> NotificationType { get; set; } = null!;

        [Input("subscribers")]
        private InputList<Inputs.BudgetsActionSubscriberArgs>? _subscribers;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-budgets-budgetsaction.html#cfn-budgets-budgetsaction-subscribers
        /// </summary>
        public InputList<Inputs.BudgetsActionSubscriberArgs> Subscribers
        {
            get => _subscribers ?? (_subscribers = new InputList<Inputs.BudgetsActionSubscriberArgs>());
            set => _subscribers = value;
        }

        public BudgetsActionArgs()
        {
        }
    }
}
