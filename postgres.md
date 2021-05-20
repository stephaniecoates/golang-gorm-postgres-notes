# Postgres
- A DB is a place to store, manipulate, and retrieve data, usually stored on a computer server
- What is a DBMS?
  - whenever we have a lot of data to store and create meaning from, we use a DBMS
  - types: Hierarchial DBMS, Network DBMS, Relational DBMS, Object-Oriented DBMS

- Postgres is the database engine, and SQL is the language (structured query language)
- SQL is a programming language to manage data in a relational DB, started in 1974
- data is stored in tables with columns and rows

# What does relational database mean?
- means there's a relation between two or more tables in the DB
- rather than having unstructured data and/or one huge table that's hard to maintain, have smaller chunks that connect to eachother

- Postgres is open source, alternative include Oracle (need license for it), MySQL (open source but owned by Oracle), and SQL Server (open source but owned by Microsoft)

- use can use a GUI to do everything you need to do to store, manipulate, and grab data from the DB, but it puts a veil on top of what's actually happening. Better to learn how to use the command line to access it! Plus if you ever need to SSH into a remote server and talk to the DB, you won't have a GUI then.

## Command Line Postgres


### setup
- Installed Postgres on Mac using Postgres.app, then defined PATH to psql in iTerm .zshrc
- (At work, is Postgres within a Docker container?)

