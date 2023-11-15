# STIR/SHAKEN Certificate Repository Compliance

## Falcon Endeavors

Name: `https://ssc.getsipnav.com/certs/121b5f2f8e2239ba93c6af3c31d4e1f30d04b832`\
Tested At: 15 Nov 23 18:05 UTC\
Time: 51ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 15 Nov 23 18:10 UTC