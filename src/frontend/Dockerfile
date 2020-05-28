FROM rust:1.43-slim AS build

WORKDIR /application

RUN rustup default nightly

COPY . ./

RUN cargo build --release

CMD ["cargo", "run", "--release"]
