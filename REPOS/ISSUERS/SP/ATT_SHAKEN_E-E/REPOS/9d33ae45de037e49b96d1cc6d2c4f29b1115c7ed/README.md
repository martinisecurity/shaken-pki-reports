# STIR/SHAKEN Certificate Repository Compliance

## ATT SHAKEN E-E

Name: `https://cert.sticr.att.net:8443/certs/att/abbf5398-e1e1-42af-96a7-092303b168ba`\
Tested At: 11 Jan 23 21:48 UTC\
Time: 2492ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 11 Jan 23 21:59 UTC