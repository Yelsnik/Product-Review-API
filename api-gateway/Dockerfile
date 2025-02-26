FROM node:23.8-slim AS Builder

WORKDIR /app
COPY package*.json ./

RUN npm ci
COPY . .
RUN npm run build

FROM node:23.8-slim AS PROD

WORKDIR /app
COPY --from=Builder /app/dist ./dist
COPY --from=Builder /app/package.json ./package.json
COPY --from=Builder /app/package-lock.json ./package-lock.json

RUN npm ci --production

EXPOSE 4000
CMD [ "node", "dist/src/main" ]