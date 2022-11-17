# STIR/SHAKEN Certificate Repository Compliance

## Bulk Solutions, LLC

Name: `https://bulkvs-cr.s3.amazonaws.com/644J_20210615001.crt`\
Tested At: 17 Nov 22 19:10 UTC\
Time: 67ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 17 Nov 22 19:21 UTC