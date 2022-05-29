// Digital Clock - Shows use of setInterval() in Go/WebAssembly
// build: GOOS=js GOARCH=wasm go build -o lib.wasm main.go

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"syscall/js"
)

func update_time(this js.Value, args []js.Value) interface{} {
	// Get the current date in this locale
	// It's done like this in JavaScript:
	//	date = new Date()
	//	s = date.toLocaleTimeString()
	date := js.Global().Get("Date").New()
	s := date.Call("toLocaleTimeString").String()

	// update the text in <div id="clock">
	// It's done like this in JavaScript:
	//	document.getElementById("clock").textContent = s
	js.Global().Get("document").Call("getElementById", "clock").Set("textContent", s)
	return nil
}

func turn_block(this js.Value, args []js.Value) interface{} {
	block := js.Global().Get("document").Call("getElementById", "block")
	class_names := block.Get("className")
	js.Global().Get("document").Call("getElementById", "word").Set("innerText", class_names)
	last_character := class_names.String()[len(class_names.String())-1:]
	if last_character == "5" {
		block.Set("className", "block rotated-0")
	} else {
		new_last_char := string(rune(last_character[0]) + 1)
		first_part_of_class_names := class_names.String()[:len(class_names.String())-1]
		new_class_names := first_part_of_class_names + new_last_char
		log.Print("CLASSNAME", new_class_names, new_last_char)
		block.Set("className", new_class_names)
	}
	return nil
}

// Could re-use this and add function and element ID value params
func set_handle_clicks() interface{} {
	document := js.Global().Get("document")

	// Create JavaScript callback connected to turn_block()
	trigger := document.Call("getElementById", "trigger")
	trigger.Call("addEventListener", "click", js.FuncOf(turn_block))

	// // Create JavaScript callback connected to add_dot()
	add_dot_button := document.Call("getElementById", "add_dot")
	add_dot_button.Call("addEventListener", "click", js.FuncOf(add_dot))
	return nil
}

func add_dot(this js.Value, args []js.Value) interface{} {

	document := js.Global().Get("document")
	parentDiv := document.Call("getElementById", "block")
	nodeList := document.Call("querySelectorAll", ".new_child")

	if nodeList.Length() == 12 {
		parentDiv.Set("innerHTML", "")
	} else {
		newDiv := document.Call("createElement", "div")
		newDiv.Set("className", "new_child")

		blockDiv := document.Call("getElementById", "block")
		blockDiv.Call("appendChild", newDiv)
	}
	return nil
}

func main() {
	// Set up a recurring timer event to call update_time() every 200 ms.
	// It's done like this in JavaScript:
	//	setInterval(update_time,200)

	// Create JavaScript callback connected to update_time()
	timer_cb := js.FuncOf(update_time)

	// Set timer to call timer_cb() every 200 ms.
	js.Global().Call("setInterval", timer_cb, "200")
	set_handle_clicks()
	// An empty select blocks, so the main() function will never exit.
	// This allows the event handler callbacks to continue operating.
	select {}
}

///////
///////
///////
///////
///////
///////
///////
///////
///////
///////
///////
func get_pokemon() {
	resp, err := http.Get("https://pokeapi.co/api/v2/pokemon")
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// json.NewDecoder(body)
	// log.Println(fmt.Sprintf("%q\n", body))
	var data interface{} // TopTracks
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(data)
}
