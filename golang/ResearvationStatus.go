package main

import (
	"fmt"
	"sync"
	"time"
)

// ReservationStatus represents the status of a reservation
type ReservationStatus string

const (
	Pending   ReservationStatus = "Pending"
	Confirmed ReservationStatus = "Confirmed"
	Canceled  ReservationStatus = "Canceled"
)

// Table struct represents a table in the restaurant
type Table struct {
	ID          int
	Capacity    int
	IsAvailable bool
	mutex       sync.Mutex
}

// ReserveTable reserves a table
func (t *Table) ReserveTable() bool {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	if t.IsAvailable {
		t.IsAvailable = false
		return true
	}
	return false
}

// ReleaseTable makes a table available again
func (t *Table) ReleaseTable() {
	t.mutex.Lock()
	t.IsAvailable = true
	t.mutex.Unlock()
}

// Customer struct represents a customer
type Customer struct {
	ID    int
	Name  string
	Email string
	Phone string
}

// Reservation struct
type Reservation struct {
	ID        int
	Customer  *Customer
	Table     *Table
	DateTime  time.Time
	Status    ReservationStatus
}

// Confirm marks the reservation as confirmed
func (r *Reservation) Confirm() {
	r.Status = Confirmed
	fmt.Printf("Reservation %d confirmed for %s\n", r.ID, r.Customer.Name)
}

// Cancel releases the table and cancels the reservation
func (r *Reservation) Cancel() {
	r.Status = Canceled
	r.Table.ReleaseTable()
	fmt.Printf("Reservation %d canceled for %s\n", r.ID, r.Customer.Name)
}

// Restaurant struct
type Restaurant struct {
	Name     string
	Location string
	Tables   []*Table
	Mutex    sync.Mutex
}

// AddTable adds a new table to the restaurant
func (rest *Restaurant) AddTable(capacity int) {
	rest.Mutex.Lock()
	defer rest.Mutex.Unlock()

	tableID := len(rest.Tables) + 1
	table := &Table{ID: tableID, Capacity: capacity, IsAvailable: true}
	rest.Tables = append(rest.Tables, table)
	fmt.Printf("Added Table %d with capacity %d\n", table.ID, capacity)
}

// CheckAvailability checks for an available table for a given party size
func (rest *Restaurant) CheckAvailability(size int) *Table {
	rest.Mutex.Lock()
	defer rest.Mutex.Unlock()

	for _, table := range rest.Tables {
		if table.IsAvailable && table.Capacity >= size {
			return table
		}
	}
	return nil
}

// CreateReservation attempts to reserve a table for a customer
func (rest *Restaurant) CreateReservation(customer *Customer, size int, dateTime time.Time) *Reservation {
	table := rest.CheckAvailability(size)
	if table == nil {
		fmt.Println("No available tables")
		return nil
	}

	if !table.ReserveTable() {
		fmt.Println("Table reservation failed")
		return nil
	}

	reservation := &Reservation{
		ID:       len(rest.Tables) + 1,
		Customer: customer,
		Table:    table,
		DateTime: dateTime,
		Status:   Pending,
	}
	fmt.Printf("Reservation created: %d for %s at %s\n", reservation.ID, customer.Name, dateTime)
	return reservation
}

// Main function to test the system
func main() {
	// Create a restaurant
	restaurant := &Restaurant{Name: "Gourmet House", Location: "Downtown"}
	restaurant.AddTable(2)
	restaurant.AddTable(4)
	restaurant.AddTable(6)

	// Create customers
	customer1 := &Customer{ID: 1, Name: "John Doe", Email: "john@example.com", Phone: "123-456-7890"}
	customer2 := &Customer{ID: 2, Name: "Alice Smith", Email: "alice@example.com", Phone: "987-654-3210"}

	// Create reservations
	res1 := restaurant.CreateReservation(customer1, 2, time.Now().Add(time.Hour))
	if res1 != nil {
		res1.Confirm()
	}

	res2 := restaurant.CreateReservation(customer2, 4, time.Now().Add(2*time.Hour))
	if res2 != nil {
		res2.Cancel() // Cancel reservation
	}

	// Check availability again
	res3 := restaurant.CreateReservation(customer2, 4, time.Now().Add(3*time.Hour))
	if res3 != nil {
		res3.Confirm()
	}
}

