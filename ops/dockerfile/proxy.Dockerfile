# syntax=docker/dockerfile:1.4

FROM node:19.4-alpine as builder

# WARNING: if container limit < MAX_OLD_SPACE_SIZE => Killed
# Docs: https://developer.ibm.com/languages/node-js/articles/nodejs-memory-management-in-container-environments/
ARG MAX_OLD_SPACE_SIZE=8192
ENV NODE_OPTIONS=--max_old_space_size=${MAX_OLD_SPACE_SIZE}

WORKDIR /app
COPY ./internal/services/proxy /app/

RUN npm ci --cache .npm --prefer-offline --force

RUN npm run build

FROM node:19.4-alpine

# Install dependencies
RUN \
  apk update && \
  apk add --no-cache curl

HEALTHCHECK \
  --interval=5s \
  --timeout=5s \
  --retries=3 \
  CMD curl -f localhost:3020/ready || exit 1

WORKDIR /app
COPY --from=builder /app/dist /app/dist
COPY --from=builder /app/package.json /app/package.json
COPY --from=builder /app/node_modules /app/node_modules

CMD ["npm", "run", "prod"]
