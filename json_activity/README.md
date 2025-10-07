* Created a struct to save the unmarshalled JSON data.
* The User struct contains name, age, email, and image URL.
* Stored the slice of structs in another struct, as the JSON data has a top-level key and doesn’t start with a slice directly.
* Fetched the response from the JSON and stored the byte information in a variable.
* Unmarshalled the bytes into the struct containing the slice of users.
* Converted the image URLs into Base64-encoded strings.
* Marshalled the updated slice of structs into JSON format.
* Saved the marshalled JSON into a file using os.WriteFile().