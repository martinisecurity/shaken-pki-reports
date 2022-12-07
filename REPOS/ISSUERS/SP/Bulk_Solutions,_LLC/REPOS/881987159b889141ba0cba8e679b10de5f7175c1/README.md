# STIR/SHAKEN Certificate Repository Compliance

## Bulk Solutions, LLC

Name: `https://bulkvs-cr.s3.amazonaws.com/644J_2022120501.pem`\
Tested At: 07 Dec 22 18:45 UTC\
Time: 47ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 07 Dec 22 18:54 UTC