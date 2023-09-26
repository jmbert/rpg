package graphics

import "github.com/veandco/go-sdl2/sdl"

type Event sdl.Event

type KeyInput map[string](func(Event))

func MainloopIter(inputMap KeyInput) uint {

	var event sdl.Event

	event = sdl.PollEvent()
	for ; event != nil; event = sdl.PollEvent() {
		switch event.GetType() {
		case sdl.WINDOWEVENT:
			winEvent := event.(*sdl.WindowEvent)
			switch winEvent.Event {
			case sdl.WINDOWEVENT_CLOSE:
				return 1
			}
		case sdl.MOUSEBUTTONDOWN:
			mEvent := event.(*sdl.MouseButtonEvent)
			if mEvent.Button == sdl.BUTTON_LEFT {
				inputMap["mbleft"](event)
			}
		}
	}

	renderer.Present()

	return 0
}
