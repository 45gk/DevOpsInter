FROM debian:11

RUN apt-get update && apt-get install -y nginx
RUN apt-get clean && rm -rf /var/lib/apt/lists/*
ADD nginx.conf /etc/nginx/conf.d/

EXPOSE 80


CMD ["nginx", "-g", "daemon off;"]
