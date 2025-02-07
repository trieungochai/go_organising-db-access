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
