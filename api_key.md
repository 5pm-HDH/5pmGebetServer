
**New Authentication Key**
----
  Insert new key into authentication table
  
  3 key types:
  * `Type 0` Master-Key
  * `Type 1` Beter-Key / Editor-Key
  * `Type 2` Beamer-Key / Monitor-Key

* **URL**

  /api/key

* **Method:**

  `POST` 
  
*  **URL Params**

   **Required:**
   `key=[string]` (master key authentication)


* **Data Params**

  **Content:** `{ key : [super_secure_key], type : [int] }`

* **Success Response:**
  
  * **Code:** 200 <br />

 
* **Error Response:**

  * **Code:** 401 UNAUTHORIZED <br />


* **Notes:**


**Delete Authentication Key**
----
  Delete a key from authentication table
  
* **URL**

  /api/key

* **Method:**

  `PUT` 
  
*  **URL Params**

   **Required:**
   `key=[string]` (master key authentication)


* **Data Params**

  **Content:** `{ key : [super_secure_key] }`

* **Success Response:**
  
  * **Code:** 200 <br />

 
* **Error Response:**

  * **Code:** 401 UNAUTHORIZED <br />


* **Notes:**

