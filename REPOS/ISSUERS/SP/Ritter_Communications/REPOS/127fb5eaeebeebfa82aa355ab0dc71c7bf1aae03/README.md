# STIR/SHAKEN Certificate Repository Compliance

## Ritter Communications

Name: `https://cdn-cr.cgah.tnsi.com/certs/ed02c72bad69b41f5a672ea7e1eb3f67a7ad3be0`\
Tested At: 21 Aug 23 20:00 UTC\
Time: 9ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 21 Aug 23 20:18 UTC