// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Create and manage wireless gateways, including LoRa gateways.
 */
export class WirelessDevice extends pulumi.CustomResource {
    /**
     * Get an existing WirelessDevice resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): WirelessDevice {
        return new WirelessDevice(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:iotwireless:WirelessDevice';

    /**
     * Returns true if the given object is an instance of WirelessDevice.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WirelessDevice {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WirelessDevice.__pulumiType;
    }

    /**
     * Wireless device arn. Returned after successful create.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Wireless device description
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Wireless device destination name
     */
    public readonly destinationName!: pulumi.Output<string>;
    /**
     * The date and time when the most recent uplink was received.
     */
    public readonly lastUplinkReceivedAt!: pulumi.Output<string | undefined>;
    /**
     * The combination of Package, Station and Model which represents the version of the LoRaWAN Wireless Device.
     */
    public readonly loRaWAN!: pulumi.Output<outputs.iotwireless.WirelessDeviceLoRaWANDevice | undefined>;
    /**
     * Wireless device name
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * A list of key-value pairs that contain metadata for the device. Currently not supported, will not create if tags are passed.
     */
    public readonly tags!: pulumi.Output<outputs.iotwireless.WirelessDeviceTag[] | undefined>;
    /**
     * Thing arn. Passed into update to associate Thing with Wireless device.
     */
    public readonly thingArn!: pulumi.Output<string | undefined>;
    /**
     * Thing Arn. If there is a Thing created, this can be returned with a Get call.
     */
    public /*out*/ readonly thingName!: pulumi.Output<string>;
    /**
     * Wireless device type, currently only Sidewalk and LoRa
     */
    public readonly type!: pulumi.Output<enums.iotwireless.WirelessDeviceType>;

    /**
     * Create a WirelessDevice resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WirelessDeviceArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.destinationName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'destinationName'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["destinationName"] = args ? args.destinationName : undefined;
            inputs["lastUplinkReceivedAt"] = args ? args.lastUplinkReceivedAt : undefined;
            inputs["loRaWAN"] = args ? args.loRaWAN : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["thingArn"] = args ? args.thingArn : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["thingName"] = undefined /*out*/;
        } else {
            inputs["arn"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["destinationName"] = undefined /*out*/;
            inputs["lastUplinkReceivedAt"] = undefined /*out*/;
            inputs["loRaWAN"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["thingArn"] = undefined /*out*/;
            inputs["thingName"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(WirelessDevice.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a WirelessDevice resource.
 */
export interface WirelessDeviceArgs {
    /**
     * Wireless device description
     */
    description?: pulumi.Input<string>;
    /**
     * Wireless device destination name
     */
    destinationName: pulumi.Input<string>;
    /**
     * The date and time when the most recent uplink was received.
     */
    lastUplinkReceivedAt?: pulumi.Input<string>;
    /**
     * The combination of Package, Station and Model which represents the version of the LoRaWAN Wireless Device.
     */
    loRaWAN?: pulumi.Input<inputs.iotwireless.WirelessDeviceLoRaWANDeviceArgs>;
    /**
     * Wireless device name
     */
    name?: pulumi.Input<string>;
    /**
     * A list of key-value pairs that contain metadata for the device. Currently not supported, will not create if tags are passed.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.iotwireless.WirelessDeviceTagArgs>[]>;
    /**
     * Thing arn. Passed into update to associate Thing with Wireless device.
     */
    thingArn?: pulumi.Input<string>;
    /**
     * Wireless device type, currently only Sidewalk and LoRa
     */
    type: pulumi.Input<enums.iotwireless.WirelessDeviceType>;
}
