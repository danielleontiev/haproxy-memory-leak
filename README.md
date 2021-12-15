# haproxy memory leak

1. Clone repository
2. Make sure `docker` and `docker-compose` are installed on the system
3. `cd haproxy-memory-leak`
4. Start containers `docker-compose up --build`
5. Open Grafana on http://localhost:3000/
6. There are two dashboards - `HAProxy 2 Full` (all HAProxy metrics) and `Docker Containers`
7. Open `Docker Containers` dashboard in Grafana, scroll to `Container Memory Usage` widget
8. Pick `haproxy23` and `haproxy24` containers
9. Observe that `haproxy24` container does not free memory
10. Stop containers `docker-compose down`
11. Comment lines 18-21 and 4 in `application/main.go` (do not return HTTP 500 errors)
12. Repeat test
13. Observe that now `haproxy24` maintains memory correctly
