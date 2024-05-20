Docker Compose Usage

**Development**:
`docker-compose --env-file .env.development up`

**Production**:
`docker-compose --env-file .env.production up -d`

**Migrations**:
**_Up Migrations_**:
`docker run --rm -v $(pwd)/migrations:/migrations --network host migrate/migrate -path=/migrations/ -database mysql://profile_user:profile_pass@tcp(localhost:3306)/social_media_profile?charset=utf8mb4&parseTime=True&loc=Local up`

**_Down Migrations_**:
`docker run --rm -v $(pwd)/migrations:/migrations --network host migrate/migrate -path=/migrations/ -database mysql://profile_user:profile_pass@tcp(localhost:3306)/social_media_profile?charset=utf8mb4&parseTime=True&loc=Local down`

