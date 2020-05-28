FROM rust:1.43-slim AS build

WORKDIR /application

RUN rustup default nightly

COPY . ./

RUN cargo build --release

ENV ROCKET_PORT=8080

CMD ["cargo", "run", "--release"]
