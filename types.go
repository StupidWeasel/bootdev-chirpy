package main

import (
  "github.com/google/uuid"
  "github.com/StupidWeasel/bootdev-chirpy/internal/database"
  "sync/atomic"
  "time"
)

type apiConfig struct {
  database *database.Queries
  platform string
  secret string
  users userFuncs
  admin adminFuncs
  messages msgFuncs
  metrics metricFuncs
}

type userFuncs struct{
  dummyHash string
  cfg *apiConfig
}
type adminFuncs struct{
  cfg *apiConfig
}
type msgFuncs struct{
  cfg *apiConfig
}
type metricFuncs struct{
  fileserverHits atomic.Int32
  cfg *apiConfig
}

type createChirpMessage struct{
  Body string `json:"body"`
}

type chirpMessage struct{
  ID uuid.UUID `json:"id"`
  CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`
  Body string `json:"body"`
  UserID uuid.UUID `json:"user_id"`
}

type getchirpMessage struct{
  ID uuid.UUID `json:"id"`
}

type chirpUser struct{
  ID uuid.UUID `json:"id"`
  CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`
  Email string `json:"email"`
}

type chirpUserLogin struct{
  ID uuid.UUID `json:"id"`
  CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`
  Email string `json:"email"`
  Token string `json:"token"`
  RefreshToken string `json:"refresh_token"`
}

type chirpRefreshAuth struct{
  Token string `json:"token"`
}

type createChirpUser struct{
  Email string `json:"email"`
  Password string `json:"password"`
}

type updateChirpUser struct{
  ID uuid.UUID
  Email string `json:"email"`
  Password string `json:"password"`
}

type loginChirpUser struct{
  Email string `json:"email"`
  Password string `json:"password"`
  ExpiresInSeconds int32 `json:"expires_in_seconds"`
}

type apiResponse struct{
  CleanedBody string `json:"cleaned_body"`
}

type apiErrorResponse struct{
  ErrorMsg string `json:"error"`
}

