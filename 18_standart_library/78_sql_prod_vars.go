// Setting SetMaxOpenConns is important for production-grade applications.
// Because the default value is unlimited, we should set it to make sure it fits what
// the underlying database can handle.

//  The value of SetMaxIdleConns (default: 2) should be increased if our applica-
// tion generates a significant number of concurrent requests. Otherwise, the
// application may experience frequent reconnects.

//  Setting SetConnMaxIdleTime is important if our application may face a burst of
// requests. When the application returns to a more peaceful state, we want to
// make sure the connections created are eventually released.

//  Setting SetConnMaxLifetime can be helpful if, for example, we connect to a
// load-balanced database server. In that case, we want to ensure that our applica-
// tion never uses a connection for too long.

// For production-grade applications, we must consider these four parameters. We can
// also use multiple connection pools if an application faces different use cases.