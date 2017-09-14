docker build -t test-static . \
&& docker run -p 8080:8080 test-static