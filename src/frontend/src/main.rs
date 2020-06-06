#![feature(proc_macro_hygiene, decl_macro)]

#[macro_use] extern crate rocket;

use rocket::response::NamedFile;
use std::io::Result;

#[get("/")]
fn index() -> Result<NamedFile> {
    NamedFile::open("static/index.html")
}

#[get("/index.js")]
fn js() -> Result<NamedFile> {
    NamedFile::open("static/index.js")
}

#[catch(404)]
fn not_found() -> Result<NamedFile> {
    NamedFile::open("static/404.html")
}

fn main() {
    rocket::ignite().mount("/", routes![index, js]).register(catchers![not_found]).launch();
}
