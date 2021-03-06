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
    EntUserEdges,
    EntUserEdgesFromJSON,
    EntUserEdgesFromJSONTyped,
    EntUserEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntUser
 */
export interface EntUser {
    /**
     * 
     * @type {EntUserEdges}
     * @memberof EntUser
     */
    edges?: EntUserEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntUser
     */
    id?: number;
    /**
     * PersonalID holds the value of the "personalID" field.
     * @type {string}
     * @memberof EntUser
     */
    personalID?: string;
    /**
     * PersonalName holds the value of the "personalName" field.
     * @type {string}
     * @memberof EntUser
     */
    personalName?: string;
}

export function EntUserFromJSON(json: any): EntUser {
    return EntUserFromJSONTyped(json, false);
}

export function EntUserFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntUser {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'edges': !exists(json, 'edges') ? undefined : EntUserEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
        'personalID': !exists(json, 'personalID') ? undefined : json['personalID'],
        'personalName': !exists(json, 'personalName') ? undefined : json['personalName'],
    };
}

export function EntUserToJSON(value?: EntUser | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'edges': EntUserEdgesToJSON(value.edges),
        'id': value.id,
        'personalID': value.personalID,
        'personalName': value.personalName,
    };
}


