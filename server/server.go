package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"server/dbconfig"
	"strconv"
	"sync"

	_ "github.com/lib/pq"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)


func main(){

	log.Print("Setting up Server...")

	router := chi.NewRouter()
	psql := dbconfig.NewConfig()

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        psql.DBHost, psql.DBPort, psql.DBUser, psql.DBPassword, psql.DBName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Error setting up database: %v", err)
	}
	defer db.Close()
	err = db.Ping()
    if err != nil {
        log.Fatalf("Error connecting to database: %v", err)
    }
	log.Print("Connected to database successfully...")

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
	}))

	handlerWrapper := HandlerWrapper{db : db, dbMutex: &sync.Mutex{}}
	
	router.Get("/items", handlerWrapper.GetAllHandler)
	router.Post("/items", handlerWrapper.CreateHandler)
	router.Put("/items/{id}", handlerWrapper.UpdateHandler)
	router.Delete("/items/{id}", handlerWrapper.DeleteHandler)

	log.Print("Listening on Port 5000...")
	http.ListenAndServe(":5000", router)

}

type HandlerWrapper struct {
	db *sql.DB
	dbMutex *sync.Mutex
}

type Item struct {
	Item_id int `json:"item_id"`
	Description string `json:"description"`
}

func (h *HandlerWrapper) GetAllHandler(w http.ResponseWriter, r *http.Request){

	h.dbMutex.Lock()
	defer h.dbMutex.Unlock()
	rows, err := h.db.Query("SELECT * FROM list ORDER BY item_id")
	if err != nil {
		log.Printf("Failed to query database: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	items := make([]Item,0)

	for rows.Next() {
		var item_id int
		var description string

		err := rows.Scan(&item_id, &description)
		if(err != nil) {
			log.Printf("Failed to read rows: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		items = append(items, Item{Item_id: item_id, Description: description})
	}

	itemJson, err := json.Marshal(items)
	if(err != nil){
		log.Printf("Failed to encode Item struct into JSON format: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(itemJson)
}

func (h *HandlerWrapper) CreateHandler(w http.ResponseWriter, r *http.Request){

	itemJson, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Failed to read request body: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var item Item

	err = json.Unmarshal(itemJson, &item)
	if err != nil {
		log.Printf("Failed to parse request body: %v", err)
	}
	h.dbMutex.Lock()
	defer h.dbMutex.Unlock()
	_, err = h.db.Exec("INSERT INTO list (description) VALUES ($1)", item.Description)
	if err != nil {
		log.Printf("Failed to create item in database: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *HandlerWrapper) UpdateHandler(w http.ResponseWriter, r *http.Request){

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Failed to read url parameter: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	item := Item{Item_id: id}
	itemJson, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Failed to read request body: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(itemJson, &item)

	if err != nil {
		log.Printf("Failed to parse request body: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	h.dbMutex.Lock()
	defer h.dbMutex.Unlock()
	_, err = h.db.Exec("UPDATE list SET description = $1 WHERE item_id = $2", item.Description, item.Item_id)

	if err != nil {
		log.Printf("Failed to update item in database: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *HandlerWrapper) DeleteHandler(w http.ResponseWriter, r *http.Request){
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Failed to read url parameter: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	item := Item{Item_id: id}

	h.dbMutex.Lock()
	defer h.dbMutex.Unlock()

	_, err = h.db.Exec("DELETE FROM list WHERE item_id = $1", item.Item_id)

	if err != nil {
		log.Printf("Failed to delete item from database: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}