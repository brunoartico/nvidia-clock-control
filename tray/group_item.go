package tray

import "github.com/getlantern/systray"

type MenuGroupItem struct {
	item *systray.MenuItem
	value int
}

func (menuGroupItem *MenuGroupItem) Hide() {
	menuGroupItem.item.Hide()
}

func (menuGroupItem *MenuGroupItem) Show() {
	menuGroupItem.item.Show()
}

func (menuGroupItem *MenuGroupItem) GetValue() int {
	return menuGroupItem.value
}

func (menuGroupItem *MenuGroupItem) Check() {
	menuGroupItem.item.Check()
}

func (menuGroupItem *MenuGroupItem) Uncheck() {
	menuGroupItem.item.Uncheck()
}
