# STIR/SHAKEN Certificate Repository Compliance

## Interbel Telephone

Name: `https://certs.iverify-aninetworks.net/interbeltelephone_20260610.crt`\
Tested At: 18 Aug 25 20:16 UTC\
Time: 169ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header has 'max-age' directive but it's value is less than 24 hours |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 18 Aug 25 21:13 UTC