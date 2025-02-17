// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * A certificate issued via a private certificate authority
 */
export class Certificate extends pulumi.CustomResource {
    /**
     * Get an existing Certificate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Certificate {
        return new Certificate(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:acmpca:Certificate';

    /**
     * Returns true if the given object is an instance of Certificate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Certificate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Certificate.__pulumiType;
    }

    /**
     * These are fields to be overridden in a certificate at the time of issuance. These requires an API_Passthrough template be used or they will be ignored.
     */
    public readonly apiPassthrough!: pulumi.Output<outputs.acmpca.CertificateApiPassthrough | undefined>;
    /**
     * The ARN of the issued certificate.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The issued certificate in base 64 PEM-encoded format.
     */
    public /*out*/ readonly certificate!: pulumi.Output<string>;
    /**
     * The Amazon Resource Name (ARN) for the private CA to issue the certificate.
     */
    public readonly certificateAuthorityArn!: pulumi.Output<string>;
    /**
     * The certificate signing request (CSR) for the Certificate.
     */
    public readonly certificateSigningRequest!: pulumi.Output<string>;
    /**
     * The name of the algorithm that will be used to sign the Certificate.
     */
    public readonly signingAlgorithm!: pulumi.Output<string>;
    /**
     * Specifies a custom configuration template to use when issuing a certificate. If this parameter is not provided, ACM Private CA defaults to the EndEntityCertificate/V1 template.
     */
    public readonly templateArn!: pulumi.Output<string | undefined>;
    /**
     * The time before which the Certificate will be valid.
     */
    public readonly validity!: pulumi.Output<outputs.acmpca.CertificateValidity>;
    /**
     * The time after which the Certificate will be valid.
     */
    public readonly validityNotBefore!: pulumi.Output<outputs.acmpca.CertificateValidity | undefined>;

    /**
     * Create a Certificate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CertificateArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.certificateAuthorityArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'certificateAuthorityArn'");
            }
            if ((!args || args.certificateSigningRequest === undefined) && !opts.urn) {
                throw new Error("Missing required property 'certificateSigningRequest'");
            }
            if ((!args || args.signingAlgorithm === undefined) && !opts.urn) {
                throw new Error("Missing required property 'signingAlgorithm'");
            }
            if ((!args || args.validity === undefined) && !opts.urn) {
                throw new Error("Missing required property 'validity'");
            }
            inputs["apiPassthrough"] = args ? args.apiPassthrough : undefined;
            inputs["certificateAuthorityArn"] = args ? args.certificateAuthorityArn : undefined;
            inputs["certificateSigningRequest"] = args ? args.certificateSigningRequest : undefined;
            inputs["signingAlgorithm"] = args ? args.signingAlgorithm : undefined;
            inputs["templateArn"] = args ? args.templateArn : undefined;
            inputs["validity"] = args ? args.validity : undefined;
            inputs["validityNotBefore"] = args ? args.validityNotBefore : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["certificate"] = undefined /*out*/;
        } else {
            inputs["apiPassthrough"] = undefined /*out*/;
            inputs["arn"] = undefined /*out*/;
            inputs["certificate"] = undefined /*out*/;
            inputs["certificateAuthorityArn"] = undefined /*out*/;
            inputs["certificateSigningRequest"] = undefined /*out*/;
            inputs["signingAlgorithm"] = undefined /*out*/;
            inputs["templateArn"] = undefined /*out*/;
            inputs["validity"] = undefined /*out*/;
            inputs["validityNotBefore"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Certificate.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Certificate resource.
 */
export interface CertificateArgs {
    /**
     * These are fields to be overridden in a certificate at the time of issuance. These requires an API_Passthrough template be used or they will be ignored.
     */
    apiPassthrough?: pulumi.Input<inputs.acmpca.CertificateApiPassthroughArgs>;
    /**
     * The Amazon Resource Name (ARN) for the private CA to issue the certificate.
     */
    certificateAuthorityArn: pulumi.Input<string>;
    /**
     * The certificate signing request (CSR) for the Certificate.
     */
    certificateSigningRequest: pulumi.Input<string>;
    /**
     * The name of the algorithm that will be used to sign the Certificate.
     */
    signingAlgorithm: pulumi.Input<string>;
    /**
     * Specifies a custom configuration template to use when issuing a certificate. If this parameter is not provided, ACM Private CA defaults to the EndEntityCertificate/V1 template.
     */
    templateArn?: pulumi.Input<string>;
    /**
     * The time before which the Certificate will be valid.
     */
    validity: pulumi.Input<inputs.acmpca.CertificateValidityArgs>;
    /**
     * The time after which the Certificate will be valid.
     */
    validityNotBefore?: pulumi.Input<inputs.acmpca.CertificateValidityArgs>;
}
