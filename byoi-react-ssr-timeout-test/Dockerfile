FROM node:20-alpine AS build
WORKDIR /app
COPY package.json package-lock.json* ./
RUN npm install --omit=dev
COPY server.js ./
COPY public ./public

FROM node:20-alpine
WORKDIR /app
COPY --from=build /app ./

# Create a user with a known UID/GID within range 10000-20000.
# This is required by Choreo to run the container as a non-root user.
RUN addgroup -g 10014 choreo && \
    adduser --disabled-password --no-create-home --uid 10014 --ingroup choreo choreouser
USER 10014

ENV PORT=8080
EXPOSE 8080
ENTRYPOINT ["node", "server.js"]
