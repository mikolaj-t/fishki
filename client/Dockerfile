FROM node:19.0-alpine3.16 as build
RUN mkdir /fishki
WORKDIR /fishki

COPY . .
RUN npm install
RUN npm run build

EXPOSE 5174
CMD ["node", "./build/index.js"]
#CMD ["ls -R"]

# FROM alpine:3.16.0
# COPY --from=build /app/public /fishki
# EXPOSE 5173
# CMD ["npm run build"]