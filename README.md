services:
  db:
    image: postgres:16.4
    volumes:
      - ./initDB:/docker-entrypoint-initdb.d
      - db_data:/var/lib/postgresql/data
    restart: always
    ports:
      - "5432:5432"
    environment:
      POSTGRES_USER: root
      POSTGRES_PASSWORD: secret
      POSTGRES_DB: university_db

  app:
    build:
      context: ./sanatorio-back
      dockerfile: Dockerfile
    depends_on:
      - db
    env_file:
      - ./sanatorio-back/.env
    restart: always
    ports:  # Cambio de expose a ports para acceso externo si es necesario
      - "8080:8080"

  nginx:
    image: nginx:latest
    ports:
      - "80:80"
    volumes:
      - ./nginx.conf:/etc/nginx/nginx.conf
    depends_on:
      - app

volumes:
  db_data:
