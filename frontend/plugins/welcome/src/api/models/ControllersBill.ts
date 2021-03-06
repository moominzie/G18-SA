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
/**
 * 
 * @export
 * @interface ControllersBill
 */
export interface ControllersBill {
    /**
     * 
     * @type {number}
     * @memberof ControllersBill
     */
    billingstatus?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersBill
     */
    employee?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersBill
     */
    price?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersBill
     */
    repairinvoice?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersBill
     */
    time?: number;
}

export function ControllersBillFromJSON(json: any): ControllersBill {
    return ControllersBillFromJSONTyped(json, false);
}

export function ControllersBillFromJSONTyped(json: any, ignoreDiscriminator: boolean): ControllersBill {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'billingstatus': !exists(json, 'billingstatus') ? undefined : json['billingstatus'],
        'employee': !exists(json, 'employee') ? undefined : json['employee'],
        'price': !exists(json, 'price') ? undefined : json['price'],
        'repairinvoice': !exists(json, 'repairinvoice') ? undefined : json['repairinvoice'],
        'time': !exists(json, 'time') ? undefined : json['time'],
    };
}

export function ControllersBillToJSON(value?: ControllersBill | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'billingstatus': value.billingstatus,
        'employee': value.employee,
        'price': value.price,
        'repairinvoice': value.repairinvoice,
        'time': value.time,
    };
}


