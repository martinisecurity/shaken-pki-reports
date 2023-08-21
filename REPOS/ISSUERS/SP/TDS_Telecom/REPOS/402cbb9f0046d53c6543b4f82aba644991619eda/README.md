# STIR/SHAKEN Certificate Repository Compliance

## TDS Telecom

Name: `https://cdn-cr.cgah.tnsi.com/certs/9b2b8bbbd5a5a28082425a21229434097aafd9b2`\
Tested At: 21 Aug 23 20:00 UTC\
Time: 8ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 21 Aug 23 20:18 UTC