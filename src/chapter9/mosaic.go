package main

import (
	"fmt"
	"html/template"
	"image"
	"image/color"
	"image/draw"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"strconv"
	"time"
)

var TILESDB map[string][3]float64

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("public"))
	mux.Handle("/static", http.StripPrefix("/static", files))
	mux.HandleFunc("/", upload)
	mux.HandleFunc("/mosaic", mosaic)
	server := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
	TILESDB = tilesDB()
	fmt.Println("Mosaic server started.")
	server.ListenAndServe()
}

func upload(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("upload.html")
	t.Execute(w, nil)
}

func mosaic(w http.ResponseWriter, r *http.Request) {
	t0 := time.Now()

	r.ParseMultipartForm(10485760)
	file, _, _ := r.FormFile("image")
	defer file.Close()

	tileSize, _ := strconv.Atoi(r.FormValue("tile_size"))
	original, _, _ := image.Decode(file)
	bounds := original.Bounds()

	newimage := image.NewNRGBA(image.Rect(bounds.Min.X, bounds.Min.X, bounds.Max.Y, bounds.Max.Y))

	db := tilesDB()

	sp := image.Point{0, 0}
	for y := bounds.Min.Y; y < bounds.Max.Y; y += tileSize {
		for x := bounds.Min.X; y < bounds.Max.X; x += tileSize {
			r, g, b, _ := original.At(x, y).RGBA()
			color := [3]float64{float64(r), float64(g), float64(b)}

			nearest :=nearest(color , &db)

			file , err := os.Open(nearest)
			if err == nil {
				img , _ , err :=image.Decode(file)

				if err == nil {
					t := resize(img , tileSize)
					tile := t.SubImage(t.Bounds())
					tileBounds := image.Rect(x, y ,x +tileSize , y + tileSize)
					draw.Draw(newimage , tileBounds , tile , sp , draw.Src)
				}else{
					fmt.Println("error:", err, nearest)
				}
			}else{
				fmt.Println("error:", nearest)
			}
			file.Close()
		}
	}

}

//拿到图片的rgb颜色的平均值
func averageColor(img image.Image) [3]float64 {
	bounds := img.Bounds()

	r, g, b := 0.0, 0.0, 0.0

	for y := bounds.Min.Y; y < bounds.Max.Y; y ++ {
		for x := bounds.Min.X; x < bounds.Max.X; x ++ {
			r1, g1, b1, _ := img.At(x, y).RGBA()
			r, g, b = r+float64(r1), g+float64(g1), b+float64(b1)
		}
	}

	totalPixels := float64(bounds.Max.X * bounds.Max.Y)

	return [3]float64{r / totalPixels, g / totalPixels, b / totalPixels}

}

func resize(in image.Image, newWidth int) image.NRGBA {
	bounds := in.Bounds()
	ratio := bounds.Dx() / newWidth //宽度
	out := image.NewNRGBA(image.Rect(bounds.Min.X/ratio, bounds.Min.Y/ratio, bounds.Max.X/ratio, bounds.Max.Y/ratio))

	for y, j := bounds.Min.Y, bounds.Min.Y; y < bounds.Max.Y; y, j = y+ratio, j+1 {

		for x, i := bounds.Min.X, bounds.Min.X; x < bounds.Max.X; x, i = x+ratio, i+1 {

			r, g, b, a := in.At(x, y).RGBA()
			out.SetNRGBA(i, j, color.NRGBA{R: uint8(r >> 8), G: uint8(g >> 8), B: uint8(b >> 8), A: uint8(a >> 8)})

		}
	}

	return *out

}

func tilesDB() map[string][3]float64 {

	fmt.Println("Start populating tiles db ...")
	db := make(map[string][3]float64)
	files, _ := ioutil.ReadDir("tiles")

	for _, f := range files {
		name := "tiles/" + f.Name()
		file, err := os.Open(name)
		if err == nil {
			img, _, err := image.Decode(file)

			if err == nil {
				db[name] = averageColor(img)
			} else {
				fmt.Println("error in populating TILEDB:", err, name)
			}
		} else {
			fmt.Println("cannot open file", name, err)

		}
		file.Close()
	}

	fmt.Println("Finished populating tiles db.")
	return db

}

func nearest(target [3]float64, db *map[string][3]float64) string {
	var filename string
	smallest := 1000000.0
	for k, v := range *db {
		dist := distance(target, v)
		if dist < smallest {
			filename, smallest = k, dist
		}
	}
	delete(*db, filename)
	return filename
}

func distance(p1 [3]float64, p2 [3]float64) float64 {
	return math.Sqrt(sq(p2[0]-p1[0]) + sq(p2[1]-p1[1]) + sq(p2[2]-p1[2]))
}

func sq(n float64) float64 {
	return n * n
}
