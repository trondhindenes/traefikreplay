# TraefikReplay

Traefik plugin for replaying a request against another url. Replayed request will include original headers and request body. The header `X-Original-Url` will be added, containing the original url.

### Configuration
`percentage`: integer. Percentage of traffic you want replayed. Set it to 100 for all traffic, 0 for none   
`replayUrl`: string. URL to invoke   
`onlyIfJson`: bool: Checks if content-type is json, skips replay if it isn't

Example:
```yaml

  middlewares:
    traefikreplay:
      plugin:
        traefikreplay:
          # Take 50% of traffic and replay it if content-type is `application/json`
          percentage: 50
          replayUrl: http://localhost:8000/replay
          OnlyIfJson: true
```