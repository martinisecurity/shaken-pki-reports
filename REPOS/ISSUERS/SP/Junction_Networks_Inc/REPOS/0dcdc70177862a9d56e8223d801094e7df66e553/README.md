# STIR/SHAKEN Certificate Repository Compliance

## Junction Networks Inc

Name: `https://stir-shaken.jnctn.net/onsip-stir-shaken.pem`\
Tested At: 27 Nov 23 23:25 UTC\
Time: 169ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 27 Nov 23 23:28 UTC