FROM node:20-alpine AS builder
WORKDIR /code
COPY web/package*.json  ./
RUN npm ci
COPY web/. .
RUN npm run build

FROM devforth/spa-to-http:latest
COPY --from=builder /code/dist/ .