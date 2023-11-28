# STIR/SHAKEN Certificate Repository Compliance

## Neustar

Name: `https://cr-partner.ccid.neustar.biz/ccid/authn/v2/certs/95.234`\
Tested At: 28 Nov 23 10:28 UTC\
Time: 230ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 28 Nov 23 10:53 UTC