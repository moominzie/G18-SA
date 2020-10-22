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
    EntReturninvoice,
    EntReturninvoiceFromJSON,
    EntReturninvoiceFromJSONTyped,
    EntReturninvoiceToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntRepairinvoiceEdges
 */
export interface EntRepairinvoiceEdges {
    /**
     * 
     * @type {EntReturninvoice}
     * @memberof EntRepairinvoiceEdges
     */
    repairinvoices?: EntReturninvoice;
}

export function EntRepairinvoiceEdgesFromJSON(json: any): EntRepairinvoiceEdges {
    return EntRepairinvoiceEdgesFromJSONTyped(json, false);
}

export function EntRepairinvoiceEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntRepairinvoiceEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'repairinvoices': !exists(json, 'repairinvoices') ? undefined : EntReturninvoiceFromJSON(json['repairinvoices']),
    };
}

export function EntRepairinvoiceEdgesToJSON(value?: EntRepairinvoiceEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'repairinvoices': EntReturninvoiceToJSON(value.repairinvoices),
    };
}


