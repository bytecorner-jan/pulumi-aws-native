// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./acceptedPortfolioShare";
export * from "./cloudFormationProduct";
export * from "./cloudFormationProvisionedProduct";
export * from "./launchNotificationConstraint";
export * from "./launchRoleConstraint";
export * from "./launchTemplateConstraint";
export * from "./portfolio";
export * from "./portfolioPrincipalAssociation";
export * from "./portfolioProductAssociation";
export * from "./portfolioShare";
export * from "./resourceUpdateConstraint";
export * from "./serviceAction";
export * from "./serviceActionAssociation";
export * from "./stackSetConstraint";
export * from "./tagOption";
export * from "./tagOptionAssociation";

// Import resources to register:
import { AcceptedPortfolioShare } from "./acceptedPortfolioShare";
import { CloudFormationProduct } from "./cloudFormationProduct";
import { CloudFormationProvisionedProduct } from "./cloudFormationProvisionedProduct";
import { LaunchNotificationConstraint } from "./launchNotificationConstraint";
import { LaunchRoleConstraint } from "./launchRoleConstraint";
import { LaunchTemplateConstraint } from "./launchTemplateConstraint";
import { Portfolio } from "./portfolio";
import { PortfolioPrincipalAssociation } from "./portfolioPrincipalAssociation";
import { PortfolioProductAssociation } from "./portfolioProductAssociation";
import { PortfolioShare } from "./portfolioShare";
import { ResourceUpdateConstraint } from "./resourceUpdateConstraint";
import { ServiceAction } from "./serviceAction";
import { ServiceActionAssociation } from "./serviceActionAssociation";
import { StackSetConstraint } from "./stackSetConstraint";
import { TagOption } from "./tagOption";
import { TagOptionAssociation } from "./tagOptionAssociation";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:servicecatalog:AcceptedPortfolioShare":
                return new AcceptedPortfolioShare(name, <any>undefined, { urn })
            case "aws-native:servicecatalog:CloudFormationProduct":
                return new CloudFormationProduct(name, <any>undefined, { urn })
            case "aws-native:servicecatalog:CloudFormationProvisionedProduct":
                return new CloudFormationProvisionedProduct(name, <any>undefined, { urn })
            case "aws-native:servicecatalog:LaunchNotificationConstraint":
                return new LaunchNotificationConstraint(name, <any>undefined, { urn })
            case "aws-native:servicecatalog:LaunchRoleConstraint":
                return new LaunchRoleConstraint(name, <any>undefined, { urn })
            case "aws-native:servicecatalog:LaunchTemplateConstraint":
                return new LaunchTemplateConstraint(name, <any>undefined, { urn })
            case "aws-native:servicecatalog:Portfolio":
                return new Portfolio(name, <any>undefined, { urn })
            case "aws-native:servicecatalog:PortfolioPrincipalAssociation":
                return new PortfolioPrincipalAssociation(name, <any>undefined, { urn })
            case "aws-native:servicecatalog:PortfolioProductAssociation":
                return new PortfolioProductAssociation(name, <any>undefined, { urn })
            case "aws-native:servicecatalog:PortfolioShare":
                return new PortfolioShare(name, <any>undefined, { urn })
            case "aws-native:servicecatalog:ResourceUpdateConstraint":
                return new ResourceUpdateConstraint(name, <any>undefined, { urn })
            case "aws-native:servicecatalog:ServiceAction":
                return new ServiceAction(name, <any>undefined, { urn })
            case "aws-native:servicecatalog:ServiceActionAssociation":
                return new ServiceActionAssociation(name, <any>undefined, { urn })
            case "aws-native:servicecatalog:StackSetConstraint":
                return new StackSetConstraint(name, <any>undefined, { urn })
            case "aws-native:servicecatalog:TagOption":
                return new TagOption(name, <any>undefined, { urn })
            case "aws-native:servicecatalog:TagOptionAssociation":
                return new TagOptionAssociation(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "servicecatalog", _module)
