package pterm_test

import (
	"reflect"
	"testing"
	"time"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
	"github.com/MarvinJWendt/testza"

	"github.com/pterm/pterm"
)

func TestInteractiveTextInputPrinter_WithDefaultText(t *testing.T) {
	p := pterm.DefaultInteractiveTextInput.WithDefaultText("default")
	testza.AssertEqual(t, p.DefaultText, "default")
}

func TestInteractiveTextInputPrinter_WithDefaultValue(t *testing.T) {
	p := pterm.DefaultInteractiveTextInput.WithDefaultValue("default")
	testza.AssertEqual(t, p.DefaultValue, "default")
}

func TestInteractiveTextInputPrinter_WithDelimiter(t *testing.T) {
	p := pterm.DefaultInteractiveTextInput.WithDelimiter(">>")
	testza.AssertEqual(t, p.Delimiter, ">>")
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
		time.Sleep(1 * time.Millisecond)
		keyboard.SimulateKeyPress('a')
		keyboard.SimulateKeyPress('b')
		keyboard.SimulateKeyPress('c')
		keyboard.SimulateKeyPress(keys.Enter)
	}()
	result, _ := pterm.DefaultInteractiveTextInput.WithMask("*").Show()
	testza.AssertEqual(t, "abc", result)
}

func TestInteractiveTextInputPrinter_WithCancel(t *testing.T) {
	go func() {
		time.Sleep(1 * time.Millisecond)
		keyboard.SimulateKeyPress(keys.CtrlC)
	}()
	result, _ := pterm.DefaultInteractiveTextInput.WithMask("*").Show()
	testza.AssertEqual(t, "", result)
}

func TestInteractiveTextInputPrinter_OnEnter(t *testing.T) {
	go func() {
		keyboard.SimulateKeyPress(keys.Enter)
	}()
	result, _ := pterm.DefaultInteractiveTextInput.WithDefaultValue("default").Show()
	testza.AssertEqual(t, "default", result)
}

func TestInteractiveTextInputPrinter_WithMultiLineOnTab(t *testing.T) {
	go func() {
		keyboard.SimulateKeyPress(keys.Tab)
	}()
	result, _ := pterm.DefaultInteractiveTextInput.
		WithMultiLine(true).
		WithDefaultValue("default").Show()
	testza.AssertEqual(t, "default", result)
}

func TestInteractiveTextInputPrinter_WithMultiLineOnUp(t *testing.T) {
	go func() {
		time.Sleep(1 * time.Millisecond)
		keyboard.SimulateKeyPress(keys.Enter)
		keyboard.SimulateKeyPress("second line")
		keyboard.SimulateKeyPress(keys.Up)
		keyboard.SimulateKeyPress("first line")
		keyboard.SimulateKeyPress(keys.Tab)
	}()
	result, _ := pterm.DefaultInteractiveTextInput.
		WithMultiLine(true).Show()
	testza.AssertEqual(t, "first line\nsecond line", result)
}

func TestInteractiveTextInputPrinter_WithMultiLineOnDown(t *testing.T) {
	go func() {
		time.Sleep(1 * time.Millisecond)
		keyboard.SimulateKeyPress("a")
		keyboard.SimulateKeyPress(keys.Enter)
		keyboard.SimulateKeyPress(keys.Enter)
		keyboard.SimulateKeyPress(keys.Up)
		keyboard.SimulateKeyPress("b")
		keyboard.SimulateKeyPress(keys.Down)
		keyboard.SimulateKeyPress("c")
		keyboard.SimulateKeyPress(keys.Tab)
	}()
	result, _ := pterm.DefaultInteractiveTextInput.
		WithMultiLine(true).Show()
	testza.AssertEqual(t, "a\nb\nc", result)
}

func TestInteractiveTextInputPrinter_WithMultiLineOnLeft(t *testing.T) {
	go func() {
		time.Sleep(1 * time.Millisecond)
		keyboard.SimulateKeyPress(keys.Enter)
		keyboard.SimulateKeyPress("a")
		keyboard.SimulateKeyPress(keys.Left)
		keyboard.SimulateKeyPress(keys.Left)
		keyboard.SimulateKeyPress("b")
		keyboard.SimulateKeyPress(keys.Tab)
	}()
	result, _ := pterm.DefaultInteractiveTextInput.
		WithMultiLine(true).Show()
	testza.AssertEqual(t, "b\na", result)
}

func TestInteractiveTextInputPrinter_WithMultiLineOnRight(t *testing.T) {
	go func() {
		time.Sleep(1 * time.Millisecond)
		keyboard.SimulateKeyPress('a')
		keyboard.SimulateKeyPress(keys.Enter)
		keyboard.SimulateKeyPress(keys.Up)
		keyboard.SimulateKeyPress(keys.Right)
		keyboard.SimulateKeyPress(keys.Right)
		keyboard.SimulateKeyPress("b")
		keyboard.SimulateKeyPress(keys.Tab)
	}()
	result, _ := pterm.DefaultInteractiveTextInput.
		WithMultiLine(true).Show()
	testza.AssertEqual(t, "a\nb", result)
}

