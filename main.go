package main

import (
	"html/template"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/length", lengthHandler)
	http.HandleFunc("/weight", weightHandler)
	http.HandleFunc("/temperature", temperatureHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	println("Server running on http://localhost:9090")
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		panic(err)
	}

}

func lengthHandler(w http.ResponseWriter, r *http.Request) {
	type PageData struct {
		Input     float64
		FromUnit  string
		ToUnit    string
		Result    float64
		Submitted bool
	}

	data := PageData{}

	if r.Method == http.MethodPost {
		r.ParseForm()
		inputVal, _ := strconv.ParseFloat(r.FormValue("input"), 64)
		from := r.FormValue("from")
		to := r.FormValue("to")

		data.Input = inputVal
		data.FromUnit = from
		data.ToUnit = to
		data.Result = convertLength(inputVal, from, to)
		data.Submitted = true
	}

	tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/length.html"))
	tmpl.ExecuteTemplate(w, "layout", data)
}

func weightHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/weight.html")
}

func temperatureHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/temperature.html")
}

func convertLength(value float64, from, to string) float64 {
	unitMap := map[string]float64{
		"millimeter": 0.001,
		"centimeter": 0.01,
		"meter":      1,
		"kilometer":  1000,
		"inch":       0.0254,
		"foot":       0.3048,
		"yard":       0.9144,
		"mile":       1609.34,
	}

	return value * unitMap[from] / unitMap[to]
}
