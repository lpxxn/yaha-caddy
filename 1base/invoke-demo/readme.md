
To disable TLS (HTTPS) in Caddy, you have a few options:

1. Use HTTP instead of HTTPS in your Caddyfile:

Instead of:
```
example.com {
    # Your directives here
}
```

Use:
```
http://example.com {
    # Your directives here
}
```

2. Explicitly set the auto_https directive to off:

```
{
    auto_https off
}

example.com {
    # Your directives here
}
```

3. Use the http_port directive to listen on port 80:

```
{
    http_port 80
}

example.com {
    # Your directives here
}
```

4. For specific domains, you can use the tls internal directive:

```
example.com {
    tls internal
    # Your other directives here
}
```

Remember that disabling TLS is generally not recommended for production environments as it compromises security. Only do this for development or testing purposes, or if you have a specific reason to serve your content over HTTP.

Also, note that some browsers and operating systems may still try to upgrade to HTTPS by default, so you might need to explicitly type "http://" in the address bar to access your site without TLS.