func TestInteractiveTextInputPrinter_OnSpace(t *testing.T) {
	go func() {
		keyboard.SimulateKeyPress(keys.Space)
		keyboard.SimulateKeyPress(keys.Tab)
	}()
	result, _ := pterm.DefaultInteractiveTextInput.
		WithMultiLine(true).
		WithDefaultValue("default").Show()
	testza.AssertEqual(t, "default ", result)
}

func TestInteractiveTextInputPrinter_OnDelete(t *testing.T) {
	go func() {
		keyboard.SimulateKeyPress(keys.Left)
		keyboard.SimulateKeyPress(keys.Delete)
		keyboard.SimulateKeyPress(keys.Enter)
	}()
	result, _ := pterm.DefaultInteractiveTextInput.
		WithDefaultValue("a").Show()
	testza.AssertEqual(t, "", result)
}

func TestInteractiveTextInputPrinter_OnBackspace(t *testing.T) {
	go func() {
		keyboard.SimulateKeyPress(keys.Backspace)
		keyboard.SimulateKeyPress(keys.Enter)
	}()
	result, _ := pterm.DefaultInteractiveTextInput.
		WithDefaultValue("a").Show()
	testza.AssertEqual(t, "", result)
}

func TestInteractiveTextInputPrinter_WithMultiLineOnDelete(t *testing.T) {
	go func() {
		keyboard.SimulateKeyPress(keys.Enter)
		keyboard.SimulateKeyPress('a')
		keyboard.SimulateKeyPress(keys.Up)
		keyboard.SimulateKeyPress(keys.Delete)
		keyboard.SimulateKeyPress(keys.Tab)
	}()
	result, _ := pterm.DefaultInteractiveTextInput.
		WithMultiLine(true).
		WithDefaultValue("a").Show()
	testza.AssertEqual(t, "aa", result)
}

func TestInteractiveTextInputPrinter_WithMultiLineOnBackspace(t *testing.T) {
	go func() {
		keyboard.SimulateKeyPress(keys.Enter)
		keyboard.SimulateKeyPress(keys.Backspace)
		keyboard.SimulateKeyPress(keys.Tab)
	}()
	result, _ := pterm.DefaultInteractiveTextInput.
		WithMultiLine(true).
		WithDefaultValue("a").Show()
	testza.AssertEqual(t, "a", result)
}

func TestInteractiveTextInputPrinter_WithMultiLineAltPressed(t *testing.T) {
	go func() {
		keyboard.SimulateKeyPress(keys.Key{AltPressed: true, Code: keys.Enter})
		keyboard.SimulateKeyPress(keys.Tab)
	}()
	result, _ := pterm.DefaultInteractiveTextInput.
		WithMultiLine(true).
		WithDefaultValue("a").Show()
	testza.AssertEqual(t, "a\n", result)
}

func TestInteractiveTextInputPrinter_OnLeftRight(t *testing.T) {
	go func() {
		keyboard.SimulateKeyPress(keys.Left)
		keyboard.SimulateKeyPress("left")
		keyboard.SimulateKeyPress(keys.Right)
		keyboard.SimulateKeyPress("right")
		keyboard.SimulateKeyPress(keys.Enter)
	}()
	result, _ := pterm.DefaultInteractiveTextInput.
		WithDefaultValue("default").Show()
	testza.AssertEqual(t, "defaullefttright", result)
}

func TestInteractiveTextInputPrinter_WithMultiLineOnLeftRight(t *testing.T) {
	go func() {
		time.Sleep(1 * time.Millisecond)
		keyboard.SimulateKeyPress("a")
		keyboard.SimulateKeyPress(keys.Enter)
		keyboard.SimulateKeyPress("b")
		keyboard.SimulateKeyPress(keys.Tab)
	}()
	result, _ := pterm.DefaultInteractiveTextInput.
		WithMultiLine(true).
		Show("Enter")
	testza.AssertEqual(t, "a\nb", result)
}

func TestInteractiveTextInputPrinter_WithOnInterruptFunc(t *testing.T) {
	// OnInterrupt function defaults to nil
	pd := pterm.InteractiveTextInputPrinter{}
	testza.AssertNil(t, pd.OnInterruptFunc)

	// Verify OnInterrupt is set
	exitfunc := func() {}
	p := pterm.DefaultInteractiveTextInput.WithOnInterruptFunc(exitfunc)
	testza.AssertEqual(t, reflect.ValueOf(p.OnInterruptFunc).Pointer(), reflect.ValueOf(exitfunc).Pointer())
}
