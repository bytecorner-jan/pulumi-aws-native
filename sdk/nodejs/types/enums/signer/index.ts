// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const SigningProfilePlatformId = {
    AWSLambdaSHA384ECDSA: "AWSLambda-SHA384-ECDSA",
} as const;

export type SigningProfilePlatformId = (typeof SigningProfilePlatformId)[keyof typeof SigningProfilePlatformId];

export const SigningProfileSignatureValidityPeriodType = {
    Days: "DAYS",
    Months: "MONTHS",
    Years: "YEARS",
} as const;

export type SigningProfileSignatureValidityPeriodType = (typeof SigningProfileSignatureValidityPeriodType)[keyof typeof SigningProfileSignatureValidityPeriodType];
