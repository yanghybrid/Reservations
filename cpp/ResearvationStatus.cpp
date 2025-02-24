#include <iostream>
#include <vector>
#include <map>
#include <mutex>
#include <ctime>

using namespace std;

// Reservation status
enum class ReservationStatus { PENDING, CONFIRMED, CANCELED };

// Table class
class Table {
private:
    int tableId;
    int capacity;
    bool isAvailable;
    mutex tableMutex;

public:
    Table(int id, int cap) : tableId(id), capacity(cap), isAvailable(true) {}

    int getId() { return tableId; }
    int getCapacity() { return capacity; }
    bool checkAvailability() { return isAvailable; }

    bool reserveTable() {
        lock_guard<mutex> lock(tableMutex);
        if (isAvailable) {
            isAvailable = false;
            return true;
        }
        return false;
    }

    void releaseTable() {
        lock_guard<mutex> lock(tableMutex);
        isAvailable = true;
    }
};

// Customer class
class Customer {
public:
    int id;
    string name;
    string email;
    string phone;

    Customer(int id, string name, string email, string phone) 
        : id(id), name(name), email(email), phone(phone) {}
};

// Reservation class
class Reservation {
public:
    int reservationId;
    Customer* customer;
    Table* table;
    time_t dateTime;
    ReservationStatus status;

    Reservation(int id, Customer* cust, Table* tbl, time_t dt) 
        : reservationId(id), customer(cust), table(tbl), dateTime(dt), status(ReservationStatus::PENDING) {}

    void confirm() {
        status = ReservationStatus::CONFIRMED;
        cout << "Reservation " << reservationId << " confirmed for " << customer->name << endl;
    }

    void cancel() {
        status = ReservationStatus::CANCELED;
        table->releaseTable();
        cout << "Reservation " << reservationId << " canceled for " << customer->name << endl;
    }
};

// Restaurant class
class Restaurant {
private:
    string name;
    string location;
    vector<Table*> tables;
    vector<Reservation*> reservations;
    mutex restaurantMutex;

public:
    Restaurant(string name, string loc) : name(name), location(loc) {}

    void addTable(int capacity) {
        lock_guard<mutex> lock(restaurantMutex);
        tables.push_back(new Table(tables.size() + 1, capacity));
        cout << "Added Table " << tables.size() << " with capacity " << capacity << endl;
    }

    Table* checkAvailability(int size) {
        lock_guard<mutex> lock(restaurantMutex);
        for (Table* table : tables) {
            if (table->checkAvailability() && table->getCapacity() >= size) {
                return table;
            }
        }
        return nullptr;
    }

    Reservation* createReservation(Customer* customer, int size, time_t dateTime) {
        Table* table = checkAvailability(size);
        if (!table || !table->reserveTable()) {
            cout << "No available tables\n";
            return nullptr;
        }

        Reservation* reservation = new Reservation(reservations.size() + 1, customer, table, dateTime);
        reservations.push_back(reservation);
        cout << "Reservation created: " << reservation->reservationId << " for " << customer->name << endl;
        return reservation;
    }
};

// Main function
int main() {
    Restaurant restaurant("Gourmet House", "Downtown");
    restaurant.addTable(2);
    restaurant.addTable(4);
    restaurant.addTable(6);

    Customer customer1(1, "John Doe", "john@example.com", "123-456-7890");
    Customer customer2(2, "Alice Smith", "alice@example.com", "987-654-3210");

    time_t now = time(0);
    Reservation* res1 = restaurant.createReservation(&customer1, 2, now);
    if (res1) res1->confirm();

    Reservation* res2 = restaurant.createReservation(&customer2, 4, now);
    if (res2) res2->cancel();

    return 0;
}
