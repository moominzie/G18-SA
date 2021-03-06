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
    EntRoom,
    EntRoomFromJSON,
    EntRoomFromJSONTyped,
    EntRoomToJSON,
    EntUser,
    EntUserFromJSON,
    EntUserFromJSONTyped,
    EntUserToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntBuildingEdges
 */
export interface EntBuildingEdges {
    /**
     * Rooms holds the value of the rooms edge.
     * @type {Array<EntRoom>}
     * @memberof EntBuildingEdges
     */
    rooms?: Array<EntRoom>;
    /**
     * UserInformations holds the value of the user_informations edge.
     * @type {Array<EntUser>}
     * @memberof EntBuildingEdges
     */
    userInformations?: Array<EntUser>;
}

export function EntBuildingEdgesFromJSON(json: any): EntBuildingEdges {
    return EntBuildingEdgesFromJSONTyped(json, false);
}

export function EntBuildingEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntBuildingEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'rooms': !exists(json, 'rooms') ? undefined : ((json['rooms'] as Array<any>).map(EntRoomFromJSON)),
        'userInformations': !exists(json, 'userInformations') ? undefined : ((json['userInformations'] as Array<any>).map(EntUserFromJSON)),
    };
}

export function EntBuildingEdgesToJSON(value?: EntBuildingEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'rooms': value.rooms === undefined ? undefined : ((value.rooms as Array<any>).map(EntRoomToJSON)),
        'userInformations': value.userInformations === undefined ? undefined : ((value.userInformations as Array<any>).map(EntUserToJSON)),
    };
}


