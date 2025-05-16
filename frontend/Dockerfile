# Build stage
FROM node:18-alpine AS build

WORKDIR /app

COPY package*.json ./
RUN npm install

COPY . .
RUN npm run build

# Production stage
FROM nginx:stable-alpine

# Install OpenSSL untuk generate sertifikat SSL
RUN apk add --no-cache openssl

# Buat direktori untuk sertifikat SSL
RUN mkdir -p /etc/nginx/ssl

# Generate self-signed certificate
RUN openssl req -x509 -nodes -days 365 -newkey rsa:2048 \
    -keyout /etc/nginx/ssl/nginx.key \
    -out /etc/nginx/ssl/nginx.crt \
    -subj "/C=ID/ST=Jakarta/L=Jakarta/O=Kos Management/CN=kos-management.com"

# Copy build output
COPY --from=build /app/dist /usr/share/nginx/html

# Copy nginx config
COPY nginx.conf /etc/nginx/conf.d/default.conf

# Expose nginx default port
EXPOSE 80 443

CMD ["nginx", "-g", "daemon off;"]
