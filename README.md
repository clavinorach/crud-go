
# Simple Go Crud App

![image](https://github.com/user-attachments/assets/011b9d3a-45c4-406c-b9af-12fa89166b05)

### Authors
- [@clavinorach](https://www.linkedin.com/in/clavinorachmadi)

## Features
- **Create Employee**: Add new employee records.
- **Read Employee**: View details of employees.
- **Update Employee**: Modify existing employee records.
- **Delete Employee**: Remove employee records from the system.

## Prerequisites
- [Go](https://golang.org/dl/) (version 1.15 or higher)
- MySQL Database

## Setup and Installation: 
1. Set up the MySQL database:

   Create a database and table named employee
   ```sh
   CREATE TABLE employee (
   id INT AUTO_INCREMENT PRIMARY KEY,
   name VARCHAR(255) NOT NULL,
   npwp VARCHAR(255) NOT NULL,
   address TEXT NOT NULL
   );


### Example Usage

- **To add a new employee**: Fill out the form at `/create`.
- **To view an employee's details**: Navigate to `/employee?id=<employee_id>`.
- **To update an employee's information**: Submit the form at `/update`.
- **To delete an employee**: Go to `/delete?id=<employee_id>`.







