# STIR/SHAKEN Certificate Repository Compliance

## ATT

Name: `https://cert.sticr.att.net:8443/certs/att/b8b9ae20-018a-4ff1-90bd-47d436341bb8`\
Tested At: 22 Aug 24 15:19 UTC\
Time: 156ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 22 Aug 24 16:06 UTC