// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const ConfigBandwidthUnits = {
    GHz: "GHz",
    MHz: "MHz",
    KHz: "kHz",
} as const;

export type ConfigBandwidthUnits = (typeof ConfigBandwidthUnits)[keyof typeof ConfigBandwidthUnits];

export const ConfigEirpUnits = {
    DBW: "dBW",
} as const;

export type ConfigEirpUnits = (typeof ConfigEirpUnits)[keyof typeof ConfigEirpUnits];

export const ConfigFrequencyUnits = {
    GHz: "GHz",
    MHz: "MHz",
    KHz: "kHz",
} as const;

export type ConfigFrequencyUnits = (typeof ConfigFrequencyUnits)[keyof typeof ConfigFrequencyUnits];

export const ConfigPolarization = {
    LeftHand: "LEFT_HAND",
    RightHand: "RIGHT_HAND",
    None: "NONE",
} as const;

export type ConfigPolarization = (typeof ConfigPolarization)[keyof typeof ConfigPolarization];

export const ConfigTrackingConfigAutotrack = {
    Required: "REQUIRED",
    Preferred: "PREFERRED",
    Removed: "REMOVED",
} as const;

export type ConfigTrackingConfigAutotrack = (typeof ConfigTrackingConfigAutotrack)[keyof typeof ConfigTrackingConfigAutotrack];
