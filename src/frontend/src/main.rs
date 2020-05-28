#![feature(proc_macro_hygiene, decl_macro)]

#[macro_use] extern crate rocket;

use rocket::response::NamedFile;
use std::io::Result;

#[get("/")]
fn index() -> Result<NamedFile> {
    NamedFile::open("static/index.html")
}

fn main() {
    rocket::ignite().mount("/", routes![index]).launch();
}
