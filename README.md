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

1. Fetch quote from `anime api` (for funny)

   - Use command to fetch random quote:

   ```sh
       go-cli quote
   ```

   - Use command to fetch quote with charactor by name

   ```sh
       go-cli quote charactor --name=naruto
   ```

   - Use command to fetch quote with anime by name

   ```sh
       go-cli quote anime --name="Dragon Ball"
   ```

   Note:

   | Flag  | Default - Required |
   | ----- | ------------------ |
   | page  | 1 - no             |
   | limit | 10 - yes           |

2. Generate password with 20 character of length and contain special character

   - Use command to generate password:

   ```sh
       go-cli generate
   ```

3. Fetch article from `dev.to` api:

   - Use command to fetch top article in `dev.to` site(default limit=10,size=1):

   ```sh
       go-cli devto top

       go-cli devto top 1 10
   ```

   Note:

   | Flag  | Default - Required |
   | ----- | ------------------ |
   | page  | 1 - no             |
   | limit | 10 - yes           |

4. Generate example environment file from `.env`:

   - Use command to generate `.env.example`:

   ```sh
      go-cli cp_env
   ```

   Note: Use command inside folder contains `.env` file.

5. Show your current `IP Address`:

   - Use command to get current `IP Address` public on `WAN`:

   ```sh
      go-cli ip public
   ```

   - Use command to get current `IP Address` public on `LAN`:

   ```sh
      go-cli ip local
   ```

6. Interact with `github`:

   - Use command create new a repo:

   ```sh
      go-cli ph create --name=repo_name
   ```

   Note:

   | Flag        | Required |
   | ----------- | -------- |
   | name        | yes      |
   | description | no       |
   | private     | no       |
   | init        | no       |

   - Use command delete the exist repo:

   ```sh
      go-cli ph delete --name=repo_name
   ```
