FROM node:16-alpine
WORKDIR /app

# Copy package.json and yarn.lock to the container
COPY package.json yarn.lock ./

# Install dependencies with Yarn
RUN yarn install

# Copy the rest of the project files to the container
COPY . .

# Build the project
RUN yarn build

ENV NODE_ENV production
EXPOSE 3000

# Start the server
CMD ["yarn", "start"]
