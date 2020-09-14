#/api/view

## Get Prayers

Get all approved and filtered prayers of today from the database.
The prayers json objects are shorten to only contain text and id

Opt.: request only Prayers with an id bigger than a given id.

### **URL**
 * **`/api/view`**

### **Method:**
 * **`GET`** 
  
### **URL Params**
  * **Required:**
   `key=[string]` (key authentication)
  * **Optional:**
   `l=[id]` (id to filter only bigger values)
   
### **Success Response:**
 *  **Code:** 200 <br />
     * **Content:** 
     array of approved and filtered prayers of today
     ```json
     [
      { 
        "id" : "<int>", 
        "prayer" : "<text>"
      },
       ...
     ]
    ```

### **Error Response:**
* **Code:** 401 UNAUTHORIZED <br />
* **Code:** 500 INTERNAL SERVER ERROR<br /> 
  * **Reason** database error
