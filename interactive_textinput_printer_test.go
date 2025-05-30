package pterm_test

import (
	"testing"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
	"github.com/MarvinJWendt/testza"

	"github.com/pterm/pterm"
)

func TestInteractiveTextInputPrinter_WithDefaultText(t *testing.T) {
	p := pterm.DefaultInteractiveTextInput.WithDefaultText("default")
	testza.AssertEqual(t, p.DefaultText, "default")
}

func TestInteractiveTextInputPrinter_WithMultiLine_true(t *testing.T) {
	p := pterm.DefaultInteractiveTextInput.WithMultiLine()
	testza.AssertTrue(t, p.MultiLine)
}

func TestInteractiveTextInputPrinter_WithMultiLine_false(t *testing.T) {
	p := pterm.DefaultInteractiveTextInput.WithMultiLine(false)
	testza.AssertFalse(t, p.MultiLine)
}

func TestInteractiveTextInputPrinter_WithTextStyle(t *testing.T) {
	style := pterm.NewStyle(pterm.FgRed)
	p := pterm.DefaultInteractiveTextInput.WithTextStyle(style)
	testza.AssertEqual(t, p.TextStyle, style)
}

func TestInteractiveTextInputPrinter_WithMask(t *testing.T) {
	go func() {
		keyboard.SimulateKeyPress('a')
		keyboard.SimulateKeyPress('b')
		keyboard.SimulateKeyPress('c')
		keyboard.SimulateKeyPress(keys.Enter)
	}()
	result, _ := pterm.DefaultInteractiveTextInput.WithMask("*").Show()
	testza.AssertEqual(t, result, "abc")
}
