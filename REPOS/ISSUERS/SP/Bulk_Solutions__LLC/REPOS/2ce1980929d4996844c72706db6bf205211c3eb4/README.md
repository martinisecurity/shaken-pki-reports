# STIR/SHAKEN Certificate Repository Compliance

## Bulk Solutions, LLC

Name: `https://bulkvs-cr.s3.amazonaws.com/608300.pem`\
Tested At: 24 Nov 23 11:00 UTC\
Time: 53ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 24 Nov 23 11:17 UTC