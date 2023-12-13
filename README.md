GoPizza is an example APP that I use for implementing new technologies I'd like to learn in Go. 

## Actors and their behaviors 
- Customer
  - Is Entity
  - Order Pizza
  - Pick Up 
  - CheckBalance
- Employee
  - Is Entity
  - Can be hired or fired 
  - SubTypes
    - Cashier
      - Can Take Order
      - Can receive Customer funds
    - Driver
      - Can Deliver Pizza
    - PizzaMaker
      - Can Make Pizza
- Pizza Store
  - Can take order 
  - Can hire/fire Employees
  - Store Pizza

## API 

### Endpoints
  - Orders: Handles all orders for a target store 
  - Customer: handles all CRUD information for customers
  - Store: Handles all store related operations
