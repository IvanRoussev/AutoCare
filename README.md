# AutoCare
AutoCare, An app to track your cars maintenance records. Keep all of your cars records in line, organized and never miss a maintenance again.
## API Reference

### Create a User

- **Use-Case:** Create yourself a user with this route

```http
  POST /users
```

| JSON Body | Type     | Requirements                | Required|
| :-------- | :------- | :------------------------- | :------|
| `full_name` | `string` | First and last name| **Yes** |
| `password` | `string` | **Min 6 Characters**|**Yes** |
| `email` | `string` | **Must be valid email format**|**Yes** |
| `country` | `string` | Your Country of residence | **Yes** |


### Login User

- **Use-Case:** Log your use into the app to receive bearer token


```http
  POST /users/login
```

| JSON Body | Type     | Requirements                | Required|
| :-------- | :------- | :------------------------- | :------|
| `username` | `string` | Input your **username**| **Yes** |
| `password` | `string` | Input your **password**|**Yes** |

# Cars
This endpoint allows you to **Create**, **View** and **Delete** you cars

### Create Car

```http
  POST /cars
```

- **Use-Case:** Create your car to log maintenances on it

#### HTTP Header

| Key | Type     | Value                       | Required|
| :-------- | :------- | :-------------------------------- | :---|
| `Authorization`      | `Bearer` | Your Bearer Token| **Yes** |

#### Request Body

| JSON Body | Type     | Requirements                | Required|
| :-------- | :------- | :-------------------------------- | :---|
| `vin` | `string` | Input your cars **VIN** number| **Yes** | 
| `username` | `string` | Input your **username**| **Yes** | 
| `model` | `string` | Input car **Model**|**Yes** |
| `make` | `string` | Input car **Make**|**Yes** |
| `year` | `int` | Input car **Year**|**Yes** |


### Get Car

```http
  POST /cars/vin/:vin
```
**Attach** vin number of car wanting to view in path

- **Use-Case:** View your car information

#### HTTP Header

| Key | Type     | Value                       | Required|
| :-------- | :------- | :-------------------------------- | :---|
| `Authorization`      | `Bearer` | Your Bearer Token| **Yes** |


### Delete Car

```http
  DELETE /cars/vin/:vin
```
**Attach** vin number of car wanting to delete in path

- **Use-Case:** Delete your car

#### HTTP Header

| Key | Type     | Value                       | Required|
| :-------- | :------- | :-------------------------------- | :---|
| `Authorization`      | `Bearer` | Your Bearer Token| **Yes** |

### Delete Car

```http
  GET /cars/users/:username
```
**Attach** username of car owner to view cars belonging to this user

- **Use-Case:** View all of your cars

#### HTTP Header

| Key | Type     | Value                       | Required|
| :-------- | :------- | :-------------------------------- | :---|
| `Authorization`      | `Bearer` | Your Bearer Token| **Yes** |

#### Query Paramters

| Key     | Value                       | Required|
| :-------- | :-------------------------------- | :---|
| `page_id`      | the first page you want to start to view at| **Yes** |
| `page_size`      | the number of cars per page you want to see| **Yes** |


## Maintenances

### Create Maintenance

This endpoint allows you to **Create** and **View** the maintenances you've done on your car

```http
  POST /maintenances
```

- **Use-Case:** Create log of maintenance done to your car
#### HTTP Header

| Key | Type     | Value                       | Required|
| :-------- | :------- | :-------------------------------- | :---|
| `Authorization`      | `Bearer` | Your Bearer Token| **Yes** |

#### Request Body

| JSON Body | Type     | Requirements                | Required|
| :-------- | :------- | :-------------------------------- | :---|
| `car_vin` | `string` | Input your cars **VIN** number| **Yes** | 
| `maintenance_type` | `string` | Input the type of maintenance done to your car. eg. **Oil Change**| **Yes** | 
| `mileage` | `int` | Input car mileage when maintenance was done|**Yes** |

### Get Maintenances
Get **all** maintenance for given car by vin

```http
  GET /maintenances/:vin
```
**Attach** username of car owner to view cars belonging to this user


- **Use-Case:** Create log of maintenance done to your car
#### HTTP Header


| Key | Type     | Value                       | Required|
| :-------- | :------- | :-------------------------------- | :---|
| `Authorization`      | `Bearer` | Your Bearer Token| **Yes** |

#### Query Paramters

| Key     | Value                       | Required|
| :-------- | :-------------------------------- | :---|
| `page_id`      | the first page you want to start to view at| **Yes** |
| `page_size`      | the number of cars per page you want to see| **Yes** |