### creating/connecting to the DB
- to create a database: `CREATE DATABASE [name];` (db commands must end with semicolon, otherwise they won't execute)
- to connect to a database:
```
  -h, --host=HOSTNAME      database server host or socket directory (default: "local socket")
  -p, --port=PORT          database server port (default: "5432")
  -U, --username=USERNAME  database user name (default: "stephaniecoates")
  -w, --no-password        never prompt for password
  -W, --password           force password prompt (should happen automatically)
  ```
- i.e. `psql -h localhost -p 5432 -U stephaniecoates test`

- to delete a database (be careful! no going back): `DROP DATABASE [name];`
  - DBs should always have backups in case one is accidently deleted

### to create a database table:

```sql
CREATE TABLE [table_name] ( 
  -- when you are in the terminal entering a multi line prompt, return after the first parenthesis and postgres won't run the command till you enter a semicolon
  CREATE TABLE person(
    id INT,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    gender VARCHAR(7),
    date_of_birth DATE );
```
- `\d` to display list of relations in DB
- `\d [table_name]` to display the table description (columns and types)

- setting constraints when creating a table
  - do this so when someone is creating a new row/inserting data, those inputs must satisfy certain constraints in order to be put in
    - BIGSERIAL is a signed int that autoincrements, PRIMARY KEY specifies that that column is the unique identifier for that input, NOT NULL means it can't be empty
```sql
CREATE TABLE person (
  id BIGSERIAL NOT NULL PRIMARY KEY,
  first_name VARCHAR(50) NOT NULL,
  last_name VARCHAR(50) NOT NULL,
  gender VARCHAR(7) NOT NULL,
  date_of_birth DATE NOT NULL,
  email VARCHAR(150) );
```

- to delete DB table: `DROP TABLE [table_name];`

### insert records into tables
```sql
INSERT INTO person (
  first_name,
  last_name,
  gender,
  date_of_birth)
VALUES ('Anne', 'Smith', 'FEMALE', DATE '1988-01-09');
-- INSERT 0 1
```

- to generate mock data, use mockaroo.com
- `\i [file_name] to execute a sql file, or any file within psql`

### Selecting/Reading Records
- All: `SELECT * FROM table_name;`
- Specific columns: `SELECT column_name, column_name FROM table_name;`

### Order by
- ASC = 1, 2, 3, 4, 5 || a, b, c(queries ascending by default)
- DESC = 5, 4, 3, 2, 1 || z, y, x
- `SELECT * FROM person ORDER BY country_of_birth DESC`
- can list two ORDER BY args in case one property has two of the same values for the column

### Distinct
- if you only want unique values, use DISTINCT keyword
- `SELECT DISTINCT country_of_birth FROM person ORDER BY country_of_birth;`

### WHERE clause, AND, & OR 
- `WHERE` allows us to filter data based on conditions
`SELECT * FROM person WHERE gender = 'female' AND (country_of_birth = 'Poland' OR country_of_birth = 'China');`
- data types in Postgres: [see youtube link]

### Comparison Operators
- arithmetic and comparison operators most popular, also bitwise and logic operators available
- `SELECT 1 = 1` will return true
- not equal sign: <>
- these are used to filter down your data within the WHERE clause

### Limit, Offset, & Fetch
- if you only want a specified number of records, use the LIMIT keyword
- `SELECT * FROM person LIMIT 10;`
- Limit begins from the first listed record, if you want a different starting point, use OFFSET
- `SELECT * FROM person OFFSET 5 LIMIT 10;` (will list 6 - 15)
- `SELECT * FROM person OFFSET 5` (will list 6 - end)
- FETCH can also be used to set a limit: (orig sql standard)
- `SELECT * FROM person OFFSET 20 FETCH FIRST 5 ROW ONLY;`
- `SELECT * FROM person OFFSET 20 FETCH FIRST ROW ONLY;` 

### IN keyword
- IN takes an array and allows you to pass in multiple values to filter by
- `SELECT * FROM person WHERE country_of_birth IN ('China', 'Brazil', 'Portugal') AND gender = 'Female' ORDER BY country_of_birth;`

### BETWEEN keyword
- if we want to select results between two values
- `SELECT * FROM person WHERE date_of_birth BETWEEN DATE '2000-01-01' and '2015-01-01' ORDER BY date_of_birth;`

### Like and iLike keyword
- used to match text against a patterns using a wildcard
- `SELECT * FROM person WHERE email LIKE '%@google.%';`
- `%` means any chars
- `____` (4 underscores) specifies 4 of any characters
- `SELECT * FROM person WHERE email LIKE '%_______@%.%';`
- ILIKE keyword is non-case-sensitive
- `SELECT * FROM person where country_of_birth ILIKE 'p%';`

### GROUP BY keyword
- `SELECT DISTINCT country_of_birth FROM person;` will return a list of countries, but what if we want the # of people from each country?
- `SELECT country_of_birth, COUNT(*) FROM person GROUP BY country_of_birth;`
  - 2rd arg (the built in func COUNT) will count the number of entries per group of the first arg
- `SELECT country_of_birth, COUNT(*) FROM person GROUP BY country_of_birth ORDER BY country_of_birth;`
- (*) just means count all the rows

### GROUP BY: HAVING
- when HAVING goes right after GROUP BY, you can filter the groupings that appear
- i.e if you only wanted to show countries that had at least 5 people:
- `SELECT country_of_birth, COUNT(*) FROM person GROUP BY country_of_birth HAVING COUNT(*) >= 5 ORDER BY country_of_birth;`
- other aggregate functions (other than `count()`) found here: https://www.postgresql.org/docs/9.5/functions-aggregate.html

### Aggregate Functions: MIN, MAX, AVG, SUM
- To find the most expensive car price in a table of cars: 
- `SELECT MAX(price) FROM car;`
- To get lowest, `SELECT MIN(price) FROM car;`
- To get average, `SELECT AVG(price) FROM car;` <!-- 54696.267550000000 -->
- To round average, `SELECT ROUND(AVG(price)) FROM car;` <!-- 54696 -->

- Grouping categories:
  - to get the minimum price per make & model: `SELECT make, model, MIN(price) FROM car GROUP BY make, model;`
  - to get average price per make: `SELECT make, ROUND(AVG(price)) FROM car GROUP BY make;`

- SUM allows us to perform addition over our dataset
- To get the sum of all the car prices:
  - `SELECT SUM(price) FROM car;`
- To get the sum grouped by car make:
  - `SELECT make, SUM(price) FROM car GROUP BY make;`

### Arithmetic Operators
- we can use our dataset to produce statistics or results
- Using basic math within Postgres:
 - `SELECT 10 + 2;`  <!-- 12 -->
 - `SELECT 10 * 5 - 8;`  <!-- 42 -->
 - `SELECT 10^3;` (exponent) <!-- 1000 -->
 - `SELECT 5!;` (factorial)  <!-- 120 -->
 - `SELECT 10 % 3;` (modulus)  <!-- 1 -->

- If we want to take 10% off of all of our car prices for a special promo, and within the table list the original price, the total discount, and the discounted price:
  - `SELECT *, ROUND(price * .10, 2), ROUND(price - price * .10, 2) FROM car;`

### Alias
- when you perform math to create new columns, Postgres will use operator name as column default (or if no operator is used, it'll show up as `?column?`)
- to customize new column names (or override existing column names), follow column math with `AS` keyword
- `SELECT id, make, model, price AS original_price, ROUND(price * .10, 2) AS discount, ROUND(price - price * .10, 2) AS discounted_price FROM car;`

## Handling NULLs in Postgres

### Coalesce

- Allows us to have a default 