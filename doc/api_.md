#/api

## Insert new Prayer

Insert a new Prayer into the Database
prayer contains a text, and a public flag

### **URL**

 * **`/api`**

### **Method:**

 * **`POST`** 
  
### **Data Params**
 *  **Content:** `{
                    "prayer": <TEXT>,
                    "is_public": <BOOL>
                  }`

### **Success Response:**
 *  **Code:** 200 <br />
 
### **Error Response:**
  * **Code:** 400 INVALID REQUEST <br /> 
    * **Reason** invalid json format
* **Code:** 500 INTERNAL SERVER ERROR<br /> 
  * **Reason** database error


## Get Prayers

Get all prayers of today from the database

### **URL**

 * **`/api`**

### **Method:**

 * **`GET`** 
  

### **URL Params**
  * **Required:**
   `key=[string]` (key authentication)
   
   
### **Success Response:**
 *  **Code:** 200 <br />
     * **Content:** 
     array of all prayers of today
     ```json
     [
      { 
        "id" : "<int>", 
        "prayer" : "<text>",
        "is_public" : "<bool>" , 
        "approved" : "<bool>", 
        "date" : "<datetime>" 
      },
       ...
     ]
    ```

### **Error Response:**
* **Code:** 401 UNAUTHORIZED <br />
* **Code:** 500 INTERNAL SERVER ERROR<br /> 
  * **Reason** database error

## Change Prayer

Insert a new Prayer into the Database
prayer contains a text, and a public flag

### **URL**

 * **`/api`**

### **Method:**

 * **`PUT`** 
 
### **URL Params**
  * **Required:**
   `key=[string]` (key authentication)

### **Data Params**
 *  **Content:** `{
                    "id": <int>,
                    "prayer" : <text>,
                    "is_public": <BOOL>
                    "approved": <BOOL>
                  }`

### **Success Response:**
 *  **Code:** 200 <br />
 
### **Error Response:**
  * **Code:** 401 UNAUTHORIZED <br />
  * **Code:** 400 INVALID REQUEST <br /> 
    * **Reason:** invalid json format
  * **Code:** 500 INTERNAL SERVER ERROR<br /> 
    * **Reason:** database error

