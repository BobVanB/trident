// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// SvmMigrationState Indicates the state of the migration.
//
// swagger:model svm_migration_state
type SvmMigrationState string

const (

	// SvmMigrationStatePrecheckStarted captures enum value "precheck_started"
	SvmMigrationStatePrecheckStarted SvmMigrationState = "precheck_started"

	// SvmMigrationStateTransferring captures enum value "transferring"
	SvmMigrationStateTransferring SvmMigrationState = "transferring"

	// SvmMigrationStateReadyForCutover captures enum value "ready_for_cutover"
	SvmMigrationStateReadyForCutover SvmMigrationState = "ready_for_cutover"

	// SvmMigrationStateCutoverTriggered captures enum value "cutover_triggered"
	SvmMigrationStateCutoverTriggered SvmMigrationState = "cutover_triggered"

	// SvmMigrationStateCutoverStarted captures enum value "cutover_started"
	SvmMigrationStateCutoverStarted SvmMigrationState = "cutover_started"

	// SvmMigrationStateCutoverComplete captures enum value "cutover_complete"
	SvmMigrationStateCutoverComplete SvmMigrationState = "cutover_complete"

	// SvmMigrationStateAbort captures enum value "abort"
	SvmMigrationStateAbort SvmMigrationState = "abort"

	// SvmMigrationStatePaused captures enum value "paused"
	SvmMigrationStatePaused SvmMigrationState = "paused"

	// SvmMigrationStateCompleted captures enum value "completed"
	SvmMigrationStateCompleted SvmMigrationState = "completed"

	// SvmMigrationStateFailed captures enum value "failed"
	SvmMigrationStateFailed SvmMigrationState = "failed"

	// SvmMigrationStateSourceCleanup captures enum value "source_cleanup"
	SvmMigrationStateSourceCleanup SvmMigrationState = "source_cleanup"
)

// for schema
var svmMigrationStateEnum []interface{}

func init() {
	var res []SvmMigrationState
	if err := json.Unmarshal([]byte(`["precheck_started","transferring","ready_for_cutover","cutover_triggered","cutover_started","cutover_complete","abort","paused","completed","failed","source_cleanup"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		svmMigrationStateEnum = append(svmMigrationStateEnum, v)
	}
}

func (m SvmMigrationState) validateSvmMigrationStateEnum(path, location string, value SvmMigrationState) error {
	if err := validate.EnumCase(path, location, value, svmMigrationStateEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this svm migration state
func (m SvmMigrationState) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateSvmMigrationStateEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this svm migration state based on the context it is used
func (m SvmMigrationState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := validate.ReadOnly(ctx, "", "body", SvmMigrationState(m)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
