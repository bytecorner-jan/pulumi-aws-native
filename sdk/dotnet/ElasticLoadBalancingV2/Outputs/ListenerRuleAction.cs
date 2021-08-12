// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ElasticLoadBalancingV2.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-action.html
    /// </summary>
    [OutputType]
    public sealed class ListenerRuleAction
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-action.html#cfn-elasticloadbalancingv2-listenerrule-action-authenticatecognitoconfig
        /// </summary>
        public readonly Outputs.ListenerRuleAuthenticateCognitoConfig? AuthenticateCognitoConfig;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-action.html#cfn-elasticloadbalancingv2-listenerrule-action-authenticateoidcconfig
        /// </summary>
        public readonly Outputs.ListenerRuleAuthenticateOidcConfig? AuthenticateOidcConfig;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-action.html#cfn-elasticloadbalancingv2-listenerrule-action-fixedresponseconfig
        /// </summary>
        public readonly Outputs.ListenerRuleFixedResponseConfig? FixedResponseConfig;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-action.html#cfn-elasticloadbalancingv2-listenerrule-action-forwardconfig
        /// </summary>
        public readonly Outputs.ListenerRuleForwardConfig? ForwardConfig;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-action.html#cfn-elasticloadbalancingv2-listenerrule-action-order
        /// </summary>
        public readonly int? Order;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-action.html#cfn-elasticloadbalancingv2-listenerrule-action-redirectconfig
        /// </summary>
        public readonly Outputs.ListenerRuleRedirectConfig? RedirectConfig;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-action.html#cfn-elasticloadbalancingv2-listenerrule-action-targetgrouparn
        /// </summary>
        public readonly string? TargetGroupArn;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-action.html#cfn-elasticloadbalancingv2-listenerrule-action-type
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private ListenerRuleAction(
            Outputs.ListenerRuleAuthenticateCognitoConfig? authenticateCognitoConfig,

            Outputs.ListenerRuleAuthenticateOidcConfig? authenticateOidcConfig,

            Outputs.ListenerRuleFixedResponseConfig? fixedResponseConfig,

            Outputs.ListenerRuleForwardConfig? forwardConfig,

            int? order,

            Outputs.ListenerRuleRedirectConfig? redirectConfig,

            string? targetGroupArn,

            string type)
        {
            AuthenticateCognitoConfig = authenticateCognitoConfig;
            AuthenticateOidcConfig = authenticateOidcConfig;
            FixedResponseConfig = fixedResponseConfig;
            ForwardConfig = forwardConfig;
            Order = order;
            RedirectConfig = redirectConfig;
            TargetGroupArn = targetGroupArn;
            Type = type;
        }
    }
}
