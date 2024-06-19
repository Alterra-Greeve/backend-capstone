<p align="center">
<img src="https://github.com/Alterra-Greeve/.github/assets/133726246/3a58ead2-7977-4f31-8f29-bb54e55dc34b" width="350" />
</p>

<p align="center">
  <b>Greeve</b><br>
  Final Project Capstone Alterra Academy<br>
  <a href="https://api.greeve.store/">Explore the docs Open API</a> 
  <br><br>
</p>

# About Project
<p align="justify">
  Greeve is an innovative app that aims to increase user awareness and participation in environmental conservation and make it easier to purchase eco-friendly goods and equipment. The app not only offers a platform to purchase environmental goods and equipment, but also provides information on the impact of users' activities on the environment as well as how to measure and reduce that impact. The following is greeve documentation from the backend using golang programming language.
</p> 

### Features available in this project:
| Feature           | Method              | Users                   |          Admin                           |
| ----------------- | ------------------- | ----------------------- | ---------------------------------------- |
| USERS FEATURE                                                                                                |
| Register          | POST                | Yes                     | No                                       |
| Login             | POST                | Yes                     | Yes                                      |
| Reset Password    | POST                | Yes                     | No                                       |
| Get All Users     | GET                 | No                      | Yes                                      |
| Edit Self Profile | PUT                 | Yes                     | Yes                                      |
| Manage User       | PUT                 | No                      | Yes                                      |
| Delete User       | DELETE              | Yes                     | Yes                                      |
| Product  FEATURE                                                                                             |
| Create Data Product | POST                | No                      | Yes                                      |
| Edit Data Product   | PUT                 | No                      | Yes                                      |
| Create Data Product | POST                | No                      | Yes                                      |
| Get Product         | GET                 | Yes                     | Yes                                      |
| Get Product by ID   | GET                 | Yes                     | Yes                                      |
| View Item from Shooping cart       | GET                 | Yes                     | No                                      |
| Add Item to Shooping cart          | POST                | Yes                     | No                                      |
| Edit Item From Shooping cart       | PUT                 | Yes                     | No                                      |
| Delete Item From Shooping cart     | DELETE              | Yes                     | No                                      |
| Transaction  FEATURE                                                                                             |
| Get All Transaction Data For User  | GET                 | Yes                     | Yes                                     |
| Get Transaction Data For User by ID| GET                 | Yes                     | Yes                                     |
| Create Transaction                 | POST                | Yes                     | No                                      |
| Cancel Transaction                 | PUT                 | Yes                     | No                                      |
| Payment For Transaction            | POST                | Yes                     | No                                      |
| Challenge Feature                                                                                                            |
| View All Challenge                 | GET                 | Yes                     | Yes                                     |
| Get Challenge by ID                | GET                 | Yes                     | Yes                                     |
| Get Challenge that has been joined                | GET                 | Yes                     | No                                  |
| Parcitipate Challenge              | POST                | Yes                     | No                                      |
| Confirmation Join Challenge        | POST                | Yes                     | No                                      |
| Viev Monthly Leaderboard           | GET                 | Yes                     | Yes                                     |
| Forum Feature                                                                                                                |
| Get All Forum        | GET                | Yes                     | Yes                                      |
| Get Forum by Id      | GET                | Yes                     | Yes                                      |
| Create Forum         | POST               | Yes                     | No                                       |
| Edit Forum           | PUT                | Yes                     | No                                       |
| Delete Forum         | DELETE             | Yes                     | Yes                                      |
| Create Message       | POST               | Yes                     | No                                       |
| Edit Message         | POST               | Yes                     | No                                       |
| Get Message by Id    | GET                | No                      | Yes                                      |
| Delete Message       | DELETE             | Yes                     | Yes                                      |
| Impact Category Feature                                                                                                |
| Get All Impact Category       | GET                | No                     | Yes                                      |
| Create Impact Category        | POST               | No                     | Yes                                      |
| Edit Impact Category          | PUT                | No                     | Yes                                      |
| Delete Impact Category        | DELETE             | No                     | Yes                                      |
| Voucher Feature                                                                                                  |
| Get All Voucher       | GET                  | Yes                    | Yes                                      |
| Create Voucher        | POST                 | No                     | Yes                                      |
| Edit Voucher          | PUT                  | No                     | Yes                                      |
| Delete Voucher        | DELETE               | No                     | Yes                                      |
| Dashboard Feature                                                                                                    |
| Get Dashboard              | GET                  | No                    | Yes                                      |
| Get Orders Product         | GET                  | No                    | Yes                                      |
| Get Orders Challenge       | GET                  | No                    | Yes                                      |

# Built With
- <a href="https://github.com/golang">Golang</a>
- <a href="https://github.com/labstack/echo">Echo</a>
- <a href="https://github.com/go-gorm/gorm">Gorm</a>
- <a href="https://github.com/mysql">MySql</a>
# Getting Started
## How to Use
Clone the project

```bash
  git clone https://github.com/Alterra-Greeve/backend-capstone.git
```
Install required dependencies
```bash
  go mod tidy
```
Setting Environmental Variable on `.env` file
```bash
  DB_HOST=
  DB_PORT=
  DB_USER=
  DB_PASS=
  DB_NAME=
  
  SMTP_USER=
  SMTP_PASS=
  SMTP_HOST=
  SMTP_PORT=
  
  JWT_SECRET=
  
  PROJECT_ID=
  BUCKET_NAME=
  GOOGLE_APPLICATION_CREDENTIALS=
  
  MIDTRANS_CLIENT_KEY=
  MIDTRANS_SERVER_KEY=
```
Run `main.go` file
```bash
go run main.go
```
Run the endpoints as in the swagger documentation

Docker & GCP
```bash
```


# System Diagram and Architecture
## ERD
<div style="border: 1px solid #000; padding: 10px; display: inline-block;">
  <img src="https://github.com/Alterra-Greeve/backend-capstone/assets/133726246/c6ddfcd2-2c8c-4d1e-a7c8-2587ec42adf6" alt="ERD_Capstone drawio">
</div>

## HLA

# Contact 
Link Project      : https://github.com/Alterra-Greeve
API Documentation : https://api.greeve.store/#/

## Contributor
**Nur Faid Praseto** 
**Chandra Wahyu Rafialdi**
