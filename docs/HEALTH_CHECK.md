# Health Check Endpoint

## Overview

The Money Manager API provides a health check endpoint for monitoring server and database status.

## Endpoint

```
GET /health
```

**No authentication required** — this endpoint is public and does not require a JWT token.

## Response

### Healthy (HTTP 200)

When the server is running and the database is reachable:

```json
{
  "status": "ok",
  "service": "money-manager",
  "database": "ok"
}
```

### Degraded (HTTP 503)

When the server is running but the database is unreachable:

```json
{
  "status": "degraded",
  "service": "money-manager",
  "database": "unreachable"
}
```

## Usage

### Command Line

```bash
curl http://localhost:8080/health
```

### Docker Health Check

Add to your `docker-compose.yml`:

```yaml
services:
  api:
    healthcheck:
      test: ["CMD", "curl", "-f", "http://localhost:8080/health"]
      interval: 30s
      timeout: 10s
      retries: 3
      start_period: 40s
```

### Kubernetes Liveness/Readiness Probe

```yaml
livenessProbe:
  httpGet:
    path: /health
    port: 8080
  initialDelaySeconds: 30
  periodSeconds: 10

readinessProbe:
  httpGet:
    path: /health
    port: 8080
  initialDelaySeconds: 5
  periodSeconds: 5
```

## Monitoring

Use this endpoint to:
- Verify the server is running
- Check database connectivity
- Configure load balancer health checks
- Set up uptime monitoring (e.g., UptimeRobot, Pingdom)
- Implement container orchestration health checks

## Notes

- The endpoint returns immediately if the database pool is not configured (testing scenarios)
- Database connectivity is verified with a simple `PING` operation
- Response time should be < 100ms under normal conditions
