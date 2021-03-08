// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/dreamvo/gilfoyle/ent/media"
	"github.com/dreamvo/gilfoyle/ent/mediaevents"
	"github.com/dreamvo/gilfoyle/ent/predicate"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
)

// MediaEventsUpdate is the builder for updating MediaEvents entities.
type MediaEventsUpdate struct {
	config
	hooks    []Hook
	mutation *MediaEventsMutation
}

// Where adds a new predicate for the builder.
func (meu *MediaEventsUpdate) Where(ps ...predicate.MediaEvents) *MediaEventsUpdate {
	meu.mutation.predicates = append(meu.mutation.predicates, ps...)
	return meu
}

// SetLevel sets the level field.
func (meu *MediaEventsUpdate) SetLevel(m mediaevents.Level) *MediaEventsUpdate {
	meu.mutation.SetLevel(m)
	return meu
}

// SetNillableLevel sets the level field if the given value is not nil.
func (meu *MediaEventsUpdate) SetNillableLevel(m *mediaevents.Level) *MediaEventsUpdate {
	if m != nil {
		meu.SetLevel(*m)
	}
	return meu
}

// SetMediaID sets the media edge to Media by id.
func (meu *MediaEventsUpdate) SetMediaID(id uuid.UUID) *MediaEventsUpdate {
	meu.mutation.SetMediaID(id)
	return meu
}

// SetMedia sets the media edge to Media.
func (meu *MediaEventsUpdate) SetMedia(m *Media) *MediaEventsUpdate {
	return meu.SetMediaID(m.ID)
}

// Mutation returns the MediaEventsMutation object of the builder.
func (meu *MediaEventsUpdate) Mutation() *MediaEventsMutation {
	return meu.mutation
}

// ClearMedia clears the "media" edge to type Media.
func (meu *MediaEventsUpdate) ClearMedia() *MediaEventsUpdate {
	meu.mutation.ClearMedia()
	return meu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (meu *MediaEventsUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(meu.hooks) == 0 {
		if err = meu.check(); err != nil {
			return 0, err
		}
		affected, err = meu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MediaEventsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = meu.check(); err != nil {
				return 0, err
			}
			meu.mutation = mutation
			affected, err = meu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(meu.hooks) - 1; i >= 0; i-- {
			mut = meu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, meu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (meu *MediaEventsUpdate) SaveX(ctx context.Context) int {
	affected, err := meu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (meu *MediaEventsUpdate) Exec(ctx context.Context) error {
	_, err := meu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (meu *MediaEventsUpdate) ExecX(ctx context.Context) {
	if err := meu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (meu *MediaEventsUpdate) check() error {
	if v, ok := meu.mutation.Level(); ok {
		if err := mediaevents.LevelValidator(v); err != nil {
			return &ValidationError{Name: "level", err: fmt.Errorf("ent: validator failed for field \"level\": %w", err)}
		}
	}
	if _, ok := meu.mutation.MediaID(); meu.mutation.MediaCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"media\"")
	}
	return nil
}

func (meu *MediaEventsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   mediaevents.Table,
			Columns: mediaevents.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: mediaevents.FieldID,
			},
		},
	}
	if ps := meu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := meu.mutation.Level(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: mediaevents.FieldLevel,
		})
	}
	if meu.mutation.MediaCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   mediaevents.MediaTable,
			Columns: []string{mediaevents.MediaColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: media.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := meu.mutation.MediaIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   mediaevents.MediaTable,
			Columns: []string{mediaevents.MediaColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: media.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, meu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{mediaevents.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// MediaEventsUpdateOne is the builder for updating a single MediaEvents entity.
type MediaEventsUpdateOne struct {
	config
	hooks    []Hook
	mutation *MediaEventsMutation
}

// SetLevel sets the level field.
func (meuo *MediaEventsUpdateOne) SetLevel(m mediaevents.Level) *MediaEventsUpdateOne {
	meuo.mutation.SetLevel(m)
	return meuo
}

// SetNillableLevel sets the level field if the given value is not nil.
func (meuo *MediaEventsUpdateOne) SetNillableLevel(m *mediaevents.Level) *MediaEventsUpdateOne {
	if m != nil {
		meuo.SetLevel(*m)
	}
	return meuo
}

// SetMediaID sets the media edge to Media by id.
func (meuo *MediaEventsUpdateOne) SetMediaID(id uuid.UUID) *MediaEventsUpdateOne {
	meuo.mutation.SetMediaID(id)
	return meuo
}

// SetMedia sets the media edge to Media.
func (meuo *MediaEventsUpdateOne) SetMedia(m *Media) *MediaEventsUpdateOne {
	return meuo.SetMediaID(m.ID)
}

// Mutation returns the MediaEventsMutation object of the builder.
func (meuo *MediaEventsUpdateOne) Mutation() *MediaEventsMutation {
	return meuo.mutation
}

// ClearMedia clears the "media" edge to type Media.
func (meuo *MediaEventsUpdateOne) ClearMedia() *MediaEventsUpdateOne {
	meuo.mutation.ClearMedia()
	return meuo
}

// Save executes the query and returns the updated entity.
func (meuo *MediaEventsUpdateOne) Save(ctx context.Context) (*MediaEvents, error) {
	var (
		err  error
		node *MediaEvents
	)
	if len(meuo.hooks) == 0 {
		if err = meuo.check(); err != nil {
			return nil, err
		}
		node, err = meuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MediaEventsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = meuo.check(); err != nil {
				return nil, err
			}
			meuo.mutation = mutation
			node, err = meuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(meuo.hooks) - 1; i >= 0; i-- {
			mut = meuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, meuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (meuo *MediaEventsUpdateOne) SaveX(ctx context.Context) *MediaEvents {
	node, err := meuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (meuo *MediaEventsUpdateOne) Exec(ctx context.Context) error {
	_, err := meuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (meuo *MediaEventsUpdateOne) ExecX(ctx context.Context) {
	if err := meuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (meuo *MediaEventsUpdateOne) check() error {
	if v, ok := meuo.mutation.Level(); ok {
		if err := mediaevents.LevelValidator(v); err != nil {
			return &ValidationError{Name: "level", err: fmt.Errorf("ent: validator failed for field \"level\": %w", err)}
		}
	}
	if _, ok := meuo.mutation.MediaID(); meuo.mutation.MediaCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"media\"")
	}
	return nil
}

func (meuo *MediaEventsUpdateOne) sqlSave(ctx context.Context) (_node *MediaEvents, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   mediaevents.Table,
			Columns: mediaevents.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: mediaevents.FieldID,
			},
		},
	}
	id, ok := meuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing MediaEvents.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := meuo.mutation.Level(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: mediaevents.FieldLevel,
		})
	}
	if meuo.mutation.MediaCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   mediaevents.MediaTable,
			Columns: []string{mediaevents.MediaColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: media.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := meuo.mutation.MediaIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   mediaevents.MediaTable,
			Columns: []string{mediaevents.MediaColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: media.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &MediaEvents{config: meuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues()
	if err = sqlgraph.UpdateNode(ctx, meuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{mediaevents.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}