# STIR/SHAKEN Certificate Repository Compliance

## Altice USA

Name: `https://cr.ccid.alticeusa.com/ccid/authn/v2/certs/11011.10003`\
Tested At: 22 Aug 24 15:29 UTC\
Time: 482ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 22 Aug 24 15:44 UTC