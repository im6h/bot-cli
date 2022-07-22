### Install

1. Install program in global, use command:

   ```
   make publish
   ```

2. For debug and build app, use command:

   ```
   make build
   ```

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

   - Use command to generate password:

   ```sh
       go-cli generate
   ```

3. Fetch article from `dev.to` api:

   - Use command to fetch top article in `dev.to` site(default limit=10,size=1):

   ```sh
       go-cli devto

       go-cli devto 1 10
   ```
