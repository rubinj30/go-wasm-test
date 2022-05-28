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

const api_url = "https://pokeapi.co/api/v2/pokemon"

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
	// update the text in <div id="clock">
	// It's done like this in JavaScript:
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
	// last_character := js.Value(class_names)[len(js.Value(class_names))-3:]
	// if last_character == "2" {
	// 	js.Value(class_names)[len(js.Value(class_names))-1:] = "1"
	// } else {
	// 	js.Value(class_names)[len(js.Value(class_names))-1:] = "2"
	// }
	// block.Set("className", js.Value(js.Value(class_names)))
	return nil
}

func set_handle_click() interface{} {
	log.Print("Setting handle click")
	trigger := js.Global().Get("document").Call("getElementById", "trigger")
	trigger.Call("addEventListener", "click", js.FuncOf(turn_block))
	return nil
}

func main() {
	// Set up a recurring timer event to call update_time() every 200 ms.
	// It's done like this in JavaScript:
	//	setInterval(update_time,200)

	// Create JavaScript callback connected to update_time()

	println("calling get_pokemon")
	timer_cb := js.FuncOf(update_time)
	// turn_cb := js.FuncOf(turn_block)

	// Set timer to call timer_cb() every 200 ms.

	js.Global().Call("setInterval", timer_cb, "200")
	// js.Global().Call("setInterval", turn_cb, "200")
	set_handle_click()

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
