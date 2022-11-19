FROM nginx:alpine

#USER nginx

COPY nginx.conf /etc/nginx/nginx.conf

WORKDIR /usr/share/nginx/html
RUN rm -rf ./*

# DEFAULT ASSETS
COPY assets/ assets/

# VOLUME FOR CUSTOM EXTENSIONS
VOLUME /usr/share/nginx/html/assets

EXPOSE 80

ENTRYPOINT ["nginx", "-g", "daemon off;"]

