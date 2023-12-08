GoPizza is an example APP that I use for implementing new technologies I'd like to learn in Go. 

## Actors and their behaviors 
- Customer
  - Is Entity
  - Order Pizza
  - Pick Up 
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
- Shop
  - Can take order 
  - Can hire/fire Employees
  - Store Pizza

