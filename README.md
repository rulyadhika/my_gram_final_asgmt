## Installation

You can clone the repository with this command, or download this [zip](https://github.com/rulyadhika/my_gram_final_asgmt/archive/refs/heads/main.zip) file.

```bash
> git clone https://github.com/rulyadhika/my_gram_final_asgmt
```

## Configuration
1. Change terminal directory to my_gram_final_asgmt folder
```bash
> cd my_gram_final_asgmt
```

2. Run this command
```bash
> go mod download
```

3. Duplicate .env.example file and rename it to .env . Or you can run this command
```bash
> copy .env.example .env
```

4. Configure your .env file

5. Duplicate makefile.example file and rename it to makefile . Or you can run this command
```bash
> copy makefile.example makefile
```

6. Configure your database in makefile file

7. Run this command to get database migrations CLI
```bash
> go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

8. Ensure your database config is setup correctly, then run this command
```bash
> make migrate-up
```

9. Run local development server
```bash
> go run main.go
```