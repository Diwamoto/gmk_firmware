//go:build tinygo

package keyboard

import (
	"machine"
	kbd "machine/usb/hid/keyboard"
)

type MatrixKeyboard struct {
	State    [][]State
	Keys     [][][]Keycode
	options  Options
	callback Callback

	Col []machine.Pin
	Row []machine.Pin
}

func (d *Device) AddMatrixKeyboard(colPins, rowPins []machine.Pin, keys [][][]Keycode, opt ...Option) *MatrixKeyboard {
	state := [][]State{}
	col := len(colPins)
	row := len(rowPins)

	for r := 0; r < row; r++ {
		column := make([]State, col)
		state = append(state, column)
	}

	for c := range colPins {
		colPins[c].Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	}
	for r := range rowPins {
		rowPins[r].Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	}

	o := Options{}
	for _, f := range opt {
		f(&o)
	}

	k := &MatrixKeyboard{
		Col:      colPins,
		Row:      rowPins,
		State:    state,
		Keys:     keys,
		options:  o,
		callback: func(layer, row, col int, state State) {},
	}

	if kbd.Keys[0][0][0] == 0 {
		for l := range keys {
			for r := range keys[l] {
				for c := range keys[l][r] {
					kbd.Keys[l][r][c] = kbd.Keycode(keys[l][r][c] & 0x0FFF)
				}
			}
		}
	}

	d.kb = append(d.kb, k)
	return k
}

func (d *MatrixKeyboard) SetCallback(fn Callback) {
	d.callback = fn
}

func (d *MatrixKeyboard) Get() [][]State {
	for c := range d.Col {
		for r := range d.Row {
			//d.State[r][c] = d.Row[r].Get()
			current := false
			if !d.options.InvertDiode {
				d.Col[c].Configure(machine.PinConfig{Mode: machine.PinOutput})
				d.Col[c].High()
				current = d.Row[r].Get()
			} else {
				d.Row[r].Configure(machine.PinConfig{Mode: machine.PinOutput})
				d.Row[r].High()
				current = d.Col[c].Get()
			}
			switch d.State[r][c] {
			case None:
				if current {
					d.State[r][c] = NoneToPress
				} else {
				}
			case NoneToPress:
				if current {
					d.State[r][c] = Press
					d.callback(0, r, c, Press)
				} else {
					d.State[r][c] = PressToRelease
					d.callback(0, r, c, Press)
					d.callback(0, r, c, PressToRelease)
				}
			case Press:
				if current {
				} else {
					d.State[r][c] = PressToRelease
					d.callback(0, r, c, PressToRelease)
				}
			case PressToRelease:
				if current {
					d.State[r][c] = NoneToPress
					d.callback(0, r, c, Press)
				} else {
					d.State[r][c] = None
				}
			}
			if !d.options.InvertDiode {
				d.Col[c].Low()
				d.Col[c].Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
			} else {
				d.Row[r].Low()
				d.Row[r].Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
			}
		}
	}

	return d.State
}

func (d *MatrixKeyboard) Key(layer, row, col int) Keycode {
	kc := Keycode(kbd.Keys[layer][row][col])
	switch kc {
	case 0x5610, 0x5101, 0x5102, 0x5103, 0x5104, 0x5105:
		return 0xFF00 | (kc & 0x00FF)
	default:
	}
	return 0xF000 | kc
}

func (d *MatrixKeyboard) Init() error {
	return nil
}
