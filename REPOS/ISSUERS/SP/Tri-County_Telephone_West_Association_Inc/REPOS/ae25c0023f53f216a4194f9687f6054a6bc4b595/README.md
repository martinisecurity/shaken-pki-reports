# STIR/SHAKEN Certificate Repository Compliance

## Tri-County Telephone West Association Inc

Name: `https://cdn-cr.cgah.tnsi.com/certs/260bd48b7dcfa43beee9f59287f1e4d87fe3fb5f`\
Tested At: 06 Jul 23 13:53 UTC\
Time: 23ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 06 Jul 23 14:08 UTC