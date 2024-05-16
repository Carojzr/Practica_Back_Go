package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Superhero representa la estructura de datos de un superhéroe
type Superhero struct {
	Name      string `json:"name"`
	Biography struct {
		FullName string `json:"fullName"`
	} `json:"biography"`
	Powerstats struct {
		Intelligence int `json:"intelligence"`
		Strength     int `json:"strength"`
		Speed        int `json:"speed"`
		Durability   int `json:"durability"`
		Power        int `json:"power"`
		Combat       int `json:"combat"`
	} `json:"powerstats"`
	Images struct {
		Xs string `json:"xs"`
		Sm string `json:"sm"`
		Md string `json:"md"`
		Lg string `json:"lg"`
	} `json:"images"`
}

// Mapa de superhéroes
var superheroes = map[string]Superhero{
	"Wolverine": {
		Name: "Wolverine",
		Biography: struct {
			FullName string `json:"fullName"`
		}{FullName: "John Logan"},
		Powerstats: struct {
			Intelligence int `json:"intelligence"`
			Strength     int `json:"strength"`
			Speed        int `json:"speed"`
			Durability   int `json:"durability"`
			Power        int `json:"power"`
			Combat       int `json:"combat"`
		}{63, 32, 50, 100, 89, 100},
		Images: struct {
			Xs string `json:"xs"`
			Sm string `json:"sm"`
			Md string `json:"md"`
			Lg string `json:"lg"`
		}{
			"https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/xs/717-wolverine.jpg",
			"https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/sm/717-wolverine.jpg",
			"https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/md/717-wolverine.jpg",
			"https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/lg/717-wolverine.jpg",
		},
	},
	"Spider-Man": {
        Name: "Spider-Man",
        Biography: struct {
            FullName string `json:"fullName"`
        }{FullName: "Peter Parker"},
        Powerstats: struct {
            Intelligence int `json:"intelligence"`
            Strength     int `json:"strength"`
            Speed        int `json:"speed"`
            Durability   int `json:"durability"`
            Power        int `json:"power"`
            Combat       int `json:"combat"`
        }{90, 55, 67, 75, 74, 85},
        Images: struct {
            Xs string `json:"xs"`
            Sm string `json:"sm"`
            Md string `json:"md"`
            Lg string `json:"lg"`
        }{
            "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/xs/620-spider-man.jpg",
            "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/sm/620-spider-man.jpg",
            "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/md/620-spider-man.jpg",
            "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/lg/620-spider-man.jpg",
        },
    },
	"Iron Man": {
        Name: "Iron Man",
        Biography: struct {
            FullName string `json:"fullName"`
        }{FullName: "Tony Stark"},
        Powerstats: struct {
            Intelligence int `json:"intelligence"`
            Strength     int `json:"strength"`
            Speed        int `json:"speed"`
            Durability   int `json:"durability"`
            Power        int `json:"power"`
            Combat       int `json:"combat"`
        }{100, 85, 58, 85, 100, 64},
        Images: struct {
            Xs string `json:"xs"`
            Sm string `json:"sm"`
            Md string `json:"md"`
            Lg string `json:"lg"`
        }{
            "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/xs/346-iron-man.jpg",
            "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/sm/346-iron-man.jpg",
            "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/md/346-iron-man.jpg",
            "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/lg/346-iron-man.jpg",
        },
    },
	"Thor": {
        Name: "Thor",
        Biography: struct {
            FullName string `json:"fullName"`
        }{FullName: "Thor Odinson"},
        Powerstats: struct {
            Intelligence int `json:"intelligence"`
            Strength     int `json:"strength"`
            Speed        int `json:"speed"`
            Durability   int `json:"durability"`
            Power        int `json:"power"`
            Combat       int `json:"combat"`
        }{69, 100, 83, 100, 100, 100},
        Images: struct {
            Xs string `json:"xs"`
            Sm string `json:"sm"`
            Md string `json:"md"`
            Lg string `json:"lg"`
        }{
            "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/xs/659-thor.jpg",
            "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/sm/659-thor.jpg",
            "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/md/659-thor.jpg",
            "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/lg/659-thor.jpg",
        },
	},
	"Aquaman": {
        Name: "Aquaman",
        Biography: struct {
            FullName string `json:"fullName"`
        }{FullName: "Orin"},
        Powerstats: struct {
            Intelligence int `json:"intelligence"`
            Strength     int `json:"strength"`
            Speed        int `json:"speed"`
            Durability   int `json:"durability"`
            Power        int `json:"power"`
            Combat       int `json:"combat"`
        }{81, 85, 79, 80, 100, 80},
        Images: struct {
            Xs string `json:"xs"`
            Sm string `json:"sm"`
            Md string `json:"md"`
            Lg string `json:"lg"`
        }{
			"https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/xs/38-aquaman.jpg",
			"https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/sm/38-aquaman.jpg",
			"https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/md/38-aquaman.jpg",
			"https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/lg/38-aquaman.jpg",
        },
    },
	"Batgirl": {
        Name: "Batgirl",
        Biography: struct {
            FullName string `json:"fullName"`
        }{FullName: "Barbara Gordon"},
        Powerstats: struct {
            Intelligence int `json:"intelligence"`
            Strength     int `json:"strength"`
            Speed        int `json:"speed"`
            Durability   int `json:"durability"`
            Power        int `json:"power"`
            Combat       int `json:"combat"`
        }{88,11,33,40,34,90},
        Images: struct {
            Xs string `json:"xs"`
            Sm string `json:"sm"`
            Md string `json:"md"`
            Lg string `json:"lg"`
        }{
			"https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/xs/63-batgirl.jpg",
       		"https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/sm/63-batgirl.jpg",
      		"https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/md/63-batgirl.jpg",
      		"https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/lg/63-batgirl.jpg",
        },
    },
	"Loki": {
        Name: "Loki",
        Biography: struct {
            FullName string `json:"fullName"`
        }{FullName: "Loki Laufeyson"},
        Powerstats: struct {
            Intelligence int `json:"intelligence"`
            Strength     int `json:"strength"`
            Speed        int `json:"speed"`
            Durability   int `json:"durability"`
            Power        int `json:"power"`
            Combat       int `json:"combat"`
        }{88,63,46,85,100,60},
        Images: struct {
            Xs string `json:"xs"`
            Sm string `json:"sm"`
            Md string `json:"md"`
            Lg string `json:"lg"`
        }{
			"https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/xs/414-loki.jpg",
			"https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/sm/414-loki.jpg",
			"https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/md/414-loki.jpg",
			"https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/lg/414-loki.jpg",
        },
    },
    "Wonder Woman": {
        Name: "Wonder Woman",
        Biography: struct {
            FullName string `json:"fullName"`
        }{FullName: "Diana Prince"},
        Powerstats: struct {
            Intelligence int `json:"intelligence"`
            Strength     int `json:"strength"`
            Speed        int `json:"speed"`
            Durability   int `json:"durability"`
            Power        int `json:"power"`
            Combat       int `json:"combat"`
        }{88,100,79,100,100,100},
        Images: struct {
            Xs string `json:"xs"`
            Sm string `json:"sm"`
            Md string `json:"md"`
            Lg string `json:"lg"`
        }{
			"https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/xs/720-wonder-woman.jpg",
			"https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/sm/720-wonder-woman.jpg",
			"https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/md/720-wonder-woman.jpg",
			"https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/lg/720-wonder-woman.jpg",
        },
    },
    "Black Window": {
        Name: "Black Window",
        Biography: struct {
            FullName string `json:"fullName"`
        }{FullName: "Natasha Romanoff"},
        Powerstats: struct {
            Intelligence int `json:"intelligence"`
            Strength     int `json:"strength"`
            Speed        int `json:"speed"`
            Durability   int `json:"durability"`
            Power        int `json:"power"`
            Combat       int `json:"combat"`
        }{75,13,33,30,36,100},
        Images: struct {
            Xs string `json:"xs"`
            Sm string `json:"sm"`
            Md string `json:"md"`
            Lg string `json:"lg"`
        }{
			"https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/xs/107-black-widow.jpg",
			"https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/sm/107-black-widow.jpg",
			"https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/md/107-black-widow.jpg",
			"https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/lg/107-black-widow.jpg",
        },
    },
    "Captain America": {
        Name: "Captain America",
        Biography: struct {
            FullName string `json:"fullName"`
        }{FullName: "Steve Rogers"},
        Powerstats: struct {
            Intelligence int `json:"intelligence"`
            Strength     int `json:"strength"`
            Speed        int `json:"speed"`
            Durability   int `json:"durability"`
            Power        int `json:"power"`
            Combat       int `json:"combat"`
        }{69,19,38,55,60,100},
        Images: struct {
            Xs string `json:"xs"`
            Sm string `json:"sm"`
            Md string `json:"md"`
            Lg string `json:"lg"`
        }{
			"https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/xs/149-captain-america.jpg",
			"https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/sm/149-captain-america.jpg",
			"https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/md/149-captain-america.jpg",
			"https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/lg/149-captain-america.jpg",
        },
    },
	
}

func main() {
	http.HandleFunc("/api/superhero", superheroHandler)
	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func superheroHandler(w http.ResponseWriter, r *http.Request) {
	// Verificar que la petición sea de tipo GET
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Obtener el nombre del superhéroe de los parámetros de la URL
	heroName := r.URL.Query().Get("hero")
	superhero, ok := superheroes[heroName]
	if !ok {
		// Si el superhéroe no se encuentra, enviar un mensaje de error
		errorMsg := fmt.Sprintf("Superhero '%s' not found", heroName)
		http.Error(w, errorMsg, http.StatusNotFound)
		return
	}

	// Setear el Content-Type de la respuesta a JSON
	w.Header().Set("Content-Type", "application/json")

	// Codificar el superhéroe en JSON y escribirlo en la respuesta
	json.NewEncoder(w).Encode(superhero)
}
