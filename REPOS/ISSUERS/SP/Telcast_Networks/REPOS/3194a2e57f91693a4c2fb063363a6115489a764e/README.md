# STIR/SHAKEN Certificate Repository Compliance

## Telcast Networks

Name: `https://ssc.getsipnav.com/certs/2ae1590ba7b804c703e2af2842d1290fe9a70fc0`\
Tested At: 28 Nov 23 20:13 UTC\
Time: 49ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 28 Nov 23 20:21 UTC