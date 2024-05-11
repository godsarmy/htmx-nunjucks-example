# htmx-nunjucks-example

## What's this?

Provide [htmx](https://htmx.org) examples applied in restAPI and [nunjucks](https://mozilla.github.io/nunjucks) template.

[Official HTMX Examples](https://htmx.org/examples/) are implemented with hypermedia calls. In this project, these examples are reimplemented with pure JSON based restAPI calls with the help of htmx [client side template](https://htmx.org/extensions/client-side-templates/).
Backend code is implemented by Golang [Gin framework](https://gin-gonic.com) for easy execution and simplicity.

## Why?

Due to various reason, it is not alway easy to convert JSON restAPIs into hypermedia calls. This project leverages client side template technology to render dynamic html pages for JSON restAPIs.

## Usage

* Make sure you have go compiler and Launch `make all`.
* Enter example subdirectory and launch http server by `./main`.
* Open `http://localhost:8080` in browser.

## Related Articles

* [Handle Json API Results in htmx](https://marcus-obst.de/blog/htmx-json-handling)
* [Creating a Spaceflight News Blog with HTMX & JSON API](https://jerrynsh.com/creating-a-spaceflight-news-blog-with-htmx-and-json-api)
