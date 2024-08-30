package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type User struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Gender  string `json:"gender"`
	Address string `json:"address"`
	Country string `json:"country"`
}

func generateRandomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func generateRandomEmail(name string) string {
	domains := []string{"example.com", "test.com", "demo.com"}
	return fmt.Sprintf("%s@%s", name, domains[rand.Intn(len(domains))])
}

func generateRandomGender() string {
	genders := []string{"Male", "Female", "Other"}
	return genders[rand.Intn(len(genders))]
}

func generateRandomAddress() string {
	streets := []string{"MG Road", "Brigade Road", "Linking Road", "Park Street", "Camac Street"}
	cities := []string{"Mumbai", "Delhi", "Bangalore", "Kolkata", "Chennai"}
	states := []string{"Maharashtra", "Delhi", "Karnataka", "West Bengal", "Tamil Nadu"}
	return fmt.Sprintf("%d %s, %s, %s, %d", rand.Intn(999), streets[rand.Intn(len(streets))], cities[rand.Intn(len(cities))], states[rand.Intn(len(states))], rand.Intn(999999))
}

func main() {
	var users []User
	for i := 0; i < 500; i++ {
		name := generateRandomString(10)
		user := User{
			ID:      generateRandomString(10),
			Name:    name,
			Email:   generateRandomEmail(name),
			Gender:  generateRandomGender(),
			Address: generateRandomAddress(),
			Country: "India",
		}
		users = append(users, user)
	}

	file, _ := json.MarshalIndent(users, "", "  ")
	_ = os.WriteFile("users.json", file, 0644)
	fmt.Println("Generated users.json with 500 dummy users")
}
