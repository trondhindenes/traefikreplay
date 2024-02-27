# TraefikReplay

Traefik plugin for replaying a request against another url

### Configuration
percentage: integer. Percentage of traffic you want replayed. Set it to 100 for all traffic, 0 for none
replayUrl: string. URL to invoke
onlyIfJson: bool: Checks if content-type is json, skips replay if it isn't

Example:
```

  middlewares:
    traefikreplay:
      plugin:
        traefikreplay:
          # Take 50% of traffic and replay it if content-type is `application/json`
          percentage: 50
          replayUrl: http://localhost:8000/replay
          OnlyIfJson: true
```