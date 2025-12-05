# Parking Lot CLI Application

## Overview

This application simulates a parking lot system that can hold up to `n` cars. Each parking slot is numbered starting from 1, increasing with distance from the entry point. The system automates ticketing, slot allocation, and fee calculation based on parking duration. Users interact with the system via a simple command-line interface (CLI).

* When a car enters, its **registration number** and **color** are recorded.
* The car is assigned the **nearest available slot**.
* On exit, the car leaves a **slot number** and **hours parked**, and the system calculates the **parking charge**:

    * $10 for the first 2 hours
    * $10 for each additional hour

The system also allows checking the **status** of the parking lot.

---

## Commands

Interactive CLI prompts the user with three options:

1. **Park a car**
   Input: `PLATE_NUMBER COLOR`
   Allocates the nearest available slot. Rejects duplicates.

2. **Leave a car**
   Input: `SLOT_NUMBER HOURS_PARKED`
   Frees the slot and calculates the parking charge.

3. **Status**
   Displays all slots with their registration numbers and colors, or marks them as free.

---

## Example Interactive Session

```
capacity:
6
1 park - Input 'PLATE NUMBER(space)COLOR'
2 leave - Input 'SLOT NUMBER(space)HOURS PARKED'
3 status
1
KA-01-HH-1234 Red
slot allocated at slot: 1
1 park - Input 'PLATE NUMBER(space)COLOR'
2 leave - Input 'SLOT NUMBER(space)HOURS PARKED'
3 status
1
KA-01-HH-1234 White
Car KA-01-HH-1234 is already parked at slot 1
1 park - Input 'PLATE NUMBER(space)COLOR'
2 leave - Input 'SLOT NUMBER(space)HOURS PARKED'
3 status
1
KA-01-HH-12346 Black
slot allocated at slot: 2
1 park - Input 'PLATE NUMBER(space)COLOR'
2 leave - Input 'SLOT NUMBER(space)HOURS PARKED'
3 status
1
KA-01-HH-12347 PINK
slot allocated at slot: 3
1 park - Input 'PLATE NUMBER(space)COLOR'
2 leave - Input 'SLOT NUMBER(space)HOURS PARKED'
3 status
1
KA-01-HH-12348 Gray
slot allocated at slot: 4
1 park - Input 'PLATE NUMBER(space)COLOR'
2 leave - Input 'SLOT NUMBER(space)HOURS PARKED'
3 status
3
Slot Registration Parking:
No              Plate Number            Color
1.              KA-01-HH-1234           Red
2.              KA-01-HH-12346          Black
3.              KA-01-HH-12347          PINK
4.              KA-01-HH-12348          Gray
5.              Free Slot
6.              Free Slot
1 park - Input 'PLATE NUMBER(space)COLOR'
2 leave - Input 'SLOT NUMBER(space)HOURS PARKED'
3 status
1
KA-01-HH-12349 Navy
slot allocated at slot: 5
1 park - Input 'PLATE NUMBER(space)COLOR'
2 leave - Input 'SLOT NUMBER(space)HOURS PARKED'
3 status
1
KA-01-HH-12340 pink
slot allocated at slot: 6
1 park - Input 'PLATE NUMBER(space)COLOR'
2 leave - Input 'SLOT NUMBER(space)HOURS PARKED'
3 status
3
Slot Registration Parking:
No              Plate Number            Color
1.              KA-01-HH-1234           Red
2.              KA-01-HH-12346          Black
3.              KA-01-HH-12347          PINK
4.              KA-01-HH-12348          Gray
5.              KA-01-HH-12349          Navy
6.              KA-01-HH-12340          pink
1 park - Input 'PLATE NUMBER(space)COLOR'
2 leave - Input 'SLOT NUMBER(space)HOURS PARKED'
3 status
1
KA-01-HH-123411 Gray
parking not available
1 park - Input 'PLATE NUMBER(space)COLOR'
2 leave - Input 'SLOT NUMBER(space)HOURS PARKED'
3 status
3
Slot Registration Parking:
No              Plate Number            Color
1.              KA-01-HH-1234           Red
2.              KA-01-HH-12346          Black
3.              KA-01-HH-12347          PINK
4.              KA-01-HH-12348          Gray
5.              KA-01-HH-12349          Navy
6.              KA-01-HH-12340          pink
1 park - Input 'PLATE NUMBER(space)COLOR'
2 leave - Input 'SLOT NUMBER(space)HOURS PARKED'
3 status
2
3 8
Registration number KA-01-HH-12347 with slot number 3 is free with Charge $70
1 park - Input 'PLATE NUMBER(space)COLOR'
2 leave - Input 'SLOT NUMBER(space)HOURS PARKED'
3 status
2
6 4
Registration number KA-01-HH-12340 with slot number 6 is free with Charge $30
1 park - Input 'PLATE NUMBER(space)COLOR'
3 status
Slot Registration Parking:
No              Plate Number            Color
1.              KA-01-HH-1234           Red
2.              KA-01-HH-12346          Black
3.              Free Slot
4.              KA-01-HH-12348          Gray
5.              KA-01-HH-12349          Navy
6.              Free Slot
```

---

## Notes

* The system allocates **slots in ascending order**.
* Duplicate registration numbers are **rejected**.
* Parking charges are calculated automatically on leaving.
* The interactive CLI is designed for **sequential commands** only.
