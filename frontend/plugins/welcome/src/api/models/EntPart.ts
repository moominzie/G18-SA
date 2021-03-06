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
    EntPartEdges,
    EntPartEdgesFromJSON,
    EntPartEdgesFromJSONTyped,
    EntPartEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntPart
 */
export interface EntPart {
    /**
     * Pname holds the value of the "Pname" field.
     * @type {string}
     * @memberof EntPart
     */
    pname?: string;
    /**
     * 
     * @type {EntPartEdges}
     * @memberof EntPart
     */
    edges?: EntPartEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntPart
     */
    id?: number;
}

export function EntPartFromJSON(json: any): EntPart {
    return EntPartFromJSONTyped(json, false);
}

export function EntPartFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntPart {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'pname': !exists(json, 'Pname') ? undefined : json['Pname'],
        'edges': !exists(json, 'edges') ? undefined : EntPartEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntPartToJSON(value?: EntPart | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'Pname': value.pname,
        'edges': EntPartEdgesToJSON(value.edges),
        'id': value.id,
    };
}


