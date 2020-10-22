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
 * @interface EntStatustEdges
 */
export interface EntStatustEdges {
    /**
     * Statusts holds the value of the statusts edge.
     * @type {Array<EntReturninvoice>}
     * @memberof EntStatustEdges
     */
    statusts?: Array<EntReturninvoice>;
}

export function EntStatustEdgesFromJSON(json: any): EntStatustEdges {
    return EntStatustEdgesFromJSONTyped(json, false);
}

export function EntStatustEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntStatustEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'statusts': !exists(json, 'statusts') ? undefined : ((json['statusts'] as Array<any>).map(EntReturninvoiceFromJSON)),
    };
}

export function EntStatustEdgesToJSON(value?: EntStatustEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'statusts': value.statusts === undefined ? undefined : ((value.statusts as Array<any>).map(EntReturninvoiceToJSON)),
    };
}

