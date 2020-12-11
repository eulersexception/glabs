package util

import (
	"fyne.io/fyne/widget"
)

type TapTabItem struct {
	*widget.TabItem

	OnTapped          func() // `json:"-"`
	OnTappedSecondary func() // `json:"-"`
}

func NewTapTabItem(text string, tappedLeft func(), tappedRight func()) *TapTabItem {
	x := &TapTabItem{
		widget.NewTabItem(text, nil),
		tappedLeft, tappedRight,
	}

	return x
}
