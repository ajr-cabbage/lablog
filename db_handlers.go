package main

import (
	"context"
	"errors"

	"github.com/ajr-cabbage/lablog/internal/database"
)

func addEntryHandler(m *MainModel, f *FormViewModel) error {
	// get form values and type assert to struct field types
	rawCategory := f.form.Get("category")
	newCategory, ok := rawCategory.(category)
	if !ok {
		return errors.New("unable to type assert category")
	}
	rawFriendlyName := f.form.Get("friendlyName")
	newFriendlyName, ok := rawFriendlyName.(string)
	if !ok {
		return errors.New("unable to type assert friendlyName")
	}
	rawHostName := f.form.Get("hostName")
	newHostName, ok := rawHostName.(string)
	if !ok {
		return errors.New("unable to type assert hostName")
	}
	rawDescription := f.form.Get("description")
	newDescription, ok := rawDescription.(string)
	if !ok {
		return errors.New("unable to type assert description")
	}
	rawIPAddress := f.form.Get("ipAddress")
	newIPAddress, ok := rawIPAddress.(string)
	if !ok {
		return errors.New("unable to type assert ipAddress")
	}
	// load entry params and pass to database
	entryParams := database.CreateEntryParams{
		Category:     int64(newCategory),
		FriendlyName: newFriendlyName,
		HostName:     newHostName,
		Description:  newDescription,
		IpAddress:    newIPAddress,
	}
	_, err := f.db.CreateEntry(context.Background(), entryParams)
	if err != nil {
		return err
	}
	// update state now that entry is done
	l, ok := m.listViewMod.(*ListViewModel)
	if ok {
		l.refreshList()
	}
	m.state = listView
	return nil
}

func deleteEntryHandler(m *MainModel, f *FormViewModel) error {
	okToDelete, ok := f.form.Get("delete").(bool)
	if !ok {
		return errors.New("unable to type assert okToDelete")
	}
	if !okToDelete {
		return nil
	}
	currentListView, ok := m.listViewMod.(*ListViewModel)
	if !ok {
		return errors.New("unable to typeassert listViewModel")
	}
	focusedListItem, ok := currentListView.lists[currentListView.focused].Items()[currentListView.lists[currentListView.focused].Index()].(Entry)
	if !ok {
		return errors.New("unable to type assert list item to Entry")
	}
	err := f.db.DeleteEntry(context.Background(), int64(focusedListItem.id))
	if err != nil {
		return err
	}
	l, ok := m.listViewMod.(*ListViewModel)
	if ok {
		l.refreshList()
	}
	m.state = listView
	return nil
}

func editEntryHandler(m *MainModel, f *FormViewModel) error {
	// get form values and type assert to struct field types
	rawCategory := f.form.Get("category")
	newCategory, ok := rawCategory.(category)
	if !ok {
		return errors.New("unable to type assert category")
	}
	rawFriendlyName := f.form.Get("friendlyName")
	newFriendlyName, ok := rawFriendlyName.(string)
	if !ok {
		return errors.New("unable to type assert friendlyName")
	}
	rawHostName := f.form.Get("hostName")
	newHostName, ok := rawHostName.(string)
	if !ok {
		return errors.New("unable to type assert hostName")
	}
	rawDescription := f.form.Get("description")
	newDescription, ok := rawDescription.(string)
	if !ok {
		return errors.New("unable to type assert description")
	}
	rawIPAddress := f.form.Get("ipAddress")
	newIPAddress, ok := rawIPAddress.(string)
	if !ok {
		return errors.New("unable to type assert ipAddress")
	}

	// get focused item to fetch ID
	currentListView, ok := m.listViewMod.(*ListViewModel)
	if !ok {
		return errors.New("unable to typeassert listViewModel")
	}
	focusedListItem, ok := currentListView.lists[currentListView.focused].Items()[currentListView.lists[currentListView.focused].Index()].(Entry)
	if !ok {
		return errors.New("unable to type assert list item to Entry")
	}

	editEntryParams := database.EditEntryByIDParams{
		Category:     int64(newCategory),
		FriendlyName: newFriendlyName,
		HostName:     newHostName,
		Description:  newDescription,
		IpAddress:    newIPAddress,
		ID:           int64(focusedListItem.id),
	}
	_, err := f.db.EditEntryByID(context.Background(), editEntryParams)
	if err != nil {
		return err
	}
	l, ok := m.listViewMod.(*ListViewModel)
	if ok {
		l.refreshList()
	}
	m.state = listView
	return nil
}
