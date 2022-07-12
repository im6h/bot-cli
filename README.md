### Feature:

1. Fetch quote from anime api (for funny)

 - Use command to fetch random quote: 
 ```sh
	go-cli anime
 ```

 - Use command to fetch quote with charactor by name
 ```sh
	go-cli anime --charactor=naruto
 ```

 - Use command to fetch quote with anime by name
 ```sh
	go-cli anime --name="Dragon Ball"
 ```

 Note: with flag `--charactor` and `--name`, program will use default value to pagination api (page=1)

2. Generate password with 20 character of length and contain special character

 	-  Use command to generate password:
 ```sh
	go-cli generate
 ```
