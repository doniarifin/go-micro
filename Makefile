# service
SERVICES = auth msg order

# dir service
AUTH_DIR = app/auth
MSG_DIR = app/msg
ORDER_DIR = app/order

# run service
.PHONY: all auth msg order

all: auth msg order
	@echo "all service running!"

auth:
	@echo "Run Auth Service..."
	@cd $(AUTH_DIR) && go run main.go &

msg:
	@echo "Run Msg Service..."
	@cd $(MSG_DIR) && go run main.go &

order:
	@echo "Run order Service..."
	@cd $(ORDER_DIR) && go run main.go &

# Stop all service
stop:
	@echo "Menghentikan semua service..."
	@pkill -f "go run main.go"
