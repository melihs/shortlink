# syntax=docker/dockerfile:1.4

ARG API_URI

# Install dependencies only when needed
FROM node:19.3-alpine as deps

# Check https://github.com/nodejs/docker-node/tree/b4117f9333da4138b03a546ec926ef50a31506c3#nodealpine to understand why libc6-compat might be needed.
RUN apk add --no-cache libc6-compat
RUN npm config set ignore-scripts false

WORKDIR /app
COPY ./ui/eslint /eslint
COPY ./ui/landing/package.json ./ui/landing/package-lock.json ./

RUN npm ci --cache .npm --prefer-offline --force

# Rebuild the source code only when needed
FROM node:19.3-alpine as builder

ARG API_URI
ENV API_URI=${API_URI}

WORKDIR /app
RUN echo @shortlink-org:registry=https://gitlab.com/api/v4/packages/npm/ >> .npmrc

COPY ./ui/landing /app/
COPY --from=deps /app/node_modules ./node_modules

RUN npm run generate

# Production image, copy all the files and run next
FROM nginxinc/nginx-unprivileged:1.23-alpine

# Delete default config
RUN rm /etc/nginx/conf.d/default.conf

WORKDIR /usr/share/nginx/html

# Use root user to copy dist folder and modify user access to specific folder
USER root

# Install dependencies
RUN \
  apk update && \
  apk add --no-cache curl

HEALTHCHECK \
  --interval=5s \
  --timeout=5s \
  --retries=3 \
  CMD curl -f localhost:8080 || exit 1

# Copy application and custom NGINX configuration
COPY --from=builder /app/out ./
COPY ./ops/dockerfile/conf/ui.local /etc/nginx/conf.d/ui.local
COPY ./ops/docker-compose/gateway/nginx/conf/nginx.conf /etc/nginx/nginx.conf
COPY ./ops/docker-compose/gateway/nginx/conf/templates /etc/nginx/template

# Setup unprivileged user 1001
RUN chown -R 1001 /usr/share/nginx/html

# Use user 1001
USER 1001
