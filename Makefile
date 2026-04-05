# ====== CONFIG ======
MODULE := github.com/lejeunel/go-image-annotator-v2
SPEC := adapters/api/openapi.yaml

OAPI := oapi-codegen

MODELS_PKG := adapters/api/models
SERVER_PKG := adapters/api/server

MODELS_OUT := $(MODELS_PKG)/models.gen.go
SERVER_OUT := $(SERVER_PKG)/server.gen.go

# ====== TARGETS ======

.PHONY: all api-code clean

all: api-code

api-code: $(MODELS_OUT) $(SERVER_OUT)

# --- Generate models (types only) ---
$(MODELS_OUT): $(SPEC)
	mkdir -p $(MODELS_PKG)
	$(OAPI) \
		-generate types \
		-package models \
		-o $(MODELS_OUT) \
		$(SPEC)

# --- Generate server (interfaces only, using models) ---
$(SERVER_OUT): $(SPEC) $(MODELS_OUT)
	mkdir -p $(SERVER_PKG)
	$(OAPI) \
		-generate types,std-http-server \
		-package server \
		-o $(SERVER_OUT) \
		-import-mapping $(MODULE)/$(MODELS_PKG):$(MODULE)/$(MODELS_PKG) \
		$(SPEC)

# --- Cleanup generated files ---
clean:
	rm -f $(MODELS_OUT) $(SERVER_OUT)
