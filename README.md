# go_organising-db-access

>In the context of a web application what would you consider a Go best practice for accessing the database in (HTTP or other) handlers?

The replies it got were a genuinely interesting mix.
- Some people advised using dependency injection
- A few favoured the simplicity of using global variables
- Others suggested putting the connection pool pointer into the request context.

The right answer depends on the project.
- What's the overall structure and size of the project?
- What's your approach to testing?
- How is it likely to grow in the future?

All these things and more should play a part when you pick an approach to take.

So in this post we're going to take a look at 4 different methods for organizing your code and structuring access to your database connection pool, and explain when they may — or may not — be a good fit for your project.

---
### 1. Using a global variable
Using a global variable to store the database connection pool like this is potentially a good fit when:
- Your application is small and simple, and keeping track of globals in your head isn't a problem.
- Your HTTP handlers are spread across multiple packages, but all your database-related code lives in one package.
- You don't need to mock the database for testing purposes.

The drawbacks of using global variables are [well-documented](https://softwareengineering.stackexchange.com/questions/148108/why-is-global-state-so-evil), but in practice I've found that for small and simple projects using a global variable like this works just fine, and it's (arguably) clearer and easier to understand than some of the other approaches we'll look at in this post.

For more complex applications — where your handlers have more dependencies beyond just the database connection pool — it's generally better to use dependency injection instead of storing everything in global variables.

The approach we've taken here also doesn't work if your database logic is spread over multiple packages, although — if you really want to — you could a separate `config` package containing an exported `DB` global variable and `import "yourproject/config"` into every file that needs it.

---
### 2. Dependency injection
In a more complex web application there are probably additional application-level objects that you want your handlers to have access to. For example, you might want your handlers to also have access to a shared logger, or a template cache, as well your database connection pool.

Rather than storing all these dependencies in global variables, a neat approach is to store them in a single custom `Env struct` like so:
```go
type Env struct {
    db *sql.DB
    logger *log.Logger
    templates *template.Template
}
```

The nice thing about this is that you can then define your handlers as methods against `Env`. This gives you a easy and idiomatic way of making the connection pool (and any other dependencies) available to your handlers.
