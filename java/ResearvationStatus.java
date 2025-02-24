import java.util.*;
import java.util.concurrent.locks.Lock;
import java.util.concurrent.locks.ReentrantLock;

// Reservation Status
enum ReservationStatus {
    PENDING, CONFIRMED, CANCELED
}

// Table class
class Table {
    private int tableId;
    private int capacity;
    private boolean isAvailable;
    private Lock lock;

    public Table(int id, int capacity) {
        this.tableId = id;
        this.capacity = capacity;
        this.isAvailable = true;
        this.lock = new ReentrantLock();
    }

    public boolean reserveTable() {
        lock.lock();
        try {
            if (isAvailable) {
                isAvailable = false;
                return true;
            }
            return false;
        } finally {
            lock.unlock();
        }
    }

    public void releaseTable() {
        lock.lock();
        try {
            isAvailable = true;
        } finally {
            lock.unlock();
        }
    }

    public int getCapacity() {
        return capacity;
    }

    public boolean isAvailable() {
        return isAvailable;
    }
}

// Customer class
class Customer {
    int id;
    String name;
    String email;
    String phone;

    public Customer(int id, String name, String email, String phone) {
        this.id = id;
        this.name = name;
        this.email = email;
        this.phone = phone;
    }
}

// Reservation class
class Reservation {
    private int reservationId;
    private Customer customer;
    private Table table;
    private Date dateTime;
    private ReservationStatus status;

    public Reservation(int id, Customer customer, Table table, Date dateTime) {
        this.reservationId = id;
        this.customer = customer;
        this.table = table;
        this.dateTime = dateTime;
        this.status = ReservationStatus.PENDING;
    }

    public void confirm() {
        status = ReservationStatus.CONFIRMED;
        System.out.println("Reservation " + reservationId + " confirmed for " + customer.name);
    }

    public void cancel() {
        status = ReservationStatus.CANCELED;
        table.releaseTable();
        System.out.println("Reservation " + reservationId + " canceled for " + customer.name);
    }
}

// Restaurant class
class Restaurant {
    private String name;
    private String location;
    private List<Table> tables;
    private List<Reservation> reservations;

    public Restaurant(String name, String location) {
        this.name = name;
        this.location = location;
        this.tables = new ArrayList<>();
        this.reservations = new ArrayList<>();
    }

    public void addTable(int capacity) {
        tables.add(new Table(tables.size() + 1, capacity));
        System.out.println("Added Table " + tables.size() + " with capacity " + capacity);
    }

    public Table checkAvailability(int size) {
        for (Table table : tables) {
            if (table.isAvailable() && table.getCapacity() >= size) {
                return table;
            }
        }
        return null;
    }

    public Reservation createReservation(Customer customer, int size, Date dateTime) {
        Table table = checkAvailability(size);
        if (table == null || !table.reserveTable()) {
            System.out.println("No available tables");
            return null;
        }

        Reservation reservation = new Reservation(reservations.size() + 1, customer, table, dateTime);
        reservations.add(reservation);
        System.out.println("Reservation created: " + reservation.reservationId + " for " + customer.name);
        return reservation;
    }
}

// Main function
public class RestaurantReservationSystem {
    public static void main(String[] args) {
        Restaurant restaurant = new Restaurant("Gourmet House", "Downtown");
        restaurant.addTable(2);
        restaurant.addTable(4);
        restaurant.addTable(6);

        Customer customer1 = new Customer(1, "John Doe", "john@example.com", "123-456-7890");
        Customer customer2 = new Customer(2, "Alice Smith", "alice@example.com", "987-654-3210");

        Reservation res1 = restaurant.createReservation(customer1, 2, new Date());
        if (res1 != null) res1.confirm();

        Reservation res2 = restaurant.createReservation(customer2, 4, new Date());
        if (res2 != null) res2.cancel();
    }
}
