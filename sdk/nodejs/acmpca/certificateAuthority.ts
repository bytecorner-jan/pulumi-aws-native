// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Private certificate authority.
 */
export class CertificateAuthority extends pulumi.CustomResource {
    /**
     * Get an existing CertificateAuthority resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): CertificateAuthority {
        return new CertificateAuthority(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:acmpca:CertificateAuthority';

    /**
     * Returns true if the given object is an instance of CertificateAuthority.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CertificateAuthority {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CertificateAuthority.__pulumiType;
    }

    /**
     * The Amazon Resource Name (ARN) of the certificate authority.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The base64 PEM-encoded certificate signing request (CSR) for your certificate authority certificate.
     */
    public /*out*/ readonly certificateSigningRequest!: pulumi.Output<string>;
    /**
     * Structure that contains CSR pass through extension information used by the CreateCertificateAuthority action.
     */
    public readonly csrExtensions!: pulumi.Output<outputs.acmpca.CertificateAuthorityCsrExtensions | undefined>;
    /**
     * Public key algorithm and size, in bits, of the key pair that your CA creates when it issues a certificate.
     */
    public readonly keyAlgorithm!: pulumi.Output<string>;
    /**
     * KeyStorageSecurityStadard defines a cryptographic key management compliance standard used for handling CA keys.
     */
    public readonly keyStorageSecurityStandard!: pulumi.Output<string | undefined>;
    /**
     * Certificate revocation information used by the CreateCertificateAuthority and UpdateCertificateAuthority actions.
     */
    public readonly revocationConfiguration!: pulumi.Output<outputs.acmpca.CertificateAuthorityRevocationConfiguration | undefined>;
    /**
     * Algorithm your CA uses to sign certificate requests.
     */
    public readonly signingAlgorithm!: pulumi.Output<string>;
    /**
     * Structure that contains X.500 distinguished name information for your CA.
     */
    public readonly subject!: pulumi.Output<outputs.acmpca.CertificateAuthoritySubject>;
    public readonly tags!: pulumi.Output<outputs.acmpca.CertificateAuthorityTag[] | undefined>;
    /**
     * The type of the certificate authority.
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a CertificateAuthority resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CertificateAuthorityArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.keyAlgorithm === undefined) && !opts.urn) {
                throw new Error("Missing required property 'keyAlgorithm'");
            }
            if ((!args || args.signingAlgorithm === undefined) && !opts.urn) {
                throw new Error("Missing required property 'signingAlgorithm'");
            }
            if ((!args || args.subject === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subject'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            inputs["csrExtensions"] = args ? args.csrExtensions : undefined;
            inputs["keyAlgorithm"] = args ? args.keyAlgorithm : undefined;
            inputs["keyStorageSecurityStandard"] = args ? args.keyStorageSecurityStandard : undefined;
            inputs["revocationConfiguration"] = args ? args.revocationConfiguration : undefined;
            inputs["signingAlgorithm"] = args ? args.signingAlgorithm : undefined;
            inputs["subject"] = args ? args.subject : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["certificateSigningRequest"] = undefined /*out*/;
        } else {
            inputs["arn"] = undefined /*out*/;
            inputs["certificateSigningRequest"] = undefined /*out*/;
            inputs["csrExtensions"] = undefined /*out*/;
            inputs["keyAlgorithm"] = undefined /*out*/;
            inputs["keyStorageSecurityStandard"] = undefined /*out*/;
            inputs["revocationConfiguration"] = undefined /*out*/;
            inputs["signingAlgorithm"] = undefined /*out*/;
            inputs["subject"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(CertificateAuthority.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a CertificateAuthority resource.
 */
export interface CertificateAuthorityArgs {
    /**
     * Structure that contains CSR pass through extension information used by the CreateCertificateAuthority action.
     */
    csrExtensions?: pulumi.Input<inputs.acmpca.CertificateAuthorityCsrExtensionsArgs>;
    /**
     * Public key algorithm and size, in bits, of the key pair that your CA creates when it issues a certificate.
     */
    keyAlgorithm: pulumi.Input<string>;
    /**
     * KeyStorageSecurityStadard defines a cryptographic key management compliance standard used for handling CA keys.
     */
    keyStorageSecurityStandard?: pulumi.Input<string>;
    /**
     * Certificate revocation information used by the CreateCertificateAuthority and UpdateCertificateAuthority actions.
     */
    revocationConfiguration?: pulumi.Input<inputs.acmpca.CertificateAuthorityRevocationConfigurationArgs>;
    /**
     * Algorithm your CA uses to sign certificate requests.
     */
    signingAlgorithm: pulumi.Input<string>;
    /**
     * Structure that contains X.500 distinguished name information for your CA.
     */
    subject: pulumi.Input<inputs.acmpca.CertificateAuthoritySubjectArgs>;
    tags?: pulumi.Input<pulumi.Input<inputs.acmpca.CertificateAuthorityTagArgs>[]>;
    /**
     * The type of the certificate authority.
     */
    type: pulumi.Input<string>;
}
