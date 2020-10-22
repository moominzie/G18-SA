/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntRepairinvoiceEdges,
    EntRepairinvoiceEdgesFromJSON,
    EntRepairinvoiceEdgesFromJSONTyped,
    EntRepairinvoiceEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntRepairinvoice
 */
export interface EntRepairinvoice {
    /**
     * Deviceid holds the value of the "deviceid" field.
     * @type {number}
     * @memberof EntRepairinvoice
     */
    deviceid?: number;
    /**
     * 
     * @type {EntRepairinvoiceEdges}
     * @memberof EntRepairinvoice
     */
    edges?: EntRepairinvoiceEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntRepairinvoice
     */
    id?: number;
    /**
     * Statusrepairid holds the value of the "statusrepairid" field.
     * @type {number}
     * @memberof EntRepairinvoice
     */
    statusrepairid?: number;
    /**
     * Symptomid holds the value of the "symptomid" field.
     * @type {number}
     * @memberof EntRepairinvoice
     */
    symptomid?: number;
    /**
     * Userid holds the value of the "userid" field.
     * @type {number}
     * @memberof EntRepairinvoice
     */
    userid?: number;
}

export function EntRepairinvoiceFromJSON(json: any): EntRepairinvoice {
    return EntRepairinvoiceFromJSONTyped(json, false);
}

export function EntRepairinvoiceFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntRepairinvoice {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'deviceid': !exists(json, 'deviceid') ? undefined : json['deviceid'],
        'edges': !exists(json, 'edges') ? undefined : EntRepairinvoiceEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
        'statusrepairid': !exists(json, 'statusrepairid') ? undefined : json['statusrepairid'],
        'symptomid': !exists(json, 'symptomid') ? undefined : json['symptomid'],
        'userid': !exists(json, 'userid') ? undefined : json['userid'],
    };
}

export function EntRepairinvoiceToJSON(value?: EntRepairinvoice | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'deviceid': value.deviceid,
        'edges': EntRepairinvoiceEdgesToJSON(value.edges),
        'id': value.id,
        'statusrepairid': value.statusrepairid,
        'symptomid': value.symptomid,
        'userid': value.userid,
    };
}

