# Use a multi-stage build to keep the image small
FROM node:20 AS builder
WORKDIR /app
COPY . .

# Final runtime image
FROM python:3.12-slim
WORKDIR /app
COPY --from=builder /app /app
CMD ["python", "python/api_key_gen.py"]
