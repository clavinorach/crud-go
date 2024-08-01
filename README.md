**Simple Go Crud App**
 
The Employee Management System is a simple web application developed using Go (Golang) that allows users to manage employee records. The application supports CRUD (Create, Read, Update, Delete) operations for employee data, including features for creating, reading, updating, and deleting employee records.
![image](https://github.com/user-attachments/assets/52c191ff-8404-4fd4-97e9-8ecf706071e2)


## Features

- **Create Employee**: Add new employee records.
- **Read Employee**: View details of employees.
- **Update Employee**: Modify existing employee records.
- **Delete Employee**: Remove employee records from the system.



### Prerequisites

- [Go](https://golang.org/dl/) (version 1.15 or higher)
- MySQL Database
- Git (optional for cloning the repository)

## Setup and Installation: 
 ```sh
CREATE TABLE employee (
  id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  npwp VARCHAR(255) NOT NULL,
  address TEXT NOT NULL
);
