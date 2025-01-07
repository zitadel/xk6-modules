# Zitadel xk6 modules

Provides helper functions for the [Zitadel load tests](https://github.com/zitadel/zitadel/tree/main/load-test).

## Provided functions

### `signJWTProfileAssertion`

Signs a JWT token request using [OIDC](github.com/zitadel/oidc) and returns it as string.

The function can be used as follows:

```ts
// @ts-ignore Import module
import zitadel from 'k6/x/zitadel';

const assertion = zitadel.signJWTProfileAssertion(
  this.keyPayload.userId,
  this.keyPayload.keyId,
  {
    audience: [Config.host],
    expiration: this.keyPayload.expiration,
    key: privateKey
});
```
