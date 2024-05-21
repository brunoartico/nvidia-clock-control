package tray

import "github.com/getlantern/systray"

type MenuGroup struct {
	items []*MenuGroupItem
}

func (group *MenuGroup) Add(item *systray.MenuItem, value int, behavior func()) {
	group.items = append(group.items, &MenuGroupItem{item, value})

	go func() {
		for {
			<-item.ClickedCh
			behavior()
			for _, groupItem := range group.items {
				groupItem.Uncheck()
			}
			item.Check()
		}
	}()
}

func (group *MenuGroup) UncheckAll() {
	for _, groupItem := range group.items {
		groupItem.Uncheck()
	}
}

func (group *MenuGroup) ShowAll() {
	for _, groupItem := range group.items {
		groupItem.Show()
	}
}

func (group *MenuGroup) GetAll() []*MenuGroupItem {
	return group.items
}
