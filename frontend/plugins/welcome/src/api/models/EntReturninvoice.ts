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
    EntReturninvoiceEdges,
    EntReturninvoiceEdgesFromJSON,
    EntReturninvoiceEdgesFromJSONTyped,
    EntReturninvoiceEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntReturninvoice
 */
export interface EntReturninvoice {
    /**
     * Addedtime holds the value of the "addedtime" field.
     * @type {string}
     * @memberof EntReturninvoice
     */
    addedtime?: string;
    /**
     * 
     * @type {EntReturninvoiceEdges}
     * @memberof EntReturninvoice
     */
    edges?: EntReturninvoiceEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntReturninvoice
     */
    id?: number;
}

export function EntReturninvoiceFromJSON(json: any): EntReturninvoice {
    return EntReturninvoiceFromJSONTyped(json, false);
}

export function EntReturninvoiceFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntReturninvoice {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'addedtime': !exists(json, 'addedtime') ? undefined : json['addedtime'],
        'edges': !exists(json, 'edges') ? undefined : EntReturninvoiceEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntReturninvoiceToJSON(value?: EntReturninvoice | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'addedtime': value.addedtime,
        'edges': EntReturninvoiceEdgesToJSON(value.edges),
        'id': value.id,
    };
}


