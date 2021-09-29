package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Product struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Price       string `json:"price"`
	Quantity    string `json:"quantity"`
	Description string `json:"desc"`
}

var productsArr []Product

func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(productsArr)
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range productsArr {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode("Sorry, item not found")

}

func createProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newProduct Product
	_ = json.NewDecoder(r.Body).Decode(&newProduct)

	// Check whether the product exists
	// If exists than increse the count else
	// add it in productArr.
	// This can be further optimized..

	for _, item := range productsArr {
		if item.Name == newProduct.Name {
			//count1, _ := strconv.ParseInt(item.Quantity, 10, 0)
			//count2, _ := strconv.ParseInt(newProduct.Quantity, 10, 0)
			//
			//item.Quantity = strconv.Itoa(int(count1) + int(count2))
			//if item.Price != newProduct.Price {
			//	item.Price = newProduct.Price
			//}
			//if item.Description != newProduct.Description {
			//	item.Description = newProduct.Description
			//}
			//json.NewEncoder(w).Encode(item)
			//productsArr[i] = item
			//json.NewEncoder(w).Encode(productsArr)
			json.NewEncoder(w).Encode("Product already exists.")

			return
		}
	}

	productarrSize := len(productsArr) + 1
	newProduct.ID = strconv.Itoa(productarrSize)
	productsArr = append(productsArr, newProduct)
	json.NewEncoder(w).Encode(newProduct)
	json.NewEncoder(w).Encode(productsArr)

}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, item := range productsArr {
		if item.Name == params["name"] {
			var newProduct Product
			_ = json.NewDecoder(r.Body).Decode(&newProduct)
			newProduct.ID = strconv.Itoa(i + 1)
			if item.Price != newProduct.Price {
				item.Price = newProduct.Price
			}
			if item.Description != newProduct.Description {
				item.Description = newProduct.Description
			}
			productsArr[i] = newProduct
			json.NewEncoder(w).Encode(productsArr)
			return
		}
	}
	var newProduct Product
	_ = json.NewDecoder(r.Body).Decode(&newProduct)
	productarrSize := len(productsArr) + 1
	newProduct.ID = strconv.Itoa(productarrSize)
	productsArr = append(productsArr, newProduct)
	json.NewEncoder(w).Encode(newProduct)
	json.NewEncoder(w).Encode(productsArr)

}

func buyProduct(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, item := range productsArr {
		if item.Name == params["name"] {
			count, _ := strconv.ParseInt(item.Quantity, 10, 0)
			if count == 0 {
				json.NewEncoder(w).Encode("Sorry, item is out of stock")
				return
			}
			json.NewEncoder(w).Encode("Congratulations, your purchase was successful")
			count--

			productsArr[i].Quantity = strconv.Itoa(int(count))
			json.NewEncoder(w).Encode(productsArr)
			return
		}
	}

	json.NewEncoder(w).Encode("Sorry, item is not available")

}

func main() {
	r := mux.NewRouter()

	productsArr = append(productsArr, Product{ID: "1", Name: "Laptop", Price: "50000", Quantity: "5", Description: "Multi Purpose device"})
	productsArr = append(productsArr, Product{ID: "2", Name: "Mobile", Price: "20000", Quantity: "10", Description: "Multi Purpose device LOL"})
	productsArr = append(productsArr, Product{ID: "3", Name: "Perfumes", Price: "500", Quantity: "50", Description: "Fragrance"})
	productsArr = append(productsArr, Product{ID: "4", Name: "Chair", Price: "1500", Quantity: "30", Description: "Furniture"})

	r.HandleFunc("/api/products", getProducts).Methods("GET")
	r.HandleFunc("/api/products/{id}", getProduct).Methods("GET")
	r.HandleFunc("/api/products", createProduct).Methods("POST")
	r.HandleFunc("/api/products/{name}", updateProduct).Methods("PUT")
	r.HandleFunc("/api/products/purchase/{name}", buyProduct).Methods("PUT")

	// r.HandleFunc("/api/products/{id}",deleteProduct).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))

	// for i := 0; i < 10; i++ {
	// 	fmt.Println("Hello World")
	// }

}
