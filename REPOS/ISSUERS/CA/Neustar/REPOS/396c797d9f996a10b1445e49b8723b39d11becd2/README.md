# STIR/SHAKEN Certificate Repository Compliance

## Neustar

Name: `https://cr-partner.ccid.neustar.biz/ccid/authn/v2/certs/176.230`\
Tested At: 01 Dec 22 19:06 UTC\
Time: 41ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 01 Dec 22 19:10 UTC