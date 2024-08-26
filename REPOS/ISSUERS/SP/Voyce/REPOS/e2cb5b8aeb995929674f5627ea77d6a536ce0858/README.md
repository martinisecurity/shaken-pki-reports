# STIR/SHAKEN Certificate Repository Compliance

## Voyce

Name: `https://ssc.getsipnav.com/certs/b1f83b2a8803ca5297ced7cb375cc6f5e2b44f85`\
Tested At: 26 Aug 24 17:57 UTC\
Time: 37ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 26 Aug 24 18:03 UTC