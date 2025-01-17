/**
 * Copyright 2023 Google Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"net/http"
)

func main() {
	// Define the routes for blue and red image handling
	http.HandleFunc("/blue", blueHandler)  // Handle blue image at "/blue"
	http.HandleFunc("/red", redHandler)    // Handle red image at "/red"
	http.ListenAndServe(":8080", nil)      // Start the server on port 8080
}

func blueHandler(w http.ResponseWriter, r *http.Request) {
	// Create a 100x100 image with blue background
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	draw.Draw(img, img.Bounds(), &image.Uniform{color.RGBA{0, 0, 255, 255}}, image.ZP, draw.Src)
	w.Header().Set("Content-Type", "image/png")  // Set response header to image/png
	png.Encode(w, img)                           // Encode image as PNG and send to client
}

func redHandler(w http.ResponseWriter, r *http.Request) {
	// Create a 100x100 image with red background
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	draw.Draw(img, img.Bounds(), &image.Uniform{color.RGBA{255, 0, 0, 255}}, image.ZP, draw.Src)
	w.Header().Set("Content-Type", "image/png")  // Set response header to image/png
	png.Encode(w, img)                           // Encode image as PNG and send to client
}
