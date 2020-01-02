package gen

import (
	"context"
	"time"

	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
	"github.com/novacloudcz/graphql-orm/events"
)

type GeneratedMutationResolver struct{ *GeneratedResolver }

type MutationEvents struct {
	Events []events.Event
}

func EnrichContextWithMutations(ctx context.Context, r *GeneratedResolver) context.Context {
	_ctx := context.WithValue(ctx, KeyMutationTransaction, r.DB.db.Begin())
	_ctx = context.WithValue(_ctx, KeyMutationEvents, &MutationEvents{})
	return _ctx
}
func FinishMutationContext(ctx context.Context, r *GeneratedResolver) (err error) {
	s := GetMutationEventStore(ctx)

	for _, event := range s.Events {
		err = r.Handlers.OnEvent(ctx, r, &event)
		if err != nil {
			return
		}
	}

	tx := GetTransaction(ctx)
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return
	}

	for _, event := range s.Events {
		err = r.EventController.SendEvent(ctx, &event)
	}

	return
}
func GetTransaction(ctx context.Context) *gorm.DB {
	return ctx.Value(KeyMutationTransaction).(*gorm.DB)
}
func GetMutationEventStore(ctx context.Context) *MutationEvents {
	return ctx.Value(KeyMutationEvents).(*MutationEvents)
}
func AddMutationEvent(ctx context.Context, e events.Event) {
	s := GetMutationEventStore(ctx)
	s.Events = append(s.Events, e)
}

func (r *GeneratedMutationResolver) CreateExport(ctx context.Context, input map[string]interface{}) (item *Export, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.CreateExport(ctx, r.GeneratedResolver, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}
func CreateExportHandler(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Export, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	now := time.Now()
	item = &Export{ID: uuid.Must(uuid.NewV4()).String(), CreatedAt: now, CreatedBy: principalID}
	tx := GetTransaction(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeCreated,
		Entity:      "Export",
		EntityID:    item.ID,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes ExportChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	if _, ok := input["id"]; ok && (item.ID != changes.ID) {
		item.ID = changes.ID
		event.EntityID = item.ID
		event.AddNewValue("id", changes.ID)
	}

	if _, ok := input["type"]; ok && (item.Type != changes.Type) && (item.Type == nil || changes.Type == nil || *item.Type != *changes.Type) {
		item.Type = changes.Type

		event.AddNewValue("type", changes.Type)
	}

	if _, ok := input["metadata"]; ok && (item.Metadata != changes.Metadata) && (item.Metadata == nil || changes.Metadata == nil || *item.Metadata != *changes.Metadata) {
		item.Metadata = changes.Metadata

		event.AddNewValue("metadata", changes.Metadata)
	}

	if _, ok := input["state"]; ok && (item.State != changes.State) && (item.State == nil || changes.State == nil || *item.State != *changes.State) {
		item.State = changes.State

		event.AddNewValue("state", changes.State)
	}

	if _, ok := input["errorDescription"]; ok && (item.ErrorDescription != changes.ErrorDescription) && (item.ErrorDescription == nil || changes.ErrorDescription == nil || *item.ErrorDescription != *changes.ErrorDescription) {
		item.ErrorDescription = changes.ErrorDescription

		event.AddNewValue("errorDescription", changes.ErrorDescription)
	}

	if _, ok := input["fileId"]; ok && (item.FileID != changes.FileID) && (item.FileID == nil || changes.FileID == nil || *item.FileID != *changes.FileID) {
		item.FileID = changes.FileID

		event.AddNewValue("fileId", changes.FileID)
	}

	err = tx.Create(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}
func (r *GeneratedMutationResolver) UpdateExport(ctx context.Context, id string, input map[string]interface{}) (item *Export, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.UpdateExport(ctx, r.GeneratedResolver, id, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}
func UpdateExportHandler(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Export, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &Export{}
	now := time.Now()
	tx := GetTransaction(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeUpdated,
		Entity:      "Export",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes ExportChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		tx.Rollback()
		return
	}

	item.UpdatedBy = principalID

	if _, ok := input["type"]; ok && (item.Type != changes.Type) && (item.Type == nil || changes.Type == nil || *item.Type != *changes.Type) {
		event.AddOldValue("type", item.Type)
		event.AddNewValue("type", changes.Type)
		item.Type = changes.Type
	}

	if _, ok := input["metadata"]; ok && (item.Metadata != changes.Metadata) && (item.Metadata == nil || changes.Metadata == nil || *item.Metadata != *changes.Metadata) {
		event.AddOldValue("metadata", item.Metadata)
		event.AddNewValue("metadata", changes.Metadata)
		item.Metadata = changes.Metadata
	}

	if _, ok := input["state"]; ok && (item.State != changes.State) && (item.State == nil || changes.State == nil || *item.State != *changes.State) {
		event.AddOldValue("state", item.State)
		event.AddNewValue("state", changes.State)
		item.State = changes.State
	}

	if _, ok := input["errorDescription"]; ok && (item.ErrorDescription != changes.ErrorDescription) && (item.ErrorDescription == nil || changes.ErrorDescription == nil || *item.ErrorDescription != *changes.ErrorDescription) {
		event.AddOldValue("errorDescription", item.ErrorDescription)
		event.AddNewValue("errorDescription", changes.ErrorDescription)
		item.ErrorDescription = changes.ErrorDescription
	}

	if _, ok := input["fileId"]; ok && (item.FileID != changes.FileID) && (item.FileID == nil || changes.FileID == nil || *item.FileID != *changes.FileID) {
		event.AddOldValue("fileId", item.FileID)
		event.AddNewValue("fileId", changes.FileID)
		item.FileID = changes.FileID
	}

	err = tx.Save(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}
func (r *GeneratedMutationResolver) DeleteExport(ctx context.Context, id string) (item *Export, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.DeleteExport(ctx, r.GeneratedResolver, id)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}
func DeleteExportHandler(ctx context.Context, r *GeneratedResolver, id string) (item *Export, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &Export{}
	now := time.Now()
	tx := GetTransaction(ctx)

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		tx.Rollback()
		return
	}

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeDeleted,
		Entity:      "Export",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	err = tx.Delete(item, TableName("exports")+".id = ?", id).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}
func (r *GeneratedMutationResolver) DeleteAllExports(ctx context.Context) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.DeleteAllExports(ctx, r.GeneratedResolver)
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}
func DeleteAllExportsHandler(ctx context.Context, r *GeneratedResolver) (bool, error) {
	tx := GetTransaction(ctx)
	err := tx.Delete(&Export{}).Error
	if err != nil {
		tx.Rollback()
		return false, err
	}
	return true, err
}
