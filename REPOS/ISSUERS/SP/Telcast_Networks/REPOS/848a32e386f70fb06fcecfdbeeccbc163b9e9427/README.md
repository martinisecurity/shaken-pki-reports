# STIR/SHAKEN Certificate Repository Compliance

## Telcast Networks

Name: `https://ssc.getsipnav.com/certs/9861b7e493ab608182e57524671889bf94cd2fc8`\
Tested At: 02 Dec 22 07:29 UTC\
Time: 61ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 02 Dec 22 07:30 UTC