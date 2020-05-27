FROM rust:1.43-slim AS base

WORKDIR /application

RUN rustup target add x86_64-unknown-linux-musl \
    && rustup default nightly

COPY Cargo.toml Cargo.lock ./
RUN cargo build --release

COPY src ./src
RUN cargo install --target x86_64-unknown-linux-musl --path .

FROM scratch

COPY --from=build /usr/local/cargo/bin/frontend .
CMD ["./frontend"]
