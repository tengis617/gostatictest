docker build -t test-static . \
&& docker run -p 8080:8080 test-static
GREEN='\033[0;32m'
NC='\033[0m'
echo -e "${GREEN}SUCCESS! server running on port :8080 ...${NC}"