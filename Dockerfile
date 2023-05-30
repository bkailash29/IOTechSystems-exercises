# Use the official Go base image
FROM golang:1.16

# Set the working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the solution
RUN go build -o solution

# Set the entrypoint to run the program
ENTRYPOINT ["./solution"]

