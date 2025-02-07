FROM node:20.12.0

ENV HOST=0.0.0.0

WORKDIR /app

COPY package*.json ./

RUN npm ci

RUN npm install -g pm2

COPY . .

RUN npm run build

CMD ["pm2-runtime", "start", "ecosystem.config.cjs"]
