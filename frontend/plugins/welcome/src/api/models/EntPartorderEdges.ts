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
    EntEmployee,
    EntEmployeeFromJSON,
    EntEmployeeFromJSONTyped,
    EntEmployeeToJSON,
    EntPart,
    EntPartFromJSON,
    EntPartFromJSONTyped,
    EntPartToJSON,
    EntRepairInvoice,
    EntRepairInvoiceFromJSON,
    EntRepairInvoiceFromJSONTyped,
    EntRepairInvoiceToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntPartorderEdges
 */
export interface EntPartorderEdges {
    /**
     * 
     * @type {EntEmployee}
     * @memberof EntPartorderEdges
     */
    employee?: EntEmployee;
    /**
     * 
     * @type {EntPart}
     * @memberof EntPartorderEdges
     */
    part?: EntPart;
    /**
     * 
     * @type {EntRepairInvoice}
     * @memberof EntPartorderEdges
     */
    repairinvoice?: EntRepairInvoice;
}

export function EntPartorderEdgesFromJSON(json: any): EntPartorderEdges {
    return EntPartorderEdgesFromJSONTyped(json, false);
}

export function EntPartorderEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntPartorderEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'employee': !exists(json, 'Employee') ? undefined : EntEmployeeFromJSON(json['Employee']),
        'part': !exists(json, 'Part') ? undefined : EntPartFromJSON(json['Part']),
        'repairinvoice': !exists(json, 'Repairinvoice') ? undefined : EntRepairInvoiceFromJSON(json['Repairinvoice']),
    };
}

export function EntPartorderEdgesToJSON(value?: EntPartorderEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'employee': EntEmployeeToJSON(value.employee),
        'part': EntPartToJSON(value.part),
        'repairinvoice': EntRepairInvoiceToJSON(value.repairinvoice),
    };
}


