package tray

import "github.com/getlantern/systray"

type MenuGroup struct {
	items []*systray.MenuItem
}

func (group *MenuGroup) Add(item *systray.MenuItem, behavior func()) {
	group.items = append(group.items, item)

